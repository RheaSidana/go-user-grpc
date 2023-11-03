package initializer

import (
	"os"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateNewGrpcServer(t *testing.T) {
	os.Setenv("NETWORK", "tcp")
	os.Setenv("PORT", "50051")

	require.Nil(t, Server)
	require.Nil(t, Lis)

	CreateNewGrpcServer()

	require.NotNil(t, Server)
	require.NotNil(t, Lis)

	defer func() {
		Server.Stop()
		Lis.Close()
	}()

	os.Unsetenv("NETWORK")
	os.Unsetenv("PORT")
}
