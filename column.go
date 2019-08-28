package morph

// Column represents a mapping between an entity field
// and a database column.
type Column struct {
	name  string
	field string
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
