package morph_test

import (
	"errors"
	"testing"

	"github.com/freerware/morph"
	"github.com/freerware/morph/mocks"
	"github.com/stretchr/testify/suite"
)

type LoadTestSuite struct {
	suite.Suite

	yamlLoader *mocks.Loader
	ymlLoader  *mocks.Loader
	jsonLoader *mocks.Loader
}

func TestLoadTestSuite(t *testing.T) {
	suite.Run(t, new(LoadTestSuite))
}

func (s *LoadTestSuite) SetupTest() {
	s.yamlLoader = &mocks.Loader{}
	s.ymlLoader = &mocks.Loader{}
	s.jsonLoader = &mocks.Loader{}
	morph.Loaders["yaml"] = s.yamlLoader
	morph.Loaders["yml"] = s.ymlLoader
	morph.Loaders["json"] = s.jsonLoader
}

func (s *LoadTestSuite) TestLoadYAML() {
	// arrange
	path := "./test_config.yaml"
	expectedConfig := morph.Configuration{
		Tables: []morph.TableConfiguration{
			{
				Name:  "user",
				Alias: "U",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "username",
						Field:         "Username",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    true,
					},
					{
						Name:          "password",
						Field:         "Password",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
					{
						Name:          "givenName",
						Field:         "GivenName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
					{
						Name:          "surname",
						Field:         "Surname",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
				},
			},
		},
	}
	s.yamlLoader.On("Load", path).Return(expectedConfig, nil)

	// action.
	actualConfig, err := morph.Load(path)

	// assert.
	s.Require().NoError(err)
	s.ElementsMatch(expectedConfig.Tables, actualConfig.Tables)
}

func (s *LoadTestSuite) TestLoadYAML_Error() {
	// arrange
	path := "./test_config.yaml"
	expectedConfig := morph.Configuration{}
	expectedErr := errors.New("whoa")
	s.yamlLoader.On("Load", path).Return(expectedConfig, expectedErr)

	// action.
	_, err := morph.Load(path)

	// assert.
	s.Require().Error(err)
}

func (s *LoadTestSuite) TestLoadJSON() {
	// arrange
	path := "./test_config.json"
	expectedConfig := morph.Configuration{
		Tables: []morph.TableConfiguration{
			{
				Name:  "user",
				Alias: "U",
				Columns: []morph.ColumnConfiguration{
					{
						Name:          "username",
						Field:         "Username",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    true,
					},
					{
						Name:          "password",
						Field:         "Password",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
					{
						Name:          "givenName",
						Field:         "GivenName",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
					{
						Name:          "surname",
						Field:         "Surname",
						FieldType:     "string",
						FieldStrategy: morph.FieldStrategyMethod,
						PrimaryKey:    false,
					},
				},
			},
		},
	}
	s.jsonLoader.On("Load", path).Return(expectedConfig, nil)

	// action.
	actualConfig, err := morph.Load(path)

	// assert.
	s.Require().NoError(err)
	s.ElementsMatch(expectedConfig.Tables, actualConfig.Tables)
}

func (s *LoadTestSuite) TestLoadJSON_Error() {
	// arrange
	path := "./test_config.json"
	expectedConfig := morph.Configuration{}
	expectedErr := errors.New("whoa")
	s.jsonLoader.On("Load", path).Return(expectedConfig, expectedErr)

	// action.
	_, err := morph.Load(path)

	// assert.
	s.Require().Error(err)
}

func (s *LoadTestSuite) TestLoad_MissingLoader() {
	// arrange
	path := "./test_config.txt"

	// action.
	_, err := morph.Load(path)

	// assert.
	s.Require().Error(err)
}

func (s *LoadTestSuite) TearDownTest() {
	morph.Loaders["yaml"] = morph.YAMLLoader{}
	morph.Loaders["yml"] = morph.YAMLLoader{}
	morph.Loaders["json"] = morph.JSONLoader{}
}

type JSONLoaderTestSuite struct {
	suite.Suite

	sut morph.Loader
}

func TestJSONLoaderTestSuite(t *testing.T) {
	suite.Run(t, new(JSONLoaderTestSuite))
}

func (s *JSONLoaderTestSuite) SetupTest() {
	s.sut = morph.JSONLoader{}
}

func (s *JSONLoaderTestSuite) TestLoad() {
	// arrange.
	path := "./test_config.json"

	// action.
	_, err := s.sut.Load(path)

	// assert.
	s.Require().NoError(err)
}

func (s *JSONLoaderTestSuite) TestLoad_MissingFile() {
	// arrange.
	path := "./missing_test_config.json"

	// action.
	_, err := s.sut.Load(path)

	// assert.
	s.Require().Error(err)
}

type YAMLLoaderTestSuite struct {
	suite.Suite

	sut morph.Loader
}

func TestYAMLLoaderTestSuite(t *testing.T) {
	suite.Run(t, new(YAMLLoaderTestSuite))
}

func (s *YAMLLoaderTestSuite) SetupTest() {
	s.sut = morph.YAMLLoader{}
}

func (s *YAMLLoaderTestSuite) TestLoad() {
	// arrange.
	path := "./test_config.yaml"

	// action.
	_, err := s.sut.Load(path)

	// assert.
	s.Require().NoError(err)
}

func (s *YAMLLoaderTestSuite) TestLoad_MissingFile() {
	// arrange.
	path := "./missing_test_config.yaml"

	// action.
	_, err := s.sut.Load(path)

	// assert.
	s.Require().Error(err)
}
