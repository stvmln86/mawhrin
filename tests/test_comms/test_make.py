"""
Tests for 'mawhrin.comms.make'.
"""

import os.path


def test_make(run):
    # success - no text
    dire, rslt = run("make", "delta")
    path = os.path.join(dire, "delta.extn")
    assert open(path).read() == "\n"
