# Shopify Backend Developer Intern Challenge

This is a REST API built in Go.

An inventory tracking app mainly with CRUD operations -> ideal for SQL databases due to their ACID nature.

For the current requirements, a SQLite or any lightweight SQL database will do it. But if we want to scale later on, we should use something "heavier" like MySQL, Postgres (which is what I've chosen here) or MSSQL.

The service here is implemented following the repository and service pattern. We have one main service defined at src/inv/pkg/api/service.go. It will hold a master repository defined in master.go. The master repo will hold all entities repository (which is only inventory at this point). Each mini repository is responsible for managing all sql transactions related to that its table. This creates loose coupling, separation of concerns and makes it extremely easy to add a new table (i.e: create a new shipments table as an append-only record of all shipments). Operations require multiple tables operations will happen in the master repo.

As of the http handling, Go's net/http library create a copy of api/service to handle every new incoming request.

Build instructions:

After cloning run these commands.

```

make

bash db.sh

./build/inv server

```

# On Replit

This is only a server so you will have to use curl to make requests :)

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

Side note about replit, if you somehow accidentally press stop, to run the server again, run

```

pg_ctl stop

```

And then reload the page or there would be some pretty nasty port conflicts.
