package f1com

import (
	"fmt"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lucas-ingemar/f1sched/internal/shared"
	"github.com/tidwall/gjson"
)

const (
	f1BaseUrl            = "https://www.formula1.com"
	f1RaceListUrlPattern = f1BaseUrl + "/en/racing/%d.html"
)

func (f F1com) RaceScheduleRaw(year int) (rs shared.RaceSchedule, err error) {
	maxYear := time.Now().Year()
	if time.Now().Month() == time.December {
		maxYear += 1
	}
	if year < 2018 || year > maxYear {
		return rs, fmt.Errorf("only years 2018 - %d are supported", maxYear)
	}

	rs.Year = year

	raceLinks, err := f.getRaceScheduleLinks(year)
	if err != nil {
		return rs, err
	}

	for idx, rSubUrl := range raceLinks {
		ri, err := f.getRaceScheduleData(fmt.Sprintf("%s/%s", f1BaseUrl, rSubUrl), idx+1)
		if err != nil {
			return shared.RaceSchedule{}, err
		}
		rs.Races = append(rs.Races, ri)
	}

	rs.FullInformation = true

	return
}

func (f F1com) getRaceScheduleLinks(year int) (links []string, err error) {
	u := fmt.Sprintf(f1RaceListUrlPattern, year)

	doc, err := openUrl(u)
	if err != nil {
		return nil, err
	}

	var raceNavigationItems string
	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if strings.Contains(s.Text(), "raceNavigationItems") {
			raceNavigationItems, _ = strings.CutSuffix(strings.TrimSpace(strings.Split(s.Text(), "=")[1]), ";")
		}
	})

	if raceNavigationItems == "" {
		return nil, fmt.Errorf("raceNavigationItems not found on url %s", u)
	}

	value := gjson.Get(raceNavigationItems, "#.path")
	for _, r := range value.Array() {
		links = append(links, r.String())
	}

	if len(links) == 0 {
		return nil, fmt.Errorf("race links not found on url %s", u)
	}

	return
}

func (f F1com) getRaceScheduleData(u string, round int) (ri shared.RaceInformation, err error) {
	doc, err := openUrl(u)
	if err != nil {
		return ri, err
	}

	var jsEventData gjson.Result
	doc.Find(`script[type="application/ld+json"]`).Each(func(_ int, s *goquery.Selection) {
		jsEventData = gjson.Parse(s.Text())
	})

	var jsLocData gjson.Result
	doc.Find("script").Each(func(_ int, s *goquery.Selection) {
		if strings.HasPrefix(strings.TrimSpace(s.Text()), "dataLayer") {
			parts := strings.Split(strings.TrimSpace(s.Text()), "=")
			if len(parts) != 2 {
				return
			}
			jsLocData = gjson.Parse(parts[1])
		}
	})

	ri.Round = round

	ri.StartTime, err = time.Parse(time.RFC3339, jsEventData.Get("startDate").String())
	if err != nil {
		return shared.RaceInformation{}, err
	}

	ri.EndTime, err = time.Parse(time.RFC3339, jsEventData.Get("endDate").String())
	if err != nil {
		return shared.RaceInformation{}, err
	}

	ri.Location.Country, err = gjsonGetString(jsLocData, "0.trackCountry")
	if err != nil {
		return shared.RaceInformation{}, err
	}

	ri.Location.City, err = gjsonGetString(jsLocData, "0.trackCity")
	if err != nil {
		return shared.RaceInformation{}, err
	}

	ri.Location.Circuit, err = gjsonGetString(jsLocData, "0.trackName")
	if err != nil {
		return shared.RaceInformation{}, err
	}

	ri.Name, err = gjsonGetString(jsLocData, "0.raceName")
	if err != nil {
		return shared.RaceInformation{}, err
	}

	subEvents := jsEventData.Get("subEvent")
	if len(subEvents.Array()) == 0 {
		return ri, fmt.Errorf("could not find subevents")
	}

	ri.FreePractice1, _ = f.getEvent(subEvents, "Practice 1")
	ri.FreePractice2, _ = f.getEvent(subEvents, "Practice 2")
	ri.FreePractice3, _ = f.getEvent(subEvents, "Practice 3")
	ri.SprintQualifying, _ = f.getEvent(subEvents, "Sprint Shootout")
	ri.Sprint, _ = f.getEvent(subEvents, "Sprint")
	ri.Qualifying, _ = f.getEvent(subEvents, "Qualifying")
	ri.Race, _ = f.getEvent(subEvents, "Race")

	if ri.SprintQualifying == nil {
		ri.SprintQualifying, _ = f.getEvent(subEvents, "Sprint Qualifying")
	}

	if ri.FreePractice1 != nil &&
		ri.FreePractice2 != nil &&
		ri.FreePractice3 != nil &&
		ri.Qualifying != nil &&
		ri.Race != nil {
		ri.Type = shared.NormalRace
	} else if ri.FreePractice1 != nil &&
		ri.SprintQualifying != nil &&
		ri.Qualifying != nil &&
		ri.Sprint != nil &&
		ri.Race != nil {
		ri.Type = shared.SprintRace
	} else {
		return ri, fmt.Errorf("could not find a race type for %s", ri.Name)
	}

	return
}

func (f F1com) getEvent(r gjson.Result, name string) (event *shared.RaceEvent, err error) {
	var matchedObject gjson.Result
	r.ForEach(func(_, value gjson.Result) bool {
		if strings.HasPrefix(strings.ReplaceAll(value.Get("name").String(), " ", ""), strings.ReplaceAll(name, " ", "")+"-") {
			matchedObject = value
			return false
		}
		return true
	})

	if matchedObject.String() == "" {
		return nil, fmt.Errorf("could not find a match for event %s", name)
	}
	e := shared.RaceEvent{}

	e.StartTime, err = time.Parse(time.RFC3339, matchedObject.Get("startDate").String())
	if err != nil {
		return nil, err
	}

	e.EndTime, err = time.Parse(time.RFC3339, matchedObject.Get("endDate").String())
	if err != nil {
		return nil, err
	}

	return &e, nil
}

func gjsonGetString(r gjson.Result, path string) (string, error) {
	s := r.Get(path).String()
	if s == "" {
		return "", fmt.Errorf("'%s' not found", path)
	}
	return s, nil
}
