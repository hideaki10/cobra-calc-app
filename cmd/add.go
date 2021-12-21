/*
Copyright © 2021 hideaki10

*/
package cmd

import (
	"calc-app/pkg/calc"
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// addCmd represents the add command
var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add two integers",
	Long:  `Add two integers`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error

		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Argument to `add` must be an intergers")
		}
		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Argument to `add` must be an intergers")
		}

		result := calc.Add(a, b, check)

		fmt.Println(result)
	},
}

func init() {
	addCmd.Flags().BoolVar(&check, "check", false, "Check controls if overflow or underflow check is performed")
	rootCmd.AddCommand(addCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// addCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// addCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
