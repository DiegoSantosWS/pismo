# PISMO
Pismo is an API for creating account transactions

## CLONE

```shell
$ git clone https://github.com/DiegoSantosWS/pismo.git && cd pismo
```
# CONFIG

To use the API you should check the file `docker-compose.yml`, on file contains the information connections with the database


#### ENV

Create the file `.env` in the principal dir.

```env
PG_HOST=<HOST_DB>
PG_USER=<USER_DB>
PG_PASS=<PASS_DB>
PG_PORT=5432
PG_DB=<DB_NAME>
```

# RUN PISMO

### GENERATE IMAGE USING COMMAND DOCKER.

To run the Pismo API you should execute the command below.


```shell
$: docker build . -t pismo:latest
```
> The command generates the image docker for the programmer.

Now you need to create the database, you can look config in the file `docker-compose.yml`, to execute the file you need to run the command below.

### To execute all:

```shell
$: docker-compose up
```

You can run each service separately, for example:

```shell
$: docker-compose up pismo
```

### To execute the database in the foreground


```shell
$ docker-compose up -d postgres
```

> Don't haveing difference on format that you execute, are only formats differents

After of run open the browser or postman and run the endpoints

### Run using the sh file

You too can execute the Pismo API using the command sh.

```shell
$: sh deploy.sh
```

> This command will to do download of image `diegosantosws/pismo`, and start all services.

#### CURL REQUEST

```shell
#create account

curl -X POST http://localhost:8080/account
   -H 'Content-Type: application/json'
   -d '{"doc_number":01425836930}'

#create transactions

curl -X POST http://localhost:8080/transactions
   -H 'Content-Type: application/json'
   -d '{"account_id":1,"operations_types_id":4,"amount":7800.00}'

curl -X GET http://localhost:8080/account/1
```
#### Result
```json
{
    "id": 1,
    "doc_number": "05769904496",
    "created_at": "2022-10-05T19:08:05.034872Z"
}
```