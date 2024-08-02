package tests

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"

	// "github.com/babylonchain/staking-queue-client/client"
	// "github.com/babylonchain/staking-queue-client/config"
	"github.com/scalarorg/staking-queue-client/client"
	"github.com/scalarorg/staking-queue-client/config"
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

func TestSlashingOrLostKeyEvent(t *testing.T) {
	numSlashingOrLostKeyEvents := 3
	slashingOrLostKeyEvents := buildNSlashingOrLostKeyEvents(numSlashingOrLostKeyEvents)
	queueCfg := config.DefaultQueueConfig()

	testServer := setupTestQueueConsumer(t, queueCfg)
	defer testServer.Stop(t)

	queueManager := testServer.QueueManager

	slashingOrLostKeyEventReceivedChan, err := queueManager.SlashingOrLostKeyQueue.ReceiveMessages()
	require.NoError(t, err)

	for _, ev := range slashingOrLostKeyEvents {
		err = queueManager.PushSlashingOrLostKeyEvent(ev)
		require.NoError(t, err)

		receivedEv := <-slashingOrLostKeyEventReceivedChan
		var slashingOrLostKeyEvent client.SlashingOrLostKeyVaultEvent
		err := json.Unmarshal([]byte(receivedEv.Body), &slashingOrLostKeyEvent)
		require.NoError(t, err)
		require.Equal(t, ev, &slashingOrLostKeyEvent)
	}
}

func TestBurnWithoutDAppEvent(t *testing.T) {
	numBurnWithoutDAppEvents := 3
	burnWithoutDAppEvents := buildNBurnWithoutDAppEvents(numBurnWithoutDAppEvents)
	queueCfg := config.DefaultQueueConfig()

	testServer := setupTestQueueConsumer(t, queueCfg)
	defer testServer.Stop(t)

	queueManager := testServer.QueueManager

	burnWithoutDAppEventReceivedChan, err := queueManager.BurnWithoutDAppQueue.ReceiveMessages()
	require.NoError(t, err)

	for _, ev := range burnWithoutDAppEvents {
		err = queueManager.PushBurnWithoutDAppEvent(ev)
		require.NoError(t, err)

		receivedEv := <-burnWithoutDAppEventReceivedChan
		var burnWithoutDAppEvent client.BurnWithoutDAppVaultEvent
		err := json.Unmarshal([]byte(receivedEv.Body), &burnWithoutDAppEvent)
		require.NoError(t, err)
		require.Equal(t, ev, &burnWithoutDAppEvent)
	}
}
