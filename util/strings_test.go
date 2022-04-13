package util

import (
	"strings"
	"testing"

	"github.com/mayudev/animethemes-cli/api"
)

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
				Overlap: "Over",
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
