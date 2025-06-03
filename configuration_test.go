package morph_test

import (
	"testing"

	"github.com/freerware/morph"
	"github.com/stretchr/testify/suite"
)

type ConfigurationTestSuite struct {
	suite.Suite
}

func TestConfigurationTestSuite(t *testing.T) {
	suite.Run(t, new(ConfigurationTestSuite))
}

func (s *ConfigurationTestSuite) TestAsMetadata() {
	// arrange.
	config := morph.Configuration{
		Tables: []morph.TableConfiguration{
			{
				TypeName: "example.User",
				Name:     "user",
				Alias:    "U",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "username",
						Field:         "Username",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    true,
					},
					{
						Name:          "business_name",
						Field:         "BusinessName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    false,
					},
				},
				References: []morph.ReferenceConfiguration{
					{
						TableName: "business",
						ForeignKey: []morph.ForeignKeyConfiguration{
							{
								ColumnName: "business_name",
							},
						},
					},
				},
			},
			{
				TypeName: "example.Business",
				Name:     "business",
				Alias:    "B",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "name",
						Field:         "Name",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    true,
					},
					{
						Name:          "user_count",
						Field:         "UserCount",
						FieldType:     "int",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    false,
					},
				},
			},
		},
	}

	// action.
	tables := config.AsMetadata()

	// assert.
	tablesByName := make(map[string]morph.Table)
	for _, table := range tables {
		tablesByName[table.Name()] = table
	}

	s.Require().Len(config.Tables, len(tables))

	userTableConfig := config.Tables[0]
	businessTableConfig := config.Tables[1]
	userTable := tablesByName[config.Tables[0].Name]
	businessTable := tablesByName[config.Tables[1].Name]

	s.Equal(userTableConfig.TypeName, userTable.TypeName())
	s.Equal(userTableConfig.Name, userTable.Name())
	s.Equal(userTableConfig.Alias, userTable.Alias())

	s.Equal(businessTableConfig.TypeName, businessTable.TypeName())
	s.Equal(businessTableConfig.Name, businessTable.Name())
	s.Equal(businessTableConfig.Alias, businessTable.Alias())

	s.Require().Len(userTableConfig.Columns, len(userTable.Columns()))
	for _, columnConfig := range userTableConfig.Columns {
		s.Require().True(userTable.HasColumn(columnConfig.Name))
		column := userTable.FindColumns(func(c morph.Column) bool { return c.Name() == columnConfig.Name })[0]
		s.Equal(columnConfig.Name, column.Name())
		s.Equal(columnConfig.Field, column.Field())
		s.Equal(columnConfig.FieldType, column.FieldType())
		s.True(column.UsingStructFieldStrategy())

		if columnConfig.PrimaryKey {
			s.True(column.PrimaryKey())
		}
	}

	for _, columnConfig := range businessTableConfig.Columns {
		s.Require().True(businessTable.HasColumn(columnConfig.Name))
		column := businessTable.FindColumns(func(c morph.Column) bool { return c.Name() == columnConfig.Name })[0]
		s.Equal(columnConfig.Name, column.Name())
		s.Equal(columnConfig.Field, column.Field())
		s.Equal(columnConfig.FieldType, column.FieldType())
		s.True(column.UsingStructFieldStrategy())

		if columnConfig.PrimaryKey {
			s.True(column.PrimaryKey())
		}
	}

	s.Require().Len(userTableConfig.References, len(userTable.ReferencesTo()))
	s.True(businessTable.IsReferencedBy(userTable))
	s.Require().Len(userTableConfig.References, len(businessTable.ReferencedBy()))
	s.True(userTable.HasReferenceTo(businessTable))
}

func (s *ConfigurationTestSuite) TestAsMetadata_InvalidTableReference() {
	// arrange.
	config := morph.Configuration{
		Tables: []morph.TableConfiguration{
			{
				TypeName: "example.User",
				Name:     "user",
				Alias:    "U",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "username",
						Field:         "Username",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    true,
					},
					{
						Name:          "business_name",
						Field:         "BusinessName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    false,
					},
				},
				References: []morph.ReferenceConfiguration{
					{
						TableName: "business",
						ForeignKey: []morph.ForeignKeyConfiguration{
							{
								ColumnName: "business_name",
							},
						},
					},
				},
			},
		},
	}

	// action.
	tables := config.AsMetadata()

	// assert.
	tablesByName := make(map[string]morph.Table)
	for _, table := range tables {
		tablesByName[table.Name()] = table
	}

	s.Require().Len(config.Tables, len(tables))

	userTableConfig := config.Tables[0]
	userTable := tablesByName[config.Tables[0].Name]

	s.Equal(userTableConfig.TypeName, userTable.TypeName())
	s.Equal(userTableConfig.Name, userTable.Name())
	s.Equal(userTableConfig.Alias, userTable.Alias())

	s.Require().Len(userTableConfig.Columns, len(userTable.Columns()))
	for _, columnConfig := range userTableConfig.Columns {
		s.Require().True(userTable.HasColumn(columnConfig.Name))
		column := userTable.FindColumns(func(c morph.Column) bool { return c.Name() == columnConfig.Name })[0]
		s.Equal(columnConfig.Name, column.Name())
		s.Equal(columnConfig.Field, column.Field())
		s.Equal(columnConfig.FieldType, column.FieldType())
		s.True(column.UsingStructFieldStrategy())

		if columnConfig.PrimaryKey {
			s.True(column.PrimaryKey())
		}
	}

	s.Require().Len(userTable.ReferencesTo(), 0)
}

func (s *ConfigurationTestSuite) TestAsMetadata_DuplicateColumn() {
	// arrange.
	config := morph.Configuration{
		Tables: []morph.TableConfiguration{
			{
				TypeName: "example.User",
				Name:     "user",
				Alias:    "U",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "username",
						Field:         "Username",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    true,
					},
					{
						Name:          "business_name",
						Field:         "BusinessName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    false,
					},
					{
						Name:          "business_name",
						Field:         "BusinessName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyStructField,
						PrimaryKey:    false,
					},
				},
			},
		},
	}

	// action.
	tables := config.AsMetadata()

	// assert.
	tablesByName := make(map[string]morph.Table)
	for _, table := range tables {
		tablesByName[table.Name()] = table
	}

	s.Require().Len(config.Tables, len(tables))

	userTableConfig := config.Tables[0]
	userTable := tablesByName[config.Tables[0].Name]

	s.Equal(userTableConfig.TypeName, userTable.TypeName())
	s.Equal(userTableConfig.Name, userTable.Name())
	s.Equal(userTableConfig.Alias, userTable.Alias())

	s.Require().Len(userTable.Columns(), len(userTableConfig.Columns)-1)
	for _, columnConfig := range userTableConfig.Columns {
		s.Require().True(userTable.HasColumn(columnConfig.Name))
		column := userTable.FindColumns(func(c morph.Column) bool { return c.Name() == columnConfig.Name })[0]
		s.Equal(columnConfig.Name, column.Name())
		s.Equal(columnConfig.Field, column.Field())
		s.Equal(columnConfig.FieldType, column.FieldType())
		s.True(column.UsingStructFieldStrategy())

		if columnConfig.PrimaryKey {
			s.True(column.PrimaryKey())
		}
	}
}
