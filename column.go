package morph

type FieldStrategy int

const (
	FieldStrategyStructField FieldStrategy = iota
	FieldStrategyMethod
)

// Column represents a mapping between an entity field
// and a database column.
type Column struct {
	name          string
	field         string
	fieldStrategy FieldStrategy
	fieldType     string
}

// Name retrieves the name of the column.
func (c *Column) Name() string {
	return c.name
}

// SetName modifies the name of the column.
func (c *Column) SetName(name string) {
	c.name = name
}

// Field retrieves the field name associated to the column.
func (c *Column) Field() string {
	return c.field
}

// SetField modifies the field name associated to the column.
func (c *Column) SetField(field string) {
	c.field = field
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
	c.fieldType = fieldType
}
