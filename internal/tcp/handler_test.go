package tcp

import (
	"context"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"github.com/Tsapen/wow/internal/config"
	"github.com/Tsapen/wow/internal/tcp/mocks"
	"github.com/Tsapen/wow/internal/wow"
)

type testCase struct {
	name string

	setUp func(conn *mocks.ConnMock, st *mocks.StorageMock, ch *mocks.ChallengerMock)
}

func TestHandle(t *testing.T) {
	testCases := []*testCase{{
		name: "success",
		setUp: func(conn *mocks.ConnMock, st *mocks.StorageMock, ch *mocks.ChallengerMock) {
			conn.On("SetReadDeadline", mock.Anything).Return(nil).Once()
			conn.On("SetWriteDeadline", mock.Anything).Return(nil).Once()

			ch.On("Generate", mock.Anything).Return("secret", nil).Once()
			conn.On("Write", mock.Anything).Return(10, nil)

			solution := "2bb80d537b1da3e38bd30361aa855686bde0eacd7162fef6a25fe97bf527a25b"
			sr := []byte("{\"solution\":\"" + solution + "\"}")
			conn.On("Read", mock.MatchedBy(func(buffer interface{}) bool {
				copy(buffer.([]byte), []byte(sr))

				return true
			})).Return(len(sr), nil).Once()
			ch.On("Verify", "secret", solution).Return(nil).Once()
			st.On("Quote", mock.Anything).Return("quote 1", nil).Once()
			conn.On("Write", mock.Anything).Return(10, nil)

			conn.On("Close", mock.Anything).Return(nil).Once()
		},
	}, {
		name: "json error",
		setUp: func(conn *mocks.ConnMock, st *mocks.StorageMock, ch *mocks.ChallengerMock) {
			conn.On("SetReadDeadline", mock.Anything).Return(nil).Once()
			conn.On("SetWriteDeadline", mock.Anything).Return(nil).Once()

			ch.On("Generate", mock.Anything).Return("secret", nil).Once()
			conn.On("Write", mock.Anything).Return(10, nil)

			sr := "incorrect json"
			conn.On("Read", mock.MatchedBy(func(buffer interface{}) bool {
				copy(buffer.([]byte), []byte(sr))

				return true
			})).Return(len(sr), nil).Once()

			conn.On("Close", mock.Anything).Return(nil).Once()
		},
	}, {
		name: "verification error",
		setUp: func(conn *mocks.ConnMock, st *mocks.StorageMock, ch *mocks.ChallengerMock) {
			conn.On("SetReadDeadline", mock.Anything).Return(nil).Once()
			conn.On("SetWriteDeadline", mock.Anything).Return(nil).Once()

			ch.On("Generate", mock.Anything).Return("secret", nil).Once()
			conn.On("Write", mock.Anything).Return(10, nil)

			sr := "{\"solution\":\"wrong_result\"}"
			conn.On("Read", mock.MatchedBy(func(buffer interface{}) bool {
				copy(buffer.([]byte), []byte(sr))

				return true
			})).Return(len(sr), nil).Once()
			ch.On("Verify", "secret", "wrong_result").Return(wow.ErrWrongSolution).Once()

			conn.On("Close", mock.Anything).Return(nil).Once()
		},
	}}

	ctx := context.Background()
	for _, tc := range testCases {
		conn := mocks.NewConnMock(t)
		storage := mocks.NewStorageMock(t)
		challenger := mocks.NewChallengerMock(t)
		tc.setUp(conn, storage, challenger)

		server, err := NewServer(
			&config.ServerConfig{Timeout: time.Second},
			challenger,
			storage,
		)
		assert.NoError(t, err)

		ctx = wow.WithReqID(ctx, uuid.NewString())
		server.handle(ctx, conn)
	}
}
