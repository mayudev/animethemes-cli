package util

import (
	"strings"
	"testing"

	"github.com/mayudev/animethemes-cli/api"
)

func TestBuildThemeString(t *testing.T) {
	tests := []struct {
		name  string
		theme *api.AnimeTheme
		check func(output string)
	}{
		{
			name: "displays information about a theme",
			theme: &api.AnimeTheme{
				Slug: "Slug",
				Song: api.Song{
					Title: "Song Title",
				},
			},
			check: func(output string) {
				if !strings.Contains(output, "Slug") {
					t.Errorf("BuildThemeString does not contain Slug")
				}

				if !strings.Contains(output, "Song Title") {
					t.Errorf("BuildThemeString does not contain song title")
				}
			},
		},
		{
			name: "displays information about video when there's only one",
			theme: &api.AnimeTheme{
				Slug: "Slug",
				Song: api.Song{
					Title: "Song Title",
				},
				Entries: []api.AnimeThemeEntry{
					{
						Episodes: "1-16",
						Nsfw:     false,
						Spoiler:  false,
						Videos: []api.Video{
							{
								ID:         0,
								Resolution: 1080,
								Source:     "BD",
								NC:         false,
								Lyrics:     false,
								Overlap:    "None",
							},
						},
					},
				},
			},
			check: func(output string) {
				if !strings.Contains(output, "1080") {
					t.Errorf("BuildThemeString does not contain information about video even though there's only one video entry")
				}
			},
		},
		{
			name: "does not display information about video when there's more than one",
			theme: &api.AnimeTheme{
				Slug: "Slug",
				Song: api.Song{
					Title: "Song Title",
				},
				Entries: []api.AnimeThemeEntry{
					{
						Episodes: "1-16",
						Nsfw:     false,
						Spoiler:  false,
						Videos: []api.Video{
							{
								ID:         0,
								Resolution: 1080,
								Source:     "BD",
								NC:         false,
								Lyrics:     false,
								Overlap:    "None",
							},
						},
					},
					{
						Episodes: "17-24",
						Nsfw:     false,
						Spoiler:  true,
						Videos: []api.Video{
							{
								ID:         1,
								Resolution: 1080,
								Source:     "BD",
								NC:         false,
								Lyrics:     false,
								Overlap:    "None",
							},
						},
					},
				},
			},
			check: func(output string) {
				if strings.Contains(output, "1080") {
					t.Errorf("BuildThemeString contains video information even though there's more than one")
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildThemeString(tt.theme)

			tt.check(got)
		})
	}
}

func TestBuildEntryChoiceString(t *testing.T) {
	tests := []struct {
		name   string
		entry  *api.AnimeThemeEntry
		want   []string
		strict bool
	}{
		{
			name: "only episode range",
			entry: &api.AnimeThemeEntry{
				Episodes: "1-2",
			},
			want:   []string{"1-2"},
			strict: false,
		},
		{
			name: "shows NSFW label",
			entry: &api.AnimeThemeEntry{
				Nsfw: true,
			},
			want:   []string{"NSFW"},
			strict: false,
		},
		{
			name: "shows Spoiler label",
			entry: &api.AnimeThemeEntry{
				Spoiler: true,
			},
			want:   []string{"Spoiler"},
			strict: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildEntryChoiceString(tt.entry)

			if tt.strict {
				if got != tt.want[0] {
					t.Errorf("BuildEntryChoiceString() = %v, want %v", got, tt.want[0])
				}

			} else {
				for _, v := range tt.want {
					if !strings.Contains(got, v) {
						t.Errorf("BuildEntryChoiceString() = %v, want %v", got, v)
					}
				}
			}
		})
	}
}

func TestBuildVideoChoiceString(t *testing.T) {
	tests := []struct {
		name  string
		video *api.Video
		want  string
	}{
		{
			name: "displays resolution",
			video: &api.Video{
				Resolution: 1080,
			},
			want: "1080",
		},
		{
			name: "displays source",
			video: &api.Video{
				Source: "BD",
			},
			want: "BD",
		},
		{
			name: "displays Creditless",
			video: &api.Video{
				NC: true,
			},
			want: "Creditless",
		},
		{
			name: "displays Lyrics",
			video: &api.Video{
				Lyrics: true,
			},
			want: "Lyrics",
		},
		{
			name: "displays Transition",
			video: &api.Video{
				Overlap: "Transition",
			},
			want: "Transition",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := BuildVideoChoiceString(tt.video)

			if !strings.Contains(got, tt.want) {
				t.Errorf("BuildVideoChoiceString() = %v, want %v", got, tt.want)
			}
		})
	}
}
