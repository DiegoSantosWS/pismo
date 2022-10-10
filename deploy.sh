#!/bin/bash

DATA_ISO=$(date +%Y-%m-%d-%H-%M-%S)

echo -e "CREATE IMAGE WITH DOCKER"

# docker build . -t pismo:latest

echo -e "FINISH RUN IMAGE WITH DOCKER"

echo "--------------------------------------------------------------------------------------------"

echo -e "UP DATABASE"

docker-compose stop postgres
docker-compose rm --force postgres
docker-compose up 

echo -e "FINISH"


# DB_HOST=${PREST_PG_HOST:-localhost}
# DB_USER=${PREST_PG_USER:-postgres}
# DB_PORT=${PREST_PG_PORT:-5432}
# DB_NAME=${PREST_PG_DATABASE:-prest} 

# psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c "DROP DATABASE IF EXISTS $DB_NAME"
# psql -h $DB_HOST -p $DB_PORT -U $DB_USER -c "create database $DB_NAME;"

# psql -h $DB_HOST -p $DB_PORT -U $DB_USER -v DB_NAME=$DB_NAME -f testdata/schema.sql