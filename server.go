package main

import (
    "github.com/zeuxisoo/go-curabitur/routes"

    "gopkg.in/macaron.v1"
)

func main() {
    m := macaron.Classic()

    m.Get("/", routes.Home)

    m.Run("127.0.0.1", 8080)
}
