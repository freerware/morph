package morph

// Configuration represents the configuration used to construct
// the table and column mappings.
type Configuration struct {
	Tables []TableConfiguration `json:"tables" yaml:"tables"`
}

// AsMetadata converts the configuration to metadata mappings.
func (c Configuration) AsMetadata() []Table {
	var tables []Table
	for _, t := range c.Tables {
		var table Table
		table.SetTypeName(t.TypeName)
		table.SetName(t.Name)
		table.SetAlias(t.Alias)
		for _, c := range t.Columns {
			var column Column
			column.SetField(c.Field)
			column.SetName(c.Name)
			table.AddColumn(column)
		}
		tables = append(tables, table)
	}
	return tables
}

// TableConfiguration represents the configuration used to construct
// a single table mapping.
type TableConfiguration struct {
	TypeName string                `json:"typeName" yaml:"typeName"`
	Name     string                `json:"name" yaml:"name"`
	Alias    string                `json:"alias" yaml:"alias"`
	Columns  []ColumnConfiguration `json:"columns" yaml:"columns"`
}

// ColumnConfiguration represents the configuration used to construct
// a single column mapping.
type ColumnConfiguration struct {
	Name  string `json:"name" yaml:"name"`
	Field string `json:"field" yaml:"field"`
}
