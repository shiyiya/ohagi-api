// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/haton14/ohagi-api/ent/recordfood"
)

// RecordFood is the model entity for the RecordFood schema.
type RecordFood struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// RecordID holds the value of the "record_id" field.
	RecordID int `json:"record_id,omitempty"`
	// FoodID holds the value of the "food_id" field.
	FoodID int `json:"food_id,omitempty"`
	// Amount holds the value of the "amount" field.
	Amount int `json:"amount,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*RecordFood) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case recordfood.FieldID, recordfood.FieldRecordID, recordfood.FieldFoodID, recordfood.FieldAmount:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type RecordFood", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the RecordFood fields.
func (rf *RecordFood) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case recordfood.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			rf.ID = int(value.Int64)
		case recordfood.FieldRecordID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field record_id", values[i])
			} else if value.Valid {
				rf.RecordID = int(value.Int64)
			}
		case recordfood.FieldFoodID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field food_id", values[i])
			} else if value.Valid {
				rf.FoodID = int(value.Int64)
			}
		case recordfood.FieldAmount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field amount", values[i])
			} else if value.Valid {
				rf.Amount = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this RecordFood.
// Note that you need to call RecordFood.Unwrap() before calling this method if this RecordFood
// was returned from a transaction, and the transaction was committed or rolled back.
func (rf *RecordFood) Update() *RecordFoodUpdateOne {
	return (&RecordFoodClient{config: rf.config}).UpdateOne(rf)
}

// Unwrap unwraps the RecordFood entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (rf *RecordFood) Unwrap() *RecordFood {
	tx, ok := rf.config.driver.(*txDriver)
	if !ok {
		panic("ent: RecordFood is not a transactional entity")
	}
	rf.config.driver = tx.drv
	return rf
}

// String implements the fmt.Stringer.
func (rf *RecordFood) String() string {
	var builder strings.Builder
	builder.WriteString("RecordFood(")
	builder.WriteString(fmt.Sprintf("id=%v", rf.ID))
	builder.WriteString(", record_id=")
	builder.WriteString(fmt.Sprintf("%v", rf.RecordID))
	builder.WriteString(", food_id=")
	builder.WriteString(fmt.Sprintf("%v", rf.FoodID))
	builder.WriteString(", amount=")
	builder.WriteString(fmt.Sprintf("%v", rf.Amount))
	builder.WriteByte(')')
	return builder.String()
}

// RecordFoods is a parsable slice of RecordFood.
type RecordFoods []*RecordFood

func (rf RecordFoods) config(cfg config) {
	for _i := range rf {
		rf[_i].config = cfg
	}
}
