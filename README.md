# machine-learning-go
Learning go(lang) via running Data Science and machine learning projects

### To run any of the examples
Make sure `go` is installed and you have installed all of the dependencies listed in the `go.mod` file. Then, navigate to the folder you wish to run and use:
`go run .`

- install go: https://golang.org/doc/install
- install dependencies: `go mod download`
- navigate to the folder you wish to run: e.g. `cd pima`
- run the project: `go run .`

### To get new dependencies
If you add a new dependency to the project, you will need to run `go mod download` to download the new dependency. This will also update the `go.sum` file.