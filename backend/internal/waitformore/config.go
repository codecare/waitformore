package waitformore

import (
	"fmt"
	"os"
	"strconv"
)

func ReadFromEnv(key, defaultValue string, secure bool) string {
	value := os.Getenv(key)
	isDefault := false

	if len(value) < 1 {
		value = defaultValue
		isDefault = true
	}

	if secure {
		fmt.Printf("%v=%v isDefault: %v\n", key, "***", isDefault)
	} else {
		fmt.Printf("%v=%v isDefault: %v\n", key, value, isDefault)
	}
	return value
}

func ReadBoolFromEnv(key string, defaultValue bool) bool {
	strValue := os.Getenv(key)
	isDefault := false

	value, err := strconv.ParseBool(strValue)
	if err != nil {
		fmt.Printf("%v: %v cannot be parsed\n", key, strValue)
		isDefault = true
		value = defaultValue
	}

	fmt.Printf("%v=%v isDefault: %v\n", key, value, isDefault)

	return value
}