package morph_test

import (
	"testing"

	"github.com/freerware/morph"
	"github.com/stretchr/testify/suite"
)

type ColumnTestSuite struct {
	suite.Suite

	sut morph.Column
}

func TestColumnTestSuite(t *testing.T) {
	suite.Run(t, new(ColumnTestSuite))
}

func (s *ColumnTestSuite) SetupTest() {
	s.sut = morph.Column{}
}

func (s *ColumnTestSuite) TestColumn_Name() {
	// arrange.
	expectedName := "username"
	s.sut.SetName(expectedName)

	// action.
	actualName := s.sut.Name()

	// assert.
	s.Equal(expectedName, actualName)
}

func (s *ColumnTestSuite) TestColumn_SetName() {
	// arrange.
	expectedName := "username"

	// action.
	s.sut.SetName(expectedName)

	// assert.
	s.Equal(expectedName, s.sut.Name())
}

func (s *ColumnTestSuite) TestColumn_Field() {
	// arrange.
	expectedField := "Username"
	s.sut.SetField(expectedField)

	// action.
	actualField := s.sut.Field()

	// assert.
	s.Equal(expectedField, actualField)
}

func (s *ColumnTestSuite) TestColumn_SetField() {
	// arrange.
	expectedField := "Username"

	// action.
	s.sut.SetField(expectedField)

	// assert.
	s.Equal(expectedField, s.sut.Field())
}
