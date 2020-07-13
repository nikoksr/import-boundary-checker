package main_test

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	"git.bytecode.nl/foss/import-boundry-checker/logging"
	"git.bytecode.nl/foss/import-boundry-checker/runner"
	"github.com/stretchr/testify/assert"
)

func TestEndToEnd(t *testing.T) {
	rootDir, err := os.Getwd()
	fmt.Println(rootDir)
	assert.NoError(t, err)

	tests := []struct {
		dir       string
		shouldErr bool
	}{
		{"go-invalid-1", true},
		{"go-valid-1", false},
	}
	for _, test := range tests {

		// cd into test dir
		err := os.Chdir("./examples/" + test.dir)
		assert.NoError(t, err)

		//REMOVE
		wd, err := os.Getwd()
		assert.NoError(t, err)
		fmt.Println(wd)

		// Load config file
		abs, err := filepath.Abs(".importrules")
		assert.NoError(t, err)
		configFile, err := ioutil.ReadFile(abs)
		fmt.Println(string(configFile))
		assert.NoError(t, err)
		config := string(configFile)

		// Create and run runner
		// TODO: Add automated end-to-end tests
		logger := logging.Logger{Verbose: false}
		r := runner.New(config, &logger)
		got := r.Run()

		// Check if we got what we expected
		assert.Equal(t, test.shouldErr, got)

		// Change back to parent directory
		err = os.Chdir(rootDir)
	}
}