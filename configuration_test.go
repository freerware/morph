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
				},
			},
		},
	}

	// action.
	tables := config.AsMetadata()

	// assert.
	s.Require().Len(config.Tables, len(tables))
	s.Equal(config.Tables[0].TypeName, tables[0].TypeName())
	s.Equal(config.Tables[0].Name, tables[0].Name())
	s.Equal(config.Tables[0].Alias, tables[0].Alias())

	s.Require().Len(config.Tables[0].Columns, len(tables[0].Columns()))
	s.Equal(config.Tables[0].Columns[0].Name, tables[0].Columns()[0].Name())
	s.Equal(config.Tables[0].Columns[0].Field, tables[0].Columns()[0].Field())
	s.Equal(config.Tables[0].Columns[0].FieldType, tables[0].Columns()[0].FieldType())
	s.True(tables[0].Columns()[0].UsingStructFieldStrategy())
	s.Equal(config.Tables[0].Columns[0].PrimaryKey, tables[0].Columns()[0].PrimaryKey())
}
