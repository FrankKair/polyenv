# coding=utf-8
from setuptools import find_packages, setup
from polyenv import __version__ as version

setup(
    author='Frank Kair',
    name='polyenv',
    version=version,
    packages=find_packages(include=['polyenv']),
    include_package_data=True,
    install_requires=['Click>=6.0','pytio',],
    author_email='frankkair@gmail.com',
    description="Run any programming language",
    long_description="Run any programming language using tryitonline (tio.run)",
    license="MIT license",
    url='https://github.com/FrankKair/polyenv',
    download_url='https://pypi.python.org/pypi/polyenv',
    platforms='any',
    test_suite='test_polyenv',
    tests_require=['pytest', ],
    setup_requires=['pytest-runner', ],
    entry_points={
        'console_scripts': [
            'polyenv=polyenv.cli:main',
        ],
    },
    classifiers=[
        'Development Status :: 2 - Pre-Alpha',
        'Intended Audience :: Developers',
        'License :: OSI Approved :: MIT License',
        'Natural Language :: English',
        "Programming Language :: Python :: 2",
        'Programming Language :: Python :: 2.7',
        'Programming Language :: Python :: 3',
        'Programming Language :: Python :: 3.4',
        'Programming Language :: Python :: 3.5',
        'Programming Language :: Python :: 3.6',
    ],
)

