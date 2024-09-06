"""
Command definition for 'show'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command()
@click.argument("NAME")
@click.pass_obj
def show(book: Book, name: str):
    """
    Print the contents of an existing note.
    """

    if note := book.get(name):
        click.echo(note.read(), nl=False)
