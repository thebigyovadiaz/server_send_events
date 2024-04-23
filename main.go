package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/events", eventsHandler)
	http.ListenAndServe(":8080", nil)
}

func eventsHandler(w http.ResponseWriter, r *http.Request) {
	// Set CORS headers to allow all origins
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Expose-Headers", "Content-Type")

	w.Header().Set("Content-Type", "text/event-stream")
	w.Header().Set("Cache-Control", "no-cache")
	w.Header().Set("Connection", "keep-alive")

	// Simulate sending events (replace with real data in the future)
	for i := 0; i < 10; i++ {
		fmt.Fprintf(w, "data: %s\n\n", fmt.Sprintf("Event %d", i+1))
		time.Sleep(2 * time.Second)

		if i == 9 {
			fmt.Fprintf(w, "data: %s\n\n", fmt.Sprint("Event sent successfully"))
		}

		w.(http.Flusher).Flush()
	}

	// Simulate closing the connection
	closeNotify := w.(http.CloseNotifier).CloseNotify()
	<-closeNotify
}
