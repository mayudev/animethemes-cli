# animethemes-cli

![GitHub Workflow Status](https://img.shields.io/github/workflow/status/mayudev/animethemes-cli/Go?style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/mayudev/animethemes-cli?style=flat-square)

CLI tool to find and listen to anime openings and endings from your terminal :notes: using [animethemes.moe](https://animethemes.moe/)

[Installation](#installation) | [Usage](#usage)

## Installation

Grab the latest [release](https://github.com/mayudev/animethemes-cli/releases/latest), make the file executable with `chmod +x animethemes` and copy it to either `/usr/bin/` or `~/.local/bin/` (if it's in your PATH)

By default, animethemes-cli uses `mpv` to play videos, but you can change that in [Configuration](#configuration)

### Build from source
```sh
# Make sure go is installed
go version
# (tested on 1.18)

# Clone this repository
git clone https://github.com/mayudev/animethemes-cli

# Compile the package
go build -ldflags "-s -w" -o animethemes

# Install the package (make sure ~/.local/bin/ is in path)
mv animethemes ~/.local/bin/
# Alternatively, if you want to install it for all users
sudo mv animethemes /usr/local/bin/
```

## Usage

Use `animethemes` to show all commands with explanations.

Flag `--player` (`-p`) allows you to specify a player to use. You can also do it through a config file.

Tip: `--player echo` will print the video URL to console instead of playing it.

### Anime

```
animethemes anime [query]
```

Note: you can omit `anime` and use just `animethemes [query]`

Looks for `[query]` in animethemes.moe database. For example, `animethemes anime toaru` will return all entries with Toaru in title.

#### Flags

- `-f` or `--forever` - asks for theme choice again once finished playing (note: do not use it with `--op` or `--ed`)
- `-p` or `--player` - specify player command to be used (default: `mpv`)
- `--openings` - show only openings
- `--endings` - show only ending themes
- `--op [n]` - automatically play opening `[n]`
- `--ed [n]` - automatically play ending `[n]`
- `-1` or `--first` - automatically selects the first entry if there's more than one result

For example, `animethemes railgun t --first --op 1` will choose `Toaru Kagaku no Railgun T` and OP1, `final phase` for you.

### Year

```
animethemes year [year]
```

Returns all anime released in year `[year]`

## Configuration

Config file for animethemes-cli is stored in `~/.config/animethemes-cli.yml`. You can use it to set your default player. For example:

```
player: vlc
```
