package config

import (
	"encoding/json"
	"os"
	"fmt"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	Current_user_name 			string	`json:"current_user_name"`
	Db_url						string	`json:"db_url"`
}

func getConfigFilePath() (string, error) {

	path, err := os.UserHomeDir()
	if err != nil {
		fmt.Println("Not found home folder")
		os.Exit(1)
	}
	full_path := fmt.Sprintf("%s/%s", path, configFileName)
	return full_path, nil
}

func write(data *Config) error {

	jsonData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Something Gone Wrong")
		os.Exit(1)
	}
	return os.WriteFile(path, jsonData, 0644)
}

func (data *Config) SetUser(user_name string) error {

	data.Current_user_name = user_name
	err := write(data)
	return err
}

func Read() Config {

	var data Config

	full_path, err := getConfigFilePath()
	if err != nil {
		fmt.Println("Something Gone Wrong")
		os.Exit(1)
	}

	file, err := os.Open(full_path)
	if err != nil {
		fmt.Println("Couldn't open the file")
		os.Exit(1)
	}
	defer file.Close()

	err = json.NewDecoder(file).Decode(&data)
	if err != nil {
		fmt.Println("Couldn't decode the json")
		os.Exit(1)
	}
	return data
}
