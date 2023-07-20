package foo

import (
	"context"
	"log"

	"github.com/wesleywu/go-lifespan/lifespan"
)

const componentName = "foo"

var initialized = false

func init() {
	lifespan.RegisterBootstrapHook(componentName, true, func(ctx context.Context) error {
		log.Printf("component %s bootstrapped", componentName)
		initialized = true
		return nil
	})
	lifespan.RegisterShutdownHook(componentName, true, func(ctx context.Context) error {
		log.Printf("component %s shut down", componentName)
		initialized = false
		return nil
	})
}

func IsInitialized() bool {
	return initialized
}
