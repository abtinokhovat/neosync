# NeoSync 
Breaking through the matrix of order statuses to ensure everything is up to date.


## Development
For bringing up the development you have to first create a docker network for the application components.
```bash
docker network create -d bridge neosync-net
```
Afterward you can run the `make dev-up` command to compose up the components.
Finally, you can run the application with `make dev` command.
```bash
make dev-up # one time command to bring up the infrastructure for development
make dev    # run the application with the compile demon for restarting after changes
make build  # exporting a binary file of the application into the /build folder 
```



## Use Case 1

`Application should periodically update the order statuses`

I implemented this use-case with a cron job:
Flow:

1. Fetching all the providers
2. Fetching all response of the provider adapters (getting all order statuses from provider)
3. Searching for status change in the providers responses
4. Updating the status of the order in the database
5. Send a notification on the ProviderSeen -> PickedUp status change