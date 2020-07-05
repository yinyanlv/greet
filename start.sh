#!/usr/bin/env bash

docker build -t prot .
docker run --name prot -p 9000:9000 --net prot_default --link mysql:mysql -d prot
