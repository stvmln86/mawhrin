"""
Tests for 'mawhrin.tools.path'.
"""

import os.path

from mawhrin.tools import path


def test_dire():
    # success
    assert path.dire("/dire/name.extn") == "/dire"


def test_extn():
    # success
    assert path.extn("/dire/name.extn") == ".extn"


def test_glob(dire):
    # success
    assert list(path.glob(dire, ".extn")) == [
        os.path.join(dire, "alpha.extn"),
        os.path.join(dire, "bravo.extn"),
    ]


def test_join():
    # success
    assert path.join("/dire", "name", ".extn") == "/dire/name.extn"


def test_match():
    # success
    assert path.match("/dire/name.extn", "NAME")
    assert not path.match("/dire/name.extn", "NOPE")


def test_name():
    # success
    assert path.name("/dire/name.extn") == "name"
