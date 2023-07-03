package lifespan

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDependencies(t *testing.T) {
	ctx := context.TODO()
	var resultNamesOrdered []string
	printNameFunc := func(name string) func(ctx context.Context) error {
		return func(ctx context.Context) error {
			fmt.Printf("%s started\n", name)
			resultNamesOrdered = append(resultNamesOrdered, name)
			return nil
		}
	}
	RegisterBootstrapHook("1.1.1", true, printNameFunc("1.1.1"), "1.1")
	RegisterBootstrapHook("1.2.1", true, printNameFunc("1.2.1"), "1.2")
	RegisterBootstrapHook("1.2.2", true, printNameFunc("1.2.2"), "1.2", "1.2.1")
	RegisterBootstrapHook("1.3.1", true, printNameFunc("1.3.1"), "1.3")
	RegisterBootstrapHook("1.3.2", true, printNameFunc("1.3.2"), "1.3.1", "1", "1.3")
	RegisterBootstrapHook("1.end", true, printNameFunc("1.end"), "1.3.2")
	RegisterBootstrapHook("1.1", true, printNameFunc("1.1"), "1")
	RegisterBootstrapHook("1", true, printNameFunc("1"))
	RegisterBootstrapHook("1.3", true, printNameFunc("1.3"), "1", "1.2.2")
	RegisterBootstrapHook("1.2", true, printNameFunc("1.2"), "1", "1.1.1")
	OnBootstrap(ctx)
	assert.Equal(t, []string{
		"1",
		"1.1",
		"1.1.1",
		"1.2",
		"1.2.1",
		"1.2.2",
		"1.3",
		"1.3.1",
		"1.3.2",
		"1.end",
	}, resultNamesOrdered)
}
