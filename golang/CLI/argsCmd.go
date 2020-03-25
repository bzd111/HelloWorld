package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/spf13/cobra"
)

func argsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "hello",
		Short: "hello",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return errors.New("requires a color argument")
			}
			return nil
		},
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("Hello, ", strings.Join(args, ""))
		},
	}
}
