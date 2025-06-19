package converters

import (
	"encoding/json"
	"time"

	"github.com/jackc/pgx/v5/pgtype"
)

// TextOrNil converts pgtype.Text to *string
func TextOrNil(t pgtype.Text) *string {
	if t.Valid {
		return &t.String
	}
	return nil
}

// StringToText converts *string to pgtype.Text
func StringToPgText(s *string) pgtype.Text {
	if s == nil || *s == "" {
		return pgtype.Text{Valid: false}
	}
	return pgtype.Text{String: *s, Valid: true}
}

// Int4OrNil converts pgtype.Int4 to *int32
func Int4OrNil(i pgtype.Int4) *int32 {
	if i.Valid {
		return &i.Int32
	}
	return nil
}

// Int32ToInt4 converts *int32 to pgtype.Int4
func IntToPgInt4(i *int) pgtype.Int4 {
	if i == nil {
		return pgtype.Int4{Valid: false}
	}
	return pgtype.Int4{Int32: int32(*i), Valid: true}
}

// Int8OrNil converts pgtype.Int8 to *int64
func Int8OrNil(i pgtype.Int8) *int64 {
	if i.Valid {
		return &i.Int64
	}
	return nil
}

// Int64ToInt8 converts *int64 to pgtype.Int8
func IntToPgInt8(i *int) pgtype.Int8 {
	if i == nil {
		return pgtype.Int8{Valid: false}
	}
	return pgtype.Int8{Int64: int64(*i), Valid: true}
}

// Float8OrNil converts pgtype.Float8 to *float64
func Float8OrNil(f pgtype.Float8) *float64 {
	if f.Valid {
		return &f.Float64
	}
	return nil
}

// Float64ToFloat8 converts *float64 to pgtype.Float8
func Float64ToPgFloat8(f *float64) pgtype.Float8 {
	if f == nil {
		return pgtype.Float8{Valid: false}
	}
	return pgtype.Float8{Float64: *f, Valid: true}
}

// BoolOrNil converts pgtype.Bool to *bool
func BoolOrNil(b pgtype.Bool) *bool {
	if b.Valid {
		return &b.Bool
	}
	return nil
}

// BoolToPgBool converts *bool to pgtype.Bool
func BoolToPgBool(b *bool) pgtype.Bool {
	if b == nil {
		return pgtype.Bool{Valid: false}
	}
	return pgtype.Bool{Bool: *b, Valid: true}
}

// TimestampOrNil converts pgtype.Timestamp to *time.Time
func TimestampOrNil(t pgtype.Timestamp) *time.Time {
	if t.Valid {
		return &t.Time
	}
	return nil
}

// TimeToTimestamp converts *time.Time to pgtype.Timestamp
func TimeToPgTimestamp(t *time.Time) pgtype.Timestamp {
	if t == nil {
		return pgtype.Timestamp{Valid: false}
	}
	return pgtype.Timestamp{Time: *t, Valid: true}
}

// DateOrNil converts pgtype.Date to *time.Time
func DateOrNil(d pgtype.Date) *time.Time {
	if d.Valid {
		return &d.Time
	}
	return nil
}

// TimeToDate converts *time.Time to pgtype.Date
func TimeToPgDate(t *time.Time) pgtype.Date {
	if t == nil {
		return pgtype.Date{Valid: false}
	}
	return pgtype.Date{Time: *t, Valid: true}
}

func ToJSONB(input map[string]interface{}) ([]byte, error) {
	if input == nil {
		return []byte("{}"), nil
	}
	return json.Marshal(input)
}

func FromJSONB(data []byte) (map[string]interface{}, error) {
	var result map[string]interface{}
	if len(data) == 0 {
		return map[string]interface{}{}, nil
	}
	err := json.Unmarshal(data, &result)
	return result, err
}
