# Bump

## A simple to use version bumping app

Bump is a quick command line utility for creating a consistant, [SEMVER](https://en.wikipedia.org/wiki/Software_versioning#Semantic_versioning) formated reference file.


## Installation

Add the binary to your path.

## Usage

Bump won't automatically bump the version for you.

```bash
bump # This will print out the current. If one doesn't exist, it will create it.
```

You can idividually bump the Major, Minor, or Patch versions. Lower points are set to 0.

```bash
# Current version is at 1.2.3

bump -p # Bumps the Patch version. New version is 1.2.4
bump -m # Bumps the Minor version. New version is 1.3.0
bump -M # Bumps the Major version. New version is 2.0.0
```

If specifying multiples the higher most point superseeds the others.

```bash
# Current version is at 1.2.3
bump -M -p # A major and Patch point flag are set but only the major version is updated.
# New version is 2.0.0
```

Git tags can automatically be created

```bash
# Current version is at 1.2.3
bump -t # Creates the Git tag v1.2.3

bump -m -t # Updates the minor version and creates the Git tag v1.3.0
```

It also supports Makefiles, and passes a SEMVER string. Just point to a 'buildall' target.

```bash
bump --make # runs the buildall target in the makefile
```

## VERSION file

The VERSION file holds the current version data as JSON formated string for easy import.

```JSON
{"MAJOR":1,"MINOR":2,"PATCH":3}
```