package null

import (
	"bytes"
	"database/sql"
	"database/sql/driver"
	"encoding/binary"
	"fmt"
	"reflect"
	"strings"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/bsontype"
)

// nullable time.Time type.
// null will set valid value to false.
// any time.Time value will be considered null when valid is set to false.
// SqlNull is for updating SQL DB value to null
type Time struct {
	Time    time.Time
	Valid   bool
	SqlNull bool
}

// create new nullable time.time
func NewTime(t time.Time, valid bool) Time {
	return Time{
		Time:  t,
		Valid: valid,
	}
}

// create valid nullable time.Time
func TimeFrom(t time.Time) Time {
	return NewTime(t, true)
}

func (t *Time) Scan(value interface{}) error {
	var sqlt sql.NullTime
	if err := sqlt.Scan(value); err != nil {
		return err
	}

	if reflect.TypeOf(value) == nil {
		*t = Time{sqlt.Time, false, false}
	} else {
		*t = Time{sqlt.Time, true, false}
	}

	return nil
}

func (t Time) Value() (driver.Value, error) {
	if !t.Valid {
		return nil, nil
	}
	return t.Time, nil
}

func (t *Time) MarshalJSON() ([]byte, error) {
	if !t.Valid {
		return nullBytes, nil
	}
	val := fmt.Sprintf("\"%s\"", t.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func (nt Time) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if !nt.Valid {
		return bson.TypeNull, nil, nil
	}
	return bson.MarshalValue(nt.Time)
}

func (nt *Time) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		return nil
	}
	s := string(b)
	s = strings.Replace(s, "\"", "", -1)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

func (nt *Time) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	if t == bson.TypeNull {
		return nil
	}

	ms := int64(binary.LittleEndian.Uint64(b))
	x := time.Unix(ms/1000, (ms%1000)*1000000).UTC()

	nt.Time = x
	nt.Valid = true
	return nil
}

func (nt *Time) UnmarshalText(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		return nil
	}
	s := string(b)
	s = strings.Replace(s, "\"", "", -1)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		nt.Valid = false
		return err
	}

	nt.Time = x
	nt.Valid = true
	return nil
}

// will return true if invalid or value is empty
func (t Time) IsNullOrZero() bool {
	return !t.Valid || t.Time.IsZero()
}

// returns true if both invalid or both have same value
func (t Time) Equal(other Time) bool {
	return t.Valid == other.Valid && (!t.Valid || t.Time.Equal(other.Time))
}

// return true if valid and both have same value
func (t Time) Is(other time.Time) bool {
	return t.Equal(TimeFrom(other))
}

// return true if valid and the argument is before the origin
func (t Time) IsBefore(other time.Time) bool {
	return t.Valid && other.Before(t.Time)
}

// return true if valid and the argument is after the origin
func (t Time) IsAfter(other time.Time) bool {
	return t.Valid && other.After(t.Time)
}

func (t Time) IsZero() bool {
	return !t.Valid
}

// Date is nullable time.Time for parsing DATE type in SQL to golang time.Time.
// SqlNull is for updating SQL DB value to null
type Date struct {
	Time    time.Time
	Valid   bool
	SqlNull bool
}

// create new nullable time.time
func NewDate(t time.Time, valid bool) Date {
	return Date{
		Time:  convertTimeToZero(t),
		Valid: valid,
	}
}

// create valid nullable time.Time
func DateFrom(t time.Time) Date {
	return NewDate(t, true)
}

func (d *Date) Scan(value interface{}) error {
	var (
		sqls sql.NullString
		t    time.Time
		err  error
	)
	if err := sqls.Scan(value); err != nil {
		return err
	}

	if sqls.String != "" && sqls.Valid {
		t, err = time.Parse(time.RFC3339, sqls.String)
		if err != nil {
			return err
		}
	}

	if reflect.TypeOf(value) == nil {
		*d = Date{t, false, false}
	} else {
		*d = Date{t, true, false}
	}

	return nil
}

func (d Date) Value() (driver.Value, error) {
	if !d.Valid {
		return nil, nil
	}
	return convertTimeToZero(d.Time), nil
}

func (d *Date) MarshalJSON() ([]byte, error) {
	if !d.Valid {
		return nullBytes, nil
	}
	d.Time = time.Date(d.Time.Year(), d.Time.Month(), d.Time.Day(), 0, 0, 0, 0, d.Time.Location())
	val := fmt.Sprintf("\"%s\"", d.Time.Format(time.RFC3339))
	return []byte(val), nil
}

func (d Date) MarshalBSONValue() (bsontype.Type, []byte, error) {
	if !d.Valid {
		return bson.TypeNull, nil, nil
	}
	d.Time = time.Date(d.Time.Year(), d.Time.Month(), d.Time.Day(), 0, 0, 0, 0, d.Time.Location())
	return bson.MarshalValue(d.Time)
}

func (d *Date) UnmarshalJSON(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		return nil
	}
	s := string(b)
	s = strings.Replace(s, "\"", "", -1)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		d.Valid = false
		return err
	}

	x = time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, x.Location())

	d.Time = x
	d.Valid = true
	return nil
}

func (d *Date) UnmarshalBSONValue(t bsontype.Type, b []byte) error {
	if t == bson.TypeNull {
		return nil
	}

	ms := int64(binary.LittleEndian.Uint64(b))
	x := time.Unix(ms/1000, (ms%1000)*1000000).UTC()

	x = time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, x.Location())

	d.Time = x
	d.Valid = true
	return nil
}

func (d *Date) UnmarshalText(b []byte) error {
	if bytes.Equal(b, nullBytes) {
		return nil
	}
	s := string(b)
	s = strings.Replace(s, "\"", "", -1)

	x, err := time.Parse(time.RFC3339, s)
	if err != nil {
		d.Valid = false
		return err
	}

	x = time.Date(x.Year(), x.Month(), x.Day(), 0, 0, 0, 0, x.Location())

	d.Time = x
	d.Valid = true
	return nil
}

// will return true if invalid or value is empty
func (d Date) IsNullOrZero() bool {
	return !d.Valid || d.Time.IsZero()
}

// returns true if both invalid or both have same value
func (d Date) Equal(other Date) bool {
	return d.Valid == other.Valid && (!d.Valid || d.Time.Equal(other.Time))
}

// return true if valid and both have same value
func (d Date) Is(other time.Time) bool {
	return d.Equal(DateFrom(other))
}

// return true if valid and the argument is before the origin
func (d Date) IsBefore(other time.Time) bool {
	return d.Valid && other.Before(d.Time)
}

// return true if valid and the argument is after the origin
func (d Date) IsAfter(other time.Time) bool {
	return d.Valid && other.After(d.Time)
}

func convertTimeToZero(t time.Time) time.Time {
	return time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location())
}

func (d Date) IsZero() bool {
	return !d.Valid
}
