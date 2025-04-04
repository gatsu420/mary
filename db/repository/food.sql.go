// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: food.sql

package repository

import (
	"context"

	"github.com/jackc/pgx/v5/pgtype"
)

const checkFoodIsRemoved = `-- name: CheckFoodIsRemoved :one
select
    removed_at is not null::bool as is_removed
from food
where id = $1
`

func (q *Queries) CheckFoodIsRemoved(ctx context.Context, id int32) (bool, error) {
	row := q.db.QueryRow(ctx, checkFoodIsRemoved, id)
	var is_removed bool
	err := row.Scan(&is_removed)
	return is_removed, err
}

const createFood = `-- name: CreateFood :exec
insert into food (
    name, type_id, intake_status_id, feeder_id, location_id, remarks
) values (
    $1, $2, $3, $4, $5, $6
)
`

type CreateFoodParams struct {
	Name           string      `db:"name"`
	TypeID         int32       `db:"type_id"`
	IntakeStatusID int32       `db:"intake_status_id"`
	FeederID       int32       `db:"feeder_id"`
	LocationID     int32       `db:"location_id"`
	Remarks        pgtype.Text `db:"remarks"`
}

func (q *Queries) CreateFood(ctx context.Context, arg *CreateFoodParams) error {
	_, err := q.db.Exec(ctx, createFood,
		arg.Name,
		arg.TypeID,
		arg.IntakeStatusID,
		arg.FeederID,
		arg.LocationID,
		arg.Remarks,
	)
	return err
}

const deleteFood = `-- name: DeleteFood :execrows
update food
set
    removed_at = current_timestamp
where id = $1
and removed_at is null
returning id, name, type_id, intake_status_id, feeder_id, location_id, remarks, created_at, updated_at, removed_at
`

func (q *Queries) DeleteFood(ctx context.Context, id int32) (int64, error) {
	result, err := q.db.Exec(ctx, deleteFood, id)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}

const getFood = `-- name: GetFood :one
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
where f.id = $1
and f.removed_at is null
`

type GetFoodRow struct {
	ID           int32              `db:"id"`
	Name         string             `db:"name"`
	Type         pgtype.Text        `db:"type"`
	IntakeStatus pgtype.Text        `db:"intake_status"`
	Feeder       pgtype.Text        `db:"feeder"`
	Location     pgtype.Text        `db:"location"`
	Remarks      pgtype.Text        `db:"remarks"`
	CreatedAt    pgtype.Timestamptz `db:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at"`
}

func (q *Queries) GetFood(ctx context.Context, id int32) (GetFoodRow, error) {
	row := q.db.QueryRow(ctx, getFood, id)
	var i GetFoodRow
	err := row.Scan(
		&i.ID,
		&i.Name,
		&i.Type,
		&i.IntakeStatus,
		&i.Feeder,
		&i.Location,
		&i.Remarks,
		&i.CreatedAt,
		&i.UpdatedAt,
	)
	return i, err
}

const listFood = `-- name: ListFood :many
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
where f.created_at between $1 and $2
and (
    $3::text is null
    or ft.name = $3
)
and (
    $4::text is null
    or fis.name = $4
)
and (
    $5::text is null
    or ff.name = $5
)
and (
    $6::text is null or
    fl.name = $6
)
`

type ListFoodParams struct {
	StartTimestamp pgtype.Timestamptz `db:"start_timestamp"`
	EndTimestamp   pgtype.Timestamptz `db:"end_timestamp"`
	Type           pgtype.Text        `db:"type"`
	IntakeStatus   pgtype.Text        `db:"intake_status"`
	Feeder         pgtype.Text        `db:"feeder"`
	Location       pgtype.Text        `db:"location"`
}

type ListFoodRow struct {
	ID           int32              `db:"id"`
	Name         string             `db:"name"`
	Type         pgtype.Text        `db:"type"`
	IntakeStatus pgtype.Text        `db:"intake_status"`
	Feeder       pgtype.Text        `db:"feeder"`
	Location     pgtype.Text        `db:"location"`
	Remarks      pgtype.Text        `db:"remarks"`
	CreatedAt    pgtype.Timestamptz `db:"created_at"`
	UpdatedAt    pgtype.Timestamptz `db:"updated_at"`
}

func (q *Queries) ListFood(ctx context.Context, arg *ListFoodParams) ([]ListFoodRow, error) {
	rows, err := q.db.Query(ctx, listFood,
		arg.StartTimestamp,
		arg.EndTimestamp,
		arg.Type,
		arg.IntakeStatus,
		arg.Feeder,
		arg.Location,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ListFoodRow
	for rows.Next() {
		var i ListFoodRow
		if err := rows.Scan(
			&i.ID,
			&i.Name,
			&i.Type,
			&i.IntakeStatus,
			&i.Feeder,
			&i.Location,
			&i.Remarks,
			&i.CreatedAt,
			&i.UpdatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateFood = `-- name: UpdateFood :execrows
update food
set
    name = coalesce($1::text, name),
    type_id = coalesce($2::integer, type_id),
    intake_status_id = coalesce($3::integer, intake_status_id),
    feeder_id = coalesce($4::integer, feeder_id),
    location_id = coalesce($5::integer, location_id),
    remarks = coalesce($6::text, remarks),
    updated_at = current_timestamp
where id = $7
and removed_at is null
returning id, name, type_id, intake_status_id, feeder_id, location_id, remarks, created_at, updated_at, removed_at
`

type UpdateFoodParams struct {
	Name           pgtype.Text `db:"name"`
	TypeID         pgtype.Int4 `db:"type_id"`
	IntakeStatusID pgtype.Int4 `db:"intake_status_id"`
	FeederID       pgtype.Int4 `db:"feeder_id"`
	LocationID     pgtype.Int4 `db:"location_id"`
	Remarks        pgtype.Text `db:"remarks"`
	ID             int32       `db:"id"`
}

func (q *Queries) UpdateFood(ctx context.Context, arg *UpdateFoodParams) (int64, error) {
	result, err := q.db.Exec(ctx, updateFood,
		arg.Name,
		arg.TypeID,
		arg.IntakeStatusID,
		arg.FeederID,
		arg.LocationID,
		arg.Remarks,
		arg.ID,
	)
	if err != nil {
		return 0, err
	}
	return result.RowsAffected(), nil
}
