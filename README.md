# Shopify Backend Developer Intern Challenge

This is a REST API built in Go.

An inventory tracking app with CRUD operations -> ideal for SQL databases due to their ACID nature.

For the current requirements, a SQLite or any lightweight SQL database will do it. But if we want to scale later on, we should use something "heavier" like MySQL, Postgres (which is what I've chosen here) or MSSQL.

The service here is implemented following the repository and service pattern. We have one main service defined at src/inv/pkg/api/service.go. It will hold a master repository defined in master.go. The master repo will hold all entities repository (which is only inventory at this point). Each mini repository is responsible for managing all sql transactions related to that its table. This creates loose coupling, separation of concerns and makes it extremely easy to add a new table (i.e: create a new shipments table as an append-only record of all shipments). Operations require multiple tables operations will happen in the master repo.

As of the http handling, Go's net/http library create a copy of api/service to handle every new incoming request.

Build instructions:

```

make // build the server

bash db.sh // build the database and run unit tests

./build/inv server

```
