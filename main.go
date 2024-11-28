package main

import (
	"flower-management/internal/core/servicesinitializer"

	"github.com/spf13/cobra"
)

var envFilename string

func main() {
	cli := &cobra.Command{
		Run: func(cmd *cobra.Command, args []string) {
			servicesinitializer.Execute(envFilename)
		},
	}
	cli.Flags().StringVarP(&envFilename, "env-filename", "e", "", "environment variables")

	err := cli.Execute()
	if err != nil {
		panic(err)
	}
}
