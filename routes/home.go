package routes

import (
    "os"

    "gopkg.in/macaron.v1"
)

func Home(ctx *macaron.Context) {
    ctx.Data["assetMode"] = os.Getenv("ASSET_MODE")

    ctx.HTML(200, "index")
}
