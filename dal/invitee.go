package dal

import "github.com/lawnchairstudios/RSVP/proto"

type InviteeGateway struct {
	adapter SQLAdapter
}

func NewInviteeGateway(sql SQLAdapter) *InviteeGateway {
	return &InviteeGateway{adapter: sql}
}

func (ig *InviteeGateway) GetFromID(id string) (proto.Invitee, error) {
	findMe := proto.Invitee{InviteeId: &id}
	db := ig.adapter.Conn().First(&findMe)

	return findMe, db.Error
}
