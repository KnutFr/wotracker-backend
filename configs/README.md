# Required configuration :


```
SERVER_PORT=xxx #Port for server start
DATABASE_URL=xxx #Postgres Database connection string
SERVER_ADDRESS=localhost #Server where adresse will start
```
Config file is only read when environment variable env is local in all other case,
config is readed from environement variables.
```shell
export env=local
make all
# Code will read configs/app.env file

```