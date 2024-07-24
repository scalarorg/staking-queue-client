package tests

import (
	"fmt"
	"time"

	// "github.com/babylonchain/staking-queue-client/client"
	"github.com/scalarorg/staking-queue-client/client"
)

func buildActiveNVaultEvents(stakerHash string, numOfEvent int) []*client.ActiveVaultEvent {
	var activeVaultEvents []*client.ActiveVaultEvent
	for i := 0; i < numOfEvent; i++ {
		activeVaultEvent := client.NewActiveVaultEvent(
			"0x1234567890abcdef"+fmt.Sprint(i),
			"0xabcdef1234567890"+fmt.Sprint(i),
			stakerHash,
			"0xabcdef1234567890"+fmt.Sprint(i),
			1+uint64(i),
			100+uint64(i),
			time.Now().Unix(),
			0,
			// 8 bytes chainID
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08},
			// 20 bytes chainIdUserAddress
			[]byte{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x10, 0x11, 0x12, 0x13, 0x14, 0x15, 0x16, 0x17, 0x18, 0x19, 0x20},
			// 20 bytes chainIdSmartContractAddress
			[]byte{0x21, 0x22, 0x23, 0x24, 0x25, 0x26, 0x27, 0x28, 0x29, 0x30, 0x31, 0x32, 0x33, 0x34, 0x35, 0x36, 0x37, 0x38, 0x39, 0x40},
			// 32 bytes amountVault
			[]byte{0x41, 0x42, 0x43, 0x44, 0x45, 0x46, 0x47, 0x48, 0x49, 0x50, 0x51, 0x52, 0x53, 0x54, 0x55, 0x56, 0x57, 0x58, 0x59, 0x60, 0x61, 0x62, 0x63, 0x64, 0x65, 0x66, 0x67, 0x68, 0x69, 0x70, 0x71, 0x72},
			false,
		)

		activeVaultEvents = append(activeVaultEvents, &activeVaultEvent)
	}
	return activeVaultEvents
}

func buildNBurningEvents(numOfEvent int) []*client.BurningVaultEvent {
	var burningEvents []*client.BurningVaultEvent
	for i := 0; i < numOfEvent; i++ {
		burningEv := client.NewBurningVaultEvent(
			"0x1234567890abcdef"+fmt.Sprint(i),
			uint64(i),
			time.Now().Unix(),
			uint64(0),
			"0xabcdef1234567890"+fmt.Sprint(i),
			"0x1234567890abcdef"+fmt.Sprint(i),
		)
		burningEvents = append(burningEvents, &burningEv)
	}

	return burningEvents
}

func buildNWithdrawVaultEvents(numOfEvent int) []*client.WithdrawVaultEvent {
	var withdrawVaultEvents []*client.WithdrawVaultEvent
	for i := 0; i < numOfEvent; i++ {
		withdrawVaultEv := client.NewWithdrawVaultEvent(
			"0x1234567890abcdef" + fmt.Sprint(i),
		)
		withdrawVaultEvents = append(withdrawVaultEvents, &withdrawVaultEv)
	}

	return withdrawVaultEvents
}
