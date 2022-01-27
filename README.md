# validate-linter

## build

    go build -o bin/lnt main.go

## run for all packages

    for file in `find ./testdata -type d`; ./bin/lnt $file
