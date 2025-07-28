// Package main provides the eniGOma command-line interface.
//
// Copyright (c) 2025 David Duarte
// Licensed under the MIT License
package main

import (
	"os"

	"github.com/coredds/eniGOma/internal/cli"
)

func main() {
	if err := cli.Execute(); err != nil {
		os.Exit(1)
	}
}
