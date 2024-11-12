package grpc

import (
	"context"
	"google.golang.org/grpc"
	"log"
	"signal-server/pkg/sfu"
)

type SFUClient struct {
	client sfu.SFUServiceClient
}

func NewSFUClient(address string) *SFUClient {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatal("Could not connect to SFU:", err)
	}

	client := sfu.NewSFUServiceClient(conn)
	return &SFUClient{client: client}
}

func (c *SFUClient) SendOffer(sessionID, userID, offer string) (*sfu.StreamResponse, error) {
	req := &sfu.StreamRequest{
		SessionId:  sessionID,
		UserId:     userID,
		StreamInfo: offer,
	}
	return c.client.RegisterStream(context.Background(), req)
}

func (c *SFUClient) SendAnswer(sessionID, userID, answer string) (*sfu.StreamResponse, error) {
	req := &sfu.StreamRequest{
		SessionId:  sessionID,
		UserId:     userID,
		StreamInfo: answer,
	}
	return c.client.RegisterStream(context.Background(), req)
}

func (c *SFUClient) SendIceCandidate(sessionID, userID, candidate string) (*sfu.IceCandidateResponse, error) {
	req := &sfu.IceCandidateRequest{
		SessionId: sessionID,
		UserId:    userID,
		Candidate: candidate,
	}
	return c.client.HandleIceCandidate(context.Background(), req)
}
