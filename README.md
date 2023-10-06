# bookstore-rest-api-go

An example of an api using Gorilla Mux and a controller-service-repository pattern.

## Execute

To run locally

```sh
make run
```

---

To run in a container (_*Requires a `.env` file in the root to load the environment variables_):

```sh
make docker-run
```

To remove docker deployment:

```sh
make docker-down
```

## Dependencies

Use `go install` to get the tools to generate the coverage outputs, dependencies graph and mocks:

- [**mockery**: to auto generate mocks](https://github.com/vektra/mockery)
- [**goda**: to show dependencies between packages](https://github.com/loov/goda)
- [**go-cover-treemap**: to get the coverage heatmap](https://github.com/nikolaydubina/go-cover-treemap)
- [**go-coverage**: to get table of coverage impact](https://github.com/gojek/go-coverage)
