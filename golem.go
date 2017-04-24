package main

import (
	"github.com/sciter-sdk/go-sciter"
	"github.com/sciter-sdk/go-sciter/window"
	"path/filepath"
	"os"
)

func main() {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic("Incorrect filepath")
	}

	rect := sciter.Rect{0,0,600,800}

	w, err := window.New(sciter.SW_TITLEBAR | sciter.SW_CONTROLS | sciter.SW_MAIN, &rect)
	if err != nil {
		panic("Sciter window not built")
	}
	w.LoadFile(dir + "/main.html")
	w.SetTitle("GolemRun")
	w.Show()
	w.Run()
}