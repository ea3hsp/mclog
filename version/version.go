package version

import "runtime"

var (
	// Package is filled at linking time
	Package = "github.com/ea3hsp/mclog"

	// Version holds the complete version number. Filled in at linking time.
	Version = "0.5.0-beta"

	// Revision is filled with the VCS (e.g. git) revision being used to build
	// the program at linking time.
	Revision = ""

	// GoVersion is Go tree's version.
	GoVersion = runtime.Version()
)
