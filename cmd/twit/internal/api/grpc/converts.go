package grpc

import (
	twit_pb "github.com/ZergsLaw/back-template/api/twit/v1"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
)

func toUser(u app.Twit) *twit_pb.Twit {
	return &twit_pb.Twit{
		Id:       u.ID.String(),
		AuthorId: u.AuthorID.String(),
		Text:     u.Text,
	}
}
