//go:generate mockery --name=(.+)Mock --case=underscore

package tcp

import (
	"net"

	"github.com/Tsapen/wow/internal/wow"
)

type ConnMock interface {
	net.Conn
}

type ChallengerMock interface {
	wow.Challenger
}

type StorageMock interface {
	wow.Storage
}
