package mongodb

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type mongodbTestL1 struct {
	suite.Suite
}

func TestMongodbSuite(t *testing.T) {
	suite.Run(t, new(mongodbTestL1))
}

func (s *mongodbTestL1) TestInit_GivenInvalidUri_ReturnError() {
	// act
	err := DB.Init("invalid uri")

	// assert
	s.Error(err)
}

func (s *mongodbTestL1) TestInit_GivenValidUri_ReturnNoError() {
	// act
	err := DB.Init("mongodb://localhost:27017")
	defer DB.Release()

	// assert
	s.Nil(err)
}

func (s *mongodbTestL1) TestInit_InitTwoTimes_ReturnNoError() {
	// act
	err := DB.Init("mongodb://localhost:27017")
	err = DB.Init("")
	defer DB.Release()

	// assert
	s.Nil(err)
}

func (s *mongodbTestL1) TestCollection_GetCollectionBeforeInit_Panic() {
	// act ans assert
	s.Panics(func() {
		DB.Collection("collection")
	})
}

func (s *mongodbTestL1) TestCollection_GetCollectionAfterInit_ReturnExpectedCollection() {
	// arrange
	collectionName := "test"
	_ = DB.Init("mongodb://localhost:27017")
	defer DB.Release()

	// act
	c := DB.Collection(collectionName)

	// assert
	s.Equal(collectionName, c.Name())
}

func (s *mongodbTestL1) TestCollection_GetExistedCollection_ReturnExpectedCollection() {
	// arrange
	collectionName := "test"
	_ = DB.Init("mongodb://localhost:27017")
	defer DB.Release()

	// act
	_ = DB.Collection(collectionName)
	c := DB.Collection(collectionName)

	// assert
	s.Equal(collectionName, c.Name())
}
