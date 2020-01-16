package main

import (
	"database/sql"
	"fmt"
	"github.com/3pings/clWallEvents/config"
	"github.com/3pings/clWallEvents/events"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"log/syslog"
	"os"
)

type eventData struct {
	name         string
	startDate    string
	venueName    string
	venueAddress string
	venueCity    string
	eventId      string
	eventUrl     string
	logoUrl      string
	venueId      string
}

func main() {

	key := os.Getenv("tktmst_apikey")
	pc := "08028"

	logwriter, e := syslog.New(syslog.LOG_NOTICE, "incident")
	if e == nil {
		log.SetOutput(logwriter)

	}

	q := events.GetEvents(pc, key)

	ed := eventData{}

	ed.name = q.Embedded.Events[0].Name
	ed.startDate = q.Embedded.Events[0].Dates.Start.LocalDate
	ed.venueName = q.Embedded.Events[0].Embedded.Venues[0].Name
	ed.venueAddress = q.Embedded.Events[0].Embedded.Venues[0].Address.Line1
	ed.venueCity = q.Embedded.Events[0].Embedded.Venues[0].City.Name
	ed.eventId = q.Embedded.Events[0].ID
	ed.eventUrl = q.Embedded.Events[0].URL
	ed.logoUrl = q.Embedded.Events[0].Embedded.Attractions[0].Images[0].URL
	ed.venueId = q.Embedded.Events[0].ID

	d := insertData(config.DB, ed)

	fmt.Println(ed)
	if d != nil {
		fmt.Println(d)

	}

	//time.Sleep(120 * time.Second)

}

func insertData(s *sql.DB, i eventData) error {

	//Insert Data into Database

	_, err := s.Exec("INSERT events(name, start_date, venue_name , venue_address, venue_city,event_id, event_url, logo_url, venue_id) VALUES(?,?,?,?,?,?,?,?,?)", i.name, i.startDate, i.venueName, i.venueAddress, i.venueCity, i.eventId, i.eventUrl, i.logoUrl, i.venueId)
	log.Print("Successfully created DB record for incident info")

	return err

}
