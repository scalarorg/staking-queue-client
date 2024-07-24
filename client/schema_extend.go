package client

type ActiveVaultEvent struct {
	EventType             EventType `json:"event_type"` // always 1. ActiveStakingEventType
	VaultTxHashHex        string    `json:"vault_tx_hash_hex"`
	StakerPkHex           string    `json:"staker_pk_hex"`
	FinalityProviderPkHex string    `json:"finality_provider_pk_hex"`
	StakingValue          uint64    `json:"staking_value"`
	StakingStartHeight    uint64    `json:"staking_start_height"`
	StakingStartTimestamp int64     `json:"staking_start_timestamp"`
	// StakingTimeLock             uint64    `json:"staking_timelock"`
	StakingOutputIndex          uint64 `json:"staking_output_index"`
	ChainID                     []byte `json:"chain_id"`
	ChainIdUserAddress          []byte `json:"chain_id_user_address"`
	ChainIdSmartContractAddress []byte `json:"chain_id_smart_contract_address"`
	MintingAmount               []byte `json:"amount_vault"`
	StakingTxHex                string `json:"staking_tx_hex"`
	IsOverflow                  bool   `json:"is_overflow"`
}

func NewActiveVaultEvent(
	vaultTxHashHex string,
	stakerPkHex string,
	finalityProviderPkHex string,
	stakingValue uint64,
	stakingStartHeight uint64,
	stakingStartTimestamp int64,
	// stakingTimeLock uint64,
	stakingOutputIndex uint64,
	chainID []byte,
	chainIdUserAddress []byte,
	chainIdSmartContractAddress []byte,
	amountVault []byte,
	stakingTxHex string,
	isOverflow bool,
) ActiveVaultEvent {
	return ActiveVaultEvent{
		EventType:             ActiveStakingEventType,
		VaultTxHashHex:        vaultTxHashHex,
		StakerPkHex:           stakerPkHex,
		FinalityProviderPkHex: finalityProviderPkHex,
		StakingValue:          stakingValue,
		StakingStartHeight:    stakingStartHeight,
		StakingStartTimestamp: stakingStartTimestamp,
		// StakingTimeLock:             stakingTimeLock,
		StakingOutputIndex:          stakingOutputIndex,
		ChainID:                     chainID,
		ChainIdUserAddress:          chainIdUserAddress,
		ChainIdSmartContractAddress: chainIdSmartContractAddress,
		MintingAmount:               amountVault,
		StakingTxHex:                stakingTxHex,
		IsOverflow:                  isOverflow,
	}
}

type BurningVaultEvent struct {
	EventType             EventType `json:"event_type"` // always 2. UnbondingStakingEventType
	VaultTxHashHex        string    `json:"vault_tx_hash_hex"`
	BurningStartHeight    uint64    `json:"burning_start_height"`
	BurningStartTimestamp int64     `json:"vault_start_timestamp"`
	// BurningTimeLock       uint64    `json:"unbonding_timelock"`
	BurningOutputIndex uint64 `json:"burning_output_index"`
	BurningTxHex       string `json:"burning_tx_hex"`
	BurningTxHashHex   string `json:"burning_tx_hash_hex"`
}

func NewBurningVaultEvent(
	vaultTxHashHex string,
	burningStartHeight uint64,
	burningStartTimestamp int64,
	// burningTimeLock uint64,
	burningOutputIndex uint64,
	burningTxHex string,
	burningTxHashHex string,
) BurningVaultEvent {
	return BurningVaultEvent{
		EventType:             UnbondingStakingEventType,
		VaultTxHashHex:        vaultTxHashHex,
		BurningStartHeight:    burningStartHeight,
		BurningStartTimestamp: burningStartTimestamp,
		// BurningTimeLock:       burningTimeLock,
		BurningOutputIndex: burningOutputIndex,
		BurningTxHex:       burningTxHex,
		BurningTxHashHex:   burningTxHashHex,
	}
}

type WithdrawVaultEvent struct {
	EventType      EventType `json:"event_type"` // always 3. WithdrawStakingEventType
	VaultTxHashHex string    `json:"vault_tx_hash_hex"`
}

func (e WithdrawVaultEvent) GetEventType() EventType {
	return VaultEventType
}

func (e WithdrawVaultEvent) GetVaultTxHashHex() string {
	return e.VaultTxHashHex
}

func NewWithdrawVaultEvent(vaultTxHashHex string) WithdrawVaultEvent {
	return WithdrawVaultEvent{
		EventType:      VaultEventType,
		VaultTxHashHex: vaultTxHashHex,
	}
}
