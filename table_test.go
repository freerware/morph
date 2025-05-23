package morph_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/freerware/morph"
	"github.com/stretchr/testify/suite"
)

type TableTestSuite struct {
	suite.Suite

	sut morph.Table
}

func TestTableTestSuite(t *testing.T) {
	suite.Run(t, new(TableTestSuite))
}

func (s *TableTestSuite) SetupTest() {
	s.sut = morph.Table{}
}

func (s *TableTestSuite) SetupSubTest() {
	s.sut = morph.Table{}
}

func (s *TableTestSuite) TestTable_TypeName() {
	// arrange.
	expectedTypeName := "example.User"
	s.sut.SetTypeName(expectedTypeName)

	// action.
	actualTypeName := s.sut.TypeName()

	// assert.
	s.Equal(expectedTypeName, actualTypeName)
}

func (s *TableTestSuite) TestTable_SetTypeName() {
	// arrange.
	expectedTypeName := "example.User"

	// action.
	s.sut.SetTypeName(expectedTypeName)

	// assert.
	s.Equal(expectedTypeName, s.sut.TypeName())
}

func (s *TableTestSuite) TestTable_SetType() {
	// arrange.
	t := struct {
		Username string
		Password string
	}{
		"fr33r",
		"m3talm1nd",
	}

	// action.
	s.sut.SetType(t)

	// assert.
	s.Equal(fmt.Sprintf("%T", t), s.sut.TypeName())
}

func (s *TableTestSuite) TestTable_Name() {
	// arrange.
	expectedName := "user"
	s.sut.SetName(expectedName)

	// action.
	actualName := s.sut.Name()

	// assert.
	s.Equal(expectedName, actualName)
}

func (s *TableTestSuite) TestTable_SetName() {
	// arrange.
	expectedName := "user"

	// action.
	s.sut.SetName(expectedName)

	// assert.
	s.Equal(expectedName, s.sut.Name())
}

func (s *TableTestSuite) TestTable_Alias() {
	// arrange.
	expectedAlias := "U"
	s.sut.SetAlias(expectedAlias)

	// action.
	actualAlias := s.sut.Alias()

	// assert.
	s.Equal(expectedAlias, actualAlias)
}

func (s *TableTestSuite) TestTable_SetAlias() {
	// arrange.
	expectedAlias := "U"

	// action.
	s.sut.SetAlias(expectedAlias)

	// assert.
	s.Equal(expectedAlias, s.sut.Alias())
}

func (s *TableTestSuite) TestTable_ColumnNames() {
	// arrange.
	usernameField := "Username"
	usernameColumnName := "username"
	passwordField := "Password"
	passwordColumnName := "password"
	usernameColumn := morph.Column{}
	usernameColumn.SetField(usernameField)
	usernameColumn.SetName(usernameColumnName)
	passwordColumn := morph.Column{}
	passwordColumn.SetField(passwordField)
	passwordColumn.SetName(passwordColumnName)
	columns := []morph.Column{usernameColumn, passwordColumn}
	expectedColumnNames := []string{usernameColumnName, passwordColumnName}
	s.Require().NoError(s.sut.AddColumns(columns...))

	// action.
	actualColumnNames := s.sut.ColumnNames()

	// assert.
	s.Require().Len(actualColumnNames, len(expectedColumnNames))
	s.ElementsMatch(expectedColumnNames, actualColumnNames)
}

func (s *TableTestSuite) TestTable_ColumnName() {
	// arrange.
	usernameField := "Username"
	usernameColumnName := "username"
	usernameColumn := morph.Column{}
	usernameColumn.SetField(usernameField)
	usernameColumn.SetName(usernameColumnName)
	s.Require().NoError(s.sut.AddColumn(usernameColumn))

	// action.
	actualColumnName, err := s.sut.ColumnName(usernameField)

	// assert.
	s.Require().NoError(err)
	s.Equal(usernameColumnName, actualColumnName)
}

func (s *TableTestSuite) TestTable_ColumnName_MissingMapping() {
	// arrange.
	usernameField := "Username"

	// action.
	_, err := s.sut.ColumnName(usernameField)

	// assert.
	s.Require().Error(err)
}

