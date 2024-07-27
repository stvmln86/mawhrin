# Mawhrin

**Mawhrin** is a command-line note file manager, written in Go 1.22 by Stephen Malone.

## Ideas

- Write entire app in pure Go (only dependency is `testify`).
- Use `orig` & `dest` path variable names to make room for `tools/path` package.
- Use `tools/meta` package to `Split`, `Parse` and `Render` custom metadata format.
- Just use `$MAWHRIN_DIR` envvar and match all text extensions (`.md`, `.txt`, etc).
