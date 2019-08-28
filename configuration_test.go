package morph

import (
	"testing"

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
	config := Configuration{
		Tables: []TableConfiguration{
			{
				TypeName: "example.User",
				Name:     "user",
				Alias:    "U",
				Columns: []ColumnConfiguration{
					{
						Name:  "username",
						Field: "Username",
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
	s.Require().Len(
		config.Tables[0].Columns, len(tables[0].Columns()))
	s.Equal(
		config.Tables[0].Columns[0].Name, tables[0].Columns()[0].Name())
	s.Equal(
		config.Tables[0].Columns[0].Field, tables[0].Columns()[0].Field())
}
