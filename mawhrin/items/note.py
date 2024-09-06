"""
Class definition for 'Note'.
"""

from collections.abc import Iterator

from mawhrin import tools


class Note:
    """
    A single plaintext note file in a directory.
    """

    def __init__(self, path: str):
        """
        Initialise the Note.
        """

        self.path = path

    def __eq__(self, note: object) -> bool:
        """
        Return True if the Note is equal to another Note.
        """

        if not isinstance(note, Note):
            return NotImplemented
        else:
            return self.path == note.path

    def __hash__(self) -> int:
        """
        Return the Note's unique hashing integer.
        """

        return hash("Note:" + self.path)

    def __iter__(self) -> Iterator[str]:
        """
        Yield each line of the Note's body as strings.
        """

        yield from self.read().splitlines(keepends=True)

    def __repr__(self) -> str:
        """
        Return the Note as a code-representative string.
        """

        return f"Note({self.path!r})"

    def __str__(self) -> str:
        """
        Return the Note as a descriptive string.
        """

        return self.name

    def delete(self):
        """
        Delete the Note by renaming it to the ".trash" extension.
        """

        tools.file.delete(self.path)

    def exists(self) -> bool:
        """
        Return True if the Note's file exists.
        """

        return tools.file.exists(self.path)

    def match(self, text: str) -> bool:
        """
        Return True if the Note's name contains a case-insensitive substring.
        """

        return tools.path.match(self.path, text)

    def read(self) -> str:
        """
        Return the Note's body as a string.
        """

        body = tools.file.read(self.path)
        return tools.data.body(body)

    def rename(self, name: str):
        """
        Rename the Note to a new name with the same directory and extension.
        """

        name = tools.data.name(name)
        tools.file.rename(self.path, name)

    def search(self, text: str) -> bool:
        """
        Return True if the Note's body contains a case-insensitive substring.
        """

        return tools.file.search(self.path, text)

    @property
    def name(self) -> str:
        """
        Return the Note's base name without the extension.
        """

        name = tools.path.name(self.path)
        return tools.data.name(name)

    def write(self, body: str):
        """
        Overwrite the Note's body with a string.
        """

        body = tools.data.body(body)
        tools.file.write(self.path, body)
