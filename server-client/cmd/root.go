/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "client",
	Short: "APP that can post and GET",
	Long:  `An APP that can post an album or recieve an albums through id or titles, it can recieve also all albums`,
}

func Execute() error {
	return rootCmd.Execute()
}
