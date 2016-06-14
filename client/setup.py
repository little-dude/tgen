import os
from setuptools import setup


def get_version():
    if os.path.isfile(os.path.join('tgenpy', '__version__')):
        with open('tgenpy/__version__', 'r') as f:
            return f.read().strip('\n')
    return 'dev'


setup(
    name='tgenpy',
    version=get_version(),
    include_package_data=True,
    description='python client for tgen',
    url='http://github.mv.usa.alcatel.com/corentih/tgen',
    author='Corentin Henry',
    author_email='corentin.henry@nokia.com',
    license='TBD',
    packages=['tgenpy'],
    package_data={'tgenpy': ['schemas.capnp']},
    install_requires=[line for line in open('requirements.txt')],
)
