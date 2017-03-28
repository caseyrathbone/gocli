package gocli

import (
  "fmt"
  "errors"
)

// Provides functionality for a developer to provide a simple map of keywords
// and functions to be executed from CLI args.
/*
package main

func main() {
  c := cli(make[]{
    command: "test",
    function: &func,
  },
  {
    help: "This is the default message"
  },
  )

  _, err = c.Run();
  if (err != nil) {
    log.Fatal(err)
  }

  os.exit(0)
}
*/

//type cmdFunc func(args ...interface{})
//type defaultFunc func(unknown string)

type SimpleFunc func(args ...interface{})

type CLI struct {
  CmdMap map[string]SimpleFunc
  HelperMsg string
}

func (c *CLI) Run(args []string) (int, error) {
  if len(args) < 1 {
    fmt.Println(c.HelperMsg)
    return 1, errors.New("Uh oh...")
  }

  command := args[0]

  if c.CmdMap[command] != nil {
    c.CmdMap[command]()
  } else {
    fmt.Println(c.HelperMsg)
  }

  return 0, nil
}
