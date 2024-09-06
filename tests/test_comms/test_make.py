"""
Tests for 'mawhrin.comms.make'.
"""

import pytest
from click import ClickException

import os.path


def test_make(run):
    # success - no text
    dire, rslt = run("make", "delta")
    path = os.path.join(dire, "delta.extn")
    assert open(path).read() == "\n"

    # success - with text
    _, rslt = run("make", "echo", "Body.")
    path = os.path.join(dire, "echo.extn")
    assert open(path).read() == "Body.\n"

    # error - note already exists
    _, rslt = run("make", "echo")
    assert rslt.output == "Error: note 'echo' already exists.\n"
