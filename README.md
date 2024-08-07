# Go For Beginners Workshop

This repository contains material for 2 hour long workshop for absolute
Go beginners.

Goal of this workshop **IS NOT** to teach solid foundation of Go, but
to give a taster to curious newcomers how Go looks like and what can be
done with it in no time (read 2 hours).

Application we are building here is a simple _microblog_ that does not
even store posts between runs. We go from _hello world_ over very basic
programming concepts and constructs all the way to webserver using
external module (`templ`).

# Setup
* Install VSCode and Go extension, optionally also Git and Git extension
* Install Go
* Install Delve
  go install github.com/go-delve/delve/cmd/dlv@latest
* Install Templ
  go install github.com/a-h/templ/cmd/templ@latest