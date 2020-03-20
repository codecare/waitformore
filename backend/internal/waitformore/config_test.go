package waitformore

import (
	"os"
	"testing"
)

func Test_config01(t *testing.T) {

	err := os.Setenv("huhu", "")
	if err != nil {
		t.Error(err)
	}
	boolFromEnv := ReadBoolFromEnv("huhu", false)
	if boolFromEnv {
		t.Errorf("failed")
	}
}

func Test_config02(t *testing.T) {

	err := os.Setenv("huhu", "False")
	if err != nil {
		t.Error(err)
	}
	boolFromEnv := ReadBoolFromEnv("huhu", true)
	if boolFromEnv {
		t.Errorf("failed")
	}
}

func Test_config03(t *testing.T) {

	err := os.Setenv("huhu", "false")
	if err != nil {
		t.Error(err)
	}
	boolFromEnv := ReadBoolFromEnv("huhu", true)
	if boolFromEnv {
		t.Errorf("failed")
	}
}

func Test_config04(t *testing.T) {

	err := os.Setenv("huhu", "True")
	if err != nil {
		t.Error(err)
	}
	boolFromEnv := ReadBoolFromEnv("huhu", false)
	if !boolFromEnv {
		t.Errorf("failed")
	}
}

func Test_config05(t *testing.T) {

	err := os.Setenv("huhu", "true")
	if err != nil {
		t.Error(err)
	}
	boolFromEnv := ReadBoolFromEnv("huhu", false)
	if !boolFromEnv {
		t.Errorf("failed")
	}
}