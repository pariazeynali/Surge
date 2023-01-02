create or replace view tehran as(
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

create or replace view ride_req_statistics as
with ride_req_last_10_mins as (
    select * from ride_request
    where req_timestamp between now() - interval '10 minutes' and now()
) select tc.osm_id , count(rr.req_timestamp)
  from tehran tc left join ride_req_last_10_mins rr
                           on ST_Contains(tc.way, rr.geom)
  group by tc.osm_id ;