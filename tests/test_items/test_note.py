"""
Tests for 'mawhrin.items.note'.
"""

import os.path

import pytest

from mawhrin.items.note import Note


@pytest.fixture(scope="function")
def note(path):
    return Note(path)


def test_init(note):
    # success
    assert note.path


def test_eq(note):
    # success
    assert note == note
    assert note != Note("/nope")
    assert note != object()


def test_hash(note):
    # success
    assert hash(note)


def test_iter(note):
    # success
    assert list(note) == ["Alpha.\n"]


def test_repr(note):
    # setup
    note.path = "alpha.extn"

    # success
    assert repr(note) == "Note('alpha.extn')"


def test_str(note):
    # success
    assert str(note) == "alpha"


def test_delete(note):
    # success
    note.delete()
    assert not os.path.isfile(note.path)


def test_exists(note):
    # success
    assert note.exists()


def test_match(note):
    # success
    assert note.match("ALPHA")


def test_read(note):
    # success
    assert note.read() == "Alpha.\n"


def test_rename(note):
    # setup
    dest = note.path.replace("alpha", "rename")

    # success
    note.rename("rename")
    assert os.path.isfile(dest)
    assert not os.path.isfile(note.path)


def test_search(note):
    # success
    assert note.search("ALPHA")


def test_name(note):
    # success
    assert note.name == "alpha"


def test_write(note):
    # success
    note.write("Write.\n")
    assert open(note.path).read() == "Write.\n"
