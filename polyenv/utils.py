# -*- coding: utf-8 -*-


def run_programming_language(language, input_file):
    import click
    from pytio import Tio, TioRequest
    from pathlib import Path

    data = Path(input_file).read_text()
    request = TioRequest(lang=language, code=data)

    tio = Tio()
    response = tio.send(request)

    if response.error:
        click.echo(response.error)
        return

    click.echo(response.result)
