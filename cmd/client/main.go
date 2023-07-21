package main

import (
	"fmt"
	"github.com/Khasmag06/gophkeeper/config"
	client2 "github.com/Khasmag06/gophkeeper/internal/client"
	decoder2 "github.com/Khasmag06/gophkeeper/pkg/decoder"
	"log"
)

func main() {
	cfg, err := config.NewConfig()
	if err != nil {
		log.Fatal(err)
	}

	decoder, err := decoder2.New(cfg.Decoder.SecretKey)
	if err != nil {
		log.Fatal(err)
	}

	client := client2.New(cfg.Server, decoder)
	fmt.Println("Start client")
	for {
		if err := client.Run(); err != nil {
			log.Println(err)
		}
	}

}
