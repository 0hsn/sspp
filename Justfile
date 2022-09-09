build:
	rm -rf dist/sspp
	go build -o dist/sspp ./src

run:
	dist/sspp
