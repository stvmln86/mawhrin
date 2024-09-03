# Mawhrin

**Mawhrin** is a command-line note file manager, written in [Python 3.12][py] by [Stephen Malone][sm].

## Installation

You can install Mawhrin using Pip...

```
pip install mawhrin
```

... or download the [latest release][rl] for your platform.

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

[is]: https://github.com/stvmln86/mawhrin/issues
[rl]: https://github.com/stvmln86/mawhrin/releases/latest
[sm]: https://github.com/stvmln86
[py]: https://www.python.org/downloads/release/python-3120/
