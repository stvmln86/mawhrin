"""
Tests for 'mawhrin.items.book'.
"""

import os.path

import pytest

from mawhrin.items.book import Book
from mawhrin.items.note import Note


@pytest.fixture(scope="function")
def book(dire):
    return Book(dire, ".extn")


def test_init(book):
    # success
    assert book.dire
    assert book.extn == ".extn"


def test_contains(book):
    # success
    assert "alpha" in book
    assert "nope" not in book


def test_eq(book):
    # success
    assert book == book
    assert book != Book("/nope", ".nope")
    assert book != object()


def test_iter(book):
    # success
    assert list(book) == [
        Note(os.path.join(book.dire, "alpha.extn")),
        Note(os.path.join(book.dire, "bravo.extn")),
    ]


def test_len(book):
    # success
    assert len(book) == 2


def test_repr(book):
    # setup
    book.dire = "/dire"

    # success
    assert repr(book) == "Book('/dire', '.extn')"


def test_create(book):
    # success
    note = book.create("delta")
    assert note.path == os.path.join(book.dire, "delta.extn")


def test_exists(book):
    # success
    assert book.exists()


def test_get(book):
    # success - Note exists
    note = book.get("alpha")
    assert note.path == os.path.join(book.dire, "alpha.extn")

    # success - Note does not exist
    note = book.get("nope")
    assert note == None
