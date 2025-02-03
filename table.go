package morph

import "fmt"

// Table represents a mapping between an entity and a database table.
type Table struct {
	typeName       string
	name           string
	alias          string
	columnsByName  map[string]Column
	columnsByField map[string]Column
}

// SetType associates the entity type to the table.
func (t *Table) SetType(entity interface{}) {
	t.SetTypeName(fmt.Sprintf("%T", entity))
}

// SetTypeName modifies the entity type name for the table.
func (t *Table) SetTypeName(typeName string) {
	t.typeName = typeName
}

// TypeName retrieves the type name of the entity associated to the table.
func (t *Table) TypeName() string {
	return t.typeName
}

// Name retrieves the the table name.
func (t *Table) Name() string {
	return t.name
}

// SetName modifies the name of the table.
func (t *Table) SetName(name string) {
	t.name = name
}

// Alias retrieves the alias for the table.
func (t *Table) Alias() string {
	return t.alias
}

// SetAlias modifies the alias of the table.
func (t *Table) SetAlias(alias string) {
	t.alias = alias
}

// ColumnNames retrieves all of the column names for the table.
func (t *Table) ColumnNames() []string {
	var columnNames []string
	for name := range t.columnsByName {
		columnNames = append(columnNames, name)
	}
	return columnNames
}

// ColumnName retrieves the column name associated to the provide field name.
func (t *Table) ColumnName(field string) (string, error) {
	if t.columnsByField == nil {
		t.columnsByField = make(map[string]Column)
	}
	if column, ok := t.columnsByField[field]; ok {
		return column.Name(), nil
	}
	return "", fmt.Errorf("morph: no mapping for field %q", field)
}

// FieldName retrieves the field name associated to the provided column name.
func (t *Table) FieldName(name string) (string, error) {
	if t.columnsByName == nil {
		t.columnsByName = make(map[string]Column)
	}
	if column, ok := t.columnsByName[name]; ok {
		return column.Field(), nil
	}
	return "", fmt.Errorf("morph: no mapping for column %q", name)
}

// Columns retrieves all of the columns for the table.
func (t *Table) Columns() []Column {
	var columns []Column
	for _, c := range t.columnsByName {
		columns = append(columns, c)
	}
	return columns
}

// AddColumn adds a column to the table.
func (t *Table) AddColumn(column Column) error {
	if t.columnsByName == nil {
		t.columnsByName = make(map[string]Column)
	}
	if _, ok := t.columnsByName[column.Name()]; ok {
		return fmt.Errorf(
			"morph: column with name %q already exists", column.Name())
	}
	if t.columnsByField == nil {
		t.columnsByField = make(map[string]Column)
	}
	if _, ok := t.columnsByField[column.Field()]; ok {
		return fmt.Errorf(
			"morph: column with field %q already exists", column.Field())
	}
	t.columnsByName[column.Name()],
		t.columnsByField[column.Field()] = column, column
	return nil
}

// AddColumns adds all of the provided columns to the table.
func (t *Table) AddColumns(columns ...Column) error {
	for _, column := range columns {
		if err := t.AddColumn(column); err != nil {
			return err
		}
	}
	return nil
}
