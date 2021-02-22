//https://medium.com/@classik19881/ci-continuous-integration-with-travis-ci-for-golang-project-532d1d1fc7b6
package main

import "fmt"

func main() {
  fmt.Println(run())
}

func run() string {
  return "Setup Travis CI for Golang project"
}
