package cmd

import (
	"fmt"
	"os"

	"github.com/deependujha/cf/cp_template"
	"github.com/deependujha/cf/os_utils"
	"github.com/spf13/cobra"
)

// var outputFilepath string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "cf",
	Short: "A cli tool for Codeforces problem solving",
	Long:  `A cli tool for Codeforces problem solving with template generation`,
	Args:  cobra.ExactArgs(1), // Enforce exactly one argument
	Run: func(cmd *cobra.Command, args []string) {
		// Get the input file path from the first argument
		inputFilepath := args[0]

		// create file
		if err := os_utils.CreateFileIfNotExists(inputFilepath); err != nil {
			fmt.Fprintf(os.Stderr, "Error creating file: %v\n", err)
			os.Exit(1)
		}

		// get complete cpp code
		code := cp_template.GetCode()

		// write to file
		if err := os_utils.AppendStringToFile(inputFilepath, code); err != nil {
			fmt.Fprintf(os.Stderr, "Error writing to file: %v\n", err)
			os.Exit(1)
		}

		// run the format command
		command := fmt.Sprintf("clang-format -i %s", inputFilepath)
		if _, err := os_utils.RunCommand(command); err != nil {
			fmt.Fprintf(os.Stderr, "Error formatting file: %v\n", err)
			os.Exit(1)
		}
	},
}

func init() {
	// Define flags
	// rootCmd.Flags().StringVarP(&outputFilepath, "output", "o", "litdata_trace.json", "Path to the output trace file")
}

// Execute adds all child commands to the root command and sets flags appropriately.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
