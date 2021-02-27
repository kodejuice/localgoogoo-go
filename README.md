<h1 align="center"><img width="512" src="./localgoogoo.png" alt="localgoogoo" /></h1>

<p align="center">
<a href="https://github.com/kodejuice/localgoogoo-go/blob/master/LICENSE"><img src="https://img.shields.io/badge/license-MIT-yellowgreen.svg?maxAge=2592000" alt="License" /></a>
<a href="https://github.com/kodejuice/localgoogoo-go/actions"><img src="https://github.com/kodejuice/localgoogoo-go/workflows/ci/badge.svg?branch=master" alt="Build Status" /></a>
</p>

<p align="center">
<a href="https://asciinema.org/a/395033">
<img src="./terminal-shot.png" alt="Asciicast" width="893" height="650"/>
</a>
</p>

A command line tool that lets you use localGoogoo from the terminal (written in Go).

Don't know what localGoogoo is?, you should <a href="https://github.com/kodejuice/localgoogoo-go"> check it out </a>



## Installation

### Requirements
  * Golang
  * <a href="https://github.com/kodejuice/localgoogoo">localGoogoo</a> installed on your system
<br><br>

#### Install using `go get`

```bash
$ go get http://github.com/kodejuice/localgoogoo-go
```

#### Install using `git clone`

```bash
$ git clone http://github.com/kodejuice/localgoogoo-go.git
$ cd localgoogoo-go
$ go install
```

This installs `localgoogoo-go` to your local machine, you can alias it to a shorter name, such as `localgoogoo` or `googoo`

Usage
-------------

Make sure localGoogoo is functioning properly in the browser, because all this package does is make http requests to localgoogoo installed on your system and render the results of any query on your terminal.



