# Surge

This application provides REST API to get called on ride request and gets geographical coordinate of source and returns price coefficient of the source area.

**Tools:**<br>
* [Postgresql link](https://www.postgresql.org/) as relational database
* [Postgis link](https://postgis.net/) tool to add geographic objects to the PostgreSQL database
* [osm2pgsql link](https://osm2pgsql.org/) to insert osm.pbf data into Postgresql database
* golang net/http package to create http server

**OSM database:**<br>
Extract open street map as osm.pdf file and insert all data into osm database to use data in application.<br>
Assuming that the map data is being updated and always displays the latest information, a routine is defined to update the map data once a day. <br>
But this method will delete the database relations. <br>
To handle this problem there are ways to update data such as using osm diffs.

**Views:**<br>
* **tehran_v** : This view returns the information of municipality districts of Tehran city.<br>
* **ride_req_statistic_v**: Views ride request count in every municipality districts.

**Tables:**<br>
* **ride_requests_t**: To simulating ride requests data there is dummy data in this table.<br>
* **coefficient_t**: To configure Thresholds/Coefficients relation.

**APIs:**<br>
* **/get-price-coefficient** : Calls on every ride request and get latitude and longitude of source address and returns price coefficient of the area. <br>
* **/ride-request** : To simulate ride request call /ride-request api to insert ride request data into database.

**Use following command to build project :** <br>
>`docker-compose up -d --build`
