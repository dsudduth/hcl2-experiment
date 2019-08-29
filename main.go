package main

import (
	"encoding/json"
	"fmt"
	"hcltest/configs"
	"log"
	"os"
)

var rootConfigPath = "demo/app.hcl"

func main() {
	c, err := loadConfig(rootConfigPath)
	if err != nil {
		os.Exit(1)
	}
	o, _ := json.MarshalIndent(c, "", "  ")
	fmt.Printf("%+v\n", string(o))
}

func loadConfig(path string) (*configs.Config, error) {
	log.Printf("[INFO] Loading Config from path: %s", path)
	if _, err := os.Stat(path); !os.IsNotExist(err) {
		cfg, err := configs.ParseFile(path)
		if err != nil {
			log.Printf("[ERROR] %s", err)
			return nil, err
		}
		return cfg, nil
	}
	return nil, nil
}
