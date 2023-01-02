FROM golang:1.19

COPY ./* /app/
WORKDIR /app

RUN apt update\
 && apt install -y osm2pgsql

CMD bash