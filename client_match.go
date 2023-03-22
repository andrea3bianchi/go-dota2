package dota2

import (
	"github.com/andrea3bianchi/go-dota2/events"
	"github.com/faceit/go-steam/protocol/gamecoordinator"
)

// handleMatchSignedOut handles an incoming steam datagram ticket.
func (d *Dota2) handleMatchSignedOut(packet *gamecoordinator.GCPacket) error {
	ev := &events.MatchSignedOut{}
	if err := d.unmarshalBody(packet, &ev.CMsgGCToClientMatchSignedOut); err != nil {
		return err
	}

	d.emit(ev)
	return nil
}
