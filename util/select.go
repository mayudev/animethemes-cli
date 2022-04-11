package util

import (
	"fmt"

	"github.com/AlecAivazis/survey/v2"
	"github.com/spf13/cobra"
)

func ShowSelection(title string, items []string) {
	prompt := &survey.Select{
		Message:  title,
		Options:  items,
		PageSize: 10,
	}

	var index int

	err := survey.AskOne(prompt, &index)

	cobra.CheckErr(err)

	fmt.Println(index)
}
