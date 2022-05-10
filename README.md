# Notice:

This is built on top of a copy of the shopify backend challenge.

# Shopify ~~Backend Developer~~ Production Intern Challenge

This is a REST API built in Go.

An inventory tracking app mainly with CRUD operations -> ideal for SQL databases due to their ACID nature.

For the current requirements, a SQLite or any lightweight SQL database will do it. But if we want to scale better later on, we should use something "heavier" like MySQL, Postgres (which is what I've chosen here) or MSSQL.

Architecture:

HTTP handlers and servers commands are defined src/inv/cmd. These will then call the invService in pkg.

The service here is implemented following the repository and service pattern. I have one main service defined at src/inv/pkg/api/service.go. It will hold a master repository defined in master.go. The master repo will hold all entities repository (which is only inventory at this point). Each mini repository is responsible for managing all sql transactions related to that its table. This creates loose coupling, separation of concerns and makes it extremely easy to add a new table (i.e: create a new shipments table as an append-only record of all shipments). Operations require multiple tables operations will happen in the master repo.

Error handling:

Any errors occurred during execution will be automatically log into stderr. I also have some custom wrapper for errors that is able to print stack trace and provide some additional metadata about query.

As of the http handling, Go's net/http library create a copy of api/service to handle every new incoming request.

# Build instructions from source:

After cloning run these commands.

```

go mod download

make

bash db.sh

./build/inv server

```

# On Replit:

Link: https://replit.com/@Posrabi/shopify-backend-project-2#Makefile

Everything on Replit is ready to go. All you need to do is press "Run".

This is only a server so you will have to use curl to make requests :)

This can be achieve forking the repl to your account and run curl in the shell. Some example commands our below.

After clicking run and wait for the "Listening on localhost:8081" message.

Try these commands in the shell

```

Create:

curl localhost:8081/create --data '{"item_id":"1","brand":"something","item_name":"test item","item_quantity":10}'

List:

curl localhost:8081/list

Edit:

curl localhost:8081/edit --data '{"item_id":"1","brand":"nothing","item_name":"test item 1","item_quantity":50}'

Delete: This uses query params not JSON

curl localhost:8081/delete?id=1 --request DELETE

Ship:

curl localhost:8081/ship --data '{"item_instance":{"item_id":"1","item_quantity":5}}'

```

A Side note about replit, if you somehow accidentally press stop, to run the server again, run

```

pg_ctl stop

```

And then reload the page or there would be some pretty nasty port conflicts coming from postgres not correctly stopped.
