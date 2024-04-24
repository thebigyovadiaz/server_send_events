package handlers

import (
	"encoding/json"
	"github.com/thebigyovadiaz/server_send_events/handlers/events"
	"net/http"
)

func InitRoutes(r *http.ServeMux) {
	handlerEvents := events.NewHandlerEvent()

	r.HandleFunc("/notify", handlerEvents.HandlerNotify)
	r.HandleFunc("/testI", HandlerTestI(handlerEvents))
	r.HandleFunc("/testII", HandlerTestII(handlerEvents))
	r.Handle("/", http.FileServer(http.Dir("./public")))
}

func HandlerTestI(notifier *events.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]any{}
		json.NewDecoder(r.Body).Decode(&data)

		notifier.Broadcast(events.EventMessage{
			EventName: "greeting",
			Data:      data,
		})
	}
}

func HandlerTestII(notifier *events.HandlerEvent) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var data = map[string]any{}
		json.NewDecoder(r.Body).Decode(&data)

		notifier.Broadcast(events.EventMessage{
			EventName: "jumping",
			Data:      data,
		})
	}
}
