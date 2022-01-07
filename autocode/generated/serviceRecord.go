package generated


type ServiceRecordMultiplayer struct {
	Additional struct {
		Filter   string `json:"filter"`
		Gamertag string `json:"gamertag"`
	} `json:"additional"`
	Data struct {
		Core struct {
			Breakdowns struct {
				Assists struct {
					Callouts int `json:"callouts"`
					Driver   int `json:"driver"`
					Emp      int `json:"emp"`
				} `json:"assists"`
				Kills struct {
					Grenades     int `json:"grenades"`
					Headshots    int `json:"headshots"`
					Melee        int `json:"melee"`
					PowerWeapons int `json:"power_weapons"`
				} `json:"kills"`
				Matches struct {
					Draws  int `json:"draws"`
					Left   int `json:"left"`
					Losses int `json:"losses"`
					Wins   int `json:"wins"`
				} `json:"matches"`
				Medals []struct {
					Count     int   `json:"count"`
					ID        int64 `json:"id"`
					ImageUrls struct {
						Large  string `json:"large"`
						Medium string `json:"medium"`
						Small  string `json:"small"`
					} `json:"image_urls"`
					Name string `json:"name"`
				} `json:"medals"`
			} `json:"breakdowns"`
			Damage struct {
				Average int `json:"average"`
				Dealt   int `json:"dealt"`
				Taken   int `json:"taken"`
			} `json:"damage"`
			Kda   float64 `json:"kda"`
			Kdr   float64 `json:"kdr"`
			Shots struct {
				Accuracy float64 `json:"accuracy"`
				Fired    int     `json:"fired"`
				Landed   int     `json:"landed"`
				Missed   int     `json:"missed"`
			} `json:"shots"`
			Summary struct {
				Assists   int `json:"assists"`
				Betrayals int `json:"betrayals"`
				Deaths    int `json:"deaths"`
				Kills     int `json:"kills"`
				Medals    int `json:"medals"`
				Suicides  int `json:"suicides"`
				Vehicles  struct {
					Destroys int `json:"destroys"`
					Hijacks  int `json:"hijacks"`
				} `json:"vehicles"`
			} `json:"summary"`
			TotalScore int `json:"total_score"`
		} `json:"core"`
		MatchesPlayed int `json:"matches_played"`
		TimePlayed    struct {
			Human   string `json:"human"`
			Seconds int    `json:"seconds"`
		} `json:"time_played"`
		WinRate float64 `json:"win_rate"`
	} `json:"data"`
}
