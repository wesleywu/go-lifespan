package lifespan

import (
	"context"
	"log"
	"sync"
)

var (
	bootstrapOnce = &sync.Once{}
	shutdownOnce  = &sync.Once{}
)

func OnBootstrap(ctx context.Context) {
	bootstrapOnce.Do(func() {
		executedHooks := make(map[string]interface{})
		for hookName, hook := range bootstrapHooks {
			err := executeHook(ctx, hook, bootstrapHooks, executedHooks)
			if err != nil {
				if hook.mandatory {
					log.Fatalf("Bootstrap hook %s error %+v\n", hookName, err)
				} else {
					log.Printf("Bootstrap hook %s error %+v\n", hookName, err)
				}
			}
		}
	})
}

func OnShutdown(ctx context.Context) {
	shutdownOnce.Do(func() {
		executedHooks := make(map[string]interface{})
		for hookName, hook := range shutdownHooks {
			err := executeHook(ctx, hook, shutdownHooks, executedHooks)
			if err != nil {
				if hook.mandatory {
					log.Fatalf("Shutdown hook %s error %+v\n", hookName, err)
				} else {
					log.Printf("Shutdown hook %s error %+v\n", hookName, err)
				}
			}
		}
	})
}

func executeHook(ctx context.Context, hook hookDef, allHooks map[string]hookDef, executedHooks map[string]interface{}) error {
	if _, ok := executedHooks[hook.name]; ok {
		return nil
	}
	dependencies := hook.dependencies
	if len(dependencies) > 0 {
		for _, depHookName := range dependencies {
			depHook, ok := allHooks[depHookName]
			if !ok {
				log.Printf("Cannot find dependent hook %s for current hook %s\n", depHookName, hook.name)
				continue
			}
			err := executeHook(ctx, depHook, allHooks, executedHooks)
			if err != nil {
				log.Fatalf("Bootstrap hook %s error %+v\n", depHookName, err)
			}
		}
	}
	err := hook.f(ctx)
	if err != nil {
		log.Fatalf("Bootstrap hook %s error %+v\n", hook.name, err)
	}
	executedHooks[hook.name] = nil
	return nil
}
