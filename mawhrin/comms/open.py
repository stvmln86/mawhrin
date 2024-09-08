"""
Command definition for 'open'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(short_help="Open a note.")
@click.argument("name", metavar="NOTE")
@click.pass_obj
def open(book: Book, name: str):
    """
    Open NOTE in $EDITOR or $VISUAL.
    """

    note = book.create(name)
    click.edit(filename=note.path)
