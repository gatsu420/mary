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
