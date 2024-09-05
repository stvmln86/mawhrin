"""
Command definition for 'list'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(name="list")
@click.argument("TEXT", default="")
@click.pass_obj
def list_(book: Book, text: str):
    """
    List all notes, or notes with names containing TEXT.
    """

    for note in book:
        if note.match(text):
            click.echo(note.name)
