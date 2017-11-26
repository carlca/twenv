package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/user"
	"path"
	
	"github.com/pkg/errors"
)

// Credentials struct is read from ~/gotwitter/config.json config file
type Credentials struct {
	ConsumerKey    string `json:"consumerkey"`
	ConsumerSecret string `json:"consumersecret"`
	AccessToken    string `json:"accesstoken"`
	AccessSecret   string `json:"accesssecret"`
}

func terminateOnError(err error) {
	if err != nil {
		fmt.Printf("%+v", errors.WithStack(err))
		os.Exit(1) // or anything else ...
	}
}

// get the current os user
func getCurrentUser() *user.User {
	usr, err := user.Current()
	terminateOnError(err)
	return usr
}

// get the current users's configuration path for the gotwitter application
func getUserConfig(usr *user.User) string {
	config := path.Join(usr.HomeDir, ".gotwitter/config.json")
	_, err := os.Stat(config)
	terminateOnError(err)
	return config
}

// open a file based on the specified config path
func openFile(configPath string) *os.File {
	file, err := os.Open(configPath)
	terminateOnError(err)
	return file
}

// read the config json file into the Credentials struct
func readCredentials(file *os.File) *Credentials {
	var creds *Credentials
	decoder := json.NewDecoder(file)
	defer file.Close()
	err := decoder.Decode(&creds)
	terminateOnError(err)
	return creds
}

func saveEnvironmantVariables(creds *Credentials) {
	os.Setenv("TWEETS_CONSUMER_KEY", creds.ConsumerKey)
	os.Setenv("TWEETS_CONSUMER_SECRET", creds.ConsumerSecret)
	os.Setenv("TWEETS_ACCESS_TOKEN", creds.AccessToken)
	os.Setenv("TWEETS_ACCESS_SECRET", creds.AccessSecret)
}

func displayEnvironmantVariables() {
	fmt.Println("TWEETS_CONSUMER_KEY="+os.Getenv("TWEETS_CONSUMER_KEY"), " \\")
	fmt.Println("TWEETS_CONSUMER_SECRET=" + os.Getenv("TWEETS_CONSUMER_SECRET") + " \\")
	fmt.Println("TWEETS_ACCESS_TOKEN=" + os.Getenv("TWEETS_ACCESS_TOKEN") + " \\")
	fmt.Println("TWEETS_ACCESS_SECRET=" + os.Getenv("TWEETS_ACCESS_SECRET") + " tweets")
}

func main() {
	currentUser := getCurrentUser()
	configPath := getUserConfig(currentUser)
	configFile := openFile(configPath)
	credentials := readCredentials(configFile)
	saveEnvironmantVariables(credentials)
	displayEnvironmantVariables()
}
