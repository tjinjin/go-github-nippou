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
	orgs, _, err := client.Organizations.List("tjinjin", nil)
	fmt.Println(orgs)
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
