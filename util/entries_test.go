package util

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mayudev/animethemes-cli/api"
)

var mockVideo = api.Video{
	ID:         0,
	Basename:   "a",
	Size:       1,
	Resolution: 1080,
	NC:         false,
	Subbed:     true,
	Lyrics:     true,
	Uncensored: false,
	Source:     "BD",
	Overlap:    "None",
	Tags:       "",
	Link:       "",
}

type mockPlayer struct{}

func (mockPlayer) Play(v *api.Video) {
	fmt.Println(v.Basename)
}
func TestAskVideos(t *testing.T) {
	rescue := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	Asker{
		Player: mockPlayer{},
	}.AskVideos([]api.Video{mockVideo, mockVideo})

	w.Close()
	out, _ := ioutil.ReadAll(r)
	os.Stdout = rescue

	output := string(out)

	if !strings.Contains(output, mockVideo.Source) ||
		!strings.Contains(output, strconv.Itoa(mockVideo.Resolution)) ||
		!strings.Contains(output, "Lyrics") {
		t.Fail()
	}
}
