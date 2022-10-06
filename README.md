# PISMO

Pismo is a api to create transactions in account

# CONFIG

To use the api you shoud check the file `docker-compose.yml`, on file contains the informations of conections with database

#### ENV

Create the file `.env` in principal dir.

```env
PG_HOST=<HOST_DB>
PG_USER=<USER_DB>
PG_PASS=<PASS_DB>
PG_PORT=5432
PG_DB=<DB_NAME>
```

# RUN PISTMO

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

### To execute database in forenground:

```bash
$: docker-compose up postgres pismo
```

### To execute database in background:

```bash
$: docker-compose up -d postgres pismo
```

After of run open the browser or postman and run the endpoints

#### CURL REQUEST

```bash
curl http://localhost:8080/account/1
```

#### Result
```json
{
    "id": 1,
    "doc_number": "05769904496",
    "created_at": "2022-10-05T19:08:05.034872Z"
}