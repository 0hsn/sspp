### Install

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
5. Put on a path that's on your OS `$PATH`. 
    ```bash
    go build -o ~/.local/bin/sspp ./src
    ```
    > On my case `~/.local/bin` is on my `$PATH`

### Usage

```bash
curl -s https://httpbin.org/get | sspp --json='headers.Host'
```
