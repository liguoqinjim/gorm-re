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
	"github.com/liguoqinjim/gorm-re/re"
	"github.com/spf13/cobra"
)

// genCmd represents the gen command
var genCmd = &cobra.Command{
	Use:   "gen",
	Short: "生成go struct",
	Long:  `mysql->go struct`,
	Run:   gen,
}

var (
	filepath string

	host     string
	port     int
	user     string
	password string
	dbName   string

	output      string
	packageName string
	mysql8      bool
	gormVersion int
)

func init() {
	rootCmd.AddCommand(genCmd)

	genCmd.Flags().StringVarP(&filepath, "file", "f", "", "配置文件")

	//数据库
	genCmd.Flags().StringVarP(&host, "host", "H", "127.0.0.1", "database host")
	genCmd.Flags().IntVarP(&port, "port", "P", 3306, "database port")
	genCmd.Flags().StringVarP(&user, "user", "u", "root", "database user")
	genCmd.Flags().StringVarP(&password, "pwd", "p", "", "database password")
	genCmd.Flags().StringVar(&dbName, "dbName", "", "database name")
	genCmd.Flags().StringVar(&output, "output", "model.go", "output file name")
	genCmd.Flags().StringVar(&packageName, "package", "model", "package name")
	genCmd.Flags().IntVarP(&gormVersion, "version", "v", 1, "gorm version")

	//if err := cmd.MarkFlagRequired("file"); err != nil {
	//	log.Errorf("make required error:%v", err)
	//}

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// genCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// genCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func gen(cmd *cobra.Command, args []string) {
	re.RE(filepath, host, port, user, password, dbName, output, packageName, "", gormVersion)
}
