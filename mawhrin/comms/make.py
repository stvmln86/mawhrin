"""
Command definition for 'make'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(short_help="Create a note.")
@click.argument("name", metavar="NOTE")
@click.argument("TEXT", default="")
@click.pass_obj
def make(book: Book, name: str, text: str):
    """
    Create NOTE as an empty file or containing TEXT.
    """

    note = book.create(name)
    note.write(text if text else "")
