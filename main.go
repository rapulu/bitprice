package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Coindesk struct {
	Time struct {
		Updated    string    `json:"updated"`
		UpdatedISO time.Time `json:"updatedISO"`
		Updateduk  string    `json:"updateduk"`
	} `json:"time"`
	Disclaimer string `json:"disclaimer"`
	ChartName  string `json:"chartName"`
	Bpi        struct {
		Usd struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"USD"`
		Gbp struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"GBP"`
		Eur struct {
			Code        string  `json:"code"`
			Symbol      string  `json:"symbol"`
			Rate        string  `json:"rate"`
			Description string  `json:"description"`
			RateFloat   float64 `json:"rate_float"`
		} `json:"EUR"`
	} `json:"bpi"`
}

func main() {
	fmt.Println("Starting app...")
	
	a := app.New()
	w := a.NewWindow("BitPrice App")
	w.Resize(fyne.NewSize(400, 400))
	img := canvas.NewImageFromFile("./bitcoin_price.jpg")
	img.FillMode = canvas.ImageFillOriginal
	UCL := widget.NewLabel("")
	URL := widget.NewLabel("")
	ECL := widget.NewLabel("")
	ERL := widget.NewLabel("")
	GCL := widget.NewLabel("")
	GRL := widget.NewLabel("")
	getPrice := widget.NewButton("Get latest price", func() {
		worker(UCL, URL, ECL, ERL, GCL, GRL)
	})
	worker(UCL, URL, ECL, ERL, GCL, GRL)

	w.SetContent(
		container.NewVBox(
			img,
			container.NewHBox(
				container.NewVBox(UCL, ECL, GCL),
				container.NewVBox(URL, ERL, GRL),
			),
			getPrice,
		),
	)
	go func ()  {
		for {
			time.Sleep(time.Minute * 2)
			worker(UCL, URL, ECL, ERL, GCL, GRL)
		}
	}()
	w.ShowAndRun()
}

func worker(UCL, URL, ECL, ERL, GCL, GRL *widget.Label) {
	res, err := http.Get("https://api.coindesk.com/v1/bpi/currentprice.json")
	if err != nil {
		fmt.Printf("Error fetching: %v", err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Printf("Error reading: %v", err)
	}
	var coindesk Coindesk

	json.Unmarshal(body, &coindesk)

	UCL.SetText(coindesk.Bpi.Usd.Code)
	URL.SetText(coindesk.Bpi.Usd.Rate)
	ECL.SetText(coindesk.Bpi.Eur.Code)
	ERL.SetText(coindesk.Bpi.Eur.Rate)
	GCL.SetText(coindesk.Bpi.Gbp.Code)
	GRL.SetText(coindesk.Bpi.Gbp.Rate)
}