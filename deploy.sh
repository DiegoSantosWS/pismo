#!/bin/bash

DATA_ISO=$(date +%Y-%m-%d-%H-%M-%S)

echo -e "-------------------------------------- Clean <none> images ---------------------------------------"

docker rmi $(docker images | grep "<none>" | awk '{print $3}') --force

echo -e "\033[0;33m######################################### pull ########################################\033[0m"
docker pull diegosantosws/pismo

echo -e "UP DATABASE"

docker-compose stop postgres pismo
docker-compose rm --force postgres pismo
docker-compose up 

echo -e "FINISH"