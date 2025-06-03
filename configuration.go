package morph

// Configuration represents the configuration used to construct
// the table and column mappings.
type Configuration struct {
	Tables []TableConfiguration `json:"tables" yaml:"tables"`
}

// AsMetadata converts the configuration to metadata mappings.
func (c Configuration) AsMetadata() []Table {
	var tables []Table

	references := map[string][]ReferenceConfiguration{}
	tablesByName := map[string]*Table{}
	columnsByName := map[string]*Column{}

	for _, t := range c.Tables {
		table := t.AsTable()
		for _, column := range table.Columns() {
			columnsByName[column.Name()] = &column
		}
		tablesByName[table.Name()] = &table

		if len(t.References) > 0 {
			if _, ok := references[table.Name()]; !ok {
				references[table.Name()] = []ReferenceConfiguration{}
			}
			references[table.Name()] = append(references[table.Name()], t.References...)
		}
	}

	for childTableName, refs := range references {
		child, ok := tablesByName[childTableName]
		if !ok {
			continue
		}
		for _, ref := range refs {
			parent, ok := tablesByName[ref.TableName]
			if !ok {
				continue
			}
			key := []Column{}
			for _, fk := range ref.ForeignKey {
				if column, ok := columnsByName[fk.ColumnName]; ok {
					key = append(key, *column)
				}
			}
			if _, err := child.References(parent, key); err != nil {
				continue
			}
		}
	}

	for _, table := range tablesByName {
		tables = append(tables, *table)
	}

	return tables
}

// TableConfiguration represents the configuration used to construct
// a single table mapping.
type TableConfiguration struct {
	TypeName   string                   `json:"typeName" yaml:"typeName"`
	Name       string                   `json:"name" yaml:"name"`
	Alias      string                   `json:"alias" yaml:"alias"`
	Columns    []ColumnConfiguration    `json:"columns" yaml:"columns"`
	References []ReferenceConfiguration `json:"references" yaml:"references"`
}

// AsTable converts the configuration to a table mapping.
func (t TableConfiguration) AsTable() Table {
	var table Table
	table.SetTypeName(t.TypeName)
	table.SetName(t.Name)
	table.SetAlias(t.Alias)

	for _, c := range t.Columns {
		if err := table.AddColumn(c.AsColumn()); err != nil {
			continue
		}
	}

	return table
}

// ReferenceConfiguration represents the configuration used to construct
// a single reference between two tables.
type ReferenceConfiguration struct {
	TableName  string                    `json:"tableName" yaml:"tableName"`
	ForeignKey []ForeignKeyConfiguration `json:"foreignKey" yaml:"foreignKey"`
}

// ForeignKeyConfiguration represents the configuration used to communicate
// which columns comprise the foreign key for a reference.
type ForeignKeyConfiguration struct {
	ColumnName string `json:"columnName" yaml:"columnName"`
}

// ColumnConfiguration represents the configuration used to construct
// a single column mapping.
type ColumnConfiguration struct {
	Name          string        `json:"name" yaml:"name"`
	Field         string        `json:"field" yaml:"field"`
	FieldType     string        `json:"fieldType" yaml:"fieldType"`
	FieldStrategy FieldStrategy `json:"fieldStrategy" yaml:"fieldStrategy"`
	PrimaryKey    bool          `json:"primaryKey" yaml:"primaryKey"`
}

// AsColumn converts the configuration to a column mapping.
func (c ColumnConfiguration) AsColumn() Column {
	var column Column
	column.SetField(c.Field)
	column.SetName(c.Name)
	column.SetFieldType(c.FieldType)
	column.SetStrategy(c.FieldStrategy)
	column.SetPrimaryKey(c.PrimaryKey)
	return column
}
