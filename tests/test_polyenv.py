#!/usr/bin/env python
# -*- coding: utf-8 -*-

"""Tests for `polyenv` package."""

from click.testing import CliRunner

from polyenv import cli

def test_calling_with_error_on_file(datadir):
    python_file = datadir / 'test_ruby_error.rb'
    runner = CliRunner()
    result = runner.invoke(cli.main, ['ruby', str(python_file)])
    assert ".code.tio:1:in `<main>': undefined method `putsssss_' for main:Object (NoMethodError)" in result.output


def test_calling_python(datadir):
    python_file = datadir / 'test_python.py'
    runner = CliRunner()
    result = runner.invoke(cli.main, ['python3', str(python_file)])
    assert result.exit_code == 0
    assert result.output == "Testing Python\n\n"


def test_calling_ruby(datadir):
    ruby_file = datadir / 'test_ruby.rb'
    runner = CliRunner()
    result = runner.invoke(cli.main, ['ruby', str(ruby_file)])
    assert result.exit_code == 0
    assert result.output == "Testing Ruby\n\n"


def test_command_line_interface():
    runner = CliRunner()
    result = runner.invoke(cli.main)
    assert result.exit_code == 2
    assert 'Usage: main' in result.output

    help_result = runner.invoke(cli.main, ['--help'])
    assert help_result.exit_code == 0
    assert '--help  Show this message and exit.' in help_result.output
