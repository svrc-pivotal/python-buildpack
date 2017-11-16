from setuptools import setup

setup(name='animals',
      version='0.1',
      description='animals test package',
      url='http://example.com',
      author='buildpack',
      author_email='buildpack@example.com',
      license='MIT',
      install_requires=['mammals'],
      packages=['animals'],
      zip_safe=False)
