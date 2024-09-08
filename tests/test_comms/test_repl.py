"""
Tests for 'mawhrin.comms.repl'.
"""


def test_repl(run):
    # success - with argument
    _, rslt = run("repl", "1 + 1")
    assert rslt.output == "2\n"

    # cannot test main repl loop
