/*
Copyright © 2022 NAME HERE monasteriosd@gmail.com>

*/ 
package cmd 

import (
	"os"
	"github.com/spf13/cobra"
    "fmt"
)



// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "restoreip",
	Short: "Transforms a string of integers into an IPv4 address if possible",
    Run: func(cmd *cobra.Command, args []string) {
        for _, arg := range args {
        fmt.Println(RestoreIp(arg))
    }
    },

}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}


