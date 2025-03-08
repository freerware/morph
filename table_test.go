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
	// arrange.
	field := "Username"
	cName := "username"
	column := morph.Column{}
	column.SetField(field)
	column.SetName(cName)

	// action.
	err := s.sut.AddColumn(column)

	// assert.
	s.NoError(err)
}

func (s *TableTestSuite) TestTable_AddColumn_AlreadyExists() {
	// arrange.
	field := "Username"
	cName := "username"
	column := morph.Column{}
	column.SetField(field)
	column.SetName(cName)
	duplicateColumn := morph.Column{}
	duplicateColumn.SetField(field)
	duplicateColumn.SetName(cName)
	expectedColumns := []morph.Column{column}
	s.Require().NoError(s.sut.AddColumns(expectedColumns...))

	// action.
	err := s.sut.AddColumn(duplicateColumn)

	// assert.
	s.Error(err)
}

func (s *TableTestSuite) TestTable_AddColumns() {
	// arrange.
	field := "Username"
	cName := "username"
	column := morph.Column{}
	column.SetField(field)
	column.SetName(cName)
	expectedColumns := []morph.Column{column}

	// action.
	err := s.sut.AddColumns(expectedColumns...)

	// assert.
	s.Require().NoError(err)
}

func (s *TableTestSuite) TestTable_AddColumns_AlreadyExists() {
	// arrange.
	field := "Username"
	cName := "username"
	column := morph.Column{}
	column.SetField(field)
	column.SetName(cName)
	duplicateColumn := morph.Column{}
	duplicateColumn.SetField(field)
	duplicateColumn.SetName(cName)
	expectedColumns := []morph.Column{column, duplicateColumn}

	// action.
	err := s.sut.AddColumns(expectedColumns...)

	// assert.
	s.Error(err)
}

