package events

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

type clients struct {
	Id          string `json:"id"`
	SendMessage chan EventMessage
}

func newClients(id string) *clients {
	return &clients{
		Id:          id,
		SendMessage: make(chan EventMessage),
	}
}

func (c *clients) OnLine(ctx context.Context, w io.Writer, flusher http.Flusher) {
	for {
		select {
		case msg := <-c.SendMessage:
			data, err := json.Marshal(msg)
			if err != nil {
				log.Println(err)
			}

			const format = "event:%s\ndata:%s\n\n"
			fmt.Fprintf(w, format, msg.EventName, string(data))
			flusher.Flush()
		case <-ctx.Done():
			return
		}
	}
}
