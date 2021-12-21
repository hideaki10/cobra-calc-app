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

var check bool

// subtractCmd represents the subtract command
var subtractCmd = &cobra.Command{
	Use:   "subtract",
	Short: "Subtract one interger form another",
	Long:  `Subtract one interger a from another interger b; result = a - b `,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		var a, b int
		var err error

		a, err = strconv.Atoi(args[0])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}

		b, err = strconv.Atoi(args[1])
		if err != nil {
			panic("Arguments to `subtract` must be integers")
		}

		calc.Subtract(a, b, check)

		fmt.Println("subtract called")
	},
}

func init() {
	subtractCmd.Flags().BoolVar(&check, "check", false, "check controls if overflow or underflow check is performed")

	rootCmd.AddCommand(subtractCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// subtractCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// subtractCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
