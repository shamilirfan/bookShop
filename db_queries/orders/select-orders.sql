select
    u.id,
    u.user_name,
    u.email,
    o.phone_number,
    o.road_number,
    o.holding_number,
    o.area,
    o.thana,
    o.district
from
    users u
    join orders o on o.user_id = u.id
order by
    u.id