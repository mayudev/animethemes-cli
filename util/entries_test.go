package util

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"testing"

	"github.com/mayudev/animethemes-cli/api"
)

// Mocks
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
	// Do nothing
}

func GenerateEntry(i int) api.AnimeThemeEntry {
	return api.AnimeThemeEntry{
		ID:       i,
		Version:  0,
		Episodes: "1",
		Nsfw:     false,
		Spoiler:  false,
		Notes:    "",
		Videos: []api.Video{
			mockVideo,
		},
		Timestamps: api.Timestamps{CreatedAt: "", DeletedAt: "", UpdatedAt: ""},
	}
}

type StdoutTest struct {
	rescue *os.File
	r      *os.File
	w      *os.File
}

func (t *StdoutTest) Setup() {
	rescue := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	t.rescue = rescue
	t.r = r
	t.w = w
}

func (t *StdoutTest) Teardown() string {
	t.w.Close()
	out, _ := ioutil.ReadAll(t.r)
	os.Stdout = t.rescue

	return string(out)
}

func ExpectOutput(t *testing.T, stdout string, expect string) {
	if !strings.Contains(stdout, expect) {
		t.Fail()
		t.Logf("Expected %v to be in output, was not found", expect)
	}
}

func TestInterface_AskVideos(t *testing.T) {
	rescue := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	// Create a mock Interface
	mockInterface := Interface{
		Player: mockPlayer{},
	}

	mockInterface.AskVideos([]api.Video{mockVideo, mockVideo})

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

func TestInterface_AskEntries(t *testing.T) {
	type args struct {
		entries []api.AnimeThemeEntry
	}
	tests := []struct {
		name string
		args args
		Run  func(i Interface, stdoutTester *StdoutTest, entries []api.AnimeThemeEntry)
	}{
		{
			name: "does not show selection when there's only one entry",
			args: args{
				entries: []api.AnimeThemeEntry{
					GenerateEntry(0),
				},
			},
			Run: func(i Interface, stdoutTester *StdoutTest, entries []api.AnimeThemeEntry) {
				i.AskEntries(entries)

				output := stdoutTester.Teardown()

				// Expect output to be empty
				if len(output) > 0 {
					t.Logf("Expected no output, got %v", output)
					t.Fail()
				}
			},
		},
		{
			name: "shows selection when there's more than one entry",
			args: args{
				entries: []api.AnimeThemeEntry{
					GenerateEntry(0),
					GenerateEntry(1),
				},
			},
			Run: func(i Interface, stdoutTester *StdoutTest, entries []api.AnimeThemeEntry) {
				i.AskEntries(entries)

				output := stdoutTester.Teardown()

				ExpectOutput(t, output, entries[0].Episodes)
				ExpectOutput(t, output, entries[1].Episodes)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Interface{
				Player: mockPlayer{},
			}

			stdoutTester := StdoutTest{}
			stdoutTester.Setup()

			tt.Run(a, &stdoutTester, tt.args.entries)
		})
	}
}
