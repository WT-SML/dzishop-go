package main

import (
	"context"
	"fmt"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"io/fs"
	"os"
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
	name     string
	isDir    bool
	fileType fs.FileMode
	fileInfo fs.FileInfo
}

func (a *App) SelectFolder() []File {
	result, err := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{})
	if err != nil {
		panic(err)
	}
	files, err := os.ReadDir(result)
	if err != nil {
		panic(err)
	}
	ret := make([]File, len(files))
	for _, v := range files {
		fileInfo, _ := v.Info()
		ret = append(ret, File{name: v.Name(), isDir: v.IsDir(), fileType: v.Type(), fileInfo: fileInfo})
	}
	println(ret)
	return ret
}
