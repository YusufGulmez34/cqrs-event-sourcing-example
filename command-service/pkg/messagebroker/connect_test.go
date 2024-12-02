package messagebroker

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRabbitMQConnect(t *testing.T) {
	t.Run("Connect", func(t *testing.T) {
		_, err := ConnectRabbitMQ()
		require.Nil(t, err)
	})
}
