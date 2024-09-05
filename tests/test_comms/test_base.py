"""
Tests for 'mawhrin.comms.base'.
"""

import click

from mawhrin.comms.base import group


@group.command()
@click.pass_obj
def mock(book):
    for note in book:
        click.echo(note)


def test_group(run):
    # success
    rslt = run("mock")
    assert rslt.output == "alpha\nbravo\n"
