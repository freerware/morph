package morph_test

import (
	"testing"
	"time"

	"github.com/freerware/morph"
	"github.com/stretchr/testify/suite"
)

type TestModel struct {
	ID          int `db:"identifier"`
	Name        *string
	Another     AnotherTestModel
	MaybeIgnore bool `db:"-"`
	UpdatedAt   time.Time
	DeletedAt   *time.Time
}

func (t *TestModel) CreatedAt() time.Time {
	return time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local)
}

func (t TestModel) Ignored() {}

func (t *TestModel) SetName(name string) {
	t.Name = &name
}

func (t TestModel) AnotherPtr() *AnotherTestModel {
	return &t.Another
}

type AnotherTestModel struct {
	ID          int
	Title       string
	Description *string
	secret      string
}

type ReflectTestSuite struct {
	suite.Suite

	obj    any
	objPtr any
}

func TestReflectTestSuite(t *testing.T) {
	suite.Run(t, new(ReflectTestSuite))
}

func (s *ReflectTestSuite) SetupTest() {
	name := "test"
	m := TestModel{
		ID:        1,
		Name:      &name,
		UpdatedAt: time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
			secret:      "shhh!",
		},
	}

	s.obj = m
	s.objPtr = &m
}

func (s *ReflectTestSuite) TestReflect_WithValue() {
	tests := []struct {
		name     string
		obj      any
		options  []morph.ReflectOption
		expected func() morph.Table
		err      error
	}{
		{
			name:    "NotStruct_Error",
			obj:     2,
			options: []morph.ReflectOption{},
			err:     morph.ErrNotStruct,
		},
		{
			name:    "WithNoOptions_AppliesDefaults",
			obj:     s.obj,
			options: []morph.ReflectOption{},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_Uppercase_SetUppercaseTableAlias",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.UpperCaseStrategy, morph.DefaultTableAliasLength)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_Lowercase_SetsLowercaseTableAlias",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.LowerCaseStrategy, morph.DefaultTableAliasLength)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("t")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_CustomLength_UsesCustomLength",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.DefaultTableAliasStrategy, 4)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("TEST")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_NegativeLength_UsesDefaultLength",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.DefaultTableAliasStrategy, -1)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableName_Singular_SetsTableNameWithSingular",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableName(morph.SnakeCaseStrategy, false)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_model")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableName_CamelCase_SetsTableNameToCamelCase",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredTableName(morph.CamelCaseStrategy, true)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("TestModels")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_Snakecase_SetsToSnakecase",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.SnakeCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name: "WithInferredColumnNames_ScreamingSnakeCase_SetsToScreamingSnakeCase",
			obj:  s.obj,
			options: []morph.ReflectOption{
				morph.WithInferredColumnNames(morph.ScreamingSnakeCaseStrategy),
				morph.WithPrimaryKeyColumn("ID"),
			},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("ID")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("NAME")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CREATED_AT")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MAYBE_IGNORE")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UPDATED_AT")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DELETED_AT")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name: "WithInferredColumnNames_Uppercase_SetsToUppercase",
			obj:  s.obj,
			options: []morph.ReflectOption{
				morph.WithInferredColumnNames(morph.UpperCaseStrategy),
				morph.WithPrimaryKeyColumn("ID"),
			},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("ID")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("NAME")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CREATEDAT")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MAYBEIGNORE")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UPDATEDAT")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DELETEDAT")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_Lowercase_SetsToLowercase",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.LowerCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("createdat")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybeignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updatedat")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deletedat")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_Camelcase_SetsToCamelcase",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.CamelCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("Id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(false)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("Name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CreatedAt")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MaybeIgnore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UpdatedAt")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DeletedAt")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTableName_SetsTableName",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithTableName("TEST_MOD")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("TEST_MOD")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTableAlias_SetsTableAlias",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithTableAlias("tm1")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("tm1")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTag_SetsTag",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithTag("db")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("identifier")
				idColumn.SetField("ID")
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMethods_SetsMethodExclusions",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithoutMethods("CreatedAt")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMatchingMethods_SetsMethodExclusionPatterns",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithoutMatchingMethods("^CreatedAt$")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutFields_SetsFieldExclusions",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithoutFields("Name")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMatchingFields_SetsFieldExclusionPatterns",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithoutMatchingFields("^Name$")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithColumnNameMapping",
			obj:     s.obj,
			options: []morph.ReflectOption{morph.WithColumnNameMapping("Name", "given_name")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.obj)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("given_name")
				nameColumn.SetField("Name")
				nameColumn.SetPrimaryKey(false)
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetPrimaryKey(false)
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// action.
			actual, err := morph.Reflect(test.obj, test.options...)

			// assert.
			if test.expected != nil {
				expected := test.expected()

				s.Equal(expected.Name(), actual.Name())
				s.Equal(expected.Alias(), actual.Alias())
				s.Equal(expected.TypeName(), actual.TypeName())
				s.ElementsMatch(expected.ColumnNames(), actual.ColumnNames())
				s.ElementsMatch(expected.Columns(), actual.Columns())
			}

			s.Equal(test.err, err)
		})
	}
}

