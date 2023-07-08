// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"netinfo/ent/record"
	"netinfo/ent/schema"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// Record is the model entity for the Record schema.
type Record struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// CreatedAt holds the value of the "createdAt" field.
	CreatedAt time.Time `json:"createdAt,omitempty"`
	// UpdatedAt holds the value of the "updatedAt" field.
	UpdatedAt time.Time `json:"updatedAt,omitempty"`
	// RequestIP holds the value of the "requestIP" field.
	RequestIP string `json:"requestIP,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// NetInterfaces holds the value of the "netInterfaces" field.
	NetInterfaces []schema.NetInterface `json:"netInterfaces,omitempty"`
	selectValues  sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Record) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case record.FieldNetInterfaces:
			values[i] = new([]byte)
		case record.FieldID, record.FieldRequestIP, record.FieldDescription:
			values[i] = new(sql.NullString)
		case record.FieldCreatedAt, record.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Record fields.
func (r *Record) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case record.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				r.ID = value.String
			}
		case record.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field createdAt", values[i])
			} else if value.Valid {
				r.CreatedAt = value.Time
			}
		case record.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updatedAt", values[i])
			} else if value.Valid {
				r.UpdatedAt = value.Time
			}
		case record.FieldRequestIP:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field requestIP", values[i])
			} else if value.Valid {
				r.RequestIP = value.String
			}
		case record.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				r.Description = value.String
			}
		case record.FieldNetInterfaces:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field netInterfaces", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &r.NetInterfaces); err != nil {
					return fmt.Errorf("unmarshal field netInterfaces: %w", err)
				}
			}
		default:
			r.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Record.
// This includes values selected through modifiers, order, etc.
func (r *Record) Value(name string) (ent.Value, error) {
	return r.selectValues.Get(name)
}

// Update returns a builder for updating this Record.
// Note that you need to call Record.Unwrap() before calling this method if this Record
// was returned from a transaction, and the transaction was committed or rolled back.
func (r *Record) Update() *RecordUpdateOne {
	return NewRecordClient(r.config).UpdateOne(r)
}

// Unwrap unwraps the Record entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (r *Record) Unwrap() *Record {
	_tx, ok := r.config.driver.(*txDriver)
	if !ok {
		panic("ent: Record is not a transactional entity")
	}
	r.config.driver = _tx.drv
	return r
}

// String implements the fmt.Stringer.
func (r *Record) String() string {
	var builder strings.Builder
	builder.WriteString("Record(")
	builder.WriteString(fmt.Sprintf("id=%v, ", r.ID))
	builder.WriteString("createdAt=")
	builder.WriteString(r.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("updatedAt=")
	builder.WriteString(r.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("requestIP=")
	builder.WriteString(r.RequestIP)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(r.Description)
	builder.WriteString(", ")
	builder.WriteString("netInterfaces=")
	builder.WriteString(fmt.Sprintf("%v", r.NetInterfaces))
	builder.WriteByte(')')
	return builder.String()
}

// Records is a parsable slice of Record.
type Records []*Record
