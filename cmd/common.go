/*
Copyright © 2025 Andriy Melnyk
*/
package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

func PersistentFlags(cmd *cobra.Command) map[string]string {
	flags := make(map[string]string)
	parent := cmd
	for parent != nil {
		parent.PersistentFlags().VisitAll(func(f *pflag.Flag) {
			flags[f.Name] = f.Value.String()
		})
		parent = parent.Parent()
	}
	return flags
}

func PrintFlags(cmd *cobra.Command) {
	for k, v := range PersistentFlags(cmd) {
		fmt.Printf("flag: --%s: %s\n", k, v)
	}
}
