package models

type countryInfo struct {
	ID   float32 `json:"_id"`
	Iso2 string  `json:"iso2"`
	Iso3 string  `json:"iso3"`
	Lat  float32 `json:"lat"`
	Long float32 `json:"long"`
	Flag string  `json:"flag"`
}

type CountryResp struct {
	Updated                float32     `json:"updated"`
	Country                string      `json:"country"`
	CountryInfo            countryInfo `json:"countryInfo"`
	Cases                  float32     `json:"cases"`
	TodayCases             float32     `json:"todayCases"`
	Deaths                 float32     `json:"deaths"`
	TodayDeaths            float32     `json:"todayDeaths"`
	Recovered              float32     `json:"recovered"`
	TodayRecovered         float32     `json:"todayRecovered"`
	Active                 float32     `json:"active"`
	Critical               float32     `json:"critical"`
	CasesPerOneMillion     float32     `json:"casesPerOneMillion"`
	DeathsPerOneMillion    float32     `json:"deathsPerOneMillion"`
	Tests                  float32     `json:"tests"`
	TestsPerOneMillion     float32     `json:"testsPerOneMillion"`
	Population             float32     `json:"population"`
	Continent              string      `json:"continent"`
	OneCasePerPeople       float32     `json:"oneCasePerPeople"`
	OneDeathPerPeople      float32     `json:"oneDeathPerPeople"`
	OneTestPerPeople       float32     `json:"oneTestPerPeople"`
	Undefined              float32     `json:"undefined"`
	ActivePerOneMillion    float32     `json:"activePerOneMillion"`
	RecoveredPerOneMillion float32     `json:"recoveredPerOneMillion"`
	CriticalPerOneMillion  float32     `json:"criticalPerOneMillion"`
}