func (s *TableTestSuite) TestTable_EvaluateValue_PointersDereferenced() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
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
	s.Equal(
		morph.EvaluationResult{
			"id":         1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_MustEvaluateValue() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result := s.sut.MustEvaluate(m)

	// assert.
	s.Equal(
		morph.EvaluationResult{
			"id":         1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_MustEvaluateValue_ErrorPanics() {
	// action + assert.
	s.PanicsWithError(morph.ErrMismatchingTypeName.Error(), func() { s.sut.MustEvaluate(TestModel{}) })
}

func (s *TableTestSuite) TestTable_EvaluateValue_WithTags() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(m, morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier"))
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Equal(
		morph.EvaluationResult{
			"identifier": 1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_EvaluateValue_NilsPreserved() {
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
	s.Equal(morph.EvaluationResult{"id": 2, "title": "another", "description": nil}, result)
}

func (s *TableTestSuite) TestTable_EvaluatePointer_PointersDereferenced() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(&m)

	// assert.
	s.NoError(err)
	s.Equal(
		morph.EvaluationResult{
			"id":         1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_MustEvaluatePointer_PointersDereferenced() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result := s.sut.MustEvaluate(&m)

	// assert.
	s.Equal(
		morph.EvaluationResult{
			"id":         1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_MustEvaluatePointer_ErrorPanics() {
	// action + assert.
	s.PanicsWithError(morph.ErrMismatchingTypeName.Error(), func() { s.sut.MustEvaluate(&TestModel{}) })
}

func (s *TableTestSuite) TestTable_EvaluatePointer_WithTags() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m, morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier"))
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(&m)

	// assert.
	s.NoError(err)
	s.Equal(
		morph.EvaluationResult{
			"identifier": 1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_EvaluatePointer_NilsPreserved() {
	// arrange.
	m := AnotherTestModel{
		ID:          2,
		Title:       "another",
		Description: nil,
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(&m)

	// assert.
	s.NoError(err)
	s.Equal(morph.EvaluationResult{"id": 2, "title": "another", "description": nil}, result)
}

func (s *TableTestSuite) TestTable_EvaluateMismatched_PointersDereferenced() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Equal(
		morph.EvaluationResult{
			"id":         1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_EvaluateMismatched_WithTags() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m, morph.WithTag("db"), morph.WithPrimaryKeyColumn("identifier"))
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Equal(
		morph.EvaluationResult{
			"identifier": 1,
			"name":       name,
			"created_at": time.Date(2024, time.February, 28, 10, 30, 0, 0, time.Local),
		},
		result,
	)
}

func (s *TableTestSuite) TestTable_EvaluateMismatched_NilsPreserved() {
	// arrange.
	m := AnotherTestModel{
		ID:          2,
		Title:       "another",
		Description: nil,
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	result, err := s.sut.Evaluate(m)

	// assert.
	s.NoError(err)
	s.Equal(morph.EvaluationResult{"id": 2, "title": "another", "description": nil}, result)
}

func (s *TableTestSuite) TestTable_Evaluate_MissingTableName() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("T")
	s.sut.SetName("")

	// action.
	_, err := s.sut.Evaluate(TestModel{})

	// assert.
	s.ErrorIs(err, morph.ErrMissingTableName)
}

func (s *TableTestSuite) TestTable_Evaluate_MissingTableAlias() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("")
	s.sut.SetName("test_models")

	// action.
	_, err := s.sut.Evaluate(TestModel{})

	// assert.
	s.ErrorIs(err, morph.ErrMissingTableAlias)
}

func (s *TableTestSuite) TestTable_Evaluate_MissingColumns() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("T")
	s.sut.SetName("test_models")

	// action.
	_, err := s.sut.Evaluate(TestModel{})

	// assert.
	s.ErrorIs(err, morph.ErrMissingColumns)
}

func (s *TableTestSuite) TestTable_Evaluate_MismatchingTypeName() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("T")
	s.sut.SetName("test_models")

	column := morph.Column{}
	column.SetField("Name")
	column.SetName("name")
	column.SetStrategy(morph.FieldStrategyStructField)

	s.sut.AddColumns(column)

	// action.
	_, err := s.sut.Evaluate(struct{}{})

	// assert.
	s.ErrorIs(err, morph.ErrMismatchingTypeName)
}

func (s *TableTestSuite) TestTable_Evaluate_MissingPrimaryKeys() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("T")
	s.sut.SetName("test_models")

	column := morph.Column{}
	column.SetField("Name")
	column.SetPrimaryKey(false)
	column.SetName("name")
	column.SetStrategy(morph.FieldStrategyStructField)

	s.sut.AddColumns(column)

	// action.
	_, err := s.sut.Evaluate(TestModel{})

	// assert.
	s.ErrorIs(err, morph.ErrMissingPrimaryKey)
}

func (s *TableTestSuite) TestTable_Evaluate_MissingNonPrimaryKeys() {
	// arrange.
	s.sut.SetType(TestModel{})
	s.sut.SetAlias("T")
	s.sut.SetName("test_models")

	column := morph.Column{}
	column.SetField("ID")
	column.SetPrimaryKey(true)
	column.SetName("id")
	column.SetStrategy(morph.FieldStrategyStructField)

	s.sut.AddColumns(column)

	// action.
	_, err := s.sut.Evaluate(TestModel{})

	// assert.
	s.ErrorIs(err, morph.ErrMissingNonPrimaryKey)
}

func (s *TableTestSuite) TestTable_InsertQuery() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.InsertQuery()

	// assert.
	s.NoError(err)
	s.Equal("INSERT INTO test_models (created_at, id, name) VALUES (?, ?, ?);", query)
}

func (s *TableTestSuite) TestTable_InsertQuery_InvalidTable() {
	// action.
	query, err := s.sut.InsertQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_InsertQuery_WithNamedParameters() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.InsertQuery(morph.WithNamedParameters())

	// assert.
	s.NoError(err)
	s.Equal("INSERT INTO test_models (created_at, id, name) VALUES (:created_at, :id, :name);", query)
}

func (s *TableTestSuite) TestTable_InsertQuery_WithPlaceholder_NoOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.InsertQuery(morph.WithPlaceholder("$", false))

	// assert.
	s.NoError(err)
	s.Equal("INSERT INTO test_models (created_at, id, name) VALUES ($, $, $);", query)
}

func (s *TableTestSuite) TestTable_InsertQuery_WithPlaceholder_WithOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.InsertQuery(morph.WithPlaceholder("$", true))

	// assert.
	s.NoError(err)
	s.Equal("INSERT INTO test_models (created_at, id, name) VALUES ($1, $2, $3);", query)
}

func (s *TableTestSuite) TestTable_UpdateQuery() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.UpdateQuery()

	// assert.
	s.NoError(err)
	s.Equal("UPDATE test_models AS T SET T.created_at = ?, T.name = ? WHERE 1=1 AND T.id = ?;", query)
}

func (s *TableTestSuite) TestTable_UpdateQuery_InvalidTable() {
	// action.
	query, err := s.sut.UpdateQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_UpdateQuery_WithoutEmptyValues() {
	// arrange.
	m := TestModel{
		ID:   1,
		Name: nil,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.UpdateQuery(morph.WithoutEmptyValues(&m))

	// assert.
	s.NoError(err)
	s.Equal("UPDATE test_models AS T SET T.created_at = ? WHERE 1=1 AND T.id = ?;", query)
}

func (s *TableTestSuite) TestTable_UpdateQuery_WithPlaceholder_NoOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.UpdateQuery(morph.WithPlaceholder("$", false))

	// assert.
	s.NoError(err)
	s.Equal("UPDATE test_models AS T SET T.created_at = $, T.name = $ WHERE 1=1 AND T.id = $;", query)
}

func (s *TableTestSuite) TestTable_UpdateQuery_WithPlaceholder_WithOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.UpdateQuery(morph.WithPlaceholder("$", true))

	// assert.
	s.NoError(err)
	s.Equal("UPDATE test_models AS T SET T.created_at = $1, T.name = $2 WHERE 1=1 AND T.id = $3;", query)
}

func (s *TableTestSuite) TestTable_UpdateQuery_WithNamedParameters() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.UpdateQuery(morph.WithNamedParameters())

	// assert.
	s.NoError(err)
	s.Equal("UPDATE test_models AS T SET T.created_at = :created_at, T.name = :name WHERE 1=1 AND T.id = :id;", query)
}

func (s *TableTestSuite) TestTable_DeleteQuery() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.DeleteQuery()

	// assert.
	s.NoError(err)
	s.Equal("DELETE FROM test_models WHERE 1=1 AND id = ?;", query)
}

func (s *TableTestSuite) TestTable_DeleteQuery_InvalidTable() {
	// action.
	query, err := s.sut.DeleteQuery()

	// assert.
	s.Error(err)
	s.Empty(query)
}

func (s *TableTestSuite) TestTable_DeleteQuery_WithPlaceholder_NoOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.DeleteQuery(morph.WithPlaceholder("$", false))

	// assert.
	s.NoError(err)
	s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $;", query)
}

func (s *TableTestSuite) TestTable_DeleteQuery_WithPlaceholder_WithOrdering() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.DeleteQuery(morph.WithPlaceholder("$", true))

	// assert.
	s.NoError(err)
	s.Equal("DELETE FROM test_models WHERE 1=1 AND id = $1;", query)
}

func (s *TableTestSuite) TestTable_DeleteQuery_WithNamedParameters() {
	// arrange.
	name := "test"
	m := TestModel{
		ID:   1,
		Name: &name,
		Another: AnotherTestModel{
			ID:          2,
			Title:       "another",
			Description: nil,
		},
	}

	var err error
	s.sut, err = morph.Reflect(&m)
	if err != nil {
		s.FailNow("unable to reflect in test", err)
	}

	// action.
	query, err := s.sut.DeleteQuery(morph.WithNamedParameters())

	// assert.
	s.NoError(err)
	s.Equal("DELETE FROM test_models WHERE 1=1 AND id = :id;", query)
}
