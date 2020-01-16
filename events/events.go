package events

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

type TMEvents struct {
	Embedded struct {
		Events []struct {
			Name   string `json:"name"`
			Type   string `json:"type"`
			ID     string `json:"id"`
			Test   bool   `json:"test"`
			URL    string `json:"url"`
			Locale string `json:"locale"`
			Images []struct {
				Ratio    string `json:"ratio"`
				URL      string `json:"url"`
				Width    int    `json:"width"`
				Height   int    `json:"height"`
				Fallback bool   `json:"fallback"`
			} `json:"images"`
			Sales struct {
				Public struct {
					StartDateTime time.Time `json:"startDateTime"`
					StartTBD      bool      `json:"startTBD"`
					EndDateTime   time.Time `json:"endDateTime"`
				} `json:"public"`
			} `json:"sales"`
			Dates struct {
				Start struct {
					LocalDate      string `json:"localDate"`
					DateTBD        bool   `json:"dateTBD"`
					DateTBA        bool   `json:"dateTBA"`
					TimeTBA        bool   `json:"timeTBA"`
					NoSpecificTime bool   `json:"noSpecificTime"`
				} `json:"start"`
				Timezone string `json:"timezone"`
				Status   struct {
					Code string `json:"code"`
				} `json:"status"`
				SpanMultipleDays bool `json:"spanMultipleDays"`
			} `json:"dates"`
			Classifications []struct {
				Primary bool `json:"primary"`
				Segment struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"segment"`
				Genre struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"genre"`
				SubGenre struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"subGenre"`
				Family bool `json:"family"`
			} `json:"classifications"`
			Promoter struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"promoter"`
			Promoters []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"promoters"`
			Links struct {
				Self struct {
					Href string `json:"href"`
				} `json:"self"`
				Attractions []struct {
					Href string `json:"href"`
				} `json:"attractions"`
				Venues []struct {
					Href string `json:"href"`
				} `json:"venues"`
			} `json:"_links"`
			Embedded struct {
				Venues []struct {
					Name       string `json:"name"`
					Type       string `json:"type"`
					ID         string `json:"id"`
					Test       bool   `json:"test"`
					URL        string `json:"url"`
					Locale     string `json:"locale"`
					PostalCode string `json:"postalCode"`
					Timezone   string `json:"timezone"`
					City       struct {
						Name string `json:"name"`
					} `json:"city"`
					State struct {
						Name string `json:"name"`
					} `json:"state"`
					Country struct {
						Name        string `json:"name"`
						CountryCode string `json:"countryCode"`
					} `json:"country"`
					Address struct {
						Line1 string `json:"line1"`
					} `json:"address"`
					Location struct {
						Longitude string `json:"longitude"`
						Latitude  string `json:"latitude"`
					} `json:"location"`
					UpcomingEvents struct {
						Total int `json:"_total"`
						MfxBe int `json:"mfx-be"`
					} `json:"upcomingEvents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"_links"`
				} `json:"venues"`
				Attractions []struct {
					Name   string `json:"name"`
					Type   string `json:"type"`
					ID     string `json:"id"`
					Test   bool   `json:"test"`
					URL    string `json:"url"`
					Locale string `json:"locale"`
					Images []struct {
						Ratio    string `json:"ratio"`
						URL      string `json:"url"`
						Width    int    `json:"width"`
						Height   int    `json:"height"`
						Fallback bool   `json:"fallback"`
					} `json:"images"`
					Classifications []struct {
						Primary bool `json:"primary"`
						Segment struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"segment"`
						Genre struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"genre"`
						SubGenre struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"subGenre"`
						Type struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"type"`
						SubType struct {
							ID   string `json:"id"`
							Name string `json:"name"`
						} `json:"subType"`
						Family bool `json:"family"`
					} `json:"classifications"`
					UpcomingEvents struct {
						Total int `json:"_total"`
						MfxBe int `json:"mfx-be"`
					} `json:"upcomingEvents"`
					Links struct {
						Self struct {
							Href string `json:"href"`
						} `json:"self"`
					} `json:"_links"`
				} `json:"attractions"`
			} `json:"_embedded"`
		} `json:"events"`
	} `json:"_embedded"`
	Links struct {
		First struct {
			Href string `json:"href"`
		} `json:"first"`
		Self struct {
			Href string `json:"href"`
		} `json:"self"`
		Next struct {
			Href string `json:"href"`
		} `json:"next"`
		Last struct {
			Href string `json:"href"`
		} `json:"last"`
	} `json:"_links"`
	Page struct {
		Size          int `json:"size"`
		TotalElements int `json:"totalElements"`
		TotalPages    int `json:"totalPages"`
		Number        int `json:"number"`
	} `json:"page"`
}

func GetEvents(postalCode, apiKey string) (t TMEvents) {

	//Set Variables
	var i TMEvents
	baseUrl := "https://app.ticketmaster.com/discovery/v2/events.json?size=1"
	url := baseUrl + "&postalCode=" + postalCode + "&apikey=" + apiKey

	req, err := http.NewRequest("Get", url, nil)
	if err != nil {
		log.Fatalln("error with GET Response", err)
	}

	//Get Response from Request
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		log.Fatalln("error with GET Response", err)
	}

	defer res.Body.Close()
	//Unmarshal Json into data Struct
	body, _ := ioutil.ReadAll(res.Body)

	err = json.Unmarshal(body, &i)
	if err != nil {
		log.Fatalln("error unmarshalling", err)
	}
	return i
}
