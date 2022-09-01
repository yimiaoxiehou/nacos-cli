/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "nacos-cli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

var NS string
var server string

func init() {
	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	// rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.nacos-cli.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.PersistentFlags().StringVarP(&NS, "namespace", "n", "", "nacos namespace")
	rootCmd.PersistentFlags().StringVarP(&server, "server", "", "http://127.0.0.1:8848", "nacos server")
}

func LoadNamespaceId() string {
	if NS != "" {
		resp, err := http.Get(server + "/nacos/v1/console/namespaces")
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		namespaceResponse := &NamespaceResponse{}
		err = json.Unmarshal(bytes, namespaceResponse)
		if err != nil {
			log.Fatal(err)
		}
		if namespaceResponse.Code != 200 {
			log.Fatal(namespaceResponse.Message)
		}
		for i := range namespaceResponse.Data {
			if namespaceResponse.Data[i].NamespaceShowName == NS {
				return namespaceResponse.Data[i].Namespace
			}
		}
		log.Fatal(fmt.Sprintf("Can't find namespace for namespaceName %s'", NS))
	}
	return ""
}
