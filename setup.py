# coding=utf-8
from setuptools import find_packages, setup
from polyenv import __version__ as version

setup(
    name='polyenv',
    version=version,
    packages=find_packages(),
    include_package_data=True,
    install_requires=[
        'pytio',
    ],
    author='Frank Kair',
    author_email='frankkair@gmail.com',
    description="Run any programming language",
    long_description="Run any programming language using tryitonline (tio.run)",
    license='MIT',
    url='https://github.com/FrankKair/polyenv',
    download_url='https://pypi.python.org/pypi/polyenv',
    platforms='any',
    test_suite='test_polyenv',
    entry_points={
        'console_scripts': ['polyenv = polyenv:run_main'],
    },
)
