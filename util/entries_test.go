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

var mockAnime = api.Anime{
	ID:       2955,
	Name:     "Toaru Kagaku no Railgun S",
	Slug:     "toaru_kagaku_no_railgun_s",
	Year:     2013,
	Season:   "Spring",
	Synopsis: " who cares",
	Themes: []api.AnimeTheme{
		{
			ID:       5287,
			Type:     "OP",
			Sequence: 1,
			Slug:     "OP1",
			Entries: []api.AnimeThemeEntry{
				{
					ID:       6052,
					Episodes: "1-16",
					Nsfw:     false,
					Spoiler:  false,
					Videos: []api.Video{
						{
							ID:         4755,
							Basename:   "ToaruKagakuNoRailgunS-OP1.webm",
							Size:       66364300,
							Resolution: 1080,
							NC:         true,
							Subbed:     false,
							Lyrics:     false,
							Uncensored: false,
							Source:     "BD",
							Overlap:    "None",
							Tags:       "NCBD1080",
							Link:       "no",
						},
					},
				},
			},
			Song: api.Song{
				ID:    5287,
				Title: "sister's noise",
			},
		},
		{
			ID:       5288,
			Type:     "OP",
			Sequence: 2,
			Slug:     "OP2",
			Entries: []api.AnimeThemeEntry{
				{
					ID:       6053,
					Episodes: "17-23",
					Nsfw:     false,
					Spoiler:  false,
					Videos: []api.Video{
						{
							ID:         4756,
							Basename:   "ToaruKagakuNoRailgunS-OP2.webm",
							Size:       66166179,
							Resolution: 1080,
							NC:         true,
							Subbed:     false,
							Lyrics:     false,
							Uncensored: false,
							Source:     "BD",
							Overlap:    "None",
							Tags:       "NCBD1080",
							Link:       "no",
						},
					},
				},
			},
			Song: api.Song{
				ID:    5288,
				Title: "eternal reality",
			},
		},
		{
			ID:       5289,
			Type:     "ED",
			Sequence: 1,
			Slug:     "ED1",
			Entries: []api.AnimeThemeEntry{
				{
					ID:       6054,
					Episodes: "2-4, 6-10, 12-13, 15-16",
					Nsfw:     false,
					Spoiler:  false,
					Videos: []api.Video{
						{
							ID:         4751,
							Basename:   "ToaruKagakuNoRailgunS-ED1.webm",
							Size:       48881436,
							Resolution: 1080,
							NC:         true,
							Subbed:     false,
							Lyrics:     false,
							Uncensored: false,
							Source:     "BD",
							Overlap:    "None",
							Tags:       "NCBD1080",
							Link:       "no",
						},
					},
				},
			},
			Song: api.Song{
				ID:    5289,
				Title: "Grow Slowly",
			},
		},
		{
			ID:       5290,
			Type:     "ED",
			Sequence: 2,
			Slug:     "ED2",
			Entries: []api.AnimeThemeEntry{
				{
					ID:       6055,
					Episodes: "11, 14",
					Nsfw:     false,
					Spoiler:  false,
					Videos: []api.Video{
						{
							ID:         4752,
							Basename:   "ToaruKagakuNoRailgunS-ED2.webm",
							Size:       20,
							Resolution: 1080,
							NC:         true,
							Subbed:     false,
							Lyrics:     false,
							Uncensored: false,
							Source:     "BD",
							Overlap:    "None",
							Tags:       "NCBD1080",
							Link:       "no",
						},
					},
				},
			},
			Song: api.Song{
				ID:    5290,
				Title: "stand still",
			},
		},
	},
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

func ExpectNoOutput(t *testing.T, stdout string, expect string) {
	if strings.Contains(stdout, expect) {
		t.Fail()
		t.Logf("Found %v in output, expected not to", expect)
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

func TestInterface_AskThemes(t *testing.T) {
	type args struct {
		anime api.Anime
	}
	tests := []struct {
		name  string
		flags Flags
		args  args
		Run   func(output string)
	}{
		{
			name:  "shows choice when no flag passed",
			flags: Flags{},
			args:  args{anime: mockAnime},
			Run: func(output string) {
				ExpectOutput(t, output, "OP1")
				ExpectOutput(t, output, "OP2")
				ExpectOutput(t, output, "ED1")
				ExpectOutput(t, output, "ED2")
			},
		},
		{
			name: "does not show ending themes when --openings flag passed",
			flags: Flags{
				OnlyOpenings: true,
				OnlyEndings:  false,
			},
			args: args{anime: mockAnime},
			Run: func(output string) {
				ExpectOutput(t, output, "OP1")
				ExpectOutput(t, output, "OP2")

				ExpectNoOutput(t, output, "ED1")
				ExpectNoOutput(t, output, "ED2")
			},
		},
		{
			name: "does not show choice when --op flag passed",
			flags: Flags{
				OpeningN: 2,
			},
			args: args{anime: mockAnime},
			Run: func(output string) {
				ExpectNoOutput(t, output, "OP1")
			},
		},
		{
			name: "does not show choice when --ed flag passed",
			flags: Flags{
				EndingN: 2,
			},
			args: args{anime: mockAnime},
			Run: func(output string) {
				ExpectNoOutput(t, output, "OP1")
			},
		},
		{
			name: "does not output anything to stdout when sequence number not found",
			flags: Flags{
				EndingN: 10,
			},
			args: args{anime: mockAnime},
			Run: func(output string) {
				ExpectOutput(t, output, "")
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			a := Interface{
				Player: mockPlayer{},
				Flags:  tt.flags,
				Exit:   func() {},
			}

			stdoutTester := StdoutTest{}
			stdoutTester.Setup()

			a.AskThemes(tt.args.anime)

			output := stdoutTester.Teardown()

			tt.Run(output)
		})
	}
}
