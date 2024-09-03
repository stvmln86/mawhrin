"""
Global unit-testing fixturs.
"""

import pytest

MOCK_FILES = {
    "alpha.extn": "Alpha.\n",
    "bravo.extn": "Bravo.\n",
    "charlie.trash": "Charlie.\n",
}


@pytest.fixture(scope="function")
def dire(tmp_path):
    """
    Return a temporary directory populated with MOCK_FILES.
    """

    for base, body in MOCK_FILES.items():
        tmp_path.joinpath(base).write_text(body)

    return str(tmp_path)
