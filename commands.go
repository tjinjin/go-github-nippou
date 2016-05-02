package main

import (
	"github.com/mitchellh/cli"
	"github.com/tjinjin/go-github-nippou/command"
)

func Commands(meta *command.Meta) map[string]cli.CommandFactory {
	return map[string]cli.CommandFactory{
		"list": func() (cli.Command, error) {
			return &command.ListCommand{
				Meta: *meta,
			}, nil
		},

		"version": func() (cli.Command, error) {
			return &command.VersionCommand{
				Meta:     *meta,
				Version:  Version,
				Revision: GitCommit,
				Name:     Name,
			}, nil
		},
	}
}
