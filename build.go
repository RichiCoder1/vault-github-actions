//+build mage

package main

import (
	"os"

	"github.com/magefile/mage/mg"
	"github.com/magefile/mage/sh"
)

// Init go and npm
func Setup() error {
	err := sh.Run("go", "mod", "download")
	if err != nil {
		return err
	}

	err = os.Chdir("web")
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	return sh.Run("npm", "ci")
}

type Build mg.Namespace

// Builds the web frontend
func (Build) Web() error {
	err := os.Chdir("web")
	if err != nil {
		return err
	}
	defer os.Chdir("..")

	return sh.Run("npm", "run", "build")
}
