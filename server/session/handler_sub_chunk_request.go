package session

import (
	"github.com/df-mc/dragonfly/server/world"
	"github.com/sandertv/gophertunnel/minecraft/protocol/packet"
)

// SubChunkRequestHandler handles sub-chunk requests from the client. The server will respond with a packet containing
// the requested sub-chunks.
type SubChunkRequestHandler struct{}

// Handle ...
func (*SubChunkRequestHandler) Handle(p packet.Packet, s *Session) error {
	// s.log.Debugf("before handle subchunk")
	pk := p.(*packet.SubChunkRequest)
	// s.log.Debugf("try handle subchunk %T", pk)
	s.ViewSubChunks(world.SubChunkPos(pk.Position), pk.Offsets)
	return nil
}
