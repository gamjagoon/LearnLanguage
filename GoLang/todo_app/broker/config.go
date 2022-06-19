package broker

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
)

type Config struct {
	HTTPPort string `json:"httpPort"`
	Host     string `json:"host"`
	Port     string `json:"port"`
}

func showHelp() {
	fmt.Printf("%s\n", "usageHelp")
	os.Exit(0)
}

func ConfigureConfig(args []string) (*Config, error) {
	config := &Config{}
	var (
		help       bool
		configFile string
	)
	fs := flag.NewFlagSet("gmqtt-broker", flag.ExitOnError)
	fs.Usage = showHelp

	fs.BoolVar(&help, "h", false, "Show this message.")
	fs.BoolVar(&help, "help", false, "Show this message.")
	fs.StringVar(&config.HTTPPort, "httpport", "8080", "Port to listen on.")
	fs.StringVar(&config.HTTPPort, "hp", "8080", "Port to listen on.")
	fs.StringVar(&config.Port, "port", "1883", "Port to listen on.")
	fs.StringVar(&config.Port, "p", "1883", "Port to listen on.")

	if err := fs.Parse(args); err != nil {
		return nil, err
	}

	if help {
		showHelp()
		return nil, nil
	}

	if configFile != "" {
		tmpConfig, e := LoadConfig(configFile)
		if e != nil {
			return nil, e
		} else {
			config = tmpConfig
		}
	}
	return config, nil
}

func LoadConfig(filename string) (*Config, error) {

	content, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	var config Config
	err = json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
