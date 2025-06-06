// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.28.0
// source: copyfrom.go

package repository

import (
	"context"
)

// iteratorForCreateEvent implements pgx.CopyFromSource.
type iteratorForCreateEvent struct {
	rows                 []CreateEventParams
	skippedFirstNextCall bool
}

func (r *iteratorForCreateEvent) Next() bool {
	if len(r.rows) == 0 {
		return false
	}
	if !r.skippedFirstNextCall {
		r.skippedFirstNextCall = true
		return true
	}
	r.rows = r.rows[1:]
	return len(r.rows) > 0
}

func (r iteratorForCreateEvent) Values() ([]interface{}, error) {
	return []interface{}{
		r.rows[0].Name,
		r.rows[0].UserID,
		r.rows[0].CreatedAt,
	}, nil
}

func (r iteratorForCreateEvent) Err() error {
	return nil
}

func (q *Queries) CreateEvent(ctx context.Context, arg []CreateEventParams) (int64, error) {
	return q.db.CopyFrom(ctx, []string{"events"}, []string{"name", "user_id", "created_at"}, &iteratorForCreateEvent{rows: arg})
}
