#!/bin/bash


docker run --rm -ti --name emq -p 18083:18083 -p 1883:1883 sneck/emqttd:latest
docker run --name ravigation-redis -p 6379:6379 -d redis redis-server --appendonly yes