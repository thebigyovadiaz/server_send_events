package events

import (
	"log"
	"net/http"
	"sync"
)

type EventMessage struct {
	EventName string
	Data      any
}

type HandlerEvent struct {
	m       sync.Mutex
	clients map[string]*clients
}

func NewHandlerEvent() *HandlerEvent {
	return &HandlerEvent{
		clients: make(map[string]*clients),
	}
}

func (hE *HandlerEvent) HandlerNotify(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	id := r.URL.Query().Get("id")
	if id == "" {
		log.Println("id is empty")
	}

	flusher, ok := w.(http.Flusher)
	if !ok {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	client := newClients(id)
	hE.addClient(client)
	log.Println("Client connected:", id)

	client.OnLine(r.Context(), w, flusher)
	log.Println("Client disconnected:", id)
	hE.removeClient(id)

}

func (hE *HandlerEvent) addClient(c *clients) {
	hE.m.Lock()
	defer hE.m.Unlock()
	hE.clients[c.Id] = c
}

func (hE *HandlerEvent) removeClient(id string) {
	hE.m.Lock()
	defer hE.m.Unlock()
	delete(hE.clients, id)
}

func (hE *HandlerEvent) Broadcast(m EventMessage) {
	hE.m.Lock()
	defer hE.m.Unlock()
	for _, c := range hE.clients {
		c.SendMessage <- m
	}
}
