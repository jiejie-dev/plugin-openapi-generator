/*
Copyright © 2024 jiejie-dev <jeremaihloo1024@gmail.com>
*/
package main

import (
	"fmt"
	"os"
	"os/exec"
)

var debug bool

func getEnv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		if debug {
			fmt.Printf("Found env var %s=%s\n", key, value)
		}
		return value
	}
	if debug {
		fmt.Printf("Env var %s not found, using default %s\n", key, def)
	}
	return def
}

func main() {
	options := []string{"generate"}
	d := getEnv("PLUGIN_DEBUG", "false")
	if d == "true" {
		debug = true
	}
	auth := getEnv("PLUGIN_AUTH", "")
	if auth != "" {
		options = append(options, "-a", auth)
	}
	apiNameSuffix := getEnv("PLUGIN_API_NAME_SUFFIX", "Api")
	if apiNameSuffix != "" {
		options = append(options, "--api-name-suffix", apiNameSuffix)
	}
	apiPackage := getEnv("PLUGIN_API_PACKAGE", "")
	if apiPackage != "" {
		options = append(options, "--api-package", apiPackage)
	}
	artifactId := getEnv("PLUGIN_ARTIFACT_ID", "")
	if artifactId != "" {
		options = append(options, "--artifact-id=", artifactId)
	}
	artifactVersion := getEnv("PLUGIN_ARTIFACT_VERSION", "")
	if artifactVersion != "" {
		options = append(options, "--artifact-version", artifactVersion)
	}
	config := getEnv("PLUGIN_CONFIG", "")
	if config != "" {
		options = append(options, "-c", config)
	}
	dryRun := getEnv("PLUGIN_DRY_RUN", "false")
	if dryRun == "true" {
		options = append(options, "--dry-run")
	}
	generatorName := getEnv("PLUGIN_GENERATOR_NAME", "")
	if generatorName != "" {
		options = append(options, "-g", generatorName)
	}
	gitHost := getEnv("PLUGIN_GIT_HOST", "")
	if gitHost != "" {
		options = append(options, "--git-host", gitHost)
	}
	gitRepoId := getEnv("PLUGIN_GIT_REPO_ID", "")
	if gitRepoId != "" {
		options = append(options, "--git-repo-id", gitRepoId)
	}
	gitUserId := getEnv("PLUGIN_GIT_USER_ID", "")
	if gitUserId != "" {
		options = append(options, "--git-user-id", gitUserId)
	}
	inputSpec := getEnv("PLUGIN_INPUT_SPEC", "")
	if inputSpec != "" {
		options = append(options, "-i", inputSpec)
	}
	output := getEnv("PLUGIN_OUTPUT", "output")
	if output != "" {
		options = append(options, "-o", output)
	}
	c := exec.Command("/usr/local/bin/docker-entrypoint.sh", options...)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	err := c.Run()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
