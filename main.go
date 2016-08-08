package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"github.com/deferpanic/deferclient/deferstats"
	"time"
	"encoding/json"
	"strconv"
)

type MLBApiResponse struct {
	Subject string `json:"subject"`
	Copyright string `json:"copyright"`
	Data struct {
			Games struct {
				      NextDayDate string `json:"next_day_date"`
				      ModifiedDate time.Time `json:"modified_date"`
				      Month string `json:"month"`
				      Year string `json:"year"`
				      Game []struct {
					      GameType string `json:"game_type"`
					      DoubleHeaderSw string `json:"double_header_sw"`
					      Location string `json:"location"`
					      AwayTime string `json:"away_time"`
					      Broadcast struct {
							       Away struct {
									    Tv string `json:"tv"`
									    Radio string `json:"radio"`
								    } `json:"away"`
							       Home struct {
									    Tv string `json:"tv"`
									    Radio string `json:"radio"`
								    } `json:"home"`
						       } `json:"broadcast"`
					      Time string `json:"time"`
					      HomeTime string `json:"home_time"`
					      HomeTeamName string `json:"home_team_name"`
					      Description string `json:"description"`
					      OriginalDate string `json:"original_date"`
					      HomeTeamCity string `json:"home_team_city"`
					      VenueID string `json:"venue_id"`
					      GamedaySw string `json:"gameday_sw"`
					      AwayWin string `json:"away_win"`
					      HomeGamesBackWildcard string `json:"home_games_back_wildcard"`
					      SavePitcher struct {
							       ID string `json:"id"`
							       Last string `json:"last"`
							       Saves string `json:"saves"`
							       Losses string `json:"losses"`
							       Era string `json:"era"`
							       NameDisplayRoster string `json:"name_display_roster"`
							       Number string `json:"number"`
							       Svo string `json:"svo"`
							       First string `json:"first"`
							       Wins string `json:"wins"`
						       } `json:"save_pitcher"`
					      AwayTeamID string `json:"away_team_id"`
					      TzHmLgGen string `json:"tz_hm_lg_gen"`
					      Status struct {
							       IsNoHitter string `json:"is_no_hitter"`
							       TopInning string `json:"top_inning"`
							       S string `json:"s"`
							       B string `json:"b"`
							       Reason string `json:"reason"`
							       Ind string `json:"ind"`
							       Status string `json:"status"`
							       IsPerfectGame string `json:"is_perfect_game"`
							       O string `json:"o"`
							       Inning string `json:"inning"`
							       InningState string `json:"inning_state"`
							       Note string `json:"note"`
						       } `json:"status"`
					      HomeLoss string `json:"home_loss"`
					      HomeGamesBack string `json:"home_games_back"`
					      HomeCode string `json:"home_code"`
					      AwaySportCode string `json:"away_sport_code"`
					      HomeWin string `json:"home_win"`
					      TimeHmLg string `json:"time_hm_lg"`
					      AwayNameAbbrev string `json:"away_name_abbrev"`
					      League string `json:"league"`
					      TimeZoneAwLg string `json:"time_zone_aw_lg"`
					      AwayGamesBack string `json:"away_games_back"`
					      HomeFileCode string `json:"home_file_code"`
					      GameDataDirectory string `json:"game_data_directory"`
					      TimeZone string `json:"time_zone"`
					      AwayLeagueID string `json:"away_league_id"`
					      HomeTeamID string `json:"home_team_id"`
					      Day string `json:"day"`
					      TimeAwLg string `json:"time_aw_lg"`
					      AwayTeamCity string `json:"away_team_city"`
					      TbdFlag string `json:"tbd_flag"`
					      TzAwLgGen string `json:"tz_aw_lg_gen"`
					      AwayCode string `json:"away_code"`
					      WinningPitcher struct {
							       ID string `json:"id"`
							       Last string `json:"last"`
							       Losses string `json:"losses"`
							       Era string `json:"era"`
							       Number string `json:"number"`
							       NameDisplayRoster string `json:"name_display_roster"`
							       First string `json:"first"`
							       Wins string `json:"wins"`
						       } `json:"winning_pitcher"`
					      GameMedia struct {
							       Media []struct {
								       Free string `json:"free,omitempty"`
								       Title string `json:"title,omitempty"`
								       Thumbnail string `json:"thumbnail"`
								       MediaState string `json:"media_state,omitempty"`
								       Start string `json:"start,omitempty"`
								       HasMlbtv string `json:"has_mlbtv,omitempty"`
								       CalendarEventID string `json:"calendar_event_id,omitempty"`
								       Enhanced string `json:"enhanced,omitempty"`
								       Type string `json:"type"`
								       Headline string `json:"headline,omitempty"`
								       ContentID string `json:"content_id,omitempty"`
								       TopicID string `json:"topic_id,omitempty"`
							       } `json:"media"`
						       } `json:"game_media"`
					      GameNbr string `json:"game_nbr"`
					      TimeDateAwLg string `json:"time_date_aw_lg"`
					      AwayGamesBackWildcard string `json:"away_games_back_wildcard"`
					      ScheduledInnings string `json:"scheduled_innings"`
					      Linescore struct {
							       Hr struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
								  } `json:"hr"`
							       E struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
								  } `json:"e"`
							       So struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
								  } `json:"so"`
							       R struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
									  Diff string `json:"diff"`
								  } `json:"r"`
							       Sb struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
								  } `json:"sb"`
							       Inning []struct {
								       Home string `json:"home,omitempty"`
								       Away string `json:"away"`
							       } `json:"inning"`
							       H struct {
									  Home string `json:"home"`
									  Away string `json:"away"`
								  } `json:"h"`
						       } `json:"linescore"`
					      VenueWChanLoc string `json:"venue_w_chan_loc"`
					      FirstPitchEt string `json:"first_pitch_et"`
					      AwayTeamName string `json:"away_team_name"`
					      HomeRuns struct {
							       Player struct {
									      StdHr string `json:"std_hr"`
									      Hr string `json:"hr"`
									      ID string `json:"id"`
									      Last string `json:"last"`
									      TeamCode string `json:"team_code"`
									      Inning string `json:"inning"`
									      Runners string `json:"runners"`
									      Number string `json:"number"`
									      NameDisplayRoster string `json:"name_display_roster"`
									      First string `json:"first"`
								      } `json:"player"`
						       } `json:"home_runs,omitempty"`
					      TimeDateHmLg string `json:"time_date_hm_lg"`
					      ID string `json:"id"`
					      HomeNameAbbrev string `json:"home_name_abbrev"`
					      TiebreakerSw string `json:"tiebreaker_sw"`
					      Ampm string `json:"ampm"`
					      HomeDivision string `json:"home_division"`
					      HomeTimeZone string `json:"home_time_zone"`
					      AwayTimeZone string `json:"away_time_zone"`
					      HmLgAmpm string `json:"hm_lg_ampm"`
					      HomeSportCode string `json:"home_sport_code"`
					      TimeDate string `json:"time_date"`
					      Links struct {
							       AwayAudio string `json:"away_audio"`
							       Wrapup string `json:"wrapup"`
							       Preview string `json:"preview"`
							       HomePreview string `json:"home_preview"`
							       AwayPreview string `json:"away_preview"`
							       TvStation string `json:"tv_station"`
							       HomeAudio string `json:"home_audio"`
							       Mlbtv string `json:"mlbtv"`
						       } `json:"links"`
					      HomeAmpm string `json:"home_ampm"`
					      GamePk string `json:"game_pk"`
					      Venue string `json:"venue"`
					      HomeLeagueID string `json:"home_league_id"`
					      VideoThumbnail string `json:"video_thumbnail"`
					      AwayLoss string `json:"away_loss"`
					      ResumeDate string `json:"resume_date"`
					      AwayFileCode string `json:"away_file_code"`
					      LosingPitcher struct {
							       ID string `json:"id"`
							       Last string `json:"last"`
							       Losses string `json:"losses"`
							       Era string `json:"era"`
							       Number string `json:"number"`
							       NameDisplayRoster string `json:"name_display_roster"`
							       First string `json:"first"`
							       Wins string `json:"wins"`
						       } `json:"losing_pitcher"`
					      AwLgAmpm string `json:"aw_lg_ampm"`
					      VideoThumbnails struct {
							       Thumbnail []struct {
								       Content string `json:"content"`
								       Height string `json:"height"`
								       Scenario string `json:"scenario"`
								       Width string `json:"width"`
							       } `json:"thumbnail"`
						       } `json:"video_thumbnails"`
					      TimeZoneHmLg string `json:"time_zone_hm_lg"`
					      AwayAmpm string `json:"away_ampm"`
					      Gameday string `json:"gameday"`
					      AwayDivision string `json:"away_division"`
				      } `json:"game"`
				      Day string `json:"day"`
			      } `json:"games"`
		} `json:"data"`
}

func unmarshallResponse(body []byte) (*MLBApiResponse, error) {
	var s = new(MLBApiResponse)
	err := json.Unmarshal(body, &s)
	if(err != nil){
		fmt.Println("whoops:", err)
	}
	return s, err
}

func mlbPing(w http.ResponseWriter, r *http.Request) {
	t := time.Now()
	d := int(t.Day())
	day := strconv.Itoa(d)
	m := int(t.Month())
	month := strconv.Itoa(m)
	y := int(t.Year())
	year := strconv.Itoa(y)
	url := "http://gd2.mlb.com/components/game/mlb/year_" + year + "/month_" + month + "/day_" + day + "/master_scoreboard.json"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}
	if len(body) != 0 {
		s, err := unmarshallResponse([]byte(body))
		fmt.Println("First game time: ", s)
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}
	}
}

func main() {
	dps := deferstats.NewClient("z57z3xsEfpqxpr0dSte0auTBItWBYa1c")
	go dps.CaptureStats()

	http.HandleFunc("/mlb", dps.HTTPHandlerFunc(mlbPing))
	http.ListenAndServe(":3000", nil)
}
