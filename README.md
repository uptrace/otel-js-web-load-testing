## Installation

To install dependecies:

```shell
pnpm install
```

To start backend (dummy Otel Collector) listening on `:14318`:

```shell
make backend
```

To start example:

```shell
UPTRACE_DSN="http://secret@localhost:14318/?grpc=14317" make
```

Then open [http://localhost:8090/](http://localhost:8090/)
