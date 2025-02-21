-- name: CreateFood :exec
insert into food (
    name, type_id, intake_status_id, feeder_id, location_id, remarks
) values (
    $1, $2, $3, $4, $5, $6
);

-- name: ListFood :many
select
    id,
    name,
    type_id,
    intake_status_id,
    feeder_id,
    location_id,
    remarks,
    created_at,
    updated_at
from food
where created_at between sqlc.arg(start_timestamp) and sqlc.arg(end_timestamp);
