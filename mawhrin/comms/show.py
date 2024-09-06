"""
Command definition for 'show'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(short_help="Print a note.")
@click.argument("name", metavar="NOTE")
@click.pass_obj
def show(book: Book, name: str):
    """
    Print NOTE, if it exists.
    """

    if note := book.get(name):
        click.echo(note.read(), nl=False)
