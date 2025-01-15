package config


import (
        "encoding/json"
        "os"
)

type Config struct {
        Database struct {
                Host     string `json:"host"`
                Port     string `json:"port"`
                User     string `json:"user"`
                Password string `json:"password"`
                Name     string `json:"dbname"`
                SSLMode  string `json:"sslmode"` // e.g., "disable" or "require"
        } `json:"database"`
}

func LoadConfig(path string) (*Config, error) {
        file, err := os.Open(path)
        if err != nil {
                return nil, err
        }
        defer file.Close()

        var config Config
        decoder := json.NewDecoder(file)
        err = decoder.Decode(&config)
        if err != nil {
                return nil, err
        }

        return &config, nil
}