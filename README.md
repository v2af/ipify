## What does ipify do?

Have you ever needed to pragmatically get your public IP address? This is quite
common for developers provisioning cloud servers, for instance, where you might
be creating servers and running bootstrapping software on them without access to
server metadata.

Being able to quickly and reliably get access to your public IP address is
essential for configuring DNS, managing external services, and a number of other
operationally related tasks.

In general, there are a number of uses for public IP address information.

## Building ipify

To develop and build ipify, you'll need to have the Go programming language
setup on your computer. If you don't, you can read more about it here:
https://golang.org/

Once you have Go installed, you'll need to clone this project into your computer's GOPATH. For me, this means I'll typically do something like:

```bash
git clone https://github.com/v2af/ipify.git $GOPATH/src/github.com/v2af/ipify
```

To build the project, change to the project directory and run:

```bash
go build
```

## Questions?

Got a question? Please create a Github issue!
