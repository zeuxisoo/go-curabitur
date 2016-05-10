package routes

import (
    "gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) string {
    return "Hello world"
}
