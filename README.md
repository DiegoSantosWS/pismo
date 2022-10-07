# PISMO

Pismo is a api to creating account transactions

## CLONE

```bash
$ git clone https://github.com/DiegoSantosWS/pismo.git && cd pismo
```
# CONFIG

To use the api you shoud check the file `docker-compose.yml`, on file contains the informations of connections with database


#### ENV

Create the file `.env` in principal dir.

```env
PG_HOST=<HOST_DB>
PG_USER=<USER_DB>
PG_PASS=<PASS_DB>
PG_PORT=5432
PG_DB=<DB_NAME>
```

# RUN PISMO

### GENERATE IMAGE USING COMMAND DOCKER.

To run the pismo api you shoud executed the command below.

```bash
$: docker build . -t pismo:latest
```

The command generate the image docker to programmer.

Now you need create the database, you can look config in file `docker-compose.yml`, to execute the file you need runed the command below.

### To execute all:

```bash
$: docker-compose up
```

### To execute database in forenground

```bash
$ docker-compose up postgres
```

### To execute database in background

```bash
$ docker-compose up -d postgres
```

> Don't haveing difference on format that you execute, are only formats differents

After of run open the browser or postman and run the endpoints

#### CURL REQUEST

```bash
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