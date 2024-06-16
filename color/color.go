package color

import (
    "regexp"
    "strings"
)

func String(text string, color string) string {
    if !strings.HasPrefix(color, "\033[") {
        color = MatchColor(color)
    }

    return color + text + colorMap["default"]
}

func Template(message string) string {
    re := regexp.MustCompile(`[{]{2}([^{|}]*)\|+([^{|}]*)[}]{2}`)
    matches := re.FindAllStringSubmatch(message, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        color := strings.TrimSpace(match[2])

        output := String(text, color)
        message = strings.ReplaceAll(message, match[0], output)
    }

    return message
}

func RemoveTemplate(message string) string {
    re := regexp.MustCompile(`[{]{2}([^{|}]*)\|+([^{|}]*)[}]{2}`)
    matches := re.FindAllStringSubmatch(message, -1)

    for _, match := range matches {
        text := strings.TrimSpace(match[1])
        message = strings.ReplaceAll(message, match[0], text)
    }

    return message
}

func RemoveColor(message string) string {
    message = Template(message)

    re := regexp.MustCompile(`(\033\[\d+m)`)
    matches := re.FindAllStringSubmatch(message, -1)

    for _, match := range matches {
        message = strings.ReplaceAll(message, match[0], "")
    }

    return message
}

func MatchColor(color string) string {
    if val, ok := colorMap[color]; ok {
        return val
    }
    return ""
}

var colorMap = map[string]string{
    "default":   "\033[0m",
    "bold":      "\033[1m",
    "italic":    "\033[3m",
    "url":       "\033[4m",
    "blink":     "\033[5m",
    "blink2":    "\033[6m",
    "selected":  "\033[7m",
    "black":     "\033[30m",
    "red":       "\033[31m",
    "green":     "\033[32m",
    "yellow":    "\033[33m",
    "blue":      "\033[34m",
    "violet":    "\033[35m",
    "beige":     "\033[36m",
    "white":     "\033[37m",
    "blackbg":   "\033[40m",
    "redbg":     "\033[41m",
    "greenbg":   "\033[42m",
    "yellowbg":  "\033[43m",
    "bluebg":    "\033[44m",
    "violetbg":  "\033[45m",
    "beigebg":   "\033[46m",
    "whitebg":   "\033[47m",
    "grey":      "\033[90m",
    "red2":      "\033[91m",
    "green2":    "\033[92m",
    "yellow2":   "\033[93m",
    "blue2":     "\033[94m",
    "violet2":   "\033[95m",
    "beige2":    "\033[96m",
    "white2":    "\033[97m",
    "greybg":    "\033[100m",
    "redbg2":    "\033[101m",
    "greenbg2":  "\033[102m",
    "yellowbg2": "\033[103m",
    "bluebg2":   "\033[104m",
    "violetbg2": "\033[105m",
    "beigebg2":  "\033[106m",
    "whitebg2":  "\033[107m",
}
