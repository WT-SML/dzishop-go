package main

import (
	"context"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

// 选择文件夹
type File struct {
	Label  string `json:"label"`
	Key    string `json:"key"`
	IsLeaf bool   `json:"isLeaf"`
}

func (a *App) WailsSelectFolder() string {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		panic(err)
	}
	return result
}
func (a *App) WailsGetDirectoryItem(dir string) string {
	files, err := os.ReadDir(dir)
	if err != nil {
		panic(err)
	}
	ret := make([]File, 0)
	for _, v := range files {
		file := File{Label: v.Name(), Key: filepath.Join(dir, v.Name()), IsLeaf: !v.IsDir()}
		fmt.Println(file)
		ret = append(ret, file)
	}
	jsonData, err := json.Marshal(ret)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(jsonData))
	return string(jsonData)
}
