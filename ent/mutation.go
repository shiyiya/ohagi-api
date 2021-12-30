// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/haton14/ohagi-api/ent/food"
	"github.com/haton14/ohagi-api/ent/predicate"
	"github.com/haton14/ohagi-api/ent/record"
	"github.com/haton14/ohagi-api/ent/recordfood"

	"entgo.io/ent"
)

const (
	// Operation types.
	OpCreate    = ent.OpCreate
	OpDelete    = ent.OpDelete
	OpDeleteOne = ent.OpDeleteOne
	OpUpdate    = ent.OpUpdate
	OpUpdateOne = ent.OpUpdateOne

	// Node types.
	TypeFood       = "Food"
	TypeRecord     = "Record"
	TypeRecordFood = "RecordFood"
)

// FoodMutation represents an operation that mutates the Food nodes in the graph.
type FoodMutation struct {
	config
	op            Op
	typ           string
	id            *int
	name          *string
	unit          *string
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*Food, error)
	predicates    []predicate.Food
}

var _ ent.Mutation = (*FoodMutation)(nil)

// foodOption allows management of the mutation configuration using functional options.
type foodOption func(*FoodMutation)

// newFoodMutation creates new mutation for the Food entity.
func newFoodMutation(c config, op Op, opts ...foodOption) *FoodMutation {
	m := &FoodMutation{
		config:        c,
		op:            op,
		typ:           TypeFood,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withFoodID sets the ID field of the mutation.
func withFoodID(id int) foodOption {
	return func(m *FoodMutation) {
		var (
			err   error
			once  sync.Once
			value *Food
		)
		m.oldValue = func(ctx context.Context) (*Food, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Food.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withFood sets the old Food of the mutation.
func withFood(node *Food) foodOption {
	return func(m *FoodMutation) {
		m.oldValue = func(context.Context) (*Food, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m FoodMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m FoodMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *FoodMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetName sets the "name" field.
func (m *FoodMutation) SetName(s string) {
	m.name = &s
}

// Name returns the value of the "name" field in the mutation.
func (m *FoodMutation) Name() (r string, exists bool) {
	v := m.name
	if v == nil {
		return
	}
	return *v, true
}

// OldName returns the old "name" field's value of the Food entity.
// If the Food object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FoodMutation) OldName(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldName is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldName requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldName: %w", err)
	}
	return oldValue.Name, nil
}

// ResetName resets all changes to the "name" field.
func (m *FoodMutation) ResetName() {
	m.name = nil
}

// SetUnit sets the "unit" field.
func (m *FoodMutation) SetUnit(s string) {
	m.unit = &s
}

// Unit returns the value of the "unit" field in the mutation.
func (m *FoodMutation) Unit() (r string, exists bool) {
	v := m.unit
	if v == nil {
		return
	}
	return *v, true
}

// OldUnit returns the old "unit" field's value of the Food entity.
// If the Food object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *FoodMutation) OldUnit(ctx context.Context) (v string, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldUnit is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldUnit requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldUnit: %w", err)
	}
	return oldValue.Unit, nil
}

// ResetUnit resets all changes to the "unit" field.
func (m *FoodMutation) ResetUnit() {
	m.unit = nil
}

// Where appends a list predicates to the FoodMutation builder.
func (m *FoodMutation) Where(ps ...predicate.Food) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *FoodMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Food).
func (m *FoodMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *FoodMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.name != nil {
		fields = append(fields, food.FieldName)
	}
	if m.unit != nil {
		fields = append(fields, food.FieldUnit)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *FoodMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case food.FieldName:
		return m.Name()
	case food.FieldUnit:
		return m.Unit()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *FoodMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case food.FieldName:
		return m.OldName(ctx)
	case food.FieldUnit:
		return m.OldUnit(ctx)
	}
	return nil, fmt.Errorf("unknown Food field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FoodMutation) SetField(name string, value ent.Value) error {
	switch name {
	case food.FieldName:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetName(v)
		return nil
	case food.FieldUnit:
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetUnit(v)
		return nil
	}
	return fmt.Errorf("unknown Food field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *FoodMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *FoodMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *FoodMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Food numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *FoodMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *FoodMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *FoodMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Food nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *FoodMutation) ResetField(name string) error {
	switch name {
	case food.FieldName:
		m.ResetName()
		return nil
	case food.FieldUnit:
		m.ResetUnit()
		return nil
	}
	return fmt.Errorf("unknown Food field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *FoodMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *FoodMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *FoodMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *FoodMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *FoodMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *FoodMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *FoodMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Food unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *FoodMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Food edge %s", name)
}

// RecordMutation represents an operation that mutates the Record nodes in the graph.
type RecordMutation struct {
	config
	op              Op
	typ             string
	id              *int
	created_at      *time.Time
	last_updated_at *time.Time
	clearedFields   map[string]struct{}
	done            bool
	oldValue        func(context.Context) (*Record, error)
	predicates      []predicate.Record
}

var _ ent.Mutation = (*RecordMutation)(nil)

// recordOption allows management of the mutation configuration using functional options.
type recordOption func(*RecordMutation)

// newRecordMutation creates new mutation for the Record entity.
func newRecordMutation(c config, op Op, opts ...recordOption) *RecordMutation {
	m := &RecordMutation{
		config:        c,
		op:            op,
		typ:           TypeRecord,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRecordID sets the ID field of the mutation.
func withRecordID(id int) recordOption {
	return func(m *RecordMutation) {
		var (
			err   error
			once  sync.Once
			value *Record
		)
		m.oldValue = func(ctx context.Context) (*Record, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().Record.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRecord sets the old Record of the mutation.
func withRecord(node *Record) recordOption {
	return func(m *RecordMutation) {
		m.oldValue = func(context.Context) (*Record, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RecordMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RecordMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RecordMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetCreatedAt sets the "created_at" field.
func (m *RecordMutation) SetCreatedAt(t time.Time) {
	m.created_at = &t
}

// CreatedAt returns the value of the "created_at" field in the mutation.
func (m *RecordMutation) CreatedAt() (r time.Time, exists bool) {
	v := m.created_at
	if v == nil {
		return
	}
	return *v, true
}

// OldCreatedAt returns the old "created_at" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldCreatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldCreatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldCreatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldCreatedAt: %w", err)
	}
	return oldValue.CreatedAt, nil
}

// ResetCreatedAt resets all changes to the "created_at" field.
func (m *RecordMutation) ResetCreatedAt() {
	m.created_at = nil
}

// SetLastUpdatedAt sets the "last_updated_at" field.
func (m *RecordMutation) SetLastUpdatedAt(t time.Time) {
	m.last_updated_at = &t
}

// LastUpdatedAt returns the value of the "last_updated_at" field in the mutation.
func (m *RecordMutation) LastUpdatedAt() (r time.Time, exists bool) {
	v := m.last_updated_at
	if v == nil {
		return
	}
	return *v, true
}

// OldLastUpdatedAt returns the old "last_updated_at" field's value of the Record entity.
// If the Record object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordMutation) OldLastUpdatedAt(ctx context.Context) (v time.Time, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldLastUpdatedAt is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldLastUpdatedAt requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldLastUpdatedAt: %w", err)
	}
	return oldValue.LastUpdatedAt, nil
}

// ResetLastUpdatedAt resets all changes to the "last_updated_at" field.
func (m *RecordMutation) ResetLastUpdatedAt() {
	m.last_updated_at = nil
}

// Where appends a list predicates to the RecordMutation builder.
func (m *RecordMutation) Where(ps ...predicate.Record) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RecordMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (Record).
func (m *RecordMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RecordMutation) Fields() []string {
	fields := make([]string, 0, 2)
	if m.created_at != nil {
		fields = append(fields, record.FieldCreatedAt)
	}
	if m.last_updated_at != nil {
		fields = append(fields, record.FieldLastUpdatedAt)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RecordMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case record.FieldCreatedAt:
		return m.CreatedAt()
	case record.FieldLastUpdatedAt:
		return m.LastUpdatedAt()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RecordMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case record.FieldCreatedAt:
		return m.OldCreatedAt(ctx)
	case record.FieldLastUpdatedAt:
		return m.OldLastUpdatedAt(ctx)
	}
	return nil, fmt.Errorf("unknown Record field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) SetField(name string, value ent.Value) error {
	switch name {
	case record.FieldCreatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetCreatedAt(v)
		return nil
	case record.FieldLastUpdatedAt:
		v, ok := value.(time.Time)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetLastUpdatedAt(v)
		return nil
	}
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RecordMutation) AddedFields() []string {
	return nil
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RecordMutation) AddedField(name string) (ent.Value, bool) {
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordMutation) AddField(name string, value ent.Value) error {
	switch name {
	}
	return fmt.Errorf("unknown Record numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RecordMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RecordMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RecordMutation) ClearField(name string) error {
	return fmt.Errorf("unknown Record nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RecordMutation) ResetField(name string) error {
	switch name {
	case record.FieldCreatedAt:
		m.ResetCreatedAt()
		return nil
	case record.FieldLastUpdatedAt:
		m.ResetLastUpdatedAt()
		return nil
	}
	return fmt.Errorf("unknown Record field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RecordMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RecordMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RecordMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RecordMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RecordMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RecordMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RecordMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown Record unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RecordMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown Record edge %s", name)
}

// RecordFoodMutation represents an operation that mutates the RecordFood nodes in the graph.
type RecordFoodMutation struct {
	config
	op            Op
	typ           string
	id            *int
	record_id     *int
	addrecord_id  *int
	food_id       *int
	addfood_id    *int
	amount        *int
	addamount     *int
	clearedFields map[string]struct{}
	done          bool
	oldValue      func(context.Context) (*RecordFood, error)
	predicates    []predicate.RecordFood
}

var _ ent.Mutation = (*RecordFoodMutation)(nil)

// recordfoodOption allows management of the mutation configuration using functional options.
type recordfoodOption func(*RecordFoodMutation)

// newRecordFoodMutation creates new mutation for the RecordFood entity.
func newRecordFoodMutation(c config, op Op, opts ...recordfoodOption) *RecordFoodMutation {
	m := &RecordFoodMutation{
		config:        c,
		op:            op,
		typ:           TypeRecordFood,
		clearedFields: make(map[string]struct{}),
	}
	for _, opt := range opts {
		opt(m)
	}
	return m
}

// withRecordFoodID sets the ID field of the mutation.
func withRecordFoodID(id int) recordfoodOption {
	return func(m *RecordFoodMutation) {
		var (
			err   error
			once  sync.Once
			value *RecordFood
		)
		m.oldValue = func(ctx context.Context) (*RecordFood, error) {
			once.Do(func() {
				if m.done {
					err = fmt.Errorf("querying old values post mutation is not allowed")
				} else {
					value, err = m.Client().RecordFood.Get(ctx, id)
				}
			})
			return value, err
		}
		m.id = &id
	}
}

// withRecordFood sets the old RecordFood of the mutation.
func withRecordFood(node *RecordFood) recordfoodOption {
	return func(m *RecordFoodMutation) {
		m.oldValue = func(context.Context) (*RecordFood, error) {
			return node, nil
		}
		m.id = &node.ID
	}
}

// Client returns a new `ent.Client` from the mutation. If the mutation was
// executed in a transaction (ent.Tx), a transactional client is returned.
func (m RecordFoodMutation) Client() *Client {
	client := &Client{config: m.config}
	client.init()
	return client
}

// Tx returns an `ent.Tx` for mutations that were executed in transactions;
// it returns an error otherwise.
func (m RecordFoodMutation) Tx() (*Tx, error) {
	if _, ok := m.driver.(*txDriver); !ok {
		return nil, fmt.Errorf("ent: mutation is not running in a transaction")
	}
	tx := &Tx{config: m.config}
	tx.init()
	return tx, nil
}

// ID returns the ID value in the mutation. Note that the ID is only available
// if it was provided to the builder or after it was returned from the database.
func (m *RecordFoodMutation) ID() (id int, exists bool) {
	if m.id == nil {
		return
	}
	return *m.id, true
}

// SetRecordID sets the "record_id" field.
func (m *RecordFoodMutation) SetRecordID(i int) {
	m.record_id = &i
	m.addrecord_id = nil
}

// RecordID returns the value of the "record_id" field in the mutation.
func (m *RecordFoodMutation) RecordID() (r int, exists bool) {
	v := m.record_id
	if v == nil {
		return
	}
	return *v, true
}

// OldRecordID returns the old "record_id" field's value of the RecordFood entity.
// If the RecordFood object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordFoodMutation) OldRecordID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldRecordID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldRecordID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldRecordID: %w", err)
	}
	return oldValue.RecordID, nil
}

// AddRecordID adds i to the "record_id" field.
func (m *RecordFoodMutation) AddRecordID(i int) {
	if m.addrecord_id != nil {
		*m.addrecord_id += i
	} else {
		m.addrecord_id = &i
	}
}

// AddedRecordID returns the value that was added to the "record_id" field in this mutation.
func (m *RecordFoodMutation) AddedRecordID() (r int, exists bool) {
	v := m.addrecord_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetRecordID resets all changes to the "record_id" field.
func (m *RecordFoodMutation) ResetRecordID() {
	m.record_id = nil
	m.addrecord_id = nil
}

// SetFoodID sets the "food_id" field.
func (m *RecordFoodMutation) SetFoodID(i int) {
	m.food_id = &i
	m.addfood_id = nil
}

// FoodID returns the value of the "food_id" field in the mutation.
func (m *RecordFoodMutation) FoodID() (r int, exists bool) {
	v := m.food_id
	if v == nil {
		return
	}
	return *v, true
}

// OldFoodID returns the old "food_id" field's value of the RecordFood entity.
// If the RecordFood object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordFoodMutation) OldFoodID(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldFoodID is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldFoodID requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldFoodID: %w", err)
	}
	return oldValue.FoodID, nil
}

// AddFoodID adds i to the "food_id" field.
func (m *RecordFoodMutation) AddFoodID(i int) {
	if m.addfood_id != nil {
		*m.addfood_id += i
	} else {
		m.addfood_id = &i
	}
}

// AddedFoodID returns the value that was added to the "food_id" field in this mutation.
func (m *RecordFoodMutation) AddedFoodID() (r int, exists bool) {
	v := m.addfood_id
	if v == nil {
		return
	}
	return *v, true
}

// ResetFoodID resets all changes to the "food_id" field.
func (m *RecordFoodMutation) ResetFoodID() {
	m.food_id = nil
	m.addfood_id = nil
}

// SetAmount sets the "amount" field.
func (m *RecordFoodMutation) SetAmount(i int) {
	m.amount = &i
	m.addamount = nil
}

// Amount returns the value of the "amount" field in the mutation.
func (m *RecordFoodMutation) Amount() (r int, exists bool) {
	v := m.amount
	if v == nil {
		return
	}
	return *v, true
}

// OldAmount returns the old "amount" field's value of the RecordFood entity.
// If the RecordFood object wasn't provided to the builder, the object is fetched from the database.
// An error is returned if the mutation operation is not UpdateOne, or the database query fails.
func (m *RecordFoodMutation) OldAmount(ctx context.Context) (v int, err error) {
	if !m.op.Is(OpUpdateOne) {
		return v, fmt.Errorf("OldAmount is only allowed on UpdateOne operations")
	}
	if m.id == nil || m.oldValue == nil {
		return v, fmt.Errorf("OldAmount requires an ID field in the mutation")
	}
	oldValue, err := m.oldValue(ctx)
	if err != nil {
		return v, fmt.Errorf("querying old value for OldAmount: %w", err)
	}
	return oldValue.Amount, nil
}

// AddAmount adds i to the "amount" field.
func (m *RecordFoodMutation) AddAmount(i int) {
	if m.addamount != nil {
		*m.addamount += i
	} else {
		m.addamount = &i
	}
}

// AddedAmount returns the value that was added to the "amount" field in this mutation.
func (m *RecordFoodMutation) AddedAmount() (r int, exists bool) {
	v := m.addamount
	if v == nil {
		return
	}
	return *v, true
}

// ResetAmount resets all changes to the "amount" field.
func (m *RecordFoodMutation) ResetAmount() {
	m.amount = nil
	m.addamount = nil
}

// Where appends a list predicates to the RecordFoodMutation builder.
func (m *RecordFoodMutation) Where(ps ...predicate.RecordFood) {
	m.predicates = append(m.predicates, ps...)
}

// Op returns the operation name.
func (m *RecordFoodMutation) Op() Op {
	return m.op
}

// Type returns the node type of this mutation (RecordFood).
func (m *RecordFoodMutation) Type() string {
	return m.typ
}

// Fields returns all fields that were changed during this mutation. Note that in
// order to get all numeric fields that were incremented/decremented, call
// AddedFields().
func (m *RecordFoodMutation) Fields() []string {
	fields := make([]string, 0, 3)
	if m.record_id != nil {
		fields = append(fields, recordfood.FieldRecordID)
	}
	if m.food_id != nil {
		fields = append(fields, recordfood.FieldFoodID)
	}
	if m.amount != nil {
		fields = append(fields, recordfood.FieldAmount)
	}
	return fields
}

// Field returns the value of a field with the given name. The second boolean
// return value indicates that this field was not set, or was not defined in the
// schema.
func (m *RecordFoodMutation) Field(name string) (ent.Value, bool) {
	switch name {
	case recordfood.FieldRecordID:
		return m.RecordID()
	case recordfood.FieldFoodID:
		return m.FoodID()
	case recordfood.FieldAmount:
		return m.Amount()
	}
	return nil, false
}

// OldField returns the old value of the field from the database. An error is
// returned if the mutation operation is not UpdateOne, or the query to the
// database failed.
func (m *RecordFoodMutation) OldField(ctx context.Context, name string) (ent.Value, error) {
	switch name {
	case recordfood.FieldRecordID:
		return m.OldRecordID(ctx)
	case recordfood.FieldFoodID:
		return m.OldFoodID(ctx)
	case recordfood.FieldAmount:
		return m.OldAmount(ctx)
	}
	return nil, fmt.Errorf("unknown RecordFood field %s", name)
}

// SetField sets the value of a field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordFoodMutation) SetField(name string, value ent.Value) error {
	switch name {
	case recordfood.FieldRecordID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetRecordID(v)
		return nil
	case recordfood.FieldFoodID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetFoodID(v)
		return nil
	case recordfood.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.SetAmount(v)
		return nil
	}
	return fmt.Errorf("unknown RecordFood field %s", name)
}

// AddedFields returns all numeric fields that were incremented/decremented during
// this mutation.
func (m *RecordFoodMutation) AddedFields() []string {
	var fields []string
	if m.addrecord_id != nil {
		fields = append(fields, recordfood.FieldRecordID)
	}
	if m.addfood_id != nil {
		fields = append(fields, recordfood.FieldFoodID)
	}
	if m.addamount != nil {
		fields = append(fields, recordfood.FieldAmount)
	}
	return fields
}

// AddedField returns the numeric value that was incremented/decremented on a field
// with the given name. The second boolean return value indicates that this field
// was not set, or was not defined in the schema.
func (m *RecordFoodMutation) AddedField(name string) (ent.Value, bool) {
	switch name {
	case recordfood.FieldRecordID:
		return m.AddedRecordID()
	case recordfood.FieldFoodID:
		return m.AddedFoodID()
	case recordfood.FieldAmount:
		return m.AddedAmount()
	}
	return nil, false
}

// AddField adds the value to the field with the given name. It returns an error if
// the field is not defined in the schema, or if the type mismatched the field
// type.
func (m *RecordFoodMutation) AddField(name string, value ent.Value) error {
	switch name {
	case recordfood.FieldRecordID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddRecordID(v)
		return nil
	case recordfood.FieldFoodID:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddFoodID(v)
		return nil
	case recordfood.FieldAmount:
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("unexpected type %T for field %s", value, name)
		}
		m.AddAmount(v)
		return nil
	}
	return fmt.Errorf("unknown RecordFood numeric field %s", name)
}

// ClearedFields returns all nullable fields that were cleared during this
// mutation.
func (m *RecordFoodMutation) ClearedFields() []string {
	return nil
}

// FieldCleared returns a boolean indicating if a field with the given name was
// cleared in this mutation.
func (m *RecordFoodMutation) FieldCleared(name string) bool {
	_, ok := m.clearedFields[name]
	return ok
}

// ClearField clears the value of the field with the given name. It returns an
// error if the field is not defined in the schema.
func (m *RecordFoodMutation) ClearField(name string) error {
	return fmt.Errorf("unknown RecordFood nullable field %s", name)
}

// ResetField resets all changes in the mutation for the field with the given name.
// It returns an error if the field is not defined in the schema.
func (m *RecordFoodMutation) ResetField(name string) error {
	switch name {
	case recordfood.FieldRecordID:
		m.ResetRecordID()
		return nil
	case recordfood.FieldFoodID:
		m.ResetFoodID()
		return nil
	case recordfood.FieldAmount:
		m.ResetAmount()
		return nil
	}
	return fmt.Errorf("unknown RecordFood field %s", name)
}

// AddedEdges returns all edge names that were set/added in this mutation.
func (m *RecordFoodMutation) AddedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// AddedIDs returns all IDs (to other nodes) that were added for the given edge
// name in this mutation.
func (m *RecordFoodMutation) AddedIDs(name string) []ent.Value {
	return nil
}

// RemovedEdges returns all edge names that were removed in this mutation.
func (m *RecordFoodMutation) RemovedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// RemovedIDs returns all IDs (to other nodes) that were removed for the edge with
// the given name in this mutation.
func (m *RecordFoodMutation) RemovedIDs(name string) []ent.Value {
	return nil
}

// ClearedEdges returns all edge names that were cleared in this mutation.
func (m *RecordFoodMutation) ClearedEdges() []string {
	edges := make([]string, 0, 0)
	return edges
}

// EdgeCleared returns a boolean which indicates if the edge with the given name
// was cleared in this mutation.
func (m *RecordFoodMutation) EdgeCleared(name string) bool {
	return false
}

// ClearEdge clears the value of the edge with the given name. It returns an error
// if that edge is not defined in the schema.
func (m *RecordFoodMutation) ClearEdge(name string) error {
	return fmt.Errorf("unknown RecordFood unique edge %s", name)
}

// ResetEdge resets all changes to the edge with the given name in this mutation.
// It returns an error if the edge is not defined in the schema.
func (m *RecordFoodMutation) ResetEdge(name string) error {
	return fmt.Errorf("unknown RecordFood edge %s", name)
}