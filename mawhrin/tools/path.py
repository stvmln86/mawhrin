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


def extn(path: str) -> str:
    """
    Return a path's file extension, with a leading dot.
    """

    base = os.path.basename(path)
    return os.path.splitext(base)[-1]


def glob(dire: str, extn: str) -> Iterator[str]:
    """
    Yield all paths in a directory matching an extension in alphabetical order.
    """

    path = os.path.join(dire, "*" + extn)
    yield from sorted(glob_.glob(path))


def join(dire: str, name: str, extn: str) -> str:
    """
    Return a joined filepath from a directory, name and extension.
    """

    return os.path.join(dire, name + extn)


def match(path: str, text: str) -> bool:
    """
    Return True if a path's name contains a case-insensitive substring.
    """

    return text.casefold() in name(path).casefold()


def name(path: str) -> str:
    """
    Return a path's base name without the extension.
    """

    base = os.path.basename(path)
    return os.path.splitext(base)[0]
