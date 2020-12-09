import os.path
from setuptools import setup, find_packages

with open(os.path.join(os.path.dirname(__file__), "requirements.txt")) as f:
    requirements = f.read().strip()

setup(
    name="reverse",
    version="0.0.0",
    description="Reverse data",
    packages=find_packages(exclude=["tests"]),
    package_data={"reverse": ["metadata/*"]},
    install_requires=requirements,
)
