package main

import (
	"embed"
	"os/exec"
	"path/filepath"
	"syscall"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// 调用 deepzoom server 服务
	relativePath := "./dzi-server/dist/main.exe"
	absPath, absPathErr := filepath.Abs(relativePath)
	if absPathErr != nil {
		panic(absPathErr)
	}
	cmd := exec.Command("cmd", "/C", absPath)
	cmd.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	cmd.Start()
	// Create an instance of the app structure
	app := NewApp()
	// Create application with options
	err := wails.Run(&options.App{
		Title:  "dzishop",
		Width:  1280,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		// BackgroundColour: &options.RGBA{R: 255, G: 255, B: 255, A: 1},
		OnStartup: app.startup,
		Bind: []interface{}{
			app,
		},
		Frameless: true,
	})
	if err != nil {
		println("Error:", err.Error())
	}
}
