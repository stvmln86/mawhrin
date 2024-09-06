"""
Command definition for 'find'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(short_help="Search all notes.")
@click.argument("TEXT")
@click.pass_obj
def find(book: Book, text: str):
    """
    List all notes containing TEXT.
    """

    for note in book:
        if note.search(text):
            click.echo(note.name)
