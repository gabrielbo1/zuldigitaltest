package config

import (
	"flag"
	"os"
)

// EnvironmentVariable - Type to environment variables.
type EnvironmentVariable string

const (
	//ApiVersion - Define API version.
	ApiVersion EnvironmentVariable = "API_VERSION"
)

type configVar struct {
	name  EnvironmentVariable
	value string
	usage string
}

var configVars []configVar = []configVar{
	{name: ApiVersion, value: "1.1", usage: "Define api version"},
}

func setVar(envVar EnvironmentVariable, value string) {
	for i := range configVars {
		if configVars[i].name == envVar {
			configVars[i].value = value
		}
	}
}

func getVar(envVar EnvironmentVariable) *configVar {
	for i := range configVars {
		if configVars[i].name == envVar {
			return &configVars[i]
		}
	}
	return nil
}

//FlagParse - Flags parsing and set values.
func FlagParse() {
	var values []*string
	if !flag.Parsed() {
		for i := range configVars {
			values = append(values, flag.String(string(configVars[i].name), configVars[i].value, configVars[i].usage))
		}
		flag.Parse()
	}
	for i := range configVars {
		configVars[i].value = *values[i]
	}
}

// GetValue - Find to environment variable value
// or return default value of variable.
func GetValue(variable EnvironmentVariable) string {
	if value := os.Getenv(string(variable)); value != "" {
		setVar(variable, value)
	}
	if conf := getVar(variable); conf != nil {
		return conf.value
	}
	return ""
}
