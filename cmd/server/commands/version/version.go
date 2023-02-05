package version

import (
	"fmt"

	"github.com/ea3hsp/mclog/version"
	"github.com/urfave/cli/v2"
)

var Command = &cli.Command{
	Name:  "version",
	Usage: "print service version",
	Action: func(ctx *cli.Context) error {
		fmt.Println("cmLog Server:")
		fmt.Println("  Version: ", version.Version)
		fmt.Println("  Revision:", version.Revision)
		fmt.Println("  Go version:", version.GoVersion)
		fmt.Println("")
		return nil
	},
}
