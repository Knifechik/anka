package grpc

import (
	twit_pb "github.com/ZergsLaw/back-template/api/twit/v1"
	"github.com/ZergsLaw/back-template/cmd/twit/internal/app"
)

func toTwit(u app.Twit) *twit_pb.Twit {
	return &twit_pb.Twit{
		Id:       u.ID.String(),
		AuthorId: u.AuthorID.String(),
		Text:     u.Text,
	}
}

func toTwits(u []app.Twit) []*twit_pb.Twit {
	var twits = make([]*twit_pb.Twit, len(u))
	for i, e := range u {
		twits[i] = toTwit(e)
	}
	return twits
}
