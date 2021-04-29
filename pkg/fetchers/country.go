package fetchers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/marcustut/go-pwa/pkg/models"
)

func FetchCountry(country string) models.CountryResp {
	resp, err := http.Get(fmt.Sprintf("https://disease.sh/v3/covid-19/countries/%s?strict=true&allowNull=true", country))
	if err != nil {
		log.Fatalln(err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}

	var ret models.CountryResp
	err = json.Unmarshal(body, &ret)
	if err != nil {
		log.Fatalln(err)
	}

	return ret
}
