package bar

import (
	"errors"
	"log"

	"github.com/wesleywu/go-lifespan/example/foo"
)

func CallFoo() error {
	initialized := foo.IsInitialized()
	if !initialized {
		return errors.New("component foo was not initialized")
	}
	log.Printf("foo component initialized: %v", foo.IsInitialized())
	return nil
}
