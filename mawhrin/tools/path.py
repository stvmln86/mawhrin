"""
Filepath derivation and manipulation functions.
"""

import glob as glob_
import os.path
from collections.abc import Iterator


def dire(path: str) -> str:
    """
    Return a path's parent directory.
    """

    return os.path.dirname(path)


def glob(dire: str, extn: str) -> Iterator[str]:
    """
    Yield all paths in a directory matching an extension.
    """

    path = os.path.join(dire, "*" + extn)
    yield from glob_.iglob(path)


def join(dire: str, slug: str, extn: str) -> str:
    """
    Return a joined filepath from a directory, slug and extension.
    """

    return os.path.join(dire, slug + extn)


def match(path: str, text: str) -> bool:
    """
    Return True if a path's slug contains a case-insensitive substring.
    """

    return text.casefold() in slug(path).casefold()


def slug(path: str) -> str:
    """
    Return a path's base name without the extension.
    """

    base = os.path.basename(path)
    return os.path.splitext(base)[0]
