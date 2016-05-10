package routes

import (
    "os"

    "gopkg.in/macaron.v1"
)

func isLocalAssetsExists(path string) (bool, error) {
    _, err := os.Stat(path)

    if err == nil {
        return true, nil
    }

    if os.IsNotExist(err) {
        return false, nil
    }

    return true, err
}

func Home(ctx *macaron.Context) {
    isLocalAssets, _ := isLocalAssetsExists("./public/build")

    ctx.Data["isLocalAssets"] = isLocalAssets
    ctx.HTML(200, "index")
}
