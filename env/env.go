package env

import (
	"fmt"
	"os"
	"strconv"
)

func MustGet(key string) string {
	val, ok := os.LookupEnv(key)
	if !ok {
		panic(fmt.Sprintf("Could nmot find this env variable, key = %s", val))
	}
	return val
}

func GetOrDefault(key, defaultVal string) string {
	val, ok := os.LookupEnv(key)
	if ok {
		return val
	}
	return defaultVal
}

func MustGetInt(key string) int {
	val := MustGet(key)
	i, err := strconv.Atoi(val)
	if err != nil {
		panic(fmt.Sprintf("Could not convert this val to int, key = %s, val = %s", key, val))
	}
	return i
}

func GetIntOrDefault(key string, defaultVal int) int {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	i, err := strconv.Atoi(val)
	if err == nil {
		return i
	}
	return defaultVal
}

func MustGetBool(key string) bool {
	val := MustGet(key)
	b, err := strconv.ParseBool(val)
	if err != nil {
		panic(fmt.Sprintf("Could not convert this val to bool, key = %s, val = %s", key, val))
	}
	return b
}

func GetBoolOrDefault(key string, defaultVal bool) bool {
	val, ok := os.LookupEnv(key)
	if !ok {
		return defaultVal
	}
	b, err := strconv.ParseBool(val)
	if err == nil {
		return b
	}
	return defaultVal
}
