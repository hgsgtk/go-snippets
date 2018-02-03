package main

import (
	"log"
	"time"

	"github.com/sclevine/agouti"
)

func main() {
	driver := agouti.ChromeDriver(agouti.Browser("chrome"))
	if err := driver.Start(); err != nil {
		log.Fatalf("Failed to start driver:%v", err)
	}
	defer driver.Stop()

	page, err := driver.NewPage()
	if err != nil {
		log.Fatalf("Failed to open page:%v", err)
	}

	if err := page.Navigate("https://qiita.com/login"); err != nil {
		log.Fatalf("Faled to navigate:%v", err)
	}

	identity := page.FindByID("identity")
	password := page.FindByID("password")
	identity.Fill("username")
	password.Fill("password")

	if err := page.FindByClass("loginSessionsForm_submit").Submit(); err != nil {
		log.Fatalf("Failed to login:%v", err)
	}

	time.Sleep(3 * time.Second)
}
