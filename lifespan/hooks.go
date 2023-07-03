package lifespan

import "context"

var (
	bootstrapHooks map[string]hookDef
	shutdownHooks  map[string]hookDef
)

type hookDef struct {
	name         string
	mandatory    bool
	f            func(ctx context.Context) error
	dependencies []string
}

func init() {
	bootstrapHooks = make(map[string]hookDef)
	shutdownHooks = make(map[string]hookDef)
}

func RegisterBootstrapHook(hookName string, mandatory bool, hookFunc func(ctx context.Context) error, dependentHookNames ...string) {
	bootstrapHooks[hookName] = hookDef{
		name:         hookName,
		mandatory:    mandatory,
		f:            hookFunc,
		dependencies: dependentHookNames,
	}
}

func RegisterShutdownHook(hookName string, mandatory bool, hookFunc func(ctx context.Context) error, dependentHookNames ...string) {
	shutdownHooks[hookName] = hookDef{
		name:         hookName,
		mandatory:    mandatory,
		f:            hookFunc,
		dependencies: dependentHookNames,
	}
}
