package node

type Service interface {
	HandleIncomingMessages(msg Message, source string)
	SendMessage(msg Message)
}

type service struct {
	repo Repository
	transport Transport
}

func NewService(repo Repository, transport Transport) Service {
	return &service {
		repo: repo,
		transport: transport,
	}
}

func (s *service) HandleIncomingMessages(msg Message, source string) {
	if s.repo.HasSeenMessage(msg.Content) {
		return
	}
	s.repo.MarkMessageSeen(msg.Content)
	log.Printf("Received message from %s: %s", source, msg.Content)

	peers:= s.repo.GetPeer()
	for _, p:= range peers {
		if p!= source {
			s.transport.SendMessage(msg, p)
		}
	}
}

func (s *service) SendMessage(msg Message) {
	s.repo.MarkMessageSeen(msg.Content)
	peers := 	s.repo.GetPeer()

	for _, p := range peers {
		s.transport.SendMessage(msg, p)
	}
}