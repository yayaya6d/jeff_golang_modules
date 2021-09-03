package set

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type SetTestSuite struct {
	suite.Suite
	sut Set
}

func TestSetTestSuite(t *testing.T) {
	suite.Run(t, new(SetTestSuite))
}

func (s *SetTestSuite) SetupTest() {
	s.sut = NewSet()
}

func (s *SetTestSuite) TestInsertAndExist_InsertValue_CheckValueExistCorrect() {
	s.sut.Insert(1)
	s.sut.Insert(true)

	s.True(s.sut.Exist(1))
	s.True(s.sut.Exist(true))
	s.False(s.sut.Exist(2))
	s.False(s.sut.Exist(false))
}

func (s *SetTestSuite) TestSize_InsertTwoValue_GetSizeIsTwo() {
	s.sut.Insert(1)
	s.sut.Insert(true)

	s.Equal(2, s.sut.Size())
}

func (s *SetTestSuite) TestDelete_InsertValueAndDeleteIt_GetValueDoesNotExist() {
	s.sut.Insert(1)
	s.sut.Delete(1)

	s.False(s.sut.Exist(1))
}

func (s *SetTestSuite) TestDelete_DeleteNonExistValue_NothingHappened() {
	s.sut.Delete(1)
}

func (s *SetTestSuite) TestValues_InsertSomeValues_GetAllValueCorrect() {
	expectedSlice := make([]interface{}, 0)

	for i := 0; i < 5; i++ {
		expectedSlice = append(expectedSlice, i)
		s.sut.Insert(i)
	}

	actualSlice := make([]interface{}, 0)
	for i := range *s.sut.Values() {
		actualSlice = append(actualSlice, i)
	}

	s.ElementsMatch(expectedSlice, actualSlice)
}
