# -*- coding: utf-8 -*-

"""Console script for polyenv."""
import sys
import click

from polyenv import utils



@click.command()
@click.argument('language')
@click.argument('file_path', type=click.Path(exists=True))
def main(language, file_path):
    """
    Language -> Language which will be used to run the file
    File Path -> "Path to the location of the file"

    """
    # TODO: Get extension and match language
    utils.run_programming_language(language, file_path)




if __name__ == "__main__":
    sys.exit(main())  # pragma: no cover
