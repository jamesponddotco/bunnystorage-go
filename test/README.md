# Integration tests

This repository contains tests against the live [bunny.net Edge Storage
API](https://docs.bunny.net/reference/storage-api) and test data used in
mock tests.

To run integration tests manually, create a `.env` file with the
required environment variables at the root of the repository and invoke
the integration tests target from `make`.

```console
make test/integration
```

The `.env` look like this:

```
BUNNY_STORAGE_ZONE=XXX
BUNNY_API_KEY=XXXX
BUNNY_READONLY_API_KEY=XXXX
```

