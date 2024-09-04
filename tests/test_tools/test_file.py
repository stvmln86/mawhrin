"""
Tests for 'mawhrin.tools.file'.
"""

import os.path
from mawhrin.tools import file


def test_delete(path):
    # setup
    dest = path.replace(".extn", ".trash")

    # success
    file.delete(path)
    assert not os.path.isfile(path)
    assert os.path.isfile(dest)


def test_exists(path):
    # setup
    dire = os.path.dirname(path)

    # success - True
    assert file.exists(path)
    assert file.exists(dire, dire=True)

    # success - False
    assert not file.exists("/nope")
    assert not file.exists("/nope", dire=True)


def test_read(path):
    # success
    assert file.read(path) == "Alpha.\n"


def test_rename(path):
    # setup
    dest = path.replace("alpha", "rename")

    # success
    file.rename(path, "rename")
    assert not os.path.isfile(path)
    assert os.path.isfile(dest)


def test_search(path):
    # success
    assert file.search(path, "ALPHA")
    assert not file.search(path, "NOPE")


def test_write(path):
    # success
    file.write(path, "Write.\n")
    assert open(path).read() == "Write.\n"
