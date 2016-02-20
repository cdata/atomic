package command

import (
	"fmt"
	"github.com/cdata/atomic/util"
	cli "github.com/codegangsta/cli"
)

// Post ...
func Post(cli *cli.Context) {
	post := cli.String("message")

	if len(post) == 0 {
		input := make(chan string)
		go util.ReceiveInputFromEditor(input)
		post = <-input
	}

	if len(post) == 0 {
		panic("A message is required to create a post.")
	}

	fmt.Println(len(post))
	fmt.Println(post)
}
