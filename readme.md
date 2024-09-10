# Mawhrin

[![](https://img.shields.io/pypi/v/mawhrin?style=flat-square)][pp]
[![](https://img.shields.io/github/last-commit/stvmln86/mawhrin?style=flat-square)][ch]
[![](https://img.shields.io/github/issues/stvmln86/mawhrin?style=flat-square)][is]
[![](https://img.shields.io/pypi/l/mawhrin?style=flat-square)][li]
[![](https://img.shields.io/pypi/pyversions/mawhrin?style=flat-square)][py]
[![](https://img.shields.io/badge/justforfunnoreally-dev-9ff?style=flat-square)][ff]


**Mawhrin** is a command-line note file manager, written in [Python 3.11][py] by [Stephen Malone][sm]. If you have a directory of plaintext note files, Mawhrin can give you a neat & tidy command-line interface to access, edit and organise them.

## Installation

You can install Mawhrin using Pip:

```
pip install mawhrin
```

Or just download the [latest release][rl] for your platform.

## Configuration

Mawhrin always operates within a directory of plaintext note files you specify, using two environment variables:

```bash
# The directory your note files are in.
export MAWHRIN_DIR = "$HOME/notes"

# The extension your note files use (with a leading dot).
export MAWHRIN_EXT = ".txt"
```

That's it. That's all you need to do.

## Commands

### Basic syntax

Your notes are always shown as lowercase names, e.g.: Mawhrin will show the note file `notes/foo.txt` as `foo`. The opposite is also true: if you ask Mawhrin to create the note `foo`, it will translate to `notes/foo.txt`.

For command-line help, call Mawhrin or a command with the the `-h` or `--help` options.

### List all notes

Use `list` to see all your notes, or notes with names matching some text:

```bash
$ mawhrin list
books_to_read
recipes_pasta
recipes_vegan
todos-2024

$ mawhrin list "2024"
todos-2024
```

- Matching uses case-insensitive substrings, so `TODO` will match `todos-2024`.
- If no matches are found, Mawhrin will print nothing.

### Search all notes

Use `find` to list notes containing some text:

```bash
$ mawhrin find "add the tofu"
recipes_vegan
```

- Searching also uses case-insensitive substrings.
- If no searches are found, Mawhrin will print nothing.

### Create a note

Use `make` to create a new empty note, or a note containing some text:

```bash
$ mawhrin make movies_to_watch
# or to add text:
$ mawhrin make movies_to_watch "- [ ] Contact"
```

- If the note already exists, Mawhrin will print an error.

### Open a note

Use `open` to open a new or existing note in your default editor:

```bash
$ mawhrin open todos-2024
```

- Your editor is determined by the environment variables `$EDITOR` or `$VISUAL`.

### Print a note

Use `show` to print the contents of a note, if it exists:

```
$ mawhrin show books_to_read
- [x] The Player of Games by Iain M. Banks
```

- If the note doesn't exist, Mawhrin will print nothing.

## Contributing

Please submit all bug reports and feature requests to the [issue tracker][is], thank you.

[ch]: https://github.com/stvmln86/mawhrin/blob/main/changes.md
[ff]: https://justforfunnoreally.dev
[is]: https://github.com/stvmln86/mawhrin/issues
[li]: https://github.com/stvmln86/mawhrin/blob/main/license.md
[rl]: https://github.com/stvmln86/mawhrin/releases/latest
[sm]: https://github.com/stvmln86
[pp]: https://pypi.org/project/mawhrin/
[py]: https://www.python.org/downloads/release/python-3110/
