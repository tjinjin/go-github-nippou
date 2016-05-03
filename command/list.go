package command

import (
	"fmt"
	"strings"

	"github.com/google/go-github/github"
)

type ListCommand struct {
	Meta
}

func (c *ListCommand) Run(args []string) int {
	// Write your code here
	client := github.NewClient(nil)

	opt := &github.ListOptions{
		Page: 2, PerPage: 2,
	}

	event, _, err := client.Activity.ListEventsPerformedByUser("tjinjin", true, opt)
	fmt.Println(event)
	fmt.Println(err)

	return 0
}

func (c *ListCommand) Synopsis() string {
	return "list"
}

func (c *ListCommand) Help() string {
	helpText := `

`
	return strings.TrimSpace(helpText)
}
