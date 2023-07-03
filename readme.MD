# Lifespan management for Golang process

This library provides lifespan management of inner components for Golang process. 

## Use case

1) We usually need to initialize some components when the process starts, and release the resources occupied by these components before the end of the process.

   For example, redis

   * at startup, initialized a redis client connected to the redis cluster
   * at shutdown, close the redis client as well as the underlying connection  

2) We need to isolate the initialization/release code of the component from the main process code and any code outside the component package.

   * prevents calling the initializing/releasing functions directly from outside the component package
   * the component should be ready to use after process bootstrapped
   * the resources of component should be released before process shutdown

3) function init() should be avoided for component initializing
   * eliminate the complex logic of component initialing in init() function

4) applying the singleton pattern
   * either the initializing or the releasing code should be executed only once

## Implementation

1) A component may register its bootstrap and shutdown hooks via these two functions:
   * lifespan.RegisterBootstrapHook(...)
   * lifespan.RegisterShutdownHook(...)

2) all bootstrap and shutdown hooks could be executed at boostrap/shutdown stage of process lifespan via these two functions:
    * lifespan.OnBootstrap(ctx context.Context)
    * lifespan.OnShutdown(ctx context.Context)
   
3) This package does not provide automatic bindings to process, because every process may have its own mechanism to implements graceful shutdown, i.e., via signal handling.
We suggested you manually calling OnBootstrap() at process startup and OnShutdown() before process end.

## Example
There is a fully runnable example in ./example directory. 

### main func
* Call lifespan.OnBootstrap() at startup
* Call bar.CallFoo() while bar is using component foo
* initialize a signal handler to listening to either context cancelling or process terminating, calling lifespan.OnShutdown() before exit

### bar
* CallFoo() use component foo, returning error if foo was not initialized.

### foo
* a sample component registering its bootstrap/shutdown hooks via lifespan.RegisterBootstrapHook and lifespan.RegisterShutdownHook when foo package was first seen by any code