/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// divideCmd represents the divide command
var divideCmd = &cobra.Command{
	Use:   "divide",
	Short: "Divide two numbers",
	Long: `Divisionn of two numbers. The second number cannot be 0. The program will also display the result`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)
 
		if num2 != 0 {
			div := num1 / num2
			fmt.Println(div)
		} else {
			fmt.Println("Don't use 0 for this!")
		}
		
	},
}

func init() {
	rootCmd.AddCommand(divideCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// divideCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// divideCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
