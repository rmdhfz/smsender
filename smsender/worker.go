package smsender

import (
	log "github.com/Sirupsen/logrus"
)

type worker struct {
	id     int
	sender *Sender
}

func (w worker) process(msg *Message) {
	log.Infof("worker '%d' process: %+v", w.id, msg)

	for _, r := range w.sender.routes {
		if r.Match(msg.To) {
			if msg.From == "" && r.From != "" {
				msg.From = r.From
			}
			msg.Route = r.Name

			r.Broker.Send(*msg)
			return
		}
	}
	w.sender.GetBroker(DefaultBroker).Send(*msg)
}
