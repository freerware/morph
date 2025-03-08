package morph

import "strings"

type FieldStrategy int

const (
	FieldStrategyStructField FieldStrategy = iota
	FieldStrategyMethod
)

// ColumnPredicate represents a condition used to match columns within find operations.
type ColumnPredicate func(Column) bool

// Column represents a mapping between an entity field
// and a database column.
type Column struct {
	name          string
	field         string
	fieldStrategy FieldStrategy
	fieldType     string
	primaryKey    bool
}

// Name retrieves the name of the column.
func (c *Column) Name() string {
	return c.name
}

// PrimaryKey indicates if the column plays a role in the primary key
// for a table. Tables with composite primary keys will have multiple
// columns that indicate a true value for this method.
func (c *Column) PrimaryKey() bool {
	return c.primaryKey
}

func (c *Column) SetPrimaryKey(pKey bool) {
	c.primaryKey = pKey
}

// SetName modifies the name of the column.
func (c *Column) SetName(name string) {
	c.name = strings.TrimSpace(name)
}

// Field retrieves the field name associated to the column.
func (c *Column) Field() string {
	return c.field
}

// SetField modifies the field name associated to the column.
func (c *Column) SetField(field string) {
	c.field = strings.TrimSpace(field)
}

// Strategy retrieves the field strategy associated to the column.
func (c *Column) Strategy() FieldStrategy {
	return c.fieldStrategy
}

// SetStrategy modifies the field strategy associated to the column.
func (c *Column) SetStrategy(strategy FieldStrategy) {
	c.fieldStrategy = strategy
}

// UsingMethodStrategy determines if the column is using the method field strategy.
func (c *Column) UsingMethodStrategy() bool {
	return c.Strategy() == FieldStrategyMethod
}

// UsingStructFieldStrategy determines if the column is using the struct field strategy.
func (c *Column) UsingStructFieldStrategy() bool {
	return c.Strategy() == FieldStrategyStructField
}

// FieldType retrieves the field type associated to the column.
func (c *Column) FieldType() string {
	return c.fieldType
}

// SetFieldType modifies the field type associated to the column.
func (c *Column) SetFieldType(fieldType string) {
	c.fieldType = strings.TrimSpace(fieldType)
}
