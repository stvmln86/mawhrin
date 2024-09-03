"""
Tests for 'mawhrin.tools.path'.
"""

from mawhrin.tools import path


def test_dire():
    # success
    assert path.dire("/dire/slug.extn") == "/dire"


def test_glob(dire):
    # success
    assert set(path.glob(dire, ".extn")) == {
        dire + "/alpha.extn",
        dire + "/bravo.extn",
    }


def test_join():
    # success
    assert path.join("/dire", "slug", ".extn") == "/dire/slug.extn"


def test_match():
    # success
    assert path.match("/dire/slug.extn", "SLUG")
    assert not path.match("/dire/slug.extn", "NOPE")


def test_slug():
    # success
    assert path.slug("/dire/slug.extn") == "slug"
