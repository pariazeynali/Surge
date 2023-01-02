#!/bin/bash

apt update
apt install -y osm2pgsql
wget https://download.geofabrik.de/asia/iran-latest.osm.pbf
osm2pgsql -c -d $OSM_DB -U $OSM_USER -H $OSM_HOST -S $STYLE_FILE $OSM_PBF_FILE