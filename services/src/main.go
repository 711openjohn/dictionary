package main

import (
	"dictionary/client"
	"dictionary/endpoints"
	"dictionary/middleware"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/exec"
	"time"

	"github.com/atotto/clipboard"
	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	if len(os.Args) > 1 && os.Args[1] == "terminal" {
		mainCli()
	} else {
		mainRest()
	}
}

func mainRest() {
	router := gin.New()
	router.Use(gin.Recovery())
	router.Use(middleware.CORS())
	router.GET("/ping", ping)
	dictionaryGroup := router.Group("/dictionary")
	dictionaryGroup.GET("lookup/", endpoints.LookupDictionary)

	if err := router.Run(); err != nil {
		log.Fatal(err)
	}
}

func mainCli() {

	prevClipboard, err := clipboard.ReadAll()
	if err != nil {
		fmt.Println("Error reading clipboard:", err)
		os.Exit(1)
	}

	// Polling loop to detect clipboard changes
	cambridge := client.NewCambridge()
	title := color.New(color.BgBlue).Add(color.FgHiYellow).Add(color.Bold)
	meaning := color.New(color.FgHiWhite)
	for {
		currentClipboard, err := clipboard.ReadAll()
		if err != nil {
			fmt.Println("Error reading :", err)
			os.Exit(1)
		}

		if currentClipboard != prevClipboard {
			prevClipboard = currentClipboard
			lookup, err := cambridge.Lookup(prevClipboard, "chinese-traditional")
			if err != nil || lookup.Symbol == "" {
				continue
			}
			title.Println(lookup.Symbol)
			if len(lookup.Definitions) > 0 {
				playMP3WithVLC(lookup.Definitions[0].PronunciationLink)
			}
			for _, definition := range lookup.Definitions {
				meaning.Println("* ", definition.POS, " | ", definition.Meaning.Text)
				if definition.Meaning.Translation != "" {
					color.Green(definition.Meaning.Translation)
				}
				for _, example := range definition.Meaning.Examples {
					color.Cyan(example.Text)
					if example.Translation != "" {
						color.Cyan(example.Translation)
					}
				}
				fmt.Println("")
			}
		}

		// Adjust the polling interval as needed
		time.Sleep(1 * time.Second)
	}
}

func ping(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "pong"})
}

func playMP3WithVLC(link string) error {
	// Use the "open" command on Mac to open the link with the default application (VLC)
	cmd := exec.Command("vlc", "--no-repeat", "--intf", "dummy", link)
	// Run the command
	err := cmd.Start()
	if err != nil {
		fmt.Println(err)
		return err
	}

	return nil
}
