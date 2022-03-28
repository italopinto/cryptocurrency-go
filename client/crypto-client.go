package client

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/italopinto/cryptocurrency-go/model"
)

// Fetch is exported, it receive the fiat and the crypto to fetch from Nomics API
func FetchCrypto(fiat string, crypto string) (string, error) {
	// build the url string
	URL := "https://api.nomics.com/v1/currencies/ticker?key=3990ec554a414b59dd85d29b2286dd85&interval=1d&ids=" + crypto + "&convert=" + fiat

	// we make http request using the Get function
	resp, err := http.Get(URL)
	if err != nil {
		log.Fatal("Ooops an error ocurred, please try again")
	}
	defer resp.Body.Close()

	// create variable of same type as our model
	var cResp model.Cryptoresponse

	//decode the data
	if err := json.NewDecoder(resp.Body).Decode(&cResp); err != nil {
		log.Fatal("Oooops! an error ocurred, please try again")
	}

	// invoke the text output function & return it with nil as the error value
	return cResp.TextOutput(), nil
}
