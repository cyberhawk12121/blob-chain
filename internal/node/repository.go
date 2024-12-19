package node

import "sync"

type Repository interface {
	AddPeer(peer string)
	GetPeer() []string

	HasSeenMessage(msg string) bool
	MarkMessageSeen(msg string)
}

type inMemoryRepository struct {
	mu           sync.RWMutex
	peers        []string
	seenMessages map[string]bool
}

func NewInMemoryRepository(initialPeers []string) Repository {
	return &inMemoryRepository{
		peers:        initialPeers,
		seenMessages: make(map[string]bool),
	}
}

func (r *inMemoryRepository) AddPeer(peer string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.peers = append(r.peers, peer)
}

func (r *inMemoryRepository) GetPeer() []string {
	r.mu.RLock()
	defer r.mu.RUnlock()
	cp:= make([]string, len(r.peers))
	copy(cp, r.peers)
	return cp
}

func (r *inMemoryRepository) HasSeenMessage(msg string) bool {
	r.mu.RLock()
	defer r.mu.RUnlock()
	return r.seenMessages[msg]
}

func (r *inMemoryRepository) MarkMessageSeen(msg string) {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.seenMessages[msg] = true
}
