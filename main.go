package main

import (
	"log"
	"path/filepath"
	"os"
	"io/ioutil"
	"sync"
	"github.com/joho/godotenv"
	"github.com/fsnotify/fsnotify"
	"ark-notify/webhook"
)

type ArkNotifier struct {
	sync.Mutex
	LogDir string
	LogSizeTable map[string]int64
	WebhookURL string
}

func (an *ArkNotifier) Initialize() {
	an.LogDir = os.Getenv("ARK_LOG_DIR")
	an.LogSizeTable = make(map[string]int64)
	an.WebhookURL = os.Getenv("DISCORD_WEBHOOK_URL")

	files, err := filepath.Glob(an.LogDir + "*.log")
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		stat, err := os.Stat(f)
		if err != nil {
			continue
		}
		log.Println(f, stat.Size())
		an.LogSizeTable[f] = stat.Size()
	}

}

func (an *ArkNotifier) onLogFileUpdated(path string){
	log.Println("modified file:", path)
	fp, err := os.Open(path)
	if err != nil {
		log.Println("error while open:", err)
		return
	}
	defer fp.Close()
	
	an.Lock()
	defer an.Unlock()
	fp.Seek(an.LogSizeTable[path], 0)
	logs, err := ioutil.ReadAll(fp)
	if err != nil {
		log.Println("error while ReadAll:", err)
		return
	}
	log.Printf("read %d bytes: %s", len(logs), string(logs))
	an.LogSizeTable[path] += int64(len(logs))
	
	webhook.SendWebhook(an.WebhookURL, string(logs))
}

func (an *ArkNotifier) Watch(){
	watcher, err := fsnotify.NewWatcher()
	if err != nil {
		log.Fatal(err)
	}
	defer watcher.Close()

	done := make(chan bool)
	go func() {
		for {
			select {
			case event, ok := <-watcher.Events:
				if !ok {
					return
				}
				log.Println("event:", event)
				if event.Op&fsnotify.Write == fsnotify.Write {
					fileName := event.Name
					an.onLogFileUpdated(fileName)
				}
			case err, ok := <-watcher.Errors:
				if !ok {
					return
				}
				log.Println("error:", err)
			}
		}
	}()

	err = watcher.Add(an.LogDir)
	if err != nil {
		log.Fatal(err)
	}
	<-done
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	var an ArkNotifier
	an.Initialize()
	an.Watch()
}