func (s *TableTestSuite) TestTable_FieldName() {
	// arrange.
	usernameField := "Username"
	usernameColumnName := "username"
	usernameColumn := morph.Column{}
	usernameColumn.SetField(usernameField)
	usernameColumn.SetName(usernameColumnName)
	s.Require().NoError(s.sut.AddColumn(usernameColumn))

	// action.
	actualFieldName, err := s.sut.FieldName(usernameColumnName)

	// assert.
	s.Require().NoError(err)
	s.Equal(usernameField, actualFieldName)
}

func (s *TableTestSuite) TestTable_FieldName_MissingMapping() {
	// arrange.
	usernameColumnName := "username"

	// action.
	_, err := s.sut.FieldName(usernameColumnName)

	// assert.
	s.Require().Error(err)
}

func (s *TableTestSuite) TestTable_Columns() {
	// arrange.
	field := "Username"
	cName := "username"
	column := morph.Column{}
	column.SetField(field)
	column.SetName(cName)
	expectedColumns := []morph.Column{column}
	s.Require().NoError(s.sut.AddColumns(expectedColumns...))

	// action.
	actualColumns := s.sut.Columns()

	// assert.
	s.Require().Len(actualColumns, len(expectedColumns))
	s.ElementsMatch(expectedColumns, actualColumns)
}

func (s *TableTestSuite) TestTable_AddColumn() {
	tests := []struct {
		name         string
		preparations func() morph.Column
		assertions   func(err error)
	}{
		{
			name: "NotPreviouslyAdded",
			preparations: func() morph.Column {
				field := "Username"
				cName := "username"
				column := morph.Column{}
				column.SetField(field)
				column.SetName(cName)
				return column
			},
			assertions: func(err error) {
				s.NoError(err)
			},
		},
		{
			name: "PreviouslyAdded",
			preparations: func() morph.Column {
				field := "Username"
				cName := "username"
				column := morph.Column{}
				column.SetField(field)
				column.SetName(cName)
				columns := []morph.Column{column}
				s.Require().NoError(s.sut.AddColumns(columns...))

				duplicateColumn := morph.Column{}
				duplicateColumn.SetField(field)
				duplicateColumn.SetName(cName)
				return duplicateColumn
			},
			assertions: func(err error) {
				s.Error(err)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			column := test.preparations()

			// action.
			err := s.sut.AddColumn(column)

			// assert.
			test.assertions(err)
		})
	}
}

func (s *TableTestSuite) TestTable_AddColumns() {
	tests := []struct {
		name         string
		preparations func() []morph.Column
		assertions   func(err error)
	}{
		{
			name: "NotPreviouslyAdded",
			preparations: func() []morph.Column {
				field := "Username"
				cName := "username"
				column := morph.Column{}
				column.SetField(field)
				column.SetName(cName)
				return []morph.Column{column}
			},
			assertions: func(err error) {
				s.NoError(err)
			},
		},
		{
			name: "PreviouslyAdded",
			preparations: func() []morph.Column {
				field := "Username"
				cName := "username"
				column := morph.Column{}
				column.SetField(field)
				column.SetName(cName)

				duplicateColumn := morph.Column{}
				duplicateColumn.SetField(field)
				duplicateColumn.SetName(cName)

				columns := []morph.Column{column, duplicateColumn}
				return columns
			},
			assertions: func(err error) {
				s.Error(err)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			columns := test.preparations()

			// action.
			err := s.sut.AddColumns(columns...)

			// assert.
			test.assertions(err)
		})
	}
}

func (s *TableTestSuite) TestTable_EvaluateWithValue() {
	tests := []struct {
		name           string
		reflectOptions []morph.ReflectOption
		preparations   func() TestModel
		assertions     func(result morph.EvaluationResult, err error)
	}{
		{
			name:           "PointersDereferenced",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         "test",
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "WithTags",
			reflectOptions: []morph.ReflectOption{morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier")},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"identifier": 1,
						"name":       "test",
						"deleted_at": nil,
						"updated_at": time.Time{},
						"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "NilsPreserved",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         nil,
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(model, test.reflectOptions...)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			result, err := s.sut.Evaluate(model)

			// assert.
			test.assertions(result, err)
		})
	}
}

