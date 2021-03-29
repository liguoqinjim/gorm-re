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
	log "github.com/sirupsen/logrus"

	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成go struct",
	Long:  `mysql->go struct`,
	Run:   gen,
}

func init() {
	rootCmd.AddCommand(genCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func gen(cmd *cobra.Command, args []string) {
	//localCmd.Flags().StringVarP(&Source, "source", "s", "", "Source directory to read from")
	var filepath string
	cmd.Flags().StringVarP(&filepath, "file", "f", "1", "配置文件")
	cmd.Flags().StringVarP(&filepath, "file2", "c", "1", "配置文件")
	//cmd.Flags().rootCmd.MarkFlagRequired("region")
	if err := cmd.MarkFlagRequired("file"); err != nil {
		log.Errorf("make required error:%v", err)
	}

	log.Println("filepath=", filepath)
}
