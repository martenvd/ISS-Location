package main

// locationData...
type locationData struct {
	Documentation string `json:"documentation"`
	Licenses      []struct {
		Name string `json:"name"`
		URL  string `json:"url"`
	} `json:"licenses"`
	Rate struct {
		Limit     int64 `json:"limit"`
		Remaining int64 `json:"remaining"`
		Reset     int64 `json:"reset"`
	} `json:"rate"`
	Results []struct {
		Annotations struct {
			DMS struct {
				Lat string `json:"lat"`
			} `json:"DMS"`
			MGRS       string `json:"MGRS"`
			Maidenhead string `json:"Maidenhead"`
			Mercator   struct {
				X float64 `json:"x"`
				Y float64 `json:"y"`
			} `json:"Mercator"`
			OSM struct {
				NoteURL string `json:"note_url"`
				URL     string `json:"url"`
			} `json:"OSM"`
			Geohash  string  `json:"geohash"`
			Qibla    float64 `json:"qibla"`
			Roadinfo struct {
				DriveOn string `json:"drive_on"`
				SpeedIn string `json:"speed_in"`
			} `json:"roadinfo"`
			Sun struct {
				Rise struct {
					Apparent     int64 `json:"apparent"`
					Astronomical int64 `json:"astronomical"`
					Civil        int64 `json:"civil"`
					Nautical     int64 `json:"nautical"`
				} `json:"rise"`
				Set struct {
					Apparent     int64 `json:"apparent"`
					Astronomical int64 `json:"astronomical"`
					Civil        int64 `json:"civil"`
					Nautical     int64 `json:"nautical"`
				} `json:"set"`
			} `json:"sun"`
			Timezone struct {
				Name         string `json:"name"`
				NowInDst     int64  `json:"now_in_dst"`
				OffsetSec    int64  `json:"offset_sec"`
				OffsetString int64  `json:"offset_string"`
				ShortName    string `json:"short_name"`
			} `json:"timezone"`
			What3words struct {
				Words string `json:"words"`
			} `json:"what3words"`
			Wikidata string `json:"wikidata"`
		} `json:"annotations"`
		Bounds struct {
			Northeast struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"northeast"`
			Southwest struct {
				Lat float64 `json:"lat"`
				Lng float64 `json:"lng"`
			} `json:"southwest"`
		} `json:"bounds"`
		Components struct {
			Category    string `json:"_category"`
			Type        string `json:"_type"`
			BodyOfWater string `json:"body_of_water"`
		} `json:"components"`
		Confidence int64  `json:"confidence"`
		Formatted  string `json:"formatted"`
		Geometry   struct {
			Lat float64 `json:"lat"`
			Lng float64 `json:"lng"`
		} `json:"geometry"`
	} `json:"results"`
	Status struct {
		Code    int64  `json:"code"`
		Message string `json:"message"`
	} `json:"status"`
	StayInformed struct {
		Blog    string `json:"blog"`
		Twitter string `json:"twitter"`
	} `json:"stay_informed"`
	Thanks    string `json:"thanks"`
	Timestamp struct {
		CreatedHTTP string `json:"created_http"`
		CreatedUnix int64  `json:"created_unix"`
	} `json:"timestamp"`
	TotalResults int64 `json:"total_results"`
}
