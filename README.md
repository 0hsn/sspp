# Structured Stream Processing and Presentation | `sspp`

## What

`sspp` is a command-line tool is a parse and presentation utility for structured documents, i.e, json. Currently it only supports json.

## Why

There is not good command line tool that works as an umbrella command to support multiple structured config / document. 

Although currently this tool is an alternative to `jq`.

## How

### How to install

1. Please install [Golang](https://go.dev/doc/install)
2. You also need [Git-scm](https://git-scm.com/downloads)
3. Clone the repository
    ```bash
    git clone https://github.com/0hsn/sspp.git
    ```
4. Change directory
    ```sh
    cd ./sspp
    ```
5. Put on a path that's on your OS `$PATH`. On my case `~/.local/bin` is on my `$PATH`
    ```bash
    go build -o ~/.local/bin/sspp ./src
    ```

### Usage

```bash
$ curl -s https://httpbin.org/get | sspp --json='headers.Host'

# you can pass data directly with --data flag
$ sspp --json='data.email' --data='{"data": {"email": "sample@example.org"}}'

# It is possible to pass default data directly with --or flag
$ sspp --json='data.nonExistent' --data='{"data": {"email": "sample@example.org"}}' --or="nil"
$ curl -s https://httpbin.org/get | sspp --json='headers.nonExistent' --or="nil"
```

### Report issue

You can report issues or ask questions on [github issues](https://github.com/0hsn/sspp/issues).