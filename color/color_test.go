package color

import (
    "testing"
)

func TestColorString(t *testing.T) {
    text := "Hello, world!"
    color := "red"
    expected := "\033[31mHello, world!\033[0m"
    result := String(text, color)
    if result != expected {
        t.Errorf("String(%q, %q) = %q, expected %q", text, color, result, expected)
    }
}

func TestColorMatchTemplate(t *testing.T) {
    message := "{{Hello|red}} {{world|green}}"
    expected := "\033[31mHello\033[0m \033[32mworld\033[0m"
    result := Template(message)
    if result != expected {
        t.Errorf("ColorMatchTemplate(%q) = %q, expected %q", message, result, expected)
    }
}

func TestRemoveColor(t *testing.T) {
    message := "\033[31mHello\033[0m, \033[32mworld\033[0m!"
    expected := "Hello, world!"
    result := RemoveColor(message)
    if result != expected {
        t.Errorf("RemoveColor(%q) = %q, expected %q", message, result, expected)
    }
}
