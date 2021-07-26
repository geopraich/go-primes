package main

import "github.com/geopraich/go-primes/internal"

func main() {
	app := internal.App{}
	app.Initialise()
	app.Run(":8083")
}
