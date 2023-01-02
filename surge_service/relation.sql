create table if not exists ride_request_t (
                              req_timestamp   timestamptz primary key default clock_timestamp()
    , geom            geometry(Point,3857) not null
) ;

create table if not exists coefficient_t (
                               req_threshold       int4range
    , price_coefficient   float
);

insert into coefficient_t
(req_threshold, price_coefficient)
values
('[0, 1000)', 1.00)
     ,('[1000, 3000)', 1.05)
     ,('[3000, 5000)', 1.10)
     ,('[5000, 10000)', 1.20)
     ,('[10000, 25000)', 2.00)
     ,('[25000, 50000)', 3.00);

create or replace view tehran_v as(
    with tehran_city as (
        select * from planet_osm_polygon
        where osm_id = -6663864
    ) , city_polygons as (
        select * from planet_osm_polygon
        where admin_level = '9'
    )
    select cp.* from city_polygons cp
    inner join tehran_city tc on ST_Contains(tc.way, cp.way)
);

create or replace view ride_req_statistics_v as
with ride_req_last_10_mins as (
    select * from ride_request
    where req_timestamp between now() - interval '10 minutes' and now()
) select tc.osm_id , count(rr.req_timestamp)
  from tehran tc left join ride_req_last_10_mins rr
                           on ST_Contains(tc.way, rr.geom)
  group by tc.osm_id ;