-- name: CreateFood :exec
insert into food (
    name, type_id, intake_status_id, feeder_id, location_id, remarks
) values (
    $1, $2, $3, $4, $5, $6
);

-- name: ListFood :many
select
    f.id,
    f.name,
    ft.name as type,
    fis.name as intake_status,
    ff.name as feeder,
    fl.name as location,
    f.remarks,
    f.created_at,
    f.updated_at
from food f
left join food_types ft on f.type_id = ft.id
left join food_intake_status fis on f.intake_status_id = fis.id
left join food_feeders ff on f.feeder_id = ff.id
left join food_locations fl on f.location_id = fl.id
where f.created_at between @start_timestamp and @end_timestamp
and (@type::text is null or ft.name = @type);
