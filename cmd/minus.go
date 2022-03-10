/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"fmt"
	"strconv"

	"github.com/spf13/cobra"
)

// minusCmd represents the minus command
var minusCmd = &cobra.Command{
	Use:   "minus",
	Short: "Substract two numbers",
	Long: `Substraction of two numbers. The program will also display the result`,
	Run: func(cmd *cobra.Command, args []string) {
		num1, _ := strconv.ParseUint(args[0], 10, 32)
		num2, _ := strconv.ParseUint(args[1], 10, 32)
		sub := num1 - num2
		fmt.Println(sub)
	},
}

func init() {
	rootCmd.AddCommand(minusCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// minusCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// minusCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
