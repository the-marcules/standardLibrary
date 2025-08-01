package main

import (
	"embed"
	"fmt"
	"mycrypto/pkg/config"
	"mycrypto/pkg/cryptokit"

	goruntime "runtime"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/menu"
	"github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

//go:embed all:frontend/dist
var assets embed.FS

func openFile(data *menu.CallbackData) {
	fmt.Println("Open file clicked")
}

func main() {

	app := NewApp(config.InitConfig())
	ckApi := cryptokit.NewCryptoKitApi(app.config)

	AppMenu := menu.NewMenu()
	FileMenu := AppMenu.AddSubmenu("File")
	FileMenu.AddText("Open", keys.CmdOrCtrl("o"), openFile)
	FileMenu.AddSeparator()
	FileMenu.AddText("Quit", keys.CmdOrCtrl("q"), func(_ *menu.CallbackData) {
		runtime.Quit(app.ctx)
	})

	if goruntime.GOOS == "darwin" {
		AppMenu.Append(menu.EditMenu()) // Für macOS, um Cmd+C, Cmd+V, Cmd+Z zu aktivieren
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "CryptoKit Desktop",
		Width:            1024,
		Height:           768,
		Menu:             AppMenu,
		AssetServer:      &assetserver.Options{Assets: assets},
		BackgroundColour: &options.RGBA{R: 0, G: 0, B: 0, A: 1},
		Mac: &mac.Options{
			// TitleBar: mac.TitleBarHiddenInset(),
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: false,
				UseToolbar:                 false,
				HideTitleBar:               false,
			},
			Appearance:           "NSAppearanceNameDarkAqua",
			WindowIsTranslucent:  true,
			WebviewIsTransparent: true,
		},
		OnStartup: app.startup,
		Bind:      []interface{}{app, ckApi},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
