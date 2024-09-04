"""
Tests for 'mawhrin.tools.data.'.
"""

from mawhrin.tools import data


def test_body():
    # success
    assert data.body("\tBody.\n\n") == "Body.\n"


def test_extn():
    # success
    assert data.extn("\t.EXTN\n") == ".extn"


def test_name():
    # success
    assert data.name("\tNAME123!!!\n") == "name123"


def test_path():
    # success
    assert data.path("\t/dire/./name.extn\n") == "/dire/name.extn"
