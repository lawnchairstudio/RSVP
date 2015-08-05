package services_test

import (
	"errors"

	"github.com/lawnchairstudios/RSVP/proto"
	"github.com/lawnchairstudios/RSVP/services"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Setup mock gateway
type FakeInviteeGateway struct {
	GetFromIDError bool
}

func (fig FakeInviteeGateway) GetFromID(id string) (proto.Invitee, error) {
	if fig.GetFromIDError {
		return proto.Invitee{}, errors.New("")
	} else {
		return proto.Invitee{
			InviteeId: &id,
		}, nil
	}
}

var _ = Describe("Invitee", func() {
	var (
		iSvrSvc *services.InviteeServerService
		iGtw    FakeInviteeGateway
	)

	BeforeEach(func() {
		iGtw = FakeInviteeGateway{}
		iSvrSvc = services.NewInviteeServerService(&iGtw)
	})

	Describe("Getting an invitee", func() {
		Context("with a valid id", func() {
			It("should return an invitee", func() {
				id := "asdf"
				obj, err := iSvrSvc.GetInvitee(nil, &proto.InviteeRequest{Id: &id})

				Expect(obj).To(Equal(&proto.Invitee{InviteeId: &id}))
				Expect(err).NotTo(HaveOccurred())
			})
		})

		Context("with an invalid id", func() {
			It("should return an error", func() {
				id := "asdf"
				iGtw.GetFromIDError = true

				_, err := iSvrSvc.GetInvitee(nil, &proto.InviteeRequest{Id: &id})

				Expect(err).To(HaveOccurred())
			})
		})
	})
})
