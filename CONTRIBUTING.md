# Contribution guidelines

This document contains guidelines for contributing code to the
`bunnystorage` package. Or it will, in the future.

## Tests

This repository contains tests against the live [bunny.net Edge Storage
API](https://docs.bunny.net/reference/storage-api) and test data used in
mock tests.

To run tests manually, create a `.env` file with the required
environment variables at the root of the repository and invoke the tests
target from `make`.

```console
make test
```

The `.env` look like this:

```
BUNNY_STORAGE_ZONE=XXX
BUNNY_API_KEY=XXXX
BUNNY_READONLY_API_KEY=XXXX
```

You can also run coverage tests if you prefer. For that, simply invoke
the coverage `make` target.

```console
make test/coverage
```
