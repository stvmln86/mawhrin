"""
Tests for 'mawhrin.comms.find'.
"""


def test_find(run):
    # success
    _, rslt = run("find", "ALPHA")
    assert rslt.output == "alpha\n"
