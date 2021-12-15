package env

import (
	"os"
	"testing"

	"github.com/stretchr/testify/suite"
)

type envTestL1 struct {
	suite.Suite
}

func (s *envTestL1) TestMustGet_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	expectedValue := "fake_value"
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key)

	// act
	actualValue := MustGet(key)

	// assert
	s.Equal(expectedValue, actualValue)
}

func (s *envTestL1) TestMustGet_GivenNonExistedEnvKey_Panic() {
	// arrange
	key := "fake_key"

	// assert
	s.Panics(func() {
		MustGet(key)
	})
}

func (s *envTestL1) TestGetOrDefault_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	expectedValue := "fake_value"
	os.Setenv(key, expectedValue)
	defer os.Unsetenv(key)

	// act
	actualValue := GetOrDefault(key, "")

	// assert
	s.Equal(expectedValue, actualValue)
}

func (s *envTestL1) TestGetOrDefault_GivenNonExistedEnvKey_ReturnDefaultVal() {
	// arrange
	key := "fake_key"
	defaultVal := "defaultVal"
	// act
	actualValue := GetOrDefault(key, defaultVal)

	// assert
	s.Equal(defaultVal, actualValue)
}

func (s *envTestL1) TestMustGetInt_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	val := "123"
	expectedInt := 123
	os.Setenv(key, val)
	defer os.Unsetenv(key)

	// act
	actualValue := MustGetInt(key)

	// assert
	s.Equal(expectedInt, actualValue)
}

func (s *envTestL1) TestMustGetInt_GivenNonExistedEnvKey_Panic() {
	// arrange
	key := "fake_key"

	// assert
	s.Panics(func() {
		MustGetInt(key)
	})
}

func (s *envTestL1) TestGetIntOrDefault_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	val := "123"
	expectedInt := 123
	os.Setenv(key, val)
	defer os.Unsetenv(key)

	// act
	actualValue := GetIntOrDefault(key, 0)

	// assert
	s.Equal(expectedInt, actualValue)
}

func (s *envTestL1) TestGetIntOrDefault_GivenNonExistedEnvKey_ReturnDefaultVal() {
	// arrange
	key := "fake_key"
	defaultVal := 123

	// act
	actualValue := GetIntOrDefault(key, defaultVal)

	// assert
	s.Equal(defaultVal, actualValue)
}

func (s *envTestL1) TestMustGetBool_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	val := "true"
	expectedBool := true
	os.Setenv(key, val)
	defer os.Unsetenv(key)

	// act
	actualValue := MustGetBool(key)

	// assert
	s.Equal(expectedBool, actualValue)
}

func (s *envTestL1) TestMustGetBool_GivenNonExistedEnvKey_Panic() {
	// arrange
	key := "fake_key"

	// assert
	s.Panics(func() {
		MustGetBool(key)
	})
}

func (s *envTestL1) TestGetBoolOrDefault_GivenExistedEnvKey_ReturnExpectedVal() {
	// arrange
	key := "fake_key"
	val := "true"
	expectedBool := true
	os.Setenv(key, val)
	defer os.Unsetenv(key)

	// act
	actualValue := GetBoolOrDefault(key, false)

	// assert
	s.Equal(expectedBool, actualValue)
}

func (s *envTestL1) TestGetBoolOrDefault_GivenNonExistedEnvKey_ReturnDefaultVal() {
	// arrange
	key := "fake_key"
	defaultVal := true

	// act
	actualValue := GetBoolOrDefault(key, defaultVal)

	// assert
	s.Equal(defaultVal, actualValue)
}

func TestEnvSuite(t *testing.T) {
	suite.Run(t, new(envTestL1))
}
