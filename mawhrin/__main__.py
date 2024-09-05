"""
Main program function.
"""

from mawhrin.comms import group


def main(args: list[str] | None = None):
    """
    Run the main Mawhrin program.
    """

    group.main(args=args)


if __name__ == "__main__":
    main()
