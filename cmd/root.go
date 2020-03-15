// Copyright © 2020 NAME HERE scbizu@gmail.com
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"io/ioutil"
	"os"

	"github.com/scbizu/vasques/internal/build"
	"github.com/scbizu/vasques/internal/reader"
	"github.com/spf13/cobra"
)

var jsonFile string
var pkgname string
var destPath string

// RootCmd represents the base command when called without any subcommands
var RootCmd = &cobra.Command{
	Use:   "vasques",
	Short: "vasques -- A tool to generate proto3 from JSON",
	RunE:  func(cmd *cobra.Command, args []string) error { return convertHandler() },
}

func convertHandler() error {
	bs, err := ioutil.ReadFile(jsonFile)
	if err != nil {
		return err
	}
	build.SetPackageName(pkgname)
	r, err := reader.NewReader(bs)
	if err != nil {
		return err
	}
	pb3, err := build.BuildFromReader(r)
	if err != nil {
		return err
	}
	return pb3.Generate(destPath)
}

// Execute adds all child commands to the root command sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(-1)
	}
}

func init() {
	RootCmd.PersistentFlags().StringVar(&jsonFile, "json", "", "json file needed to be converted to proto")
	RootCmd.PersistentFlags().StringVarP(&pkgname, "package", "p", "", "generated proto3 package")
	RootCmd.PersistentFlags().StringVarP(&destPath, "dest", "d", "", "generated file destination path")
}
