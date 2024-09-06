"""
Tests for 'mawhrin.comms.show'.
"""


def test_show(run):
    # success
    _, rslt = run("show", "alpha")
    assert rslt.output == "Alpha.\n"
