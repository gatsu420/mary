// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: events.sql

package repository

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type CreateEventParams struct {
	Name      string
	UserID    string
	CreatedAt pgtype.Timestamptz
}
