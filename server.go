package main

import (
    "github.com/zeuxisoo/go-curabitur/routes"
    "github.com/zeuxisoo/go-curabitur/public"
    "github.com/zeuxisoo/go-curabitur/templates"

    "gopkg.in/macaron.v1"
    "github.com/go-macaron/bindata"
)

func main() {
    m := macaron.Classic()

    m.Use(macaron.Static("public",
        macaron.StaticOptions{
            FileSystem: bindata.Static(bindata.Options{
                Asset:      public.Asset,
                AssetDir:   public.AssetDir,
                AssetInfo:  public.AssetInfo,
                AssetNames: public.AssetNames,
                Prefix:     "",
            }),
        },
    ))

    m.Use(macaron.Renderer(macaron.RenderOptions{
        TemplateFileSystem: bindata.Templates(bindata.Options{
            Asset:      templates.Asset,
            AssetDir:   templates.AssetDir,
            AssetInfo:  templates.AssetInfo,
            AssetNames: templates.AssetNames,
            Prefix:     "",
        }),
    }))

    m.Get("/", routes.Home)

    m.Run("127.0.0.1", 8080)
}
