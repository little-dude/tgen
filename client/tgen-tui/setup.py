from setuptools import setup


setup(
    name='tgen_tui',
    version='0.1',
    include_package_data=True,
    description='a basic TUI for tgen',
    url='http://github.com/corentih/tgen',
    author='Corentin Henry',
    author_email='corentin.henry@nokia.com',
    license='MIT',
    packages=['tui'],
    install_requires=[line for line in open('requirements.txt')],
    entry_points={'console_scripts': ['tgen-tui = tui.main:run']}
)
