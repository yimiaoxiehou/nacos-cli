/*
Copyright © 2022 NAME HERE <EMAIL ADDRESS>

*/
package cmd

import (
	"encoding/json"
	"github.com/spf13/cobra"
	"io"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"
)

// loginCmd represents the login command
var loginCmd = &cobra.Command{
	Use:        "login",
	Aliases:    nil,
	SuggestFor: nil,
	Short:      "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Example:                "",
	ValidArgs:              nil,
	ValidArgsFunction:      nil,
	Args:                   nil,
	ArgAliases:             nil,
	BashCompletionFunction: "",
	Deprecated:             "",
	Annotations:            nil,
	Version:                "",
	PersistentPreRun:       nil,
	PersistentPreRunE:      nil,
	PreRun:                 nil,
	PreRunE:                nil,
	Run: func(cmd *cobra.Command, args []string) {
		doLogin()
	},
	RunE:                       nil,
	PostRun:                    nil,
	PostRunE:                   nil,
	PersistentPostRun:          nil,
	PersistentPostRunE:         nil,
	FParseErrWhitelist:         cobra.FParseErrWhitelist{},
	CompletionOptions:          cobra.CompletionOptions{},
	TraverseChildren:           false,
	Hidden:                     false,
	SilenceErrors:              false,
	SilenceUsage:               false,
	DisableFlagParsing:         false,
	DisableAutoGenTag:          false,
	DisableFlagsInUseLine:      false,
	DisableSuggestions:         false,
	SuggestionsMinimumDistance: 0,
}

var username string
var password string

func init() {
	rootCmd.AddCommand(loginCmd)
	loginCmd.Flags().StringVarP(&username, "username", "u", "nacos", "nacos username")
	loginCmd.Flags().StringVarP(&password, "password", "p", "nacos", "nacos password")
}

func checkSessionAlive() {
	_, err := os.Stat(".nacos/session")
	if err != nil {
		log.Fatal("session file not exist. please login.")
	}
	now := time.Now().Format(time.RFC3339)
	bytes, err := os.ReadFile(".nacos/expireTime")
	if err != nil {
		os.Remove(".nacos")
		log.Fatal(err)
	}
	if string(bytes) <= now {
		log.Fatal("login expire. please reLogin.")
	}
}

func LoadAccessToken() string {
	checkSessionAlive()
	bytes, err := os.ReadFile(".nacos/session")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

func doLogin() {
	resp, err := http.PostForm(server+"/nacos/v1/auth/users/login", url.Values{"username": {username}, "password": {password}})
	if err != nil {
		log.Fatal(err)
	}
	loginResp := make(map[string]interface{})
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	err = json.Unmarshal(bytes, &loginResp)
	if err != nil {
		log.Fatal(err)
	}
	_, err = os.Stat(".nacos")
	if err == nil {
		err = os.Remove(".nacos")
		if err != nil {
			log.Fatal(err)
		}
	}

	err = os.Mkdir(".nacos", os.FileMode(0755))
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(".nacos/session", []byte(loginResp["accessToken"].(string)), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}
	err = os.WriteFile(".nacos/expireTime", []byte(time.Now().Add(18000*time.Second).Format(time.RFC3339)), os.ModePerm)
	if err != nil {
		log.Fatal(err)
	}

}
