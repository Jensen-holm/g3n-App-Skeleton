# G3N graphics engine wrapper 

A wrapper for the G3N computer graphics library with built in physics simulation package.

![G3N wrapper physics sim demo](https://github.com/Jensen-holm/g3n-Wrapper/blob/main/demo.gif)

right now running the following to run the program:

`$ go run main.go` <br>

will bring up the following issue:

go run main.go                                       01:57:04 PM
../../../go/pkg/mod/github.com/golang/freetype@v0.0.0-20170609003504-e2365dfdc4a0/truetype/face.go:13:2: no required module provides package golang.org/x/image/font; to add it:
    go get golang.org/x/image/font

after updating the package, this project no longer works as there is some form of <br>
issue with the newer version of the font package and my application. <br>

