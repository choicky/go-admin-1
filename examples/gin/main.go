package main

import (
	"github.com/chenhg5/go-admin/adapter"
	"github.com/chenhg5/go-admin/engine"
	"github.com/chenhg5/go-admin/examples/datamodel"
	"github.com/chenhg5/go-admin/modules/config"
	"github.com/chenhg5/go-admin/plugins/admin"
	"github.com/chenhg5/go-admin/plugins/example"
	"github.com/gin-gonic/gin"
	"github.com/chenhg5/go-admin/template/types"
)

func main() {
	r := gin.Default()

	eng := engine.Default()

	cfg := config.Config{
		DATABASE: []config.Database{
			{
				HOST:         "127.0.0.1",
				PORT:         "3306",
				USER:         "root",
				PWD:          "root",
				NAME:         "godmin",
				MAX_IDLE_CON: 50,
				MAX_OPEN_CON: 150,
				DRIVER:       "mysql",
			},
		},
		DOMAIN: "localhost",
		PREFIX: "admin",
		STORE: config.Store{
			PATH:   "./uploads",
			PREFIX: "uploads",
		},
		LANGUAGE: "cn",
		INDEX: "/",
	}

	adminPlugin := admin.NewAdmin(datamodel.TableFuncConfig)

	// you can custom a plugin like:

	examplePlugin := example.NewExample()

	if err := eng.AddConfig(cfg).AddPlugins(adminPlugin, examplePlugin).AddAdapter(new(adapter.Gin)).Use(r); err != nil {
		panic(err)
	}

	r.Static("/uploads", "./uploads")

	// you can custom your pages like:

	r.GET("/" + cfg.PREFIX + "/custom", func(ctx *gin.Context) {
		adapter.GinContent(ctx, func() types.Panel {
			return datamodel.GetContent()
		})
	})

	r.Run(":9033")
}
