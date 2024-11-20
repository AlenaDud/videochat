package services

type SignalService struct {
}

func NewSignalService() *SignalService {
	return &SignalService{}
}

//
//func (s *SignalService) HandleSDP(sdp []byte) error {
//	// Отправка SDP в SFU сервис
//	return s.sfuClient.SendSDP(sdp)
//}