func (s *TableTestSuite) TestTable_EvaluateWithPointer() {
	tests := []struct {
		name           string
		reflectOptions []morph.ReflectOption
		preparations   func() TestModel
		assertions     func(result morph.EvaluationResult, err error)
	}{
		{
			name:           "PointersDereferenced",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         "test",
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "WithTags",
			reflectOptions: []morph.ReflectOption{morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier")},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"identifier": 1,
						"name":       "test",
						"deleted_at": nil,
						"updated_at": time.Time{},
						"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "NilsPreserved",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         nil,
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model, test.reflectOptions...)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			result, err := s.sut.Evaluate(&model)

			// assert.
			test.assertions(result, err)
		})
	}
}

func (s *TableTestSuite) TestTable_EvaluateMismatched() {
	tests := []struct {
		name           string
		reflectOptions []morph.ReflectOption
		preparations   func() TestModel
		assertions     func(result morph.EvaluationResult, err error)
	}{
		{
			name:           "PointersDereferenced",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         "test",
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "WithTags",
			reflectOptions: []morph.ReflectOption{morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier")},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"identifier": 1,
						"name":       "test",
						"deleted_at": nil,
						"updated_at": time.Time{},
						"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
		{
			name:           "NilsPreserved",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult, err error) {
				s.NoError(err)
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         nil,
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model, test.reflectOptions...)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			result, err := s.sut.Evaluate(model)

			// assert.
			test.assertions(result, err)
		})
	}
}

func (s *TableTestSuite) TestTable_MustEvaluateValue() {
	tests := []struct {
		name           string
		reflectOptions []morph.ReflectOption
		preparations   func() TestModel
		assertions     func(result morph.EvaluationResult)
		panics         bool
		err            error
	}{
		{
			name:   "ErrorPanics",
			panics: true,
			err:    morph.ErrMismatchingTypeName,
		},
		{
			name:           "PointersDereferenced",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult) {
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         "test",
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			if test.panics {
				s.PanicsWithError(test.err.Error(), func() { s.sut.MustEvaluate(TestModel{}) })

				return
			}

			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(model, test.reflectOptions...)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			result := s.sut.MustEvaluate(model)

			// assert.
			test.assertions(result)
		})
	}
}
func (s *TableTestSuite) TestTable_MustEvaluatePointer() {
	tests := []struct {
		name           string
		reflectOptions []morph.ReflectOption
		preparations   func() TestModel
		assertions     func(result morph.EvaluationResult)
		panics         bool
		err            error
	}{
		{
			name:   "ErrorPanics",
			panics: true,
			err:    morph.ErrMismatchingTypeName,
		},
		{
			name:           "PointersDereferenced",
			reflectOptions: []morph.ReflectOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(result morph.EvaluationResult) {
				s.Equal(
					morph.EvaluationResult{
						"id":           1,
						"name":         "test",
						"maybe_ignore": false,
						"deleted_at":   nil,
						"updated_at":   time.Time{},
						"created_at":   time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
					},
					result,
				)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			if test.panics {
				s.PanicsWithError(test.err.Error(), func() { s.sut.MustEvaluate(&TestModel{}) })
				return
			}

			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model, test.reflectOptions...)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			result := s.sut.MustEvaluate(&model)

			// assert.
			test.assertions(result)
		})
	}
}

