package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Money struct{
	Symbol string `json:"symbol"`
	Price string `json:"price"`
}

func main() {
	cYlw := "\033[33m"
	cRst := "\033[0m"

	symbol := flag.String("s","BTCUSDT","Exchange rates and currency conversion.") 
	flag.Parse()

	res, err := http.Get("https://api.binance.com/api/v3/ticker/price?symbol="+*symbol)
	if err != nil{
	log.Fatalln(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil{
		log.Fatalln(err)
	}

	money := Money {}
	jsonErr := json.Unmarshal([]byte(body), &money)
	if jsonErr != nil{
		log.Fatalln(err)
	}
	
	price:= string(cYlw)+money.Price+string(cRst)
	fmt.Printf("Currency CLI\n-----------\n%s\n%s\n-----------\n",money.Symbol,price)
}
