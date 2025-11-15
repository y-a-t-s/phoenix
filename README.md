# Phoenix

A simple wrapper for [firebird](https://github.com/y-a-t-s/firebird).

Designed with JS-free tor use in mind. You can route requests through tor by setting the SOCKS5 proxy details in your environment (i.e. `export ALL_PROXY="socks5://127.0.0.1:9050"`).

<hr>

You can set the domain this program connects to by setting the `KF_HOST` environment variable. So if you want to use this for clearnet access, you would set `KF_HOST="kiwifarms.st"`.
Example:
```sh
$ KF_HOST="kiwifarms.st" go run .
```