func (s *TableTestSuite) TestTable_EvaluateErrors() {
	tests := []struct {
		name         string
		obj          any
		preparations func()
		err          error
	}{
		{
			name: "MissingTableName",
			obj:  &TestModel{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("T")
				s.sut.SetName("")
			},
			err: morph.ErrMissingTableName,
		},
		{
			name: "MissingColumns",
			obj:  &TestModel{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("T")
				s.sut.SetName("test_models")
			},
			err: morph.ErrMissingColumns,
		},
		{
			name: "MissingTableAlias",
			obj:  &TestModel{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("")
				s.sut.SetName("test_models")
			},
			err: morph.ErrMissingTableAlias,
		},
		{
			name: "MismatchingTypeName",
			obj:  struct{}{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("T")
				s.sut.SetName("test_models")

				column := morph.Column{}
				column.SetField("Name")
				column.SetName("name")
				column.SetStrategy(morph.FieldStrategyStructField)

				if err := s.sut.AddColumns(column); err != nil {
					s.FailNow("failed to prepare test: %v", err.Error())
				}
			},
			err: morph.ErrMismatchingTypeName,
		},
		{
			name: "MissingPrimaryKeys",
			obj:  &TestModel{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("T")
				s.sut.SetName("test_models")

				column := morph.Column{}
				column.SetField("Name")
				column.SetPrimaryKey(false)
				column.SetName("name")
				column.SetStrategy(morph.FieldStrategyStructField)

				if err := s.sut.AddColumns(column); err != nil {
					s.FailNow("failed to prepare test: %v", err.Error())
				}
			},
			err: morph.ErrMissingPrimaryKey,
		},
		{
			name: "MissingNonPrimaryKeys",
			obj:  &TestModel{},
			preparations: func() {
				s.sut.SetType(TestModel{})
				s.sut.SetAlias("T")
				s.sut.SetName("test_models")

				column := morph.Column{}
				column.SetField("ID")
				column.SetPrimaryKey(true)
				column.SetName("id")
				column.SetStrategy(morph.FieldStrategyStructField)

				if err := s.sut.AddColumns(column); err != nil {
					s.FailNow("failed to prepare test: %v", err.Error())
				}
			},
			err: morph.ErrMissingNonPrimaryKey,
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			test.preparations()

			// action.
			_, err := s.sut.Evaluate(test.obj)

			// assert.
			s.ErrorIs(err, test.err)
		})
	}
}

func (s *TableTestSuite) TestTable_InsertQuery_InvalidTable() {
	// action.
	query, err := s.sut.InsertQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_InsertQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES (?, ?, ?, ?, ?, ?);", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES (:created_at, :deleted_at, :id, :maybe_ignore, :name, :updated_at);", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($, $, $, $, $, $);", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($1, $2, $3, $4, $5, $6);", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, err := s.sut.InsertQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_MustInsertQuery_InvalidTable() {
	// action + assert.
	s.Panics(func() { s.sut.MustInsertQuery() })
}

func (s *TableTestSuite) TestTable_MustInsertQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES (?, ?, ?, ?, ?, ?);", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES (:created_at, :deleted_at, :id, :maybe_ignore, :name, :updated_at);", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($, $, $, $, $, $);", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($1, $2, $3, $4, $5, $6);", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query := s.sut.MustInsertQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_InsertQueryWithArgs() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(obj TestModel, query string, args []any, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES (?, ?, ?, ?, ?, ?);", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), obj.ID, *obj.Name, obj.MaybeIgnore, obj.UpdatedAt}, args)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($, $, $, $, $, $);", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), obj.ID, *obj.Name, obj.MaybeIgnore, obj.UpdatedAt}, args)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("INSERT INTO test_models (created_at, deleted_at, id, maybe_ignore, name, updated_at) VALUES ($1, $2, $3, $4, $5, $6);", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), obj.ID, *obj.Name, obj.MaybeIgnore, obj.UpdatedAt}, args)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, args, err := s.sut.InsertQueryWithArgs(&model, test.queryOptions...)

			// assert.
			test.assertions(model, query, args, err)
		})
	}
}

func (s *TableTestSuite) TestTable_InsertQueryWithArgs_InvalidTable() {
	// action.
	query, args, err := s.sut.InsertQueryWithArgs(&TestModel{})

	// assert.
	s.Error(err)
	s.Empty(query)
	s.Empty(args)
}

