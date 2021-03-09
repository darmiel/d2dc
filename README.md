# d2dc
d2dc converts a `docker run`-command into a docker-compose service (see [example](#example)).
> Allan, please add details!

## Example
```bash
$ d2dc docker run -d --name ftpd_server -p 21:21 -p 30000-30009:30000-30009 -e "PUBLICHOST=localhost" stilliard/pure-ftpd

# [ d2dc v. 1.0.0-alpha ]
# Generated docker-compose [service] from:
# d2dc docker run -d --name ftpd_server -p 21:21 -p 30000-30009:30000-30009 -e PUBLICHOST=localhost stilliard/pure-ftpd
services:
  ftpd_server:
    image: stilliard/pure-ftpd
    environment:
      PUBLICHOST: localhost
    publish:
    - "21:21"
    - 30000-30009:30000-30009
```

## Installation
``` bash
$ go get github.com/darmiel/d2dc
```

## Build
### Prerequisites
* Go
* Git

### Build
Build builds the project and outputs the executable file in the root folder

```bash
$ git clone https://github.com/darmiel/d2dc
$ go build ./cmd/d2dc
```

### Install
Install builds the project and saves the executable file in the GOPATH

```bash
$ git clone https://github.com/darmiel/d2dc
$ go install ./cmd/d2dc
```

## Why?
I wanted to learn a bit of the commands for `docker run` and so I thought this would be a good project to learn those commands.