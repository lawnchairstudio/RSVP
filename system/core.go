package system

import (
	"github.com/lawnchairstudios/RSVP/dal"
	"github.com/lawnchairstudios/RSVP/proto"
	"github.com/lawnchairstudios/RSVP/services"
)

type Application struct {
	ServerServices  *ServerServices
	postgresAdapter *dal.PostgresAdapter
}

type ServerServices struct {
	Invitees proto.RSVPServer
}

func (a *Application) Init() {
	dbURL := "postgres://localhost:5432/rsvp-dev?sslmode=disable"
	a.postgresAdapter = dal.NewPostgresAdapter(dbURL)

	// invitee gateway and service
	ig := dal.NewInviteeGateway(a.postgresAdapter)

	a.ServerServices = &ServerServices{}
	a.ServerServices.Invitees = services.NewInviteeServerService(ig)
}

func (a *Application) DeInit() {
	a.postgresAdapter.Conn().Close()
}
