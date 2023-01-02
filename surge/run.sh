#!/bin/bash

osm2pgsql -c -d $OSM_DB -U $OSM_USER -H $OSM_HOST -S $STYLE_FILE $OSM_PBF_FILE
go run .