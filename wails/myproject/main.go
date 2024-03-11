package main

import (
	"context"
	"embed"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"myproject/pkg/fileOperation"
	"myproject/pkg/storage"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()
	store := storage.NewStorage("testfile.txt")
	fileOperations := fileOperation.NewFileOperation()

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "myproject",
		Width:  1024,
		Height: 768,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup: func(ctx context.Context) {
			app.startup(ctx)
			fileOperations.SetCtx(ctx)
		},
		Bind: []interface{}{
			app,
			store,
			fileOperations,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
