package main

import (
	"log"

	vlc "github.com/adrg/libvlc-go/v3"
)

func main() {
	log.Printf("start")

	if err := vlc.Init( "--quiet"); err != nil {
		log.Fatal(err)
	}
	defer vlc.Release()

	//Create a new player
	player, err := vlc.NewPlayer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		player.Stop()
		player.Release()
	}()

	media, err := player.LoadMediaFromURL("http://stream-uk1.radioparadise.com/mp3-32")
	if err != nil {
		log.Fatal(err)
	}
	defer media.Release()

	//Retrieve player event manager
	manager, err := player.EventManager()
	if err != nil {
		log.Fatal(err)
	}

	quit := make(chan struct{})
	eventCallback := func(event vlc.Event, userData interface{}){
		close(quit)
	}

	eventID, err := manager.Attach(vlc.MediaPlayerEndReached, eventCallback, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer manager.Detach(eventID)

	err = player.Play()
	if err != nil {
		log.Fatal(err)
	}

	<-quit
}


