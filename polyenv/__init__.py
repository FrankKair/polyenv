# coding=utf-8
__title__ = 'polyenv'
__author__ = 'Frank Kair'
__version__ = '0.1.0'
__copyright__ = '2018 Frank Kair'
__license__ = 'MIT'

import sys
from os import path
from pytio import Tio, TioRequest
# from .polyenv import run_programming_language


def run_programming_language(language, input_file):
    with open(input_file, 'r') as f:
        data = f.read()
        tio = Tio()
        request = TioRequest(lang=language, code=data)
        response = tio.send(request)
        if response.error:
            print(response.error)
            sys.exit(1)

        return response.result


def main(args):
    if len(args) < 3:
        print('Missing language and/or input file.')
        print('$ polyenv language path_to_file')
        sys.exit(1)

    # TODO: Get extension and match language
    language = args[1]
    input_file = args[2]
    output = run_programming_language(language, input_file)
    print(output)


def run_main():
    main(sys.argv)


if __name__ == '__main__':
    run_main()
