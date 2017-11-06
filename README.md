[![Build Status](https://travis-ci.org/alekskivuls/jsonSearchEngine.svg?branch=master)](https://travis-ci.org/alekskivuls/jsonSearchEngine)

# JSON Search Engine
Engine to index JSON files and search for properties.

### Building
```
go build
```

### Running
Pass in individual files or folders as arguments for the engine to index those files for search.
```
./jsonSearchEngine <files>
```

## Development

### Formatting
```
goimports -w -l -e .
```