package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	// "github.com/babylonchain/staking-queue-client/client"
	"github.com/babylonchain/staking-queue-client/config"
	"github.com/scalarorg/staking-queue-client/client"
)

const (
	mockStakerHash_v2 = "0x1234567890abcdef"
)

func TestVaultEvent(t *testing.T) {
	numVaultEvents := 3
	activeVaultEvents := buildActiveNVaultEvents(mockStakerHash_v2, numVaultEvents)
	queueCfg := config.DefaultQueueConfig()

	testServer := setupTestQueueConsumer(t, queueCfg)
	defer testServer.Stop(t)

	queueManager := testServer.QueueManager

	vaultEventReceivedChan, err := queueManager.VaultQueue.ReceiveMessages()
	require.NoError(t, err)

	for _, ev := range activeVaultEvents {
		err = queueManager.PushVaultEvent(ev)
		require.NoError(t, err)

		receivedEv := <-vaultEventReceivedChan
		var vaultEv client.ActiveVaultEvent
		err := json.Unmarshal([]byte(receivedEv.Body), &vaultEv)
		require.NoError(t, err)
		require.Equal(t, ev, &vaultEv)
	}
}

func TestBurningEvent(t *testing.T) {
	numBurningEvents := 3
	burningEvents := buildNBurningEvents(numBurningEvents)
	queueCfg := config.DefaultQueueConfig()

	testServer := setupTestQueueConsumer(t, queueCfg)
	defer testServer.Stop(t)

	queueManager := testServer.QueueManager

	burningEvReceivedChan, err := queueManager.BurningQueue.ReceiveMessages()
	require.NoError(t, err)

	for _, ev := range burningEvents {
		err = queueManager.PushBurningEvent(ev)
		require.NoError(t, err)

		receivedEv := <-burningEvReceivedChan
		var burningEv client.BurningVaultEvent
		err := json.Unmarshal([]byte(receivedEv.Body), &burningEv)
		require.NoError(t, err)
		require.Equal(t, ev, &burningEv)
	}
}

func TestWithdrawVaultEvent(t *testing.T) {
	numWithdrawVaultEvents := 3
	withdrawEvents := buildNWithdrawVaultEvents(numWithdrawVaultEvents)
	queueCfg := config.DefaultQueueConfig()

	testServer := setupTestQueueConsumer(t, queueCfg)
	defer testServer.Stop(t)

	queueManager := testServer.QueueManager

	withdrawVaultEventsReceivedChan, err := queueManager.WithdrawVaultQueue.ReceiveMessages()
	require.NoError(t, err)

	for _, ev := range withdrawEvents {
		err = queueManager.PushWithdrawVaultEvent(ev)
		require.NoError(t, err)

		receivedEv := <-withdrawVaultEventsReceivedChan
		var withdrawVaultEv client.WithdrawVaultEvent
		err := json.Unmarshal([]byte(receivedEv.Body), &withdrawVaultEv)
		require.NoError(t, err)
		require.Equal(t, ev, &withdrawVaultEv)
	}
}
