# Mawhrin

**Mawhrin** is a command-line note file manager, written in [Python 3.11][py] by [Stephen Malone][sm]. It provides a simple, universal command-line interface to your notes, no matter where you are in the terminal.

## Installation

You can install Mawhrin using Pip...

```
pip install mawhrin
```

... or download the [latest release][rl] for your platform.

## Commands

### `list [TEXT]`

List all notes, or notes with names containing `TEXT`.

<details><summary>Example:</summary>

```
$ mawhrin list
alpha
bravo
charlie

$ mawhrin list ch
charlie
```

</details>

## Contributing

Please submit all bug reports and feature requests to the [issue tracker][is], thank you. For developers, run `pip install mawhrin[test]` to install testing dependencies.

[is]: https://github.com/stvmln86/mawhrin/issues
[rl]: https://github.com/stvmln86/mawhrin/releases/latest
[sm]: https://github.com/stvmln86
[py]: https://www.python.org/downloads/release/python-3110/
