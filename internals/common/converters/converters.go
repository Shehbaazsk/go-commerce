package converters

import (
	"database/sql"
	"time"
)

// StringOrNil converts sql.NullString to *string
func StringOrNil(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// Int64OrNil converts sql.NullInt64 to *int64
func Int64OrNil(ni sql.NullInt64) *int64 {
	if ni.Valid {
		return &ni.Int64
	}
	return nil
}

// BoolOrNil converts sql.NullBool to *bool
func BoolOrNil(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}

// TimeOrNil converts sql.NullTime to *time.Time
func TimeOrNil(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}

// StringToNullString safely converts *string to sql.NullString
func StringToNullString(s *string) sql.NullString {
	if s != nil {
		return sql.NullString{
			String: *s,
			Valid:  true,
		}
	}
	return sql.NullString{}
}

// BoolToNullBool safely converts *bool to sql.NullBool
func BoolToNullBool(b *bool) sql.NullBool {
	if b != nil {
		return sql.NullBool{
			Bool:  *b,
			Valid: true,
		}
	}
	return sql.NullBool{}
}

// TimeToNullTime safely converts *time.Time to sql.NullTime
func TimeToNullTime(t *time.Time) sql.NullTime {
	if t != nil {
		return sql.NullTime{
			Time:  *t,
			Valid: true,
		}
	}
	return sql.NullTime{}
}

// NullStringToPointer converts sql.NullString to *string
func NullStringToPointer(ns sql.NullString) *string {
	if ns.Valid {
		return &ns.String
	}
	return nil
}

// NullBoolToPointer converts sql.NullBool to *bool
func NullBoolToPointer(nb sql.NullBool) *bool {
	if nb.Valid {
		return &nb.Bool
	}
	return nil
}

// NullTimeToPointer converts sql.NullTime to *time.Time
func NullTimeToPointer(nt sql.NullTime) *time.Time {
	if nt.Valid {
		return &nt.Time
	}
	return nil
}
