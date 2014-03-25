Go implementations of various algorithms introduced by the Algorithms course from the university of Princeton. This is still a work in progress.

In case you are new to the Go ecosystem, this is a quick start:

    $ cd ~
    $ mkdir "goroot"
    $ export GOPATH="$HOME/goroot"
    $ go install github.com/seri/goalgo
    $ cd goroot/src/github.com/seri/goalgo/examples
    $ go run sort_client.go
