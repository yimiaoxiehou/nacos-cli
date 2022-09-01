/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/scylladb/termtables"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		accessToken := LoadAccessToken()
		resp, err := http.Get(server + fmt.Sprintf("/nacos/v1/ns/catalog/services?accessToken=%s&namespaceId=%s&pageNo=1&pageSize=100", accessToken, LoadNamespaceId()))
		if err != nil {
			log.Fatal(err)
		}
		bytes, err := io.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		servicesResponse := &ServicesResponse{}
		err = json.Unmarshal(bytes, &servicesResponse)
		if err != nil {
			log.Fatal(err)
		}
		t := termtables.CreateTable()
		t.AddHeaders("Name", "HealthyInstanceCount")
		for i := range servicesResponse.ServiceList {
			t.AddRow(servicesResponse.ServiceList[i].Name, servicesResponse.ServiceList[i].HealthyInstanceCount)
		}
		t.Style.BorderI = ""
		t.Style.BorderX = ""
		t.Style.BorderY = ""
		fmt.Println(t.Render())
	},
}

func init() {
	rootCmd.AddCommand(serviceCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// serviceCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// serviceCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
