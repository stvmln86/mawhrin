"""
Low-level data generation & sanitisation functions.
"""

import os.path
from string import ascii_lowercase, digits

NAME_CHARS = ascii_lowercase + digits + "-_"


def body(body: str) -> str:
    """
    Return a trimmed note body string with an ending newline.
    """

    return body.strip() + "\n"


def extn(extn: str) -> str:
    """
    Return a trimmed lowercase file extension string with a leading dot.
    """

    extn = extn.strip().lower()
    return "." + extn.lstrip(".")


def name(name: str) -> str:
    """
    Return a trimmed lowercase alphanumeric note name string.
    """

    return "".join(char for char in name.strip().lower() if char in NAME_CHARS)


def path(path: str) -> str:
    """
    Return a trimmed normalised path string.
    """

    return os.path.normpath(path.strip())
