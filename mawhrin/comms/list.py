"""
Command definition for 'list'.
"""

import click

from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(name="list", short_help="List all notes.")
@click.argument("NAME", default="")
@click.pass_obj
def list_(book: Book, name: str):
    """
    List all notes, or notes matching NAME.
    """

    for note in book:
        if note.match(name):
            click.echo(note.name)
