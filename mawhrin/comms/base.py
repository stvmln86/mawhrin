"""
Click base group definition.
"""

import click
from click import ClickException

from mawhrin.items import Book

SETTINGS = {
    "help_option_names": ["-h", "--help"],
}


@click.group(name="mawhrin", context_settings=SETTINGS)
@click.option("--dire", envvar="MAWHRIN_DIR", hidden=True)
@click.option("--extn", envvar="MAWHRIN_EXT", hidden=True)
@click.pass_context
def group(ctx: click.Context, dire: str, extn: str):
    """
    A plaintext note file manager in Python.
    """

    if not dire:
        raise ClickException("environment variable 'MAWHRIN_DIR' is unset")

    if not extn:
        raise ClickException("environment variable 'MAWHRIN_EXT' is unset")

    ctx.obj = Book(dire, extn)
