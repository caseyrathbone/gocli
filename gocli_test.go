package gocli

import (
  "fmt"
  "testing"
)

func TestUnknownCommand(t *testing.T) {
  // Default error message should be returned.
  fmt.Println("TestUnknownCommand()")
}

func TestFirstCommandInMap(t *testing.T) {
  fmt.Println("TestFirstCommandInMap()")
}

func TestLastCommandInMap(t *testing.T) {
  fmt.Println("TestLastCommandInMap()")
}

func TestNilDefault(t *testing.T) {
  fmt.Println("TestNilDefault()")
}

func TestNilCommandMap(t *testing.T) {
  fmt.Println("TestNilCommandMap()")
}

func TestHelp(t *testing.T) {
  // Setup the command map
  /*
  for _, c := range []struct {
    in, want string
  }{
    {"hello, world", "dlrow ,olleh"},
  } {
    got := Simple(c.in)
    if got != c.want {
      t.Errorf("Reversing error")
    }
  }
  */
}

func TestMultiArgs(t *testing.T) {
  fmt.Println("another test")
}
