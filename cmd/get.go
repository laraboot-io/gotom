/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "Read a TOML document and extract an specific value",
	Args:  cobra.ExactArgs(2),
	Long:  `Use this command to extract a value from a TOML document.`,
	Run: func(cmd *cobra.Command, args []string) {

		//path, err := os.Getwd()
		//
		//if err != nil {
		//	log.Println(err)
		//	panic(err)
		//}

		file_name := fmt.Sprintf("%s", args[1])

		file, err := os.Open(file_name)

		if err != nil {
			fmt.Printf("'%s' was not found", args[1])
			panic(err)
		}

		defer func() {
			if err = file.Close(); err != nil {
				log.Fatal(err)
			}
		}()

		content, err := ioutil.ReadAll(file)

		config, _ := toml.Load(string(content))

		fmt.Printf("Content: %s\n", string(content))
		fmt.Printf("Getting: %s\n", args[0])

		// From: https://github.com/pelletier/go-toml#usage-example
		//		config, _ := toml.Load(`
		//[postgres]
		//user = "pelletier"
		//password = "mypassword"`)
		//		// retrieve data directly
		//		user := config.Get("postgres.user").(string)
		//
		//		// or using an intermediate object
		//		postgresConfig := config.Get("postgres").(*toml.Tree)
		//		password := postgresConfig.Get("password").(string)

		expValue := config.Get(args[0]).(string)
		fmt.Printf("%s\n", expValue)
	},
}

func init() {

	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
