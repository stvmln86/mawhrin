"""
Class definition for 'Book'.
"""

from collections.abc import Callable, Iterator

from mawhrin import tools
from mawhrin.items.note import Note


class Book:
    """
    A directory containing multiple plaintext note files.
    """

    def __init__(self, dire: str, extn: str):
        """
        Initialise the Book.
        """

        self.dire = tools.data.path(dire)
        self.extn = tools.data.extn(extn)

    def __contains__(self, name: str) -> bool:
        """
        Return True if the Book contains a named Note.
        """

        return self.get(name).exists()

    def __eq__(self, book: object) -> bool:
        """
        Return True if the Book is equal to another Book.
        """

        if not isinstance(book, Book):
            return NotImplemented
        else:
            return self.dire == book.dire and self.extn == book.extn

    def __getitem__(self, name: str) -> Note:
        """
        Return an existing Note in the Book by name.
        """

        note = self.get(name)
        if note.exists():
            return note
        else:
            raise KeyError(f"Note {note.name!r} does not exist")

    def __iter__(self) -> Iterator[Note]:
        """
        Yield each Note in the Book in alphabetical order.
        """

        for path in tools.path.glob(self.dire, self.extn):
            yield Note(path)

    def __len__(self) -> int:
        """
        Return the number of Notes in the Book.
        """

        return len(list(tools.path.glob(self.dire, self.extn)))

    def __repr__(self) -> str:
        """
        Return the Book as a code-representative string.
        """

        return f"Book({self.dire!r}, {self.extn!r})"

    def exists(self) -> bool:
        """
        Return True if the Book's directory exists.
        """

        return tools.file.exists(self.dire, dire=True)

    def get(self, name: str) -> Note:
        """
        Return a new or existing Note in the Book.
        """

        path = tools.path.join(self.dire, tools.data.name(name), self.extn)
        return Note(path)
