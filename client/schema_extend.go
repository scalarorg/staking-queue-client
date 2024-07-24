package client

const (
	// Scalar
	ActiveVaultQueueName   string = "active_vault_queue"
	BurningVaultQueueName  string = "scalar_burning_queue"
	WithdrawVaultQueueName string = "scalar_withdraw_vault_queue"
)

const (
	ActiveVaultEventType   EventType = 7
	BurningVaultEventType  EventType = 8
	WithdrawVaultEventType EventType = 9
)

type ActiveVaultEvent struct {
	EventType             EventType `json:"event_type"`
	VaultTxHashHex        string    `json:"vault_tx_hash_hex"`
	VaultTxHex            string    `json:"vault_tx_hex"`
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
	MintingAmount               []byte `json:"amount_minting"`
	IsOverflow                  bool   `json:"is_overflow"`
}

func (e ActiveVaultEvent) GetEventType() EventType {
	return ActiveVaultEventType
}

func (e ActiveVaultEvent) GetVaultTxHashHex() string {
	return e.VaultTxHashHex
}

func NewActiveVaultEvent(
	vaultTxHashHex string,
	vaultTxHex string,
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
	amountMinting []byte,
	isOverflow bool,
) ActiveVaultEvent {
	return ActiveVaultEvent{
		EventType:             ActiveVaultEventType,
		VaultTxHashHex:        vaultTxHashHex,
		VaultTxHex:            vaultTxHex,
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
		MintingAmount:               amountMinting,
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

func (e BurningVaultEvent) GetEventType() EventType {
	return BurningVaultEventType
}

func (e BurningVaultEvent) GetVaultTxHashHex() string {
	return e.VaultTxHashHex
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
		EventType:             BurningVaultEventType,
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
	return ActiveVaultEventType
}

func (e WithdrawVaultEvent) GetVaultTxHashHex() string {
	return e.VaultTxHashHex
}

func NewWithdrawVaultEvent(vaultTxHashHex string) WithdrawVaultEvent {
	return WithdrawVaultEvent{
		EventType:      WithdrawVaultEventType,
		VaultTxHashHex: vaultTxHashHex,
	}
}
