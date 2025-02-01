# social-connect-golang-api

To run the api locally, follow these steps:-

Create a file `scripts/set-env.sh` following the example file provided and replace 'xxx' with confidential values:

```bash
export API_DATABASE_NAME=social-connect
export API_DATABASE_USER=xxx
export API_DATABASE_PASSWORD=xxx
export API_DATABASE_PORT=5432

export PGADMIN_EMAIL=xxx
export PGADMIN_PASSWORD=xxx
```

Set necessary environment variables by executing the script:

```bash
. scripts/set-env.sh
```

Open up terminal from root dir & run:

```bash
go run api/main.go
```
