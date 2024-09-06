"""
Tests for 'mawhrin.comms.list'.
"""


def test_list(run):
    # success - no text
    _, rslt = run("list")
    assert rslt.output == "alpha\nbravo\n"

    # success - with text
    _, rslt = run("list", "ALPHA")
    assert rslt.output == "alpha\n"
