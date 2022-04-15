package api

const (
	API_ENDPOINT = "https://staging.animethemes.moe"
	RESOURCE_URL = "https://animethemes.moe/video/"
)

type Timestamps struct {
	// ISO strings
	CreatedAt string `json:"created_at"`
	UpdatedAt string `json:"updated_at"`
	DeletedAt string `json:"deleted_at,omitempty"`
}

type Anime struct {
	ID       int          `json:"id"`
	Name     string       `json:"name"`
	Slug     string       `json:"slug"`
	Year     int          `json:"year,omitempty"`
	Season   string       `json:"season,omitempty"`
	Synopsis string       `json:"synopsis,omitempty"`
	Themes   []AnimeTheme `json:"animethemes,omitempty"`

	Timestamps
}

type AnimeTheme struct {
	ID       int               `json:"id"`
	Type     string            `json:"type,omitempty"`
	Sequence int               `json:"sequence,omitempty"`
	Group    string            `json:"group,omitempty"`
	Slug     string            `json:"slug"`
	Entries  []AnimeThemeEntry `json:"animethemeentries,omitempty"`
	Song     Song              `json:"song"`

	Timestamps
}

type AnimeThemeEntry struct {
	ID       int     `json:"id"`
	Version  int     `json:"version,omitempty"`
	Episodes string  `json:"episodes,omitempty"`
	Nsfw     bool    `json:"nsfw"`
	Spoiler  bool    `json:"spoiler"`
	Notes    string  `json:"notes,omitempty"`
	Videos   []Video `json:"videos,omitempty"`

	Timestamps
}

type Song struct {
	ID    int    `json:"id"`
	Title string `json:"title,omitempty"`
	As    string `json:"string,omitempty"`

	Timestamps
}

type Video struct {
	ID         int    `json:"id"`
	Basename   string `json:"basename"`
	Size       int    `json:"size"`
	Resolution int    `json:"resolution,omitempty"`
	NC         bool   `json:"nc"`
	Subbed     bool   `json:"subbed"`
	Lyrics     bool   `json:"lyrics"`
	Uncensored bool   `json:"uncensored"`
	Source     string `json:"source,omitempty"`
	Overlap    string `json:"overlap"`
	Tags       string `json:"tags"`
	Link       string `json:"link"`
}

type Links struct {
	First string `json:"first"`
	Last  string `json:"last"`
	Prev  string `json:"prev"`
	Next  string `json:"next"`
}

type Meta struct {
	CurrentPage int `json:"current_page"`
	From        int `json:"from"`
	PerPage     int `json:"per_page"`
	To          int `json:"to"`
	Total       int `json:"total"`
}

type AnimeSearch struct {
	Anime []Anime `json:"anime"`
	Links Links   `json:"links"`
	Meta  Meta    `json:"meta"`
}
