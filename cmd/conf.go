/*
Copyright © 2024 hvturingga hvturingga@disism.com
*/
package cmd

import (
	"fmt"
	"github.com/disism/godis/internal/nats"
	"github.com/spf13/cobra"
	"os"
)

// confCmd represents the conf command
var confCmd = &cobra.Command{
	Use:     "conf",
	Short:   "",
	Long:    ``,
	Aliases: []string{"c"},
}

var addConfCmd = &cobra.Command{
	Use:   "add",
	Short: "",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		name, _ := cmd.Flags().GetString("name")
		path, _ := cmd.Flags().GetString("path")
		fmt.Println(name, path)
		v, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		if err := nats.KVPut(name, v); err != nil {
			fmt.Println(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(confCmd)
	confCmd.AddCommand(addConfCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// confCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// confCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	addConfCmd.Flags().StringP("name", "n", "", "conf name")
	addConfCmd.Flags().StringP("path", "p", "", "config file path")

	if err := addConfCmd.MarkFlagRequired("name"); err != nil {
		fmt.Println(err)
		return
	}

	if err := addConfCmd.MarkFlagRequired("path"); err != nil {
		fmt.Println(err)
		return
	}

}
