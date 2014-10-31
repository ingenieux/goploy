package main

import (
	"errors"
	"github.com/docopt/docopt-go"
	"strings"
	"testing"
)

var invalidCli = errors.New("Invalid CLI Invocation")

func invokeCli(t *testing.T, cmd string) (map[string]interface{}, error) {
	args := strings.Split(cmd, " ")[1:]

	t.Logf("Calling with args: %s\n", args)

	result, err := docopt.Parse(usage, args, true, "", false, false)

	t.Logf("(result, err): %q, %q\n", result, err)

	if 0 == len(result) {
		return nil, invalidCli
	}

	return result, err
}

func TestMissingApplication(t *testing.T) {
	_, err := invokeCli(t, "goploy push")

	if nil == err {
		t.Error("Should have triggered an error.")
	} else {
		t.Log(err)
	}
}

func TestPositionalArgumentWithDefaults(t *testing.T) {
	values, err := invokeCli(t, "goploy push myapp")

	if nil != err {
		t.Errorf("Unexpected error %q\n", err)
	}

	if app, ok := values["APPLICATION"]; !ok {
		t.Error("Value of Application wasn't found")
	} else {
		t.Logf("Using application %s\n", app)
		t.Logf("Using environment %s\n", values["ENVIRONMENT"])
	}
}

func TestWithAppAndEnvironment(t *testing.T) {
	values, err := invokeCli(t, "goploy push myapp myenv")

	if nil != err {
		t.Errorf("Unexpected error %q\n", err)
	}

	if env, ok := values["ENVIRONMENT"]; !ok {
		t.Error("Value of Environment wasn't found")
	} else {
		app := values["APPLICATION"]
		t.Logf("Using application %s\n", app)
		t.Logf("Using environment %s\n", env)
	}
}

func TestWithOptionalArgument(t *testing.T) {
	lines := []string{
		"goploy -b other push myapp myenv",
		"goploy --branch other push myapp myenv",
	}

	for _, v := range lines {
		values, err := invokeCli(t, v)

		if nil != err {
			t.Errorf("Unexpected error %q\n", err)
		}

		v, ok := values["--branch"]

		if !ok {
			t.Error("Invalid value for --branch")
		}

		branch, ok := v.(string)

		if !ok {
			t.Error("Invalid value for --branch")
		}

		if "other" != branch {
			t.Error("Invalid value for --branch")
		}
	}
}
