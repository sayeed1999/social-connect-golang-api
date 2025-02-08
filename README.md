# social-connect-golang-api

To run the api locally, follow these steps:-

Create a file `scripts/set-env.sh` following the example file provided and replace 'xxx' with confidential values:

```bash
export API_DATABASE_NAME=social-connect
export API_DATABASE_USER=xxx
export API_DATABASE_PASSWORD=xxx
export API_DATABASE_PORT=5432

export API_REDIS_PORT=6379

export PGADMIN_EMAIL=xxx
export PGADMIN_PASSWORD=xxx
```

Set necessary environment variables by executing the script:

```bash
. scripts/set-env.sh
```

Get the docker containers up for dependent services e.g postgres, redis, .. by running:

```bash
docker compose up -d
```

Open up terminal from root dir & run:

```bash
go run api/main.go
```

To run all tests in the project, run from root dir:

```bash
go test ./...
```

## Additional Documentations

- [ARCHITECTURE](documentations/ARCHITECTURE.md)
