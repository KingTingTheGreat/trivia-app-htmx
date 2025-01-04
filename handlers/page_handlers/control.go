package page_handlers

import (
	"log"
	"net/http"
	"trivia-app/dlog"
	"trivia-app/handlers"
	"trivia-app/handlers/ws_handlers"
)

type ControlData struct {
	Error       string
	Leaderboard ws_handlers.LeaderboardData
	BuzzedIn    ws_handlers.LeaderboardData
	PlayerList  []string
}

func Control(w http.ResponseWriter, r *http.Request) {
	dlog.DLog("control handler")
	playerList := ws_handlers.AllPlayerNames()
	data := ControlData{
		Leaderboard: ws_handlers.LeaderboardData{
			TableId: "leaderboard",
			Title:   "Leaderboard",
			Headers: []string{
				"Name",
				"Score",
			},
			RowData:  ws_handlers.MakeLeaderboard(),
			Endpoint: "/leaderboard-ws",
		},
		BuzzedIn: ws_handlers.LeaderboardData{
			TableId: "buzzed-in",
			Title:   "Buzzed In",
			Headers: []string{
				"Name",
				"Time",
			},
			RowData:  ws_handlers.MakeBuzzedIn(),
			Endpoint: "/buzzed-in-ws",
		},
		PlayerList: playerList,
	}

	log.Println(playerList)

	handlers.RenderTemplate(w, "control.html", data)
}
