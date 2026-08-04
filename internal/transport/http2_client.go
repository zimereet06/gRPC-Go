package transport

import (
	"context"
	"errors"
	"io"
)

// handleGoAway processes the GOAWAY frame and ensures all active streams are closed,
// unblocking any goroutines waiting on flow control or write buffers.
func (t *http2Client) handleGoAway() {
	t.mu.Lock()
	defer t.mu.Unlock()

	for _, s := range t.activeStreams {
		// Signal the stream that the transport is closing.
		s.writeQuota.close() // Ensure blocked writers are unblocked
		s.close(errors.New("transport is closing due to GOAWAY"))
	}
}
