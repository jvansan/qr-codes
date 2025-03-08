package main

import (
	"fmt"

	"github.com/yeqown/go-qrcode/v2"
	"github.com/yeqown/go-qrcode/writer/standard"
)

type QRCode struct {
	url  string
	path string
}

var CODES = []QRCode{
	{
		url:  "https://nanaimomountainbikeclub.com/sponsorship/",
		path: "nmbc-sponsorship.jpeg",
	},
	{
		url:  "https://nmbc.tidyhq.com/public/membership_levels",
		path: "nmbc-tidhq.jpeg",
	},
	{
		url:  "https://bit.ly/3ZJP3bq",
		path: "nmbc-votes.jpeg",
	},
	{
		url:  "https://www.eventbrite.ca/e/nmbc-presents-murderhorn-maiden-part-the-fourth-tickets-1236792753419",
		path: "nmbc-murderhorn-eventbrite2025.jpeg",
	},
}

func create_qrcode(code QRCode) {
	qrc, err := qrcode.New(code.url)
	if err != nil {
		fmt.Printf("could not generate QRCode: %v", err)
		return
	}

	w, err := standard.New(code.path)
	if err != nil {
		fmt.Printf("standard.New failed: %v", err)
		return
	}

	// save file
	if err = qrc.Save(w); err != nil {
		fmt.Printf("could not save image: %v", err)
	}
}

func main() {
	for _, code := range CODES {
		create_qrcode(code)
	}
}
