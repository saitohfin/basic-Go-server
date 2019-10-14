package poker

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"html/template"
	"net/http"
)

type PlayerServer struct {
	store PlayerStore
	http.Handler
	template *template.Template
}

type PlayerStore interface {
	GetPlayerScore(name string) int

	RecordWin(name string)

	GetLeague() League
}

type Player struct {
	Name string
	Wins int
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}
const htmlTemplatePath = "game.html"

func PlayerServerFactory(store PlayerStore) (*PlayerServer, error) {

	p := new(PlayerServer)

	tmpl, err := template.ParseFiles("game.html")

	if err != nil {
		return nil, fmt.Errorf("problem opening %s %v", htmlTemplatePath, err)
	}

	p.template = tmpl
	p.store = store

	router := http.NewServeMux()
	router.Handle("/league", http.HandlerFunc(p.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(p.playersHandler))
	router.Handle("/game", http.HandlerFunc(p.game))
	router.Handle("/ws", http.HandlerFunc(p.webSocket))

	p.Handler = router

	return p, nil
}

func (p *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	league := p.store.GetLeague()
	w.Header().Set("content-type", "application/json")
	json.NewEncoder(w).Encode(league)
	w.WriteHeader(http.StatusOK)
}

func (p *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	player := r.URL.Path[len("/players/"):]

	switch r.Method {
	case http.MethodPost:
		p.processWin(w, player)
	case http.MethodGet:
		p.showScore(w, player)
	}
}

func (p *PlayerServer) game(w http.ResponseWriter, r *http.Request) {
	p.template.Execute(w, nil)
}

func (p *PlayerServer) webSocket(w http.ResponseWriter, r *http.Request) {
	conn, _ := wsUpgrader.Upgrade(w, r, nil)
	_, winnerMsg, _ := conn.ReadMessage()
	p.store.RecordWin(string(winnerMsg))
}

func (p *PlayerServer) processWin(w http.ResponseWriter, player string) {
	p.store.RecordWin(player)
	w.WriteHeader(http.StatusAccepted)
}

func (p *PlayerServer) showScore(w http.ResponseWriter, player string) {

	score := p.store.GetPlayerScore(player)

	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}

	fmt.Fprint(w, score)
}
