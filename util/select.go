package util

import (
	"io"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

// SimpleSelection shows a choice of items and returns the index of the selected item
func SimpleSelection(title string, items []string) int {
	prompt := &survey.Select{
		Message:  title,
		Options:  items,
		PageSize: 10,
	}

	var index int

	err := survey.AskOne(prompt, &index)

	if err != nil {
		if err == io.EOF {
			return 0
		} else {
			cobra.CheckErr(err)
		}
	}

	return index
}
