package utils

import (
    "encoding/json"
    "os"
)

type Config struct {
    ServerPort string `json:"server_port"`
    LogLevel   string `json:"log_level"`
}

func LoadConfig(filePath string) (*Config, error) {
    file, err := os.Open(filePath)
    if err != nil {
        return nil, err
    }
    defer file.Close()

    config := &Config{}
    decoder := json.NewDecoder(file)
    if err := decoder.Decode(config); err != nil {
        return nil, err
    }

    return config, nil
}

