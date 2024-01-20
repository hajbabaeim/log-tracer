# log-tracer
This repository demonstrates showcases of Sentry &amp; Jaeger for tracking the system performance through its logs

### Sentry
1. #### Install
    To install sentry on your machine, go to [this address](https://develop.sentry.dev/self-hosted/) and download the latest version of its source code.

2. #### Use 
    You might reach your Sentry server by this address: `http://localhost:9000/` on your local machine, then you can create your own project.

### Jaeger

#### Use badger for local storage
Download [BadgerV3](https://github.com/dgraph-io/badger/archive/refs/tags/v1.6.2.tar.gz) and run `go install`

Run the Jaeger with badger:
```
docker run \
  -e SPAN_STORAGE_TYPE=badger \
  -e BADGER_EPHEMERAL=false \
  -e BADGER_DIRECTORY_VALUE=/badger/data \
  -e BADGER_DIRECTORY_KEY=/badger/key \
  -v <storage_dir_on_host>:/badger \
  -p 16686:16686 \
  jaegertracing/all-in-one:1.53
```