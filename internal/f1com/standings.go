package f1com

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/PuerkitoBio/goquery"
	"github.com/lucas-ingemar/f1sched/internal/shared"
)

const (
	f1DriverStandingsUrlPattern = f1BaseUrl + "/en/results.html/%d/drivers.html"
)

func (f F1com) DriverStandingsRaw(rs shared.RaceSchedule) (ds shared.DriverStandings, err error) {
	u := fmt.Sprintf(f1DriverStandingsUrlPattern, rs.Year)

	doc, err := openUrl(u)
	if err != nil {
		return
	}

	parseErrs := []error{}
	doc.Find("tbody tr").Each(func(_ int, s *goquery.Selection) {
		sd, err := f.parseDriverRow(s)
		parseErrs = append(parseErrs, err)
		ds.Entries = append(ds.Entries, sd)
	})

	ds.UpdatedAt = time.Now()

	if errors.Join(parseErrs...) != nil {
		return ds, errors.Join(parseErrs...)
	}

	return
}

func (f F1com) parseDriverRow(s *goquery.Selection) (sd shared.StandingDriver, err error) {
	errs := []error{}
	s.Find("td:not(.limiter)").Each(func(idx int, ss *goquery.Selection) {
		switch idx {
		case 0:
			sd.Position, err = strconv.Atoi(ss.Text())
			errs = append(errs, err)

		case 1:
			sd.DriverFirstName = strings.TrimSpace(ss.Find(".hide-for-tablet").Text())
			sd.DriverLastName = strings.TrimSpace(ss.Find(".hide-for-mobile").Text())
			sd.DriverShort = strings.TrimSpace(ss.Find(".hide-for-desktop").Text())

		case 2:
			sd.Nationality = strings.TrimSpace(ss.Text())

		case 3:
			sd.Team = strings.TrimSpace(ss.Find("a").Text())

		case 4:
			sd.Points, err = strconv.Atoi(ss.Text())
			errs = append(errs, err)
		}
	})
	err = errors.Join(errs...)
	return
}
