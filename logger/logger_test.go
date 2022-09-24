package logger

import (
	"bytes"
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/suite"
)

type LoggerTestSuite struct {
	suite.Suite
}

type logOutput struct {
	Level string
	Msg   string
}

func TestLoggerTestSuite(t *testing.T) {
	suite.Run(t, new(LoggerTestSuite))
}

func (s *LoggerTestSuite) TestNewLogger_GivenInvalidLevel_UseInfoAsDefault() {
	// arrange
	l := NewLogger("dont care")

	// act
	level := l.GetLevel().String()

	// assert
	s.Equal("info", level)
}

func (s *LoggerTestSuite) TestSetLoggerLevel_SetLoggerLevelToError_SetLevelCorrectly() {
	// arrange
	Log.SetLoggerLevel("error")

	// act
	level := Log.GetLevel().String()

	// assert
	s.Equal("error", level)
}

func (s *LoggerTestSuite) TestSetLoggerLevel_GivenInvalidLevel_ReturnError() {
	// act
	err := Log.SetLoggerLevel("dont care")

	// assert
	s.Error(err)
}

func (s *LoggerTestSuite) TestInfo_GivenMsg_PrintLogCorrectly() {
	// arrange
	var buf bytes.Buffer
	Log.SetOutput(&buf)
	expecteOutput := logOutput{
		"info", "test msg",
	}

	// act
	Log.Info("test msg")

	// assert
	var actualOutput logOutput
	err := json.Unmarshal(buf.Bytes(), &actualOutput)

	s.NoError(err)
	s.Equal(expecteOutput, actualOutput)
}

func (s *LoggerTestSuite) TestWithField_GivenField_PrintLogCorrectly() {
	// arrange
	var buf bytes.Buffer
	Log.SetOutput(&buf)

	type logEithField struct {
		logOutput
		Key string
	}

	expectedOutput := logEithField{
		logOutput: logOutput{
			Level: "info",
			Msg:   "test msg",
		},
		Key: "value",
	}

	// act
	Log.WithField("key", "value").Info("test msg")

	// assert
	var actualOutput logEithField
	err := json.Unmarshal(buf.Bytes(), &actualOutput)

	s.NoError(err)
	s.Equal(expectedOutput, actualOutput)
}

func (s *LoggerTestSuite) TestWithField_GivenFields_PrintLogCorrectly() {
	// arrange
	var buf bytes.Buffer
	Log.SetOutput(&buf)

	type logEithFields struct {
		logOutput
		Int1 int
		Key1 string
		Key2 string
	}

	expectedOutput := logEithFields{
		logOutput: logOutput{
			Level: "info",
			Msg:   "test msg",
		},
		Int1: 123,
		Key1: "value1",
		Key2: "value2",
	}

	inputFields := Fields{
		"Int1": 123,
		"key1": "value1",
		"key2": "value2",
	}

	// act
	Log.WithFields(inputFields).Info("test msg")

	// assert
	var actualOutput logEithFields
	err := json.Unmarshal(buf.Bytes(), &actualOutput)

	s.NoError(err)
	s.Equal(expectedOutput, actualOutput)
}
