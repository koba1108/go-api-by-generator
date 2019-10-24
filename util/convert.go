package util

import "database/sql"

func NullStringToStringP(ns sql.NullString) *string {
	var res *string
	if ns.Valid {
		res = &ns.String
	}
	return res
}

func NullInt64ToInt32P(ni64 sql.NullInt64) *int32 {
	var res *int32
	if ni64.Valid {
		i32 := int32(ni64.Int64)
		res = &i32
	}
	return res
}
