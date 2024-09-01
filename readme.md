# MAWHRIN

**Mawhrin** is a command-line note file manager, written in [Go 1.22][go] by [Stephen Malone][sm].

## Installation

You can install Mawhrin using your Go tools...

```
go get github.com/stvmln86/mawhrin
```

... or download the [latest release][rl] for your platform.

## Configuration

Mawhrin will automatically create a database file at one of three locations, depending on how your environment variables are set:

Variable                | Path
----------------------- | ----
`MAWHRIN_DB`            | `$MAWHRIN_DB`
`XDG_CONFIG_HOME`       | `$XDG_CONFIG_HOME/mawhrin/mawhrin.db`
`HOME` or `USERPROFILE` | `$HOME/.mawhrin`

## Commands

### `list [TEXT]`

List all notes, or notes with names containing `TEXT`.

### `find [TEXT]`

List all notes, or notes with bodies containing `TEXT`.

### `make NAME [BODY]`

Create `NOTE`, either empty or containing `BODY`.

### `show NOTE`

Print `NOTE`'s body, if it exists.

### `edit NOTE`

Open `NOTE` in `$EDITOR` or `$VISUAL`.

### `kill NOTE`

Delete `NOTE`, if it exists.

## Contributing

Please submit all bug reports and feature requests to the [issue tracker][is], thank you.

[go]: https://golang.org/doc/go1.22
[is]: https://github.com/stvmln86/mawhrin/issues
[rl]: https://github.com/stvmln86/mawhrin/releases/latest
[sm]: https://github.com/stvmln86
