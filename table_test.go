package morph_test

import (
	"testing"

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