func (s *ReflectTestSuite) TestReflect_WithPointer() {
	notStruct := 2

	tests := []struct {
		name     string
		obj      any
		options  []morph.ReflectOption
		expected func() morph.Table
		err      error
	}{
		{
			name:    "NotStruct_Error",
			obj:     &notStruct,
			options: []morph.ReflectOption{},
			err:     morph.ErrNotStruct,
		},
		{
			name:    "WithNoOptions_AppliesDefaults",
			obj:     s.objPtr,
			options: []morph.ReflectOption{},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_Uppercase_SetUppercaseTableAlias",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.UpperCaseStrategy, morph.DefaultTableAliasLength)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_Lowercase_SetsLowercaseTableAlias",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.LowerCaseStrategy, morph.DefaultTableAliasLength)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("t")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_CustomLength_UsesCustomLength",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.DefaultTableAliasStrategy, 4)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("TEST")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableAlias_NegativeLength_UsesDefaultLength",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableAlias(morph.DefaultTableAliasStrategy, -1)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableName_Singular_SetsTableNameWithSingular",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableName(morph.SnakeCaseStrategy, false)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_model")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredTableName_CamelCase_SetsTableNameToCamelCase",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredTableName(morph.CamelCaseStrategy, true)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("TestModels")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_Snakecase_SetsToSnakecase",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.SnakeCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name: "WithInferredColumnNames_ScreamingSnakeCase_SetsToScreamingSnakeCase",
			obj:  s.objPtr,
			options: []morph.ReflectOption{
				morph.WithInferredColumnNames(morph.ScreamingSnakeCaseStrategy),
				morph.WithPrimaryKeyColumn("ID"),
			},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("ID")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("NAME")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CREATED_AT")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MAYBE_IGNORE")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UPDATED_AT")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DELETED_AT")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name: "WithInferredColumnNames_Uppercase_SetsToUppercase",
			obj:  s.objPtr,
			options: []morph.ReflectOption{
				morph.WithInferredColumnNames(morph.UpperCaseStrategy),
				morph.WithPrimaryKeyColumn("ID"),
			},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("ID")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("NAME")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CREATEDAT")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MAYBEIGNORE")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UPDATEDAT")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DELETEDAT")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_LowerCase_SetsToLowerCase",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.LowerCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("createdat")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybeignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updatedat")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deletedat")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithInferredColumnNames_Camelcase_SetsToCamelcase",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithInferredColumnNames(morph.CamelCaseStrategy)},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("Id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(false)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("Name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("CreatedAt")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("MaybeIgnore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("UpdatedAt")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("DeletedAt")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTableName_SetsTableName",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithTableName("TEST_MOD")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("TEST_MOD")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTableAlias_SetsTableAlias",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithTableAlias("tm1")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("tm1")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithTag_SetsTag",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithTag("db")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("identifier")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(false)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, createdAtColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMethods_SetsMethodExclusions",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithoutMethods("CreatedAt")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMatchingMethods_SetsMethodExclusionPatterns",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithoutMatchingMethods("^CreatedAt$")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutFields_SetsFieldExclusions",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithoutFields("Name")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithoutMatchingFields_SetsFieldExclusionPatterns",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithoutMatchingFields("^Name$")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithPrimaryKeyColumn",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithPrimaryKeyColumn("id")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetPrimaryKey(false)
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetPrimaryKey(false)
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithPrimaryKeyColumns",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithPrimaryKeyColumns("id", "created_at")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("name")
				nameColumn.SetField("Name")
				nameColumn.SetPrimaryKey(false)
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetPrimaryKey(true)
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
		{
			name:    "WithColumnNameMapping",
			obj:     s.objPtr,
			options: []morph.ReflectOption{morph.WithColumnNameMapping("Name", "given_name")},
			expected: func() morph.Table {
				t := morph.Table{}
				t.SetType(s.objPtr)
				t.SetName("test_models")
				t.SetAlias("T")

				columns := []morph.Column{}

				var idColumn morph.Column
				idColumn.SetName("id")
				idColumn.SetField("ID")
				idColumn.SetPrimaryKey(true)
				idColumn.SetStrategy(morph.FieldStrategyStructField)
				idColumn.SetFieldType("int")

				var nameColumn morph.Column
				nameColumn.SetName("given_name")
				nameColumn.SetField("Name")
				nameColumn.SetPrimaryKey(false)
				nameColumn.SetStrategy(morph.FieldStrategyStructField)
				nameColumn.SetFieldType("*string")

				var createdAtColumn morph.Column
				createdAtColumn.SetName("created_at")
				createdAtColumn.SetField("CreatedAt")
				createdAtColumn.SetPrimaryKey(false)
				createdAtColumn.SetStrategy(morph.FieldStrategyMethod)
				createdAtColumn.SetFieldType("time.Time")

				var maybeIgnoreColumn morph.Column
				maybeIgnoreColumn.SetName("maybe_ignore")
				maybeIgnoreColumn.SetField("MaybeIgnore")
				maybeIgnoreColumn.SetStrategy(morph.FieldStrategyStructField)
				maybeIgnoreColumn.SetFieldType("bool")

				var updatedAtColumn morph.Column
				updatedAtColumn.SetName("updated_at")
				updatedAtColumn.SetField("UpdatedAt")
				updatedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				updatedAtColumn.SetFieldType("time.Time")

				var deletedAtColumn morph.Column
				deletedAtColumn.SetName("deleted_at")
				deletedAtColumn.SetField("DeletedAt")
				deletedAtColumn.SetStrategy(morph.FieldStrategyStructField)
				deletedAtColumn.SetFieldType("*time.Time")

				columns = append(columns, idColumn, createdAtColumn, nameColumn, maybeIgnoreColumn, updatedAtColumn, deletedAtColumn)
				if err := t.AddColumns(columns...); err != nil {
					s.FailNow("failed to setup expectations for test: %v", err.Error())
				}
				return t
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// action.
			actual, err := morph.Reflect(test.obj, test.options...)

			// assert.
			if test.expected != nil {
				expected := test.expected()

				s.Equal(expected.Name(), actual.Name())
				s.Equal(expected.Alias(), actual.Alias())
				s.Equal(expected.TypeName(), actual.TypeName())
				s.ElementsMatch(expected.ColumnNames(), actual.ColumnNames())
				s.ElementsMatch(expected.Columns(), actual.Columns())
			}

			s.Equal(test.err, err)
		})
	}
}
