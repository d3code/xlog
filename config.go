package xlog

import (
    "bufio"
    "gopkg.in/yaml.v3"
    "os"
)

type Config struct {
    Enabled bool   `yaml:"enabled"`
    Type    string `yaml:"type"`
    Level   string `yaml:"level"`
    Format  string `yaml:"format"`
    Color   bool   `yaml:"color"`
    Caller  string `yaml:"caller"`
    Path    string `yaml:"path"`
}

var (
    logConfig        map[string]Config
    logChannels      = make(map[string]chan logItem)
    consoleConfig    *Config
    writerOutConsole = bufio.NewWriter(os.Stdout)
    writerErrConsole = bufio.NewWriter(os.Stderr)
)

func init() {
    err := readConfig()
    if err != nil {
        logConfig = map[string]Config{
            "info": {
                Enabled: true,
                Type:    "console",
                Level:   "info",
                Format:  "text",
                Color:   true,
                Caller:  "short",
            },
        }
    }

    for name, config := range logConfig {
        if config.Enabled {
            switch {
            case config.Type == "console":
                consoleConfig = &config
            case config.Type == "file":
                logChannels[name] = make(chan logItem)
                go fileWriter(name, config)
            case config.Type == "database":
                logChannels[name] = make(chan logItem)
                go databaseWriter(name, config)
            }
        }
    }
}

func readConfig() error {
    file, err := os.ReadFile("./config/log.yaml")
    if err != nil {
        return err
    }

    err = yaml.Unmarshal(file, &logConfig)
    if err != nil {
        return err
    }

    return nil
}
