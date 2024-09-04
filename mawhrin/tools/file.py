"""
Filesystem I/O functions.
"""

import os.path

from mawhrin import tools


def delete(path: str):
    """
    Delete a file by renaming it to the ".trash" extension.
    """

    dire = tools.path.dire(path)
    name = tools.path.name(path)
    dest = tools.path.join(dire, name, ".trash")
    os.rename(path, dest)


def exists(path: str, *, dire: bool = False) -> bool:
    """
    Return True if a file or directory exists.
    """

    if dire:
        return os.path.isdir(path)
    else:
        return os.path.isfile(path)


def read(path: str) -> str:
    """
    Return a file's body as a string.
    """

    with open(path, "r", encoding="utf-8") as fobj:
        return fobj.read()


def rename(path: str, name: str):
    """
    Rename a file to a new name with the same directory and extension.
    """

    dire = tools.path.dire(path)
    extn = tools.path.extn(path)
    dest = tools.path.join(dire, name, extn)
    os.rename(path, dest)


def search(path: str, text: str) -> bool:
    """
    Return True if a file's body contains a case-insensitive substring.
    """

    return text.casefold() in read(path).casefold()


def write(path: str, body: str):
    """
    Overwrite a new or existing file's body with a string.
    """

    with open(path, "w", encoding="utf-8") as fobj:
        fobj.write(body)
