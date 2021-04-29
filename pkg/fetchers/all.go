package fetchers

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/marcustut/go-pwa/pkg/models"
)

func FetchAllCountries() models.AllResp {
	resp, err := http.Get("https://disease.sh/v3/covid-19/all?allowNull=true")
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var ret models.AllResp
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Fatalln(err)
	}

	return ret
}
