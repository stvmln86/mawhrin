"""
Global unit-testing fixturs.
"""

import pytest
from click.testing import CliRunner

from mawhrin.comms.base import group

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


@pytest.fixture(scope="function")
def path(tmp_path):
    """
    Return a temporary file populated from the 'alpha' entry of MOCK_FILES.
    """

    path = tmp_path.joinpath("alpha.extn")
    path.write_text(MOCK_FILES["alpha.extn"])
    return str(path)


@pytest.fixture(scope="function")
def run(dire):
    """
    Return a function that runs a command in the Click base group and returns the test
    directory and result.
    """

    def run(*args):
        envs = {"MAWHRIN_DIR": dire, "MAWHRIN_EXT": ".extn"}
        clir = CliRunner(env=envs)
        return dire, clir.invoke(group, args)

    return run