func (s *TableTestSuite) TestTable_UpdateQuery_InvalidTable() {
	// action.
	query, err := s.sut.UpdateQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_UpdateQuery() {
	tests := []struct {
		name         string
		queryOptions func(TestModel) []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.deleted_at = ?, T.maybe_ignore = ?, T.name = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithoutEmptyValues",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithoutEmptyValues(&m)} },
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.maybe_ignore = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", false)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $, T.deleted_at = $, T.maybe_ignore = $, T.name = $, T.updated_at = $ WHERE 1=1 AND T.id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", true)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $1, T.deleted_at = $2, T.maybe_ignore = $3, T.name = $4, T.updated_at = $5 WHERE 1=1 AND T.id = $6;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithNamedParameters()} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = :created_at, T.deleted_at = :deleted_at, T.maybe_ignore = :maybe_ignore, T.name = :name, T.updated_at = :updated_at WHERE 1=1 AND T.id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, err := s.sut.UpdateQuery(test.queryOptions(model)...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_MustUpdateQuery_InvalidTable() {
	// action + assert.
	s.Panics(func() { s.sut.MustUpdateQuery() })
}

func (s *TableTestSuite) TestTable_MustUpdateQuery() {
	tests := []struct {
		name         string
		queryOptions func(TestModel) []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.deleted_at = ?, T.maybe_ignore = ?, T.name = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithoutEmptyValues",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithoutEmptyValues(&m)} },
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.maybe_ignore = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", false)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $, T.deleted_at = $, T.maybe_ignore = $, T.name = $, T.updated_at = $ WHERE 1=1 AND T.id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", true)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $1, T.deleted_at = $2, T.maybe_ignore = $3, T.name = $4, T.updated_at = $5 WHERE 1=1 AND T.id = $6;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithNamedParameters()} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = :created_at, T.deleted_at = :deleted_at, T.maybe_ignore = :maybe_ignore, T.name = :name, T.updated_at = :updated_at WHERE 1=1 AND T.id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query := s.sut.MustUpdateQuery(test.queryOptions(model)...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_UpdateQueryWithArgs() {
	tests := []struct {
		name         string
		queryOptions func(m TestModel) []morph.QueryOption
		preparations func() TestModel
		assertions   func(obj TestModel, query string, args []any, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.deleted_at = ?, T.maybe_ignore = ?, T.name = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), obj.ID, *obj.Name, obj.MaybeIgnore, obj.UpdatedAt}, args)
			},
		},
		{
			name:         "WithoutEmptyValues",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithoutEmptyValues(&m)} },
			preparations: func() TestModel {
				return TestModel{
					ID:   1,
					Name: nil,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.maybe_ignore = ?, T.updated_at = ? WHERE 1=1 AND T.id = ?;", query)
				s.ElementsMatch([]any{obj.CreatedAt(), obj.MaybeIgnore, obj.UpdatedAt, obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", false)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $, T.deleted_at = $, T.maybe_ignore = $, T.name = $, T.updated_at = $ WHERE 1=1 AND T.id = $;", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), *obj.Name, obj.MaybeIgnore, obj.UpdatedAt, obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: func(m TestModel) []morph.QueryOption { return []morph.QueryOption{morph.WithPlaceholder("$", true)} },
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("UPDATE test_models AS T SET T.created_at = $1, T.deleted_at = $2, T.maybe_ignore = $3, T.name = $4, T.updated_at = $5 WHERE 1=1 AND T.id = $6;", query)
				s.ElementsMatch([]any{obj.CreatedAt(), any(nil), *obj.Name, obj.MaybeIgnore, obj.UpdatedAt, obj.ID}, args)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, args, err := s.sut.UpdateQueryWithArgs(model, test.queryOptions(model)...)

			// assert.
			test.assertions(model, query, args, err)
		})
	}
}

func (s *TableTestSuite) TestTable_UpdateQueryWithArgs_InvalidTable() {
	// action.
	query, args, err := s.sut.UpdateQueryWithArgs(&TestModel{})

	// assert.
	s.Error(err)
	s.Empty(query)
	s.Empty(args)
}

