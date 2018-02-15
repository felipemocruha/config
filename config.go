package config

import (
	"fmt"
	"io/ioutil"
	"os"
	"reflect"
	"strings"
)

func isStruct(conf interface{}) bool {
	if reflect.ValueOf(conf).Kind() != reflect.Struct {
		return false
	}
	return true
}

func strToPair(envVar string) (string, string) {
	keyValue := strings.SplitN(envVar, "=", 2)
	return keyValue[0], keyValue[1]
}

func getenv(env map[string]string, key, fallback string) string {
	value := env[key]
	if len(value) == 0 {
		return fallback
	}
	return value
}

func fileToBytes(path string) ([]byte, error) {
	raw, err := ioutil.ReadFile(path)
	if err != nil || len(raw) == 0 {
		return []byte{}, fmt.Errorf("Unable to read file")
	}

	return raw, nil
}

func GetEnvironment() map[string]string {
	values := os.Environ()
	environment := make(map[string]string)

	for _, s := range values {
		k, v := strToPair(s)
		environment[k] = v
	}

	return environment
}

func GetConfigFromEnv(conf interface{}) error {
	
	return nil
}

func GetConfigFromEnvFile(conf interface{}) error {

	return nil
}

func GetConfigFromJson(conf interface{}) error {

	return nil
}

func envToConfig(env map[string]string, conf interface{}) error {
	// if !isStruct(conf) {
	// 	return fmt.Errorf("config object must be a struct")
	// }

	varType := reflect.TypeOf(conf).Elem()
	for i := 0; i < varType.NumField(); i++ {
		field := varType.Field(i)
		tag := parseTag(field.Tag.Get("cfg"))

		fallback := ""
		if len(tag) >= 2 {
			fallback = tag[1]
		}
		val := getenv(env, tag[0], fallback)

		reflect.ValueOf(conf).Elem().Field(i).SetString(val)
	}

	return nil
}

func parseTag(tag string) []string {
	parsed := strings.SplitN(tag, ",", -1)
	if len(tag) >= 2 {
		return parsed
	}
	return []string{tag}
}
