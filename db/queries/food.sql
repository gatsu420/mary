-- name: CreateFood :exec
insert into food (
    name, type, intake_status, feeder, location, remarks
) values (
    $1, $2, $3, $4, $5, $6
);
