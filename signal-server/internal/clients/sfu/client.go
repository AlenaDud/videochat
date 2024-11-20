package sfu

//
//import (
//	"context"
//	"google.golang.org/grpc"
//	"signal-server/pkg/sfu"
//)
//
//type SFUClient struct {
//	client sfu.SFUServiceClient
//}
//
//func NewSFUClient(addr string) (*SFUClient, error) {
//	conn, err := grpc.Dial(addr, grpc.WithInsecure())
//	if err != nil {
//		return nil, err
//	}
//	return &SFUClient{client: sfu.NewSFUServiceClient(conn)}, nil
//}
//
//func (c *SFUClient) SendSDP(sdp []byte) error {
//	req := &sfu.SDPRequest{Sdp: string(sdp)}
//	_, err := c.client.SendSDP(context.Background(), req)
//	return err
//}
