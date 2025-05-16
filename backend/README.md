# Structure

`api/` -- all api information and handling (payloads, handlers, controller)
`service/` -- only business logic here (**NO external dependency should be imported here**)
`repository/` -- used from service layer to inject external dependency and actually do all the operations
`model/` -- maybe remove but it can be used to transfer data between layers or to standarize a way of processing types
`main.go` -- starts everything

## Flow

All routes are implemented on `api` pacakge, and the callback for each route is should be on `handlers.go` following the same pattern (they're all methods from the `Controller` type). These handlers forward the task to the service layer (as it's done in the example route). The service layer only does business logic, so it should not inject any external dependencies. By business logic, we mean that it orchestrates every call to the repository layer without directly depending on any dependency (for apis, database, data processing, etc). This ensures that the layers are decoupled and more modular.
The repository layer injects every necessary external dependency. In other words, it actually does the job, by effectively fetching an endpoint, inserting or reading some data from somewhere, and so on.

So, the repository layer should offer small functions that does specific and unique tasks so that the service layer can orchestrate these calls and treat the error properly. 