func (s *TableTestSuite) TestTable_DeleteQuery_InvalidTable() {
	// action.
	query, err := s.sut.DeleteQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_DeleteQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $1;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, err := s.sut.DeleteQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_MustDeleteQuery_InvalidTable() {
	// action + assert.
	s.Panics(func() { s.sut.MustDeleteQuery() })
}

func (s *TableTestSuite) TestTable_MustDeleteQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $1;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query := s.sut.MustDeleteQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_DeleteQueryWithArgs() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(obj TestModel, query string, args []any, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = ?;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $1;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, args, err := s.sut.DeleteQueryWithArgs(model, test.queryOptions...)

			// assert.
			test.assertions(model, query, args, err)
		})
	}
}

func (s *TableTestSuite) TestTable_DeleteQueryWithArgs_InvalidTable() {
	// action.
	query, args, err := s.sut.DeleteQueryWithArgs(&TestModel{})

	// assert.
	s.Error(err)
	s.Empty(query)
	s.Empty(args)
}

func (s *TableTestSuite) TestTable_SelectQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $1;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, err := s.sut.SelectQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_MustSelectQuery_InvalidTable() {
	// action + assert.
	s.Panics(func() { s.sut.MustSelectQuery() })
}

func (s *TableTestSuite) TestTable_MustSelectQuery() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(query string, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = ?;", query)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $;", query)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $1;", query)
			},
		},
		{
			name:         "WithNamedParameters",
			queryOptions: []morph.QueryOption{morph.WithNamedParameters()},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(query string, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = :id;", query)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query := s.sut.MustSelectQuery(test.queryOptions...)

			// assert.
			test.assertions(query, err)
		})
	}
}

func (s *TableTestSuite) TestTable_SelectQueryWithArgs() {
	tests := []struct {
		name         string
		queryOptions []morph.QueryOption
		preparations func() TestModel
		assertions   func(obj TestModel, query string, args []any, err error)
	}{
		{
			name:         "NoOptions",
			queryOptions: []morph.QueryOption{},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = ?;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_NoOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", false)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
		{
			name:         "WithPlaceholder_WithOrdering",
			queryOptions: []morph.QueryOption{morph.WithPlaceholder("$", true)},
			preparations: func() TestModel {
				name := "test"
				return TestModel{
					ID:   1,
					Name: &name,
					Another: AnotherTestModel{
						ID:          2,
						Title:       "another",
						Description: nil,
					},
				}
			},
			assertions: func(obj TestModel, query string, args []any, err error) {
				s.Require().NoError(err)
				s.Equal("SELECT T.created_at, T.deleted_at, T.id, T.maybe_ignore, T.name, T.updated_at FROM test_models AS T WHERE 1=1 AND T.id = $1;", query)
				s.ElementsMatch([]any{obj.ID}, args)
			},
		},
	}

	for _, test := range tests {
		s.Run(test.name, func() {
			// arrange.
			model := test.preparations()

			var err error
			s.sut, err = morph.Reflect(&model)
			if err != nil {
				s.FailNow("unable to reflect in test", err)
			}

			// action.
			query, args, err := s.sut.SelectQueryWithArgs(model, test.queryOptions...)

			// assert.
			test.assertions(model, query, args, err)
		})
	}
}

func (s *TableTestSuite) TestTable_SelectQueryWithArgs_InvalidTable() {
	// action.
	query, args, err := s.sut.SelectQueryWithArgs(&TestModel{})

	// assert.
	s.Error(err)
	s.Empty(query)
	s.Empty(args)
}

func (s *TableTestSuite) TestTable_EvaluationResults_Empties() {
	// arrange.
	m := AnotherTestModel{
		ID:          2,
		Title:       "another",
		Description: nil,
	}

	var err error
	s.sut, err = morph.Reflect(m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Len(result.Empties(), 1)
	s.Equal("description", result.Empties()[0])
}

func (s *TableTestSuite) TestTable_EvaluationResults_NonEmpties() {
	// arrange.
	m := AnotherTestModel{
		ID:          2,
		Title:       "another",
		Description: nil,
	}

	var err error
	s.sut, err = morph.Reflect(m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Len(result.NonEmpties(), 2)
	s.ElementsMatch(result.NonEmpties(), []string{"id", "title"})
}
