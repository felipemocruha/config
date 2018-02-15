package config

import (
	"testing"
)

func TestIsStructTrue(t *testing.T) {
	type Test struct{}

	if !isStruct(Test{}) {
		t.Errorf("A struct is expected as input")
	}
}

func TestIsStructFalse(t *testing.T) {
	if isStruct(1) {
		t.Errorf("The output must be false")
	}
}

func TestGetenv(t *testing.T) {
	expected := "test"
	env := map[string]string{"TEST": expected}
	result := getenv(env, "TEST", "fallback")

	if result != expected {
		t.Errorf("Unexpected output, want: (%v), got: (%v)", expected, result)
	}
}

func TestGetenvFallback(t *testing.T) {
	expected := "test"
	env := map[string]string{}
	result := getenv(env, "TEST", expected)

	if result != expected {
		t.Errorf("Unexpected output, want: (%v), got: (%v)", expected, result)
	}
}

func TestStrToPair(t *testing.T) {
	key := "TEST"
	val := "test"
	input := key + "=" + val

	k, v := strToPair(input)
	if k != key {
		t.Errorf("Unexpected key, want: (%v), got: (%v)", key, k)
	}
	if v != val {
		t.Errorf("Unexpected val, want: (%v), got: (%v)", val, v)
	}
}

func TestStrToPairUnescapedEq(t *testing.T) {
	key := "TEST"
	val := "tes=t"
	input := key + "=" + val

	k, v := strToPair(input)
	if k != key {
		t.Errorf("Unexpected key, want: (%v), got: (%v)", key, k)
	}
	if v != val {
		t.Errorf("Unexpected val, want: (%v), got: (%v)", val, v)
	}
}

func TestFileToBytesNotFound(t *testing.T) {
	_, err := fileToBytes("you will never find this file")
	if err == nil {
		t.Errorf("It must raise an error")
	}
}

func TestEnvToConfig(t *testing.T) {
	type ServerConfig struct {
		Host     string `cfg:"HOST,localhost"`
		Port     string `cfg:"PORT,:8080"`
		Password string `cfg:"PASSWORD"`
	}
	s := ServerConfig{}
	envToConfig(map[string]string{}, &s)

	if s.Host != "localhost" {
		t.Errorf("Unexpected value, want: (%v), got: (%v)", "localhost", s.Host)
	}
	
	if s.Port != ":8080" {
		t.Errorf("Unexpected value, want: (%v), got: (%v)", ":8080", s.Port)
	}

	if s.Password != "" {
		t.Errorf("Unexpected value, want: (%v), got: (%v)", "", s.Password)
	}
}

func TestEnvToConfigNotAStruct(t *testing.T) {
	v := 1
	err := envToConfig(map[string]string{}, v) 
	if err == nil {
		t.Errorf("It must raise an error")
	}
}
