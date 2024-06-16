package xlog

import (
    "bufio"
    "fmt"
    "log"
    "os"
    "strings"
)

func fileWriter(name string, config Config) {
    var writerOutFile *bufio.Writer
    var writerErrFile *bufio.Writer

    file, err := os.OpenFile(config.Path+"application.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    writerOutFile = bufio.NewWriter(file)

    file, err = os.OpenFile(config.Path+"application-err.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
    if err != nil {
        log.Fatalf("error: %v", err)
    }
    writerErrFile = bufio.NewWriter(file)

    for {
        select {
        case msg := <-logChannels[name]:

            var writer *bufio.Writer
            if msg.Level == LevelError || msg.Level == LevelFatal {
                writer = writerErrFile
            } else {
                writer = writerOutFile
            }

            var timestamp = msg.Timestamp.Format("2006-01-02 15:04:05")
            var level = msg.Level.name()
            var caller = msg.Caller
            var line = fmt.Sprintf("(%d)", msg.Line)
            var message = msg.Message

            if config.Caller == "short" {
                caller = caller[strings.LastIndex(caller, "/")+1:]
            }

            var output string
            if config.Caller == "short" || config.Caller == "long" {
                output = fmt.Sprintf("%s [ %s ] %s %s %s\n", timestamp, level, caller, line, message)
            } else {
                output = fmt.Sprintf("%s [ %s ] %s\n", timestamp, level, message)
            }

            writer.WriteString(output)
            writer.Flush()
        }
    }
}
