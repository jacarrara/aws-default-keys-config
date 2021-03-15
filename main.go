package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
)

type awsKeys struct {
	aws_access_key_id     string
	aws_secret_access_key string
	aws_session_token     string
}

type AwsKeys interface {
	getKeys()
	setKeys()
}

func main() {
	k := awsKeys{}

	k.getKeys()
	k.setKeys()

	fmt.Println("Done!")
}

func (k *awsKeys) getKeys() {
	k.aws_access_key_id = os.Getenv("AWS_ACCESS_KEY_ID")
	k.aws_secret_access_key = os.Getenv("AWS_SECRET_ACCESS_KEY")
	k.aws_session_token = os.Getenv("AWS_SESSION_TOKEN")
}

func (k *awsKeys) setKeys() {
	config := fmt.Sprintf("[default]\naws_access_key_id = %s\naws_secret_access_key = %s\naws_session_token = %s", k.aws_access_key_id, k.aws_secret_access_key, k.aws_session_token)

	homeDir, err := os.UserHomeDir()
	if err != nil {
		fmt.Println(err)
	}

	awsDirPath := filepath.Join(homeDir, ".aws")

	if _, err := os.Stat(awsDirPath); os.IsNotExist(err) {
		err := os.Mkdir(awsDirPath, 0755)
		if err != nil {
			fmt.Printf("No se pudo crear el directorio en %s", awsDirPath)
			os.Exit(1)
		}
	}

	filename := filepath.Join(homeDir, ".aws", "credentials")

	configBytes := []byte(config)
	errF := ioutil.WriteFile(filename, configBytes, 0644)
	if errF != nil {
		fmt.Println(errF)
	}

}
