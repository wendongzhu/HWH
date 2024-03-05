package main

import (
	"embed"
	"github.com/labstack/gommon/log"
	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/mac"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
	"github.com/wendongzhu/hwh/backend/end"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()             // app interface
	sys := end.NewRealTime()    // real time data interface
	backend := end.NewBackend() // backend interface

	// Create application with options
	err := wails.Run(&options.App{
		Title:     "HWH",
		MinWidth:  1300,
		MinHeight: 700,
		Frameless: true,
		//DisableResize: true,
		EnableFraudulentWebsiteDetection: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		OnDomReady:       sys.OnDomReady,
		Bind: []interface{}{
			app,
			sys,
			backend,
		},
		Windows: &windows.Options{
			WebviewIsTransparent: true, // 网页透明
			WindowIsTranslucent:  true, // 窗口透明
			Theme:                windows.Dark,
		},
		Mac: &mac.Options{
			TitleBar: &mac.TitleBar{
				TitlebarAppearsTransparent: true,
				HideTitle:                  false,
				HideTitleBar:               false,
				FullSizeContent:            false,
				UseToolbar:                 false,
				HideToolbarSeparator:       true,
			},
			Appearance:           mac.NSAppearanceNameDarkAqua,
			WebviewIsTransparent: true,
			WindowIsTranslucent:  false,
			About: &mac.AboutInfo{
				Title:   "HWH",
				Message: "© 2023 wendong.zhu",
				//Icon:    icon,
			},
		},
		Linux: &linux.Options{
			//Icon: icon,
			WindowIsTranslucent: false,
			WebviewGpuPolicy:    linux.WebviewGpuPolicyAlways,
			ProgramName:         "HWH",
		},
		Debug: options.Debug{
			OpenInspectorOnStartup: false,
		},
	})

	if err != nil {
		log.Error("Error:", err.Error())

	}
}
