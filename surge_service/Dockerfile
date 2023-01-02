FROM golang:1.19

WORKDIR /usr/src/app

RUN apt update\
 && apt install -y osm2pgsql

COPY . .

