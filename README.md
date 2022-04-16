# animethemes-cli

A CLI for animethemes.moe

## Installation

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
mv animethemes /usr/local/bin/
```

## Usage

Use `animethemes` to show all commands with explanation.

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
