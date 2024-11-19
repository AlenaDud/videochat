package services

import (
	"signal-server/internal/clients/sfu"
)

type SignalService struct {
	sfuClient *sfu.SFUClient
}

func NewSignalService(sfuClient *sfu.SFUClient) *SignalService {
	return &SignalService{sfuClient: sfuClient}
}

func (s *SignalService) HandleSDP(sdp []byte) error {
	// Отправка SDP в SFU сервис
	return s.sfuClient.SendSDP(sdp)
}
