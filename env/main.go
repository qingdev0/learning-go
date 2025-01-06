package main

import (
	"fmt"
	"os"
	"runtime"
)

func checkAndPrintEnvVar(envVarName, runtimeValue string) {
	envVarValue := os.Getenv(envVarName)
	var output string

	if envVarValue == "" {
		output = fmt.Sprintf("%s: not set (runtime: %s)", envVarName, runtimeValue)
	} else {
		output = fmt.Sprintf("%s: %s (runtime: %s)", envVarName, envVarValue, runtimeValue)
	}

	fmt.Println(output)
}

func main() {
	checkAndPrintEnvVar("GOROOT", runtime.GOROOT())
	checkAndPrintEnvVar("GOPATH", "no runtime equivalent")
	checkAndPrintEnvVar("GOOS", runtime.GOOS)
	checkAndPrintEnvVar("GOARCH", runtime.GOARCH)
}
