package utils

import (
        "encoding/json"
        "log"
        "os"
)

func LoadProperties(path string, target interface{}) error {
        file, err := os.Open(path)
        if err != nil {
                log.Fatal("Error opening properties file:", err)
                return err
        }
        defer file.Close()

        decoder := json.NewDecoder(file)
        err = decoder.Decode(target)
        if err != nil {
                log.Fatal("Error decoding properties file:", err)
                return err
        }
        return nil
}