"""
Tests for 'mawhrin.comms.list'.
"""


def test_list(run):
    # success - no text
    rslt = run("list")
    assert rslt.output == "alpha\nbravo\n"

    # success - with text
    rslt = run("list", "ALPHA")
    assert rslt.output == "alpha\n"
