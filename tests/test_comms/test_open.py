"""
Tests for 'mawhrin.comms.open'.
"""

import os


def test_list(capfd, run):
    # setup
    os.environ["EDITOR"] = "echo"
    os.environ["VISUAL"] = "echo"

    # success
    dire, _ = run("open", "alpha")
    path = os.path.join(dire, "alpha.extn")
    assert capfd.readouterr()[0] == f"{path}\n"
