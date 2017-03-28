package gocli

import (
  "fmt"
  "errors"
)

// Intended to provide a simple mapping of keywords
// to functions from CLI args.

type KeywordFunc func(args ...string)

type CLI struct {
  CmdMap map[string] KeywordFunc
  HelperMsg string
}

func (c *CLI) Run(args ...string) (int, error) {
  if len(args) < 1 {
    fmt.Println(c.HelperMsg)
    return 1, errors.New("Invalid commandline input: Expecting additional arguments")
  }

  // Argument to lookup in map.
  command := args[0]

  if c.CmdMap[command] != nil {
    c.CmdMap[command](args[1:]...)
  } else {
    fmt.Println(c.HelperMsg)
  }

  return 0, nil
}
