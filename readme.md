# Mawhrin

**Mawhrin** is a command-line note file manager, written in [Python 3.11][py] by [Stephen Malone][sm]. If you have a directory of plaintext note files, Mawhrin can give you a neat & tidy command-line interface to access, edit and organise them.

## Installation

You can install Mawhrin using Pip:

```
pip install mawhrin
```

Or just download the [latest release][rl] for your platform.

## Configuration

Mawhrin always operates within a directory of plaintext note files you specify with two environment variables:

```bash
# The directory your note files are in.
export MAWHRIN_DIR = "$HOME/notes"

# The extension your note files use (with a leading dot).
export MAWHRIN_EXT = ".txt"
```

That's it. That's all you need to do. 

## Commands

### `list [NAME]`

List all notes, or notes matching NAME.

<details><summary>Example:</summary>

```
$ mawhrin list
books_to_read
recipes_pasta
recipes_vegan
todos-2024

$ mawhrin list 2024
todos-2024
```

</details>

### `find TEXT`

List all notes containing `TEXT`.

<details><summary>Example:</summary>

```
$ mawhrin find broccoli 
recipes_pasta
recipes_vegan
```

</details>

### `make NOTE [TEXT]`

Create `NOTE` as an empty file or containing `TEXT`.

<details><summary>Example:</summary>

```
$ mawhrin make todos_2025 "- [ ] Import old todos."
$ mawhrin show todos_2025
- [ ] Import old todos.
```

</details>

### `show [NOTE]`

Print NOTE, if it exists.

<details><summary>Example:</summary>

```
$ mawhrin show books_to_read
- [x] The Player of Games by Iain M. Banks
...
```

</details>

## Contributing

Please submit all bug reports and feature requests to the [issue tracker][is], thank you.  

[is]: https://github.com/stvmln86/mawhrin/issues
[rl]: https://github.com/stvmln86/mawhrin/releases/latest
[sm]: https://github.com/stvmln86
[py]: https://www.python.org/downloads/release/python-3110/
