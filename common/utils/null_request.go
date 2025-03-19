package utils

import (
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/wrapperspb"
)

func NullStringWrapperToPGText(wrapper *wrapperspb.StringValue) pgtype.Text {
	if wrapper == nil {
		return pgtype.Text{Valid: false}
	}

	return pgtype.Text{String: wrapper.Value, Valid: true}
}

func NullInt32WrapperToPGInt4(wrapper *wrapperspb.Int32Value) pgtype.Int4 {
	if wrapper == nil {
		return pgtype.Int4{Valid: false}
	}

	return pgtype.Int4{Int32: wrapper.Value, Valid: true}
}
