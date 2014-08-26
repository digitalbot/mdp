package main

import (
	"github.com/lxn/walk"
	. "github.com/lxn/walk/declarative"
)

import (
	"fmt"
	"log"
	"time"
	"gopkg.in/fsnotify.v1"
)

func main() {
	filename := "C:\\Users\\kosuke\\.gitconfig"
	fmt.Println(filename)
	
	var mw *walk.MainWindow
	var wv *walk.WebView
	if err := (MainWindow{
		AssignTo: &mw,
		Title:   "mdp",
		MinSize: Size{800, 600},
		Layout:  VBox{},
		Children: []Widget{
			WebView{
				AssignTo: &wv,
				Name:     "wv",
				URL:      "file://" + filename,
			},
		},
	}.Create()); err != nil {
		log.Fatal(err)
	}

	fmt.Println("test")
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()
	
	time.Sleep(10 * time.Millisecond)
	go func() {
		time.Sleep(10 * time.Millisecond)
		for {
			select {
			case event := <-watcher.Events:
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					log.Println("modified:", event.Name)
					wv.SetURL(filename)
					wv.URLChanged()
					watcher.Remove(filename)
					watcher.Add(filename)
				}
			case err := <-watcher.Errors:
				log.Println("error:", err)
			}
		}
	}()

	watcher.Add(filename)
	if err != nil {
		log.Fatal(err)
	}
	mw.Run()
}
