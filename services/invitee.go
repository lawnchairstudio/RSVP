package services

import (
	"github.com/lawnchairstudios/RSVP/proto"
	"golang.org/x/net/context"
)

type InviteeRepository interface {
	// GetFromID gets a single invitee based on the id that is passed in. It returns the
	// invitee related to the id as well as an error if something goes wrong.
	GetFromID(string) (proto.Invitee, error)
}

type InviteeServerService struct {
	inviteeGtw InviteeRepository
}

func NewInviteeServerService(gateway InviteeRepository) *InviteeServerService {
	return &InviteeServerService{
		inviteeGtw: gateway,
	}
}

func (iss InviteeServerService) GetInvitee(ctx context.Context, ir *proto.InviteeRequest) (*proto.Invitee, error) {
	invitee, err := iss.inviteeGtw.GetFromID(ir.GetId())

	return &invitee, err
}
