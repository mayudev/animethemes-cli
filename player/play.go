package player

import (
	"bytes"
	"io"
	"os"
	"os/exec"

	"github.com/mayudev/animethemes-cli/api"
	"github.com/spf13/cobra"
)

var Player string

func PlayVideo(video *api.Video) {
	url := api.RESOURCE_URL + video.Basename

	cmd := exec.Command(Player, url)

	// Print player output to console, because why not
	var stdBuffer bytes.Buffer
	mw := io.MultiWriter(os.Stdout, &stdBuffer)

	cmd.Stdout = mw
	cmd.Stderr = mw

	err := cmd.Run()
	cobra.CheckErr(err)
}
