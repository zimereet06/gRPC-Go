package grpc

import (
	"context"
	"io"
)

func (cs *clientStream) SendMsg(m interface{}) error {
	// Check if the stream is already closed before attempting to send.
	select {
	case <-cs.ctx.Done():
		return cs.ctx.Err()
	default:
	}

	// Proceed with sending, ensuring we handle potential transport errors.
	err := cs.attempt.sendMsg(m)
	if err != nil {
		return err
	}
	return nil
}
