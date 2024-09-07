"""
Command definition for 'repl'.
"""

import readline
from code import InteractiveConsole

import click

from mawhrin import VERSION_TEXT
from mawhrin.comms.base import group
from mawhrin.items import Book


@group.command(hidden=True)
@click.argument("code", default="")
@click.pass_obj
def repl(book: Book, code: str):
    """
    Open the notes directory in an interactive REPL.
    """

    cons = InteractiveConsole(locals={"book": book})

    if code:
        cons.runsource(code)
    else:
        print(VERSION_TEXT)
        print(">>> book")
        cons.runsource("book")
        cons.interact("", exitmsg="")
