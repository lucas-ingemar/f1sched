package countries

import "github.com/lucas-ingemar/f1sched/internal/shared"

var (
	countryAltNames = map[string]string{
		"United States": "USA",
	}

	countries = []shared.Country{
		{
			Name: "Bahrain",
			BgColor: shared.Color{
				Color1: "#CE1126",
				Color2: "#FFFFFF",
				Color3: "#CE1126",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Saudi Arabia",
			BgColor: shared.Color{
				Color1: "#165d31",
				Color2: "#FFFFFF",
				Color3: "#165d31",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Australia",
			BgColor: shared.Color{
				Color1: "#012169",
				Color2: "#FFFFFF",
				Color3: "#E4002B",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Azerbaijan",
			BgColor: shared.Color{
				Color1: "#0092BC",
				Color2: "#E4002B",
				Color3: "#00AF66",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#FFFFFF",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "USA",
			BgColor: shared.Color{
				Color1: "#B31942",
				Color2: "#FFFFFF",
				Color3: "#0A3161",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Monaco",
			BgColor: shared.Color{
				Color1: "#CE1126",
				Color2: "#FFFFFF",
				Color3: "#CE1126",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Spain",
			BgColor: shared.Color{
				Color1: "#AA151B",
				Color2: "#F1BF00",
				Color3: "#AA151B",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Canada",
			BgColor: shared.Color{
				Color1: "#D80621",
				Color2: "#FFFFFF",
				Color3: "#D80621",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Austria",
			BgColor: shared.Color{
				Color1: "#EF3340",
				Color2: "#FFFFFF",
				Color3: "#EF3340",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "UK",
			BgColor: shared.Color{
				Color1: "#012169",
				Color2: "#FFFFFF",
				Color3: "#C8102E",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Hungary",
			BgColor: shared.Color{
				Color1: "#CE2939",
				Color2: "#FFFFFF",
				Color3: "#477050",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Belgium",
			BgColor: shared.Color{
				Color1: "#2D2926",
				Color2: "#FFCD00",
				Color3: "#C8102E",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Netherlands",
			BgColor: shared.Color{
				Color1: "#C8102E",
				Color2: "#FFFFFF",
				Color3: "#003DA5",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Italy",
			BgColor: shared.Color{
				Color1: "#008C45",
				Color2: "#F4F9FF",
				Color3: "#CD212A",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Singapore",
			BgColor: shared.Color{
				Color1: "#C73b3C",
				Color2: "#FFFFFF",
				Color3: "#C73b3C",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Japan",
			BgColor: shared.Color{
				Color1: "#BC002D",
				Color2: "#FFFFFF",
				Color3: "#BC002D",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Qatar",
			BgColor: shared.Color{
				Color1: "#8A1538",
				Color2: "#FFFFFF",
				Color3: "#8A1538",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Mexico",
			BgColor: shared.Color{
				Color1: "#006341",
				Color2: "#FFFFFF",
				Color3: "#C8102E",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "Brazil",
			BgColor: shared.Color{
				Color1: "#009739",
				Color2: "#FEDD00",
				Color3: "#012169",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "UAE",
			BgColor: shared.Color{
				Color1: "#EF3340",
				Color2: "#009739",
				Color3: "#000000",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#FFFFFF",
				Color3: "#FFFFFF",
			},
		},
		{
			Name: "China",
			BgColor: shared.Color{
				Color1: "#EE1C25",
				Color2: "#FFFF00",
				Color3: "#EE1C25",
			},
			FgColor: shared.Color{
				Color1: "#FFFFFF",
				Color2: "#000000",
				Color3: "#FFFFFF",
			},
		},
	}
)

func GetCountry(countryName string) shared.Country {
	cname, ok := countryAltNames[countryName]
	if !ok {
		cname = countryName
	}

	for _, c := range countries {
		if c.Name == cname {
			return c
		}
	}
	return shared.Country{
		Name: "Unknown",
		BgColor: shared.Color{
			Color1: "#FFFFFF",
			Color2: "#FFFFFF",
			Color3: "#FFFFFF",
		},
		FgColor: shared.Color{
			Color1: "#000000",
			Color2: "#000000",
			Color3: "#000000",
		},
	}
}
