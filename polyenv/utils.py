# -*- coding: utf-8 -*-


def available_languages():
    from requests import get
    url = "https://tio.run/languages.json"
    data = get(url).json()
    return [f"{data[key]['name']} : {key}" for key, value in data.items()]


def search_languages(target):
    all_langs = available_languages()
    langs = [lang for lang in all_langs if target.lower() in lang.lower()]
    click.echo('Available languages with your query:')
    click.echo("Language : command\n")
    for lang in langs:
        click.echo(f"{lang}\n")


def run_programming_language(language, input_file):
    import click
    from pytio import Tio, TioRequest
    from pathlib import Path

    data = Path(input_file).read_text()
    request = TioRequest(lang=language, code=data)
    response = Tio().send(request)
    if response.error:
        click.echo(response.error)
        return

    not_found = f"The language '{language}' could not be found on the server.\n"
    if response.result == not_found:
        search_languages(language)

    click.echo(response.result)
