// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ICrossChainSyncSnippet is an auto generated low-level Go binding around an user-defined struct.
type ICrossChainSyncSnippet struct {
	RemoteBlockId uint64
	SyncedInBlock uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
}

// ITierProviderTier is an auto generated low-level Go binding around an user-defined struct.
type ITierProviderTier struct {
	VerifierName      [32]byte
	ValidityBond      *big.Int
	ContestBond       *big.Int
	CooldownWindow    *big.Int
	ProvingWindow     uint16
	MaxBlocksToVerify uint8
}

// TaikoDataBlock is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlock struct {
	MetaHash             [32]byte
	AssignedProver       common.Address
	LivenessBond         *big.Int
	BlockId              uint64
	ProposedAt           uint64
	ProposedIn           uint64
	NextTransitionId     uint32
	VerifiedTransitionId uint32
	Reserved             [7][32]byte
}

// TaikoDataBlockMetadata is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataBlockMetadata struct {
	L1Hash           [32]byte
	Difficulty       [32]byte
	BlobHash         [32]byte
	ExtraData        [32]byte
	DepositsHash     [32]byte
	Coinbase         common.Address
	Id               uint64
	GasLimit         uint32
	Timestamp        uint64
	L1Height         uint64
	TxListByteOffset *big.Int
	TxListByteSize   *big.Int
	MinTier          uint16
	BlobUsed         bool
	ParentMetaHash   [32]byte
}

// TaikoDataConfig is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataConfig struct {
	ChainId                      uint64
	BlockMaxProposals            uint64
	BlockRingBufferSize          uint64
	MaxBlocksToVerifyPerProposal uint64
	BlockMaxGasLimit             uint32
	BlockMaxTxListBytes          *big.Int
	BlobExpiry                   *big.Int
	BlobAllowedForDA             bool
	LivenessBond                 *big.Int
	EthDepositRingBufferSize     *big.Int
	EthDepositMinCountPerBlock   uint64
	EthDepositMaxCountPerBlock   uint64
	EthDepositMinAmount          *big.Int
	EthDepositMaxAmount          *big.Int
	EthDepositGas                *big.Int
	EthDepositMaxFee             *big.Int
}

// TaikoDataEthDeposit is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataEthDeposit struct {
	Recipient common.Address
	Amount    *big.Int
	Id        uint64
}

// TaikoDataSlotA is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataSlotA struct {
	GenesisHeight           uint64
	GenesisTimestamp        uint64
	NumEthDeposits          uint64
	NextEthDepositToProcess uint64
}

// TaikoDataSlotB is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataSlotB struct {
	NumBlocks           uint64
	LastVerifiedBlockId uint64
	ProvingPaused       bool
}

// TaikoDataTransition is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataTransition struct {
	ParentHash [32]byte
	BlockHash  [32]byte
	SignalRoot [32]byte
	Graffiti   [32]byte
}

// TaikoDataTransitionState is an auto generated low-level Go binding around an user-defined struct.
type TaikoDataTransitionState struct {
	Key           [32]byte
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Prover        common.Address
	ValidityBond  *big.Int
	Contester     common.Address
	ContestBond   *big.Int
	Timestamp     uint64
	Tier          uint16
	Contestations uint8
	Reserved      [4][32]byte
}

// TaikoL1ClientMetaData contains all meta data concerning the TaikoL1Client contract.
var TaikoL1ClientMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"addressManager\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canDepositEthToL2\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"depositEtherToL2\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"getBlock\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"blk\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.Block\",\"components\":[{\"name\":\"metaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"livenessBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposedAt\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"proposedIn\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nextTransitionId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"verifiedTransitionId\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"__reserved\",\"type\":\"bytes32[7]\",\"internalType\":\"bytes32[7]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getConfig\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.Config\",\"components\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockMaxProposals\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockRingBufferSize\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"maxBlocksToVerifyPerProposal\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockMaxGasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"blockMaxTxListBytes\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"blobExpiry\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"blobAllowedForDA\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"livenessBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositRingBufferSize\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethDepositMinCountPerBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ethDepositMaxCountPerBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"ethDepositMinAmount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositMaxAmount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"ethDepositGas\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"ethDepositMaxFee\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getMinTier\",\"inputs\":[{\"name\":\"rand\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getStateVariables\",\"inputs\":[],\"outputs\":[{\"name\":\"a\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.SlotA\",\"components\":[{\"name\":\"genesisHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"genesisTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numEthDeposits\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"b\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.SlotB\",\"components\":[{\"name\":\"numBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"provingPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSyncedSnippet\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structICrossChainSync.Snippet\",\"components\":[{\"name\":\"remoteBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTier\",\"inputs\":[{\"name\":\"tierId\",\"type\":\"uint16\",\"internalType\":\"uint16\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structITierProvider.Tier\",\"components\":[{\"name\":\"verifierName\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"validityBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"contestBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"cooldownWindow\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"provingWindow\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"maxBlocksToVerify\",\"type\":\"uint8\",\"internalType\":\"uint8\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTierIds\",\"inputs\":[],\"outputs\":[{\"name\":\"ids\",\"type\":\"uint16[]\",\"internalType\":\"uint16[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getTransition\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.TransitionState\",\"components\":[{\"name\":\"key\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"prover\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"validityBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"contester\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"contestBond\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"tier\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"contestations\",\"type\":\"uint8\",\"internalType\":\"uint8\"},{\"name\":\"__reserved\",\"type\":\"bytes32[4]\",\"internalType\":\"bytes32[4]\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"init\",\"inputs\":[{\"name\":\"_addressManager\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_genesisBlockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"isBlobReusable\",\"inputs\":[{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"isConfigValid\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"pause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"pauseProving\",\"inputs\":[{\"name\":\"pause\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"paused\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeBlock\",\"inputs\":[{\"name\":\"params\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"txList\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"meta\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.BlockMetadata\",\"components\":[{\"name\":\"l1Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"difficulty\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"coinbase\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l1Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txListByteOffset\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"txListByteSize\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"minTier\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"blobUsed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"parentMetaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"depositsProcessed\",\"type\":\"tuple[]\",\"internalType\":\"structTaikoData.EthDeposit[]\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"proveBlock\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"input\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"proxiableUUID\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"resolve\",\"inputs\":[{\"name\":\"name\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"allowZeroAddress\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[{\"name\":\"addr\",\"type\":\"address\",\"internalType\":\"addresspayable\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"state\",\"inputs\":[],\"outputs\":[{\"name\":\"slotA\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.SlotA\",\"components\":[{\"name\":\"genesisHeight\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"genesisTimestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"numEthDeposits\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"nextEthDepositToProcess\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]},{\"name\":\"slotB\",\"type\":\"tuple\",\"internalType\":\"structTaikoData.SlotB\",\"components\":[{\"name\":\"numBlocks\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"lastVerifiedBlockId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"provingPaused\",\"type\":\"bool\",\"internalType\":\"bool\"}]}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"unpause\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeTo\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"upgradeToAndCall\",\"inputs\":[{\"name\":\"newImplementation\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"data\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"verifyBlocks\",\"inputs\":[{\"name\":\"maxBlocksToVerify\",\"type\":\"uint64\",\"internalType\":\"uint64\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"AdminChanged\",\"inputs\":[{\"name\":\"previousAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"newAdmin\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BeaconUpgraded\",\"inputs\":[{\"name\":\"beacon\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlobCached\",\"inputs\":[{\"name\":\"blobHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlockProposed\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"livenessBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"meta\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.BlockMetadata\",\"components\":[{\"name\":\"l1Hash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"difficulty\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blobHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"extraData\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"depositsHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"coinbase\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"gasLimit\",\"type\":\"uint32\",\"internalType\":\"uint32\"},{\"name\":\"timestamp\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"l1Height\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"txListByteOffset\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"txListByteSize\",\"type\":\"uint24\",\"internalType\":\"uint24\"},{\"name\":\"minTier\",\"type\":\"uint16\",\"internalType\":\"uint16\"},{\"name\":\"blobUsed\",\"type\":\"bool\",\"internalType\":\"bool\"},{\"name\":\"parentMetaHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"depositsProcessed\",\"type\":\"tuple[]\",\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit[]\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlockVerified\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"contestations\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BlockVerified\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"assignedProver\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"},{\"name\":\"contestations\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainSynced\",\"inputs\":[{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CrossChainSynced\",\"inputs\":[{\"name\":\"syncedInBlock\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockId\",\"type\":\"uint64\",\"indexed\":true,\"internalType\":\"uint64\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"EthDeposited\",\"inputs\":[{\"name\":\"deposit\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.EthDeposit\",\"components\":[{\"name\":\"recipient\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint96\",\"internalType\":\"uint96\"},{\"name\":\"id\",\"type\":\"uint64\",\"internalType\":\"uint64\"}]}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Initialized\",\"inputs\":[{\"name\":\"version\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"uint8\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Paused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ProvingPaused\",\"inputs\":[{\"name\":\"paused\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenCredited\",\"inputs\":[{\"name\":\"to\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDebited\",\"inputs\":[{\"name\":\"from\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenDeposited\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TokenWithdrawn\",\"inputs\":[{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransitionContested\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tran\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.Transition\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"graffiti\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"contester\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"contestBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"TransitionProved\",\"inputs\":[{\"name\":\"blockId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"tran\",\"type\":\"tuple\",\"indexed\":false,\"internalType\":\"structTaikoData.Transition\",\"components\":[{\"name\":\"parentHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"blockHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"signalRoot\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"graffiti\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}]},{\"name\":\"prover\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"validityBond\",\"type\":\"uint96\",\"indexed\":false,\"internalType\":\"uint96\"},{\"name\":\"tier\",\"type\":\"uint16\",\"indexed\":false,\"internalType\":\"uint16\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Unpaused\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Upgraded\",\"inputs\":[{\"name\":\"implementation\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_ALREADY_CONTESTED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_ALREADY_PROVED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_ASSIGNED_PROVER_NOT_ALLOWED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOB_FOR_DA_DISABLED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOB_NOT_FOUND\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOB_NOT_REUSEABLE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOCK_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_BLOCK_MISMATCH\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INSUFFICIENT_TOKEN\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_ADDRESS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_AMOUNT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_BLOCK_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_CONFIG\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_ETH_DEPOSIT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_HOOK\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PARAM\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PAUSE_STATUS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PROOF\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_PROVER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_TIER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_INVALID_TRANSITION\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_LIVENESS_BOND_NOT_RECEIVED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_NOT_ASSIGNED_PROVER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_PROPOSER_NOT_EOA\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_PROVING_PAUSED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_RECEIVE_DISABLED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TOO_MANY_BLOCKS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TOO_MANY_TIERS\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TRANSITION_ID_ZERO\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TRANSITION_ID_ZERO\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TRANSITION_NOT_FOUND\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TXLIST_OFFSET_SIZE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_TXLIST_TOO_LARGE\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNAUTHORIZED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNEXPECTED_PARENT\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNEXPECTED_TRANSITION_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNEXPECTED_TRANSITION_ID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"L1_UNEXPECTED_TRANSITION_TIER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"REENTRANT_CALL\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_DENIED\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_INVALID_MANAGER\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_UNEXPECTED_CHAINID\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"RESOLVER_ZERO_ADDR\",\"inputs\":[{\"name\":\"chainId\",\"type\":\"uint64\",\"internalType\":\"uint64\"},{\"name\":\"name\",\"type\":\"string\",\"internalType\":\"string\"}]}]",
}

// TaikoL1ClientABI is the input ABI used to generate the binding from.
// Deprecated: Use TaikoL1ClientMetaData.ABI instead.
var TaikoL1ClientABI = TaikoL1ClientMetaData.ABI

// TaikoL1Client is an auto generated Go binding around an Ethereum contract.
type TaikoL1Client struct {
	TaikoL1ClientCaller     // Read-only binding to the contract
	TaikoL1ClientTransactor // Write-only binding to the contract
	TaikoL1ClientFilterer   // Log filterer for contract events
}

// TaikoL1ClientCaller is an auto generated read-only Go binding around an Ethereum contract.
type TaikoL1ClientCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TaikoL1ClientTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TaikoL1ClientFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TaikoL1ClientSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TaikoL1ClientSession struct {
	Contract     *TaikoL1Client    // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TaikoL1ClientCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TaikoL1ClientCallerSession struct {
	Contract *TaikoL1ClientCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts        // Call options to use throughout this session
}

// TaikoL1ClientTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TaikoL1ClientTransactorSession struct {
	Contract     *TaikoL1ClientTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts        // Transaction auth options to use throughout this session
}

// TaikoL1ClientRaw is an auto generated low-level Go binding around an Ethereum contract.
type TaikoL1ClientRaw struct {
	Contract *TaikoL1Client // Generic contract binding to access the raw methods on
}

// TaikoL1ClientCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TaikoL1ClientCallerRaw struct {
	Contract *TaikoL1ClientCaller // Generic read-only contract binding to access the raw methods on
}

// TaikoL1ClientTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TaikoL1ClientTransactorRaw struct {
	Contract *TaikoL1ClientTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTaikoL1Client creates a new instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1Client(address common.Address, backend bind.ContractBackend) (*TaikoL1Client, error) {
	contract, err := bindTaikoL1Client(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TaikoL1Client{TaikoL1ClientCaller: TaikoL1ClientCaller{contract: contract}, TaikoL1ClientTransactor: TaikoL1ClientTransactor{contract: contract}, TaikoL1ClientFilterer: TaikoL1ClientFilterer{contract: contract}}, nil
}

// NewTaikoL1ClientCaller creates a new read-only instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientCaller(address common.Address, caller bind.ContractCaller) (*TaikoL1ClientCaller, error) {
	contract, err := bindTaikoL1Client(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCaller{contract: contract}, nil
}

// NewTaikoL1ClientTransactor creates a new write-only instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientTransactor(address common.Address, transactor bind.ContractTransactor) (*TaikoL1ClientTransactor, error) {
	contract, err := bindTaikoL1Client(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransactor{contract: contract}, nil
}

// NewTaikoL1ClientFilterer creates a new log filterer instance of TaikoL1Client, bound to a specific deployed contract.
func NewTaikoL1ClientFilterer(address common.Address, filterer bind.ContractFilterer) (*TaikoL1ClientFilterer, error) {
	contract, err := bindTaikoL1Client(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientFilterer{contract: contract}, nil
}

// bindTaikoL1Client binds a generic wrapper to an already deployed contract.
func bindTaikoL1Client(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TaikoL1ClientMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1Client *TaikoL1ClientRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1Client.Contract.TaikoL1ClientCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1Client *TaikoL1ClientRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TaikoL1ClientTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1Client *TaikoL1ClientRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TaikoL1ClientTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TaikoL1Client *TaikoL1ClientCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TaikoL1Client.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TaikoL1Client *TaikoL1ClientTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TaikoL1Client *TaikoL1ClientTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.contract.Transact(opts, method, params...)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) AddressManager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "addressManager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) AddressManager() (common.Address, error) {
	return _TaikoL1Client.Contract.AddressManager(&_TaikoL1Client.CallOpts)
}

// AddressManager is a free data retrieval call binding the contract method 0x3ab76e9f.
//
// Solidity: function addressManager() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) AddressManager() (common.Address, error) {
	return _TaikoL1Client.Contract.AddressManager(&_TaikoL1Client.CallOpts)
}

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) CanDepositEthToL2(opts *bind.CallOpts, amount *big.Int) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "canDepositEthToL2", amount)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) CanDepositEthToL2(amount *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.CanDepositEthToL2(&_TaikoL1Client.CallOpts, amount)
}

// CanDepositEthToL2 is a free data retrieval call binding the contract method 0xcf151d9a.
//
// Solidity: function canDepositEthToL2(uint256 amount) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) CanDepositEthToL2(amount *big.Int) (bool, error) {
	return _TaikoL1Client.Contract.CanDepositEthToL2(&_TaikoL1Client.CallOpts, amount)
}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint64,uint32,uint32,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientCaller) GetBlock(opts *bind.CallOpts, blockId uint64) (TaikoDataBlock, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getBlock", blockId)

	if err != nil {
		return *new(TaikoDataBlock), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataBlock)).(*TaikoDataBlock)

	return out0, err

}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint64,uint32,uint32,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientSession) GetBlock(blockId uint64) (TaikoDataBlock, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetBlock is a free data retrieval call binding the contract method 0x5fa15e79.
//
// Solidity: function getBlock(uint64 blockId) view returns((bytes32,address,uint96,uint64,uint64,uint64,uint32,uint32,bytes32[7]) blk)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetBlock(blockId uint64) (TaikoDataBlock, error) {
	return _TaikoL1Client.Contract.GetBlock(&_TaikoL1Client.CallOpts, blockId)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
func (_TaikoL1Client *TaikoL1ClientCaller) GetConfig(opts *bind.CallOpts) (TaikoDataConfig, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getConfig")

	if err != nil {
		return *new(TaikoDataConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataConfig)).(*TaikoDataConfig)

	return out0, err

}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
func (_TaikoL1Client *TaikoL1ClientSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetConfig is a free data retrieval call binding the contract method 0xc3f909d4.
//
// Solidity: function getConfig() view returns((uint64,uint64,uint64,uint64,uint32,uint24,uint24,bool,uint96,uint256,uint64,uint64,uint96,uint96,uint256,uint256))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetConfig() (TaikoDataConfig, error) {
	return _TaikoL1Client.Contract.GetConfig(&_TaikoL1Client.CallOpts)
}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientCaller) GetMinTier(opts *bind.CallOpts, rand *big.Int) (uint16, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getMinTier", rand)

	if err != nil {
		return *new(uint16), err
	}

	out0 := *abi.ConvertType(out[0], new(uint16)).(*uint16)

	return out0, err

}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientSession) GetMinTier(rand *big.Int) (uint16, error) {
	return _TaikoL1Client.Contract.GetMinTier(&_TaikoL1Client.CallOpts, rand)
}

// GetMinTier is a free data retrieval call binding the contract method 0x59ab4e23.
//
// Solidity: function getMinTier(uint256 rand) view returns(uint16)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetMinTier(rand *big.Int) (uint16, error) {
	return _TaikoL1Client.Contract.GetMinTier(&_TaikoL1Client.CallOpts, rand)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64) a, (uint64,uint64,bool) b)
func (_TaikoL1Client *TaikoL1ClientCaller) GetStateVariables(opts *bind.CallOpts) (struct {
	A TaikoDataSlotA
	B TaikoDataSlotB
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getStateVariables")

	outstruct := new(struct {
		A TaikoDataSlotA
		B TaikoDataSlotB
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.A = *abi.ConvertType(out[0], new(TaikoDataSlotA)).(*TaikoDataSlotA)
	outstruct.B = *abi.ConvertType(out[1], new(TaikoDataSlotB)).(*TaikoDataSlotB)

	return *outstruct, err

}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64) a, (uint64,uint64,bool) b)
func (_TaikoL1Client *TaikoL1ClientSession) GetStateVariables() (struct {
	A TaikoDataSlotA
	B TaikoDataSlotB
}, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetStateVariables is a free data retrieval call binding the contract method 0xdde89cf5.
//
// Solidity: function getStateVariables() view returns((uint64,uint64,uint64,uint64) a, (uint64,uint64,bool) b)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetStateVariables() (struct {
	A TaikoDataSlotA
	B TaikoDataSlotB
}, error) {
	return _TaikoL1Client.Contract.GetStateVariables(&_TaikoL1Client.CallOpts)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL1Client *TaikoL1ClientCaller) GetSyncedSnippet(opts *bind.CallOpts, blockId uint64) (ICrossChainSyncSnippet, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getSyncedSnippet", blockId)

	if err != nil {
		return *new(ICrossChainSyncSnippet), err
	}

	out0 := *abi.ConvertType(out[0], new(ICrossChainSyncSnippet)).(*ICrossChainSyncSnippet)

	return out0, err

}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL1Client *TaikoL1ClientSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL1Client.Contract.GetSyncedSnippet(&_TaikoL1Client.CallOpts, blockId)
}

// GetSyncedSnippet is a free data retrieval call binding the contract method 0x8cfb0459.
//
// Solidity: function getSyncedSnippet(uint64 blockId) view returns((uint64,uint64,bytes32,bytes32))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetSyncedSnippet(blockId uint64) (ICrossChainSyncSnippet, error) {
	return _TaikoL1Client.Contract.GetSyncedSnippet(&_TaikoL1Client.CallOpts, blockId)
}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16,uint8))
func (_TaikoL1Client *TaikoL1ClientCaller) GetTier(opts *bind.CallOpts, tierId uint16) (ITierProviderTier, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTier", tierId)

	if err != nil {
		return *new(ITierProviderTier), err
	}

	out0 := *abi.ConvertType(out[0], new(ITierProviderTier)).(*ITierProviderTier)

	return out0, err

}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16,uint8))
func (_TaikoL1Client *TaikoL1ClientSession) GetTier(tierId uint16) (ITierProviderTier, error) {
	return _TaikoL1Client.Contract.GetTier(&_TaikoL1Client.CallOpts, tierId)
}

// GetTier is a free data retrieval call binding the contract method 0x576c3de7.
//
// Solidity: function getTier(uint16 tierId) view returns((bytes32,uint96,uint96,uint24,uint16,uint8))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTier(tierId uint16) (ITierProviderTier, error) {
	return _TaikoL1Client.Contract.GetTier(&_TaikoL1Client.CallOpts, tierId)
}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[] ids)
func (_TaikoL1Client *TaikoL1ClientCaller) GetTierIds(opts *bind.CallOpts) ([]uint16, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTierIds")

	if err != nil {
		return *new([]uint16), err
	}

	out0 := *abi.ConvertType(out[0], new([]uint16)).(*[]uint16)

	return out0, err

}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[] ids)
func (_TaikoL1Client *TaikoL1ClientSession) GetTierIds() ([]uint16, error) {
	return _TaikoL1Client.Contract.GetTierIds(&_TaikoL1Client.CallOpts)
}

// GetTierIds is a free data retrieval call binding the contract method 0xd8cde1c6.
//
// Solidity: function getTierIds() view returns(uint16[] ids)
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTierIds() ([]uint16, error) {
	return _TaikoL1Client.Contract.GetTierIds(&_TaikoL1Client.CallOpts)
}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,uint8,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientCaller) GetTransition(opts *bind.CallOpts, blockId uint64, parentHash [32]byte) (TaikoDataTransitionState, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "getTransition", blockId, parentHash)

	if err != nil {
		return *new(TaikoDataTransitionState), err
	}

	out0 := *abi.ConvertType(out[0], new(TaikoDataTransitionState)).(*TaikoDataTransitionState)

	return out0, err

}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,uint8,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientSession) GetTransition(blockId uint64, parentHash [32]byte) (TaikoDataTransitionState, error) {
	return _TaikoL1Client.Contract.GetTransition(&_TaikoL1Client.CallOpts, blockId, parentHash)
}

// GetTransition is a free data retrieval call binding the contract method 0xfd257e29.
//
// Solidity: function getTransition(uint64 blockId, bytes32 parentHash) view returns((bytes32,bytes32,bytes32,address,uint96,address,uint96,uint64,uint16,uint8,bytes32[4]))
func (_TaikoL1Client *TaikoL1ClientCallerSession) GetTransition(blockId uint64, parentHash [32]byte) (TaikoDataTransitionState, error) {
	return _TaikoL1Client.Contract.GetTransition(&_TaikoL1Client.CallOpts, blockId, parentHash)
}

// IsBlobReusable is a free data retrieval call binding the contract method 0xb67d7638.
//
// Solidity: function isBlobReusable(bytes32 blobHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsBlobReusable(opts *bind.CallOpts, blobHash [32]byte) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isBlobReusable", blobHash)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsBlobReusable is a free data retrieval call binding the contract method 0xb67d7638.
//
// Solidity: function isBlobReusable(bytes32 blobHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsBlobReusable(blobHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsBlobReusable(&_TaikoL1Client.CallOpts, blobHash)
}

// IsBlobReusable is a free data retrieval call binding the contract method 0xb67d7638.
//
// Solidity: function isBlobReusable(bytes32 blobHash) view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsBlobReusable(blobHash [32]byte) (bool, error) {
	return _TaikoL1Client.Contract.IsBlobReusable(&_TaikoL1Client.CallOpts, blobHash)
}

// IsConfigValid is a free data retrieval call binding the contract method 0xe3f1bdc5.
//
// Solidity: function isConfigValid() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) IsConfigValid(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "isConfigValid")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsConfigValid is a free data retrieval call binding the contract method 0xe3f1bdc5.
//
// Solidity: function isConfigValid() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) IsConfigValid() (bool, error) {
	return _TaikoL1Client.Contract.IsConfigValid(&_TaikoL1Client.CallOpts)
}

// IsConfigValid is a free data retrieval call binding the contract method 0xe3f1bdc5.
//
// Solidity: function isConfigValid() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) IsConfigValid() (bool, error) {
	return _TaikoL1Client.Contract.IsConfigValid(&_TaikoL1Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientSession) Owner() (common.Address, error) {
	return _TaikoL1Client.Contract.Owner(&_TaikoL1Client.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Owner() (common.Address, error) {
	return _TaikoL1Client.Contract.Owner(&_TaikoL1Client.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "paused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientSession) Paused() (bool, error) {
	return _TaikoL1Client.Contract.Paused(&_TaikoL1Client.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() view returns(bool)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Paused() (bool, error) {
	return _TaikoL1Client.Contract.Paused(&_TaikoL1Client.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCaller) ProxiableUUID(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "proxiableUUID")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientSession) ProxiableUUID() ([32]byte, error) {
	return _TaikoL1Client.Contract.ProxiableUUID(&_TaikoL1Client.CallOpts)
}

// ProxiableUUID is a free data retrieval call binding the contract method 0x52d1902d.
//
// Solidity: function proxiableUUID() view returns(bytes32)
func (_TaikoL1Client *TaikoL1ClientCallerSession) ProxiableUUID() ([32]byte, error) {
	return _TaikoL1Client.Contract.ProxiableUUID(&_TaikoL1Client.CallOpts)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve(opts *bind.CallOpts, chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve", chainId, name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve(chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve is a free data retrieval call binding the contract method 0x3eb6b8cf.
//
// Solidity: function resolve(uint64 chainId, bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve(chainId uint64, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve(&_TaikoL1Client.CallOpts, chainId, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCaller) Resolve0(opts *bind.CallOpts, name [32]byte, allowZeroAddress bool) (common.Address, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "resolve0", name, allowZeroAddress)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// Resolve0 is a free data retrieval call binding the contract method 0xa86f9d9e.
//
// Solidity: function resolve(bytes32 name, bool allowZeroAddress) view returns(address addr)
func (_TaikoL1Client *TaikoL1ClientCallerSession) Resolve0(name [32]byte, allowZeroAddress bool) (common.Address, error) {
	return _TaikoL1Client.Contract.Resolve0(&_TaikoL1Client.CallOpts, name, allowZeroAddress)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,bool) slotB)
func (_TaikoL1Client *TaikoL1ClientCaller) State(opts *bind.CallOpts) (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
}, error) {
	var out []interface{}
	err := _TaikoL1Client.contract.Call(opts, &out, "state")

	outstruct := new(struct {
		SlotA TaikoDataSlotA
		SlotB TaikoDataSlotB
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.SlotA = *abi.ConvertType(out[0], new(TaikoDataSlotA)).(*TaikoDataSlotA)
	outstruct.SlotB = *abi.ConvertType(out[1], new(TaikoDataSlotB)).(*TaikoDataSlotB)

	return *outstruct, err

}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,bool) slotB)
func (_TaikoL1Client *TaikoL1ClientSession) State() (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// State is a free data retrieval call binding the contract method 0xc19d93fb.
//
// Solidity: function state() view returns((uint64,uint64,uint64,uint64) slotA, (uint64,uint64,bool) slotB)
func (_TaikoL1Client *TaikoL1ClientCallerSession) State() (struct {
	SlotA TaikoDataSlotA
	SlotB TaikoDataSlotB
}, error) {
	return _TaikoL1Client.Contract.State(&_TaikoL1Client.CallOpts)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) DepositEtherToL2(opts *bind.TransactOpts, recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "depositEtherToL2", recipient)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) DepositEtherToL2(recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts, recipient)
}

// DepositEtherToL2 is a paid mutator transaction binding the contract method 0x047a289d.
//
// Solidity: function depositEtherToL2(address recipient) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) DepositEtherToL2(recipient common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.DepositEtherToL2(&_TaikoL1Client.TransactOpts, recipient)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Init(opts *bind.TransactOpts, _addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "init", _addressManager, _genesisBlockHash)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash)
}

// Init is a paid mutator transaction binding the contract method 0x2cc0b254.
//
// Solidity: function init(address _addressManager, bytes32 _genesisBlockHash) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Init(_addressManager common.Address, _genesisBlockHash [32]byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Init(&_TaikoL1Client.TransactOpts, _addressManager, _genesisBlockHash)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL1Client *TaikoL1ClientSession) Pause() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Pause(&_TaikoL1Client.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Pause() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Pause(&_TaikoL1Client.TransactOpts)
}

// PauseProving is a paid mutator transaction binding the contract method 0xff00c391.
//
// Solidity: function pauseProving(bool pause) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) PauseProving(opts *bind.TransactOpts, pause bool) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "pauseProving", pause)
}

// PauseProving is a paid mutator transaction binding the contract method 0xff00c391.
//
// Solidity: function pauseProving(bool pause) returns()
func (_TaikoL1Client *TaikoL1ClientSession) PauseProving(pause bool) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.PauseProving(&_TaikoL1Client.TransactOpts, pause)
}

// PauseProving is a paid mutator transaction binding the contract method 0xff00c391.
//
// Solidity: function pauseProving(bool pause) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) PauseProving(pause bool) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.PauseProving(&_TaikoL1Client.TransactOpts, pause)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes params, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientTransactor) ProposeBlock(opts *bind.TransactOpts, params []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proposeBlock", params, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes params, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientSession) ProposeBlock(params []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, params, txList)
}

// ProposeBlock is a paid mutator transaction binding the contract method 0xef16e845.
//
// Solidity: function proposeBlock(bytes params, bytes txList) payable returns((bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProposeBlock(params []byte, txList []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProposeBlock(&_TaikoL1Client.TransactOpts, params, txList)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) ProveBlock(opts *bind.TransactOpts, blockId uint64, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "proveBlock", blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientSession) ProveBlock(blockId uint64, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, input)
}

// ProveBlock is a paid mutator transaction binding the contract method 0x10d008bd.
//
// Solidity: function proveBlock(uint64 blockId, bytes input) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) ProveBlock(blockId uint64, input []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.ProveBlock(&_TaikoL1Client.TransactOpts, blockId, input)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.RenounceOwnership(&_TaikoL1Client.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.RenounceOwnership(&_TaikoL1Client.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TransferOwnership(&_TaikoL1Client.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.TransferOwnership(&_TaikoL1Client.TransactOpts, newOwner)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL1Client *TaikoL1ClientSession) Unpause() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Unpause(&_TaikoL1Client.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Unpause() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Unpause(&_TaikoL1Client.TransactOpts)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) UpgradeTo(opts *bind.TransactOpts, newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "upgradeTo", newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL1Client *TaikoL1ClientSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.UpgradeTo(&_TaikoL1Client.TransactOpts, newImplementation)
}

// UpgradeTo is a paid mutator transaction binding the contract method 0x3659cfe6.
//
// Solidity: function upgradeTo(address newImplementation) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) UpgradeTo(newImplementation common.Address) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.UpgradeTo(&_TaikoL1Client.TransactOpts, newImplementation)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) UpgradeToAndCall(opts *bind.TransactOpts, newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "upgradeToAndCall", newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.UpgradeToAndCall(&_TaikoL1Client.TransactOpts, newImplementation, data)
}

// UpgradeToAndCall is a paid mutator transaction binding the contract method 0x4f1ef286.
//
// Solidity: function upgradeToAndCall(address newImplementation, bytes data) payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) UpgradeToAndCall(newImplementation common.Address, data []byte) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.UpgradeToAndCall(&_TaikoL1Client.TransactOpts, newImplementation, data)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocksToVerify) returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) VerifyBlocks(opts *bind.TransactOpts, maxBlocksToVerify uint64) (*types.Transaction, error) {
	return _TaikoL1Client.contract.Transact(opts, "verifyBlocks", maxBlocksToVerify)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocksToVerify) returns()
func (_TaikoL1Client *TaikoL1ClientSession) VerifyBlocks(maxBlocksToVerify uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.VerifyBlocks(&_TaikoL1Client.TransactOpts, maxBlocksToVerify)
}

// VerifyBlocks is a paid mutator transaction binding the contract method 0x8778209d.
//
// Solidity: function verifyBlocks(uint64 maxBlocksToVerify) returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) VerifyBlocks(maxBlocksToVerify uint64) (*types.Transaction, error) {
	return _TaikoL1Client.Contract.VerifyBlocks(&_TaikoL1Client.TransactOpts, maxBlocksToVerify)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TaikoL1Client.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientSession) Receive() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Receive(&_TaikoL1Client.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_TaikoL1Client *TaikoL1ClientTransactorSession) Receive() (*types.Transaction, error) {
	return _TaikoL1Client.Contract.Receive(&_TaikoL1Client.TransactOpts)
}

// TaikoL1ClientAdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the TaikoL1Client contract.
type TaikoL1ClientAdminChangedIterator struct {
	Event *TaikoL1ClientAdminChanged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientAdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientAdminChanged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientAdminChanged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientAdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientAdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientAdminChanged represents a AdminChanged event raised by the TaikoL1Client contract.
type TaikoL1ClientAdminChanged struct {
	PreviousAdmin common.Address
	NewAdmin      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterAdminChanged(opts *bind.FilterOpts) (*TaikoL1ClientAdminChangedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientAdminChangedIterator{contract: _TaikoL1Client.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientAdminChanged) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientAdminChanged)
				if err := _TaikoL1Client.contract.UnpackLog(event, "AdminChanged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseAdminChanged is a log parse operation binding the contract event 0x7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f.
//
// Solidity: event AdminChanged(address previousAdmin, address newAdmin)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseAdminChanged(log types.Log) (*TaikoL1ClientAdminChanged, error) {
	event := new(TaikoL1ClientAdminChanged)
	if err := _TaikoL1Client.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBeaconUpgradedIterator is returned from FilterBeaconUpgraded and is used to iterate over the raw logs and unpacked data for BeaconUpgraded events raised by the TaikoL1Client contract.
type TaikoL1ClientBeaconUpgradedIterator struct {
	Event *TaikoL1ClientBeaconUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBeaconUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBeaconUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBeaconUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBeaconUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBeaconUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBeaconUpgraded represents a BeaconUpgraded event raised by the TaikoL1Client contract.
type TaikoL1ClientBeaconUpgraded struct {
	Beacon common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterBeaconUpgraded is a free log retrieval operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBeaconUpgraded(opts *bind.FilterOpts, beacon []common.Address) (*TaikoL1ClientBeaconUpgradedIterator, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBeaconUpgradedIterator{contract: _TaikoL1Client.contract, event: "BeaconUpgraded", logs: logs, sub: sub}, nil
}

// WatchBeaconUpgraded is a free log subscription operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBeaconUpgraded(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBeaconUpgraded, beacon []common.Address) (event.Subscription, error) {

	var beaconRule []interface{}
	for _, beaconItem := range beacon {
		beaconRule = append(beaconRule, beaconItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BeaconUpgraded", beaconRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBeaconUpgraded)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBeaconUpgraded is a log parse operation binding the contract event 0x1cf3b03a6cf19fa2baba4df148e9dcabedea7f8a5c07840e207e5c089be95d3e.
//
// Solidity: event BeaconUpgraded(address indexed beacon)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBeaconUpgraded(log types.Log) (*TaikoL1ClientBeaconUpgraded, error) {
	event := new(TaikoL1ClientBeaconUpgraded)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BeaconUpgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlobCachedIterator is returned from FilterBlobCached and is used to iterate over the raw logs and unpacked data for BlobCached events raised by the TaikoL1Client contract.
type TaikoL1ClientBlobCachedIterator struct {
	Event *TaikoL1ClientBlobCached // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlobCachedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlobCached)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlobCached)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlobCachedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlobCachedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlobCached represents a BlobCached event raised by the TaikoL1Client contract.
type TaikoL1ClientBlobCached struct {
	BlobHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterBlobCached is a free log retrieval operation binding the contract event 0xb303828b7c63a3e480df6d8239eb7be1d4a1eb1c8878d6fa461103b94f4ce852.
//
// Solidity: event BlobCached(bytes32 blobHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlobCached(opts *bind.FilterOpts) (*TaikoL1ClientBlobCachedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlobCached")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlobCachedIterator{contract: _TaikoL1Client.contract, event: "BlobCached", logs: logs, sub: sub}, nil
}

// WatchBlobCached is a free log subscription operation binding the contract event 0xb303828b7c63a3e480df6d8239eb7be1d4a1eb1c8878d6fa461103b94f4ce852.
//
// Solidity: event BlobCached(bytes32 blobHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlobCached(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlobCached) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlobCached")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlobCached)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlobCached", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlobCached is a log parse operation binding the contract event 0xb303828b7c63a3e480df6d8239eb7be1d4a1eb1c8878d6fa461103b94f4ce852.
//
// Solidity: event BlobCached(bytes32 blobHash)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlobCached(log types.Log) (*TaikoL1ClientBlobCached, error) {
	event := new(TaikoL1ClientBlobCached)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlobCached", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockProposedIterator is returned from FilterBlockProposed and is used to iterate over the raw logs and unpacked data for BlockProposed events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposedIterator struct {
	Event *TaikoL1ClientBlockProposed // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockProposed)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockProposed)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockProposed represents a BlockProposed event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockProposed struct {
	BlockId           *big.Int
	AssignedProver    common.Address
	LivenessBond      *big.Int
	Meta              TaikoDataBlockMetadata
	DepositsProcessed []TaikoDataEthDeposit
	Raw               types.Log // Blockchain specific contextual infos
}

// FilterBlockProposed is a free log retrieval operation binding the contract event 0xa62cea5af360b010ef0d23472a2a7493b54175fd9fd2f9c2aa2bb427d2f4d3ca.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockProposed(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address) (*TaikoL1ClientBlockProposedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockProposedIterator{contract: _TaikoL1Client.contract, event: "BlockProposed", logs: logs, sub: sub}, nil
}

// WatchBlockProposed is a free log subscription operation binding the contract event 0xa62cea5af360b010ef0d23472a2a7493b54175fd9fd2f9c2aa2bb427d2f4d3ca.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockProposed(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockProposed, blockId []*big.Int, assignedProver []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockProposed", blockIdRule, assignedProverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockProposed)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockProposed is a log parse operation binding the contract event 0xa62cea5af360b010ef0d23472a2a7493b54175fd9fd2f9c2aa2bb427d2f4d3ca.
//
// Solidity: event BlockProposed(uint256 indexed blockId, address indexed assignedProver, uint96 livenessBond, (bytes32,bytes32,bytes32,bytes32,bytes32,address,uint64,uint32,uint64,uint64,uint24,uint24,uint16,bool,bytes32) meta, (address,uint96,uint64)[] depositsProcessed)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockProposed(log types.Log) (*TaikoL1ClientBlockProposed, error) {
	event := new(TaikoL1ClientBlockProposed)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockVerifiedIterator is returned from FilterBlockVerified and is used to iterate over the raw logs and unpacked data for BlockVerified events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerifiedIterator struct {
	Event *TaikoL1ClientBlockVerified // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockVerifiedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockVerified)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockVerified)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockVerifiedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockVerifiedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockVerified represents a BlockVerified event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified struct {
	BlockId        *big.Int
	AssignedProver common.Address
	Prover         common.Address
	BlockHash      [32]byte
	SignalRoot     [32]byte
	Tier           uint16
	Contestations  uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified is a free log retrieval operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (*TaikoL1ClientBlockVerifiedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerifiedIterator{contract: _TaikoL1Client.contract, event: "BlockVerified", logs: logs, sub: sub}, nil
}

// WatchBlockVerified is a free log subscription operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockVerified)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerified is a log parse operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockVerified(log types.Log) (*TaikoL1ClientBlockVerified, error) {
	event := new(TaikoL1ClientBlockVerified)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientBlockVerified0Iterator is returned from FilterBlockVerified0 and is used to iterate over the raw logs and unpacked data for BlockVerified0 events raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified0Iterator struct {
	Event *TaikoL1ClientBlockVerified0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientBlockVerified0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientBlockVerified0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientBlockVerified0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientBlockVerified0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientBlockVerified0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientBlockVerified0 represents a BlockVerified0 event raised by the TaikoL1Client contract.
type TaikoL1ClientBlockVerified0 struct {
	BlockId        *big.Int
	AssignedProver common.Address
	Prover         common.Address
	BlockHash      [32]byte
	SignalRoot     [32]byte
	Tier           uint16
	Contestations  uint8
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBlockVerified0 is a free log retrieval operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterBlockVerified0(opts *bind.FilterOpts, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (*TaikoL1ClientBlockVerified0Iterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "BlockVerified0", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientBlockVerified0Iterator{contract: _TaikoL1Client.contract, event: "BlockVerified0", logs: logs, sub: sub}, nil
}

// WatchBlockVerified0 is a free log subscription operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchBlockVerified0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientBlockVerified0, blockId []*big.Int, assignedProver []common.Address, prover []common.Address) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}
	var assignedProverRule []interface{}
	for _, assignedProverItem := range assignedProver {
		assignedProverRule = append(assignedProverRule, assignedProverItem)
	}
	var proverRule []interface{}
	for _, proverItem := range prover {
		proverRule = append(proverRule, proverItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "BlockVerified0", blockIdRule, assignedProverRule, proverRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientBlockVerified0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseBlockVerified0 is a log parse operation binding the contract event 0xaeba6e73abba9419294b1017075cf8dc2e7de6f2d7fd3b336b3ba882a2acfca5.
//
// Solidity: event BlockVerified(uint256 indexed blockId, address indexed assignedProver, address indexed prover, bytes32 blockHash, bytes32 signalRoot, uint16 tier, uint8 contestations)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseBlockVerified0(log types.Log) (*TaikoL1ClientBlockVerified0, error) {
	event := new(TaikoL1ClientBlockVerified0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "BlockVerified0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientCrossChainSyncedIterator is returned from FilterCrossChainSynced and is used to iterate over the raw logs and unpacked data for CrossChainSynced events raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSyncedIterator struct {
	Event *TaikoL1ClientCrossChainSynced // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientCrossChainSyncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientCrossChainSynced)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientCrossChainSynced)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientCrossChainSyncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientCrossChainSyncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientCrossChainSynced represents a CrossChainSynced event raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced struct {
	SyncedInBlock uint64
	BlockId       uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced is a free log retrieval operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced(opts *bind.FilterOpts, syncedInBlock []uint64, blockId []uint64) (*TaikoL1ClientCrossChainSyncedIterator, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCrossChainSyncedIterator{contract: _TaikoL1Client.contract, event: "CrossChainSynced", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced is a free log subscription operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced, syncedInBlock []uint64, blockId []uint64) (event.Subscription, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "CrossChainSynced", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientCrossChainSynced)
				if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainSynced is a log parse operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseCrossChainSynced(log types.Log) (*TaikoL1ClientCrossChainSynced, error) {
	event := new(TaikoL1ClientCrossChainSynced)
	if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientCrossChainSynced0Iterator is returned from FilterCrossChainSynced0 and is used to iterate over the raw logs and unpacked data for CrossChainSynced0 events raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced0Iterator struct {
	Event *TaikoL1ClientCrossChainSynced0 // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientCrossChainSynced0)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientCrossChainSynced0)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientCrossChainSynced0Iterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientCrossChainSynced0 represents a CrossChainSynced0 event raised by the TaikoL1Client contract.
type TaikoL1ClientCrossChainSynced0 struct {
	SyncedInBlock uint64
	BlockId       uint64
	BlockHash     [32]byte
	SignalRoot    [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterCrossChainSynced0 is a free log retrieval operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterCrossChainSynced0(opts *bind.FilterOpts, syncedInBlock []uint64, blockId []uint64) (*TaikoL1ClientCrossChainSynced0Iterator, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "CrossChainSynced0", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientCrossChainSynced0Iterator{contract: _TaikoL1Client.contract, event: "CrossChainSynced0", logs: logs, sub: sub}, nil
}

// WatchCrossChainSynced0 is a free log subscription operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchCrossChainSynced0(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientCrossChainSynced0, syncedInBlock []uint64, blockId []uint64) (event.Subscription, error) {

	var syncedInBlockRule []interface{}
	for _, syncedInBlockItem := range syncedInBlock {
		syncedInBlockRule = append(syncedInBlockRule, syncedInBlockItem)
	}
	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "CrossChainSynced0", syncedInBlockRule, blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientCrossChainSynced0)
				if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced0", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseCrossChainSynced0 is a log parse operation binding the contract event 0xf35ec3b262cf74881db1b8051c635496bccb1497a1e776dacb463d0e0e2b0f51.
//
// Solidity: event CrossChainSynced(uint64 indexed syncedInBlock, uint64 indexed blockId, bytes32 blockHash, bytes32 signalRoot)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseCrossChainSynced0(log types.Log) (*TaikoL1ClientCrossChainSynced0, error) {
	event := new(TaikoL1ClientCrossChainSynced0)
	if err := _TaikoL1Client.contract.UnpackLog(event, "CrossChainSynced0", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientEthDepositedIterator is returned from FilterEthDeposited and is used to iterate over the raw logs and unpacked data for EthDeposited events raised by the TaikoL1Client contract.
type TaikoL1ClientEthDepositedIterator struct {
	Event *TaikoL1ClientEthDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientEthDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientEthDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientEthDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientEthDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientEthDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientEthDeposited represents a EthDeposited event raised by the TaikoL1Client contract.
type TaikoL1ClientEthDeposited struct {
	Deposit TaikoDataEthDeposit
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterEthDeposited is a free log retrieval operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterEthDeposited(opts *bind.FilterOpts) (*TaikoL1ClientEthDepositedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientEthDepositedIterator{contract: _TaikoL1Client.contract, event: "EthDeposited", logs: logs, sub: sub}, nil
}

// WatchEthDeposited is a free log subscription operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchEthDeposited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientEthDeposited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "EthDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientEthDeposited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseEthDeposited is a log parse operation binding the contract event 0x7120a3b075ad25974c5eed76dedb3a217c76c9c6d1f1e201caeba9b89de9a9d9.
//
// Solidity: event EthDeposited((address,uint96,uint64) deposit)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseEthDeposited(log types.Log) (*TaikoL1ClientEthDeposited, error) {
	event := new(TaikoL1ClientEthDeposited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "EthDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the TaikoL1Client contract.
type TaikoL1ClientInitializedIterator struct {
	Event *TaikoL1ClientInitialized // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientInitialized)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientInitialized)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientInitialized represents a Initialized event raised by the TaikoL1Client contract.
type TaikoL1ClientInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterInitialized(opts *bind.FilterOpts) (*TaikoL1ClientInitializedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientInitializedIterator{contract: _TaikoL1Client.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientInitialized) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientInitialized)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseInitialized(log types.Log) (*TaikoL1ClientInitialized, error) {
	event := new(TaikoL1ClientInitialized)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the TaikoL1Client contract.
type TaikoL1ClientOwnershipTransferredIterator struct {
	Event *TaikoL1ClientOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientOwnershipTransferred represents a OwnershipTransferred event raised by the TaikoL1Client contract.
type TaikoL1ClientOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*TaikoL1ClientOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientOwnershipTransferredIterator{contract: _TaikoL1Client.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientOwnershipTransferred)
				if err := _TaikoL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseOwnershipTransferred(log types.Log) (*TaikoL1ClientOwnershipTransferred, error) {
	event := new(TaikoL1ClientOwnershipTransferred)
	if err := _TaikoL1Client.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientPausedIterator is returned from FilterPaused and is used to iterate over the raw logs and unpacked data for Paused events raised by the TaikoL1Client contract.
type TaikoL1ClientPausedIterator struct {
	Event *TaikoL1ClientPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientPaused represents a Paused event raised by the TaikoL1Client contract.
type TaikoL1ClientPaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterPaused is a free log retrieval operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterPaused(opts *bind.FilterOpts) (*TaikoL1ClientPausedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientPausedIterator{contract: _TaikoL1Client.contract, event: "Paused", logs: logs, sub: sub}, nil
}

// WatchPaused is a free log subscription operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchPaused(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientPaused) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Paused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientPaused)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Paused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParsePaused is a log parse operation binding the contract event 0x62e78cea01bee320cd4e420270b5ea74000d11b0c9f74754ebdbfc544b05a258.
//
// Solidity: event Paused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParsePaused(log types.Log) (*TaikoL1ClientPaused, error) {
	event := new(TaikoL1ClientPaused)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Paused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientProvingPausedIterator is returned from FilterProvingPaused and is used to iterate over the raw logs and unpacked data for ProvingPaused events raised by the TaikoL1Client contract.
type TaikoL1ClientProvingPausedIterator struct {
	Event *TaikoL1ClientProvingPaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientProvingPausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientProvingPaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientProvingPaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientProvingPausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientProvingPausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientProvingPaused represents a ProvingPaused event raised by the TaikoL1Client contract.
type TaikoL1ClientProvingPaused struct {
	Paused bool
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterProvingPaused is a free log retrieval operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterProvingPaused(opts *bind.FilterOpts) (*TaikoL1ClientProvingPausedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "ProvingPaused")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientProvingPausedIterator{contract: _TaikoL1Client.contract, event: "ProvingPaused", logs: logs, sub: sub}, nil
}

// WatchProvingPaused is a free log subscription operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchProvingPaused(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientProvingPaused) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "ProvingPaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientProvingPaused)
				if err := _TaikoL1Client.contract.UnpackLog(event, "ProvingPaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseProvingPaused is a log parse operation binding the contract event 0xed64db85835d07c3c990b8ebdd55e32d64e5ed53143b6ef2179e7bfaf17ddc3b.
//
// Solidity: event ProvingPaused(bool paused)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseProvingPaused(log types.Log) (*TaikoL1ClientProvingPaused, error) {
	event := new(TaikoL1ClientProvingPaused)
	if err := _TaikoL1Client.contract.UnpackLog(event, "ProvingPaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenCreditedIterator is returned from FilterTokenCredited and is used to iterate over the raw logs and unpacked data for TokenCredited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCreditedIterator struct {
	Event *TaikoL1ClientTokenCredited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTokenCreditedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenCredited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTokenCredited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTokenCreditedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenCreditedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenCredited represents a TokenCredited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenCredited struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenCredited is a free log retrieval operation binding the contract event 0xcc91b0b567e69e32d26830c50d1078f0baec319d458fa847f2633c1c2f71dd74.
//
// Solidity: event TokenCredited(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenCredited(opts *bind.FilterOpts) (*TaikoL1ClientTokenCreditedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenCredited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenCreditedIterator{contract: _TaikoL1Client.contract, event: "TokenCredited", logs: logs, sub: sub}, nil
}

// WatchTokenCredited is a free log subscription operation binding the contract event 0xcc91b0b567e69e32d26830c50d1078f0baec319d458fa847f2633c1c2f71dd74.
//
// Solidity: event TokenCredited(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenCredited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenCredited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenCredited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenCredited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenCredited is a log parse operation binding the contract event 0xcc91b0b567e69e32d26830c50d1078f0baec319d458fa847f2633c1c2f71dd74.
//
// Solidity: event TokenCredited(address to, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenCredited(log types.Log) (*TaikoL1ClientTokenCredited, error) {
	event := new(TaikoL1ClientTokenCredited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenCredited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDebitedIterator is returned from FilterTokenDebited and is used to iterate over the raw logs and unpacked data for TokenDebited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebitedIterator struct {
	Event *TaikoL1ClientTokenDebited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTokenDebitedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDebited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTokenDebited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTokenDebitedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDebitedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDebited represents a TokenDebited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDebited struct {
	From   common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenDebited is a free log retrieval operation binding the contract event 0x9414c932608e522aea77e1b5fd1cdb441e5d57daf49d9a9a0b7409ef8d9e0070.
//
// Solidity: event TokenDebited(address from, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDebited(opts *bind.FilterOpts) (*TaikoL1ClientTokenDebitedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDebited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDebitedIterator{contract: _TaikoL1Client.contract, event: "TokenDebited", logs: logs, sub: sub}, nil
}

// WatchTokenDebited is a free log subscription operation binding the contract event 0x9414c932608e522aea77e1b5fd1cdb441e5d57daf49d9a9a0b7409ef8d9e0070.
//
// Solidity: event TokenDebited(address from, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDebited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDebited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDebited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDebited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenDebited is a log parse operation binding the contract event 0x9414c932608e522aea77e1b5fd1cdb441e5d57daf49d9a9a0b7409ef8d9e0070.
//
// Solidity: event TokenDebited(address from, uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDebited(log types.Log) (*TaikoL1ClientTokenDebited, error) {
	event := new(TaikoL1ClientTokenDebited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDebited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenDepositedIterator is returned from FilterTokenDeposited and is used to iterate over the raw logs and unpacked data for TokenDeposited events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDepositedIterator struct {
	Event *TaikoL1ClientTokenDeposited // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTokenDepositedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenDeposited)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTokenDeposited)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTokenDepositedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenDepositedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenDeposited represents a TokenDeposited event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenDeposited struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenDeposited is a free log retrieval operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenDeposited(opts *bind.FilterOpts) (*TaikoL1ClientTokenDepositedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenDepositedIterator{contract: _TaikoL1Client.contract, event: "TokenDeposited", logs: logs, sub: sub}, nil
}

// WatchTokenDeposited is a free log subscription operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenDeposited(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenDeposited) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenDeposited")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenDeposited)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenDeposited is a log parse operation binding the contract event 0x26a49ee784523ce049bcbe276a63c7c9dbd9f428b1aa53633e679c9c046e8858.
//
// Solidity: event TokenDeposited(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenDeposited(log types.Log) (*TaikoL1ClientTokenDeposited, error) {
	event := new(TaikoL1ClientTokenDeposited)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenDeposited", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTokenWithdrawnIterator is returned from FilterTokenWithdrawn and is used to iterate over the raw logs and unpacked data for TokenWithdrawn events raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawnIterator struct {
	Event *TaikoL1ClientTokenWithdrawn // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTokenWithdrawnIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTokenWithdrawn)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTokenWithdrawn)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTokenWithdrawnIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTokenWithdrawnIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTokenWithdrawn represents a TokenWithdrawn event raised by the TaikoL1Client contract.
type TaikoL1ClientTokenWithdrawn struct {
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTokenWithdrawn is a free log retrieval operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTokenWithdrawn(opts *bind.FilterOpts) (*TaikoL1ClientTokenWithdrawnIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTokenWithdrawnIterator{contract: _TaikoL1Client.contract, event: "TokenWithdrawn", logs: logs, sub: sub}, nil
}

// WatchTokenWithdrawn is a free log subscription operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTokenWithdrawn(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTokenWithdrawn) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TokenWithdrawn")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTokenWithdrawn)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTokenWithdrawn is a log parse operation binding the contract event 0xc172f6497c150fc242267f743e8e4034b31b16ee123408d6d5f75a81128de114.
//
// Solidity: event TokenWithdrawn(uint256 amount)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTokenWithdrawn(log types.Log) (*TaikoL1ClientTokenWithdrawn, error) {
	event := new(TaikoL1ClientTokenWithdrawn)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TokenWithdrawn", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionContestedIterator is returned from FilterTransitionContested and is used to iterate over the raw logs and unpacked data for TransitionContested events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContestedIterator struct {
	Event *TaikoL1ClientTransitionContested // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTransitionContestedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionContested)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTransitionContested)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTransitionContestedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionContestedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionContested represents a TransitionContested event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionContested struct {
	BlockId     *big.Int
	Tran        TaikoDataTransition
	Contester   common.Address
	ContestBond *big.Int
	Tier        uint16
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransitionContested is a free log retrieval operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionContested(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionContestedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionContestedIterator{contract: _TaikoL1Client.contract, event: "TransitionContested", logs: logs, sub: sub}, nil
}

// WatchTransitionContested is a free log subscription operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionContested(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionContested, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionContested", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionContested)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransitionContested is a log parse operation binding the contract event 0xb4c0a86c1ff239277697775b1e91d3375fd3a5ef6b345aa4e2f6001c890558f6.
//
// Solidity: event TransitionContested(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address contester, uint96 contestBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionContested(log types.Log) (*TaikoL1ClientTransitionContested, error) {
	event := new(TaikoL1ClientTransitionContested)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionContested", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientTransitionProvedIterator is returned from FilterTransitionProved and is used to iterate over the raw logs and unpacked data for TransitionProved events raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProvedIterator struct {
	Event *TaikoL1ClientTransitionProved // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientTransitionProvedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientTransitionProved)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientTransitionProved)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientTransitionProvedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientTransitionProvedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientTransitionProved represents a TransitionProved event raised by the TaikoL1Client contract.
type TaikoL1ClientTransitionProved struct {
	BlockId      *big.Int
	Tran         TaikoDataTransition
	Prover       common.Address
	ValidityBond *big.Int
	Tier         uint16
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterTransitionProved is a free log retrieval operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterTransitionProved(opts *bind.FilterOpts, blockId []*big.Int) (*TaikoL1ClientTransitionProvedIterator, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientTransitionProvedIterator{contract: _TaikoL1Client.contract, event: "TransitionProved", logs: logs, sub: sub}, nil
}

// WatchTransitionProved is a free log subscription operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchTransitionProved(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientTransitionProved, blockId []*big.Int) (event.Subscription, error) {

	var blockIdRule []interface{}
	for _, blockIdItem := range blockId {
		blockIdRule = append(blockIdRule, blockIdItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "TransitionProved", blockIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientTransitionProved)
				if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransitionProved is a log parse operation binding the contract event 0xc195e4be3b936845492b8be4b1cf604db687a4d79ad84d979499c136f8e6701f.
//
// Solidity: event TransitionProved(uint256 indexed blockId, (bytes32,bytes32,bytes32,bytes32) tran, address prover, uint96 validityBond, uint16 tier)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseTransitionProved(log types.Log) (*TaikoL1ClientTransitionProved, error) {
	event := new(TaikoL1ClientTransitionProved)
	if err := _TaikoL1Client.contract.UnpackLog(event, "TransitionProved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientUnpausedIterator is returned from FilterUnpaused and is used to iterate over the raw logs and unpacked data for Unpaused events raised by the TaikoL1Client contract.
type TaikoL1ClientUnpausedIterator struct {
	Event *TaikoL1ClientUnpaused // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientUnpausedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientUnpaused)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientUnpaused)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientUnpausedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientUnpausedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientUnpaused represents a Unpaused event raised by the TaikoL1Client contract.
type TaikoL1ClientUnpaused struct {
	Account common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterUnpaused is a free log retrieval operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterUnpaused(opts *bind.FilterOpts) (*TaikoL1ClientUnpausedIterator, error) {

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientUnpausedIterator{contract: _TaikoL1Client.contract, event: "Unpaused", logs: logs, sub: sub}, nil
}

// WatchUnpaused is a free log subscription operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchUnpaused(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientUnpaused) (event.Subscription, error) {

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Unpaused")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientUnpaused)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Unpaused", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUnpaused is a log parse operation binding the contract event 0x5db9ee0a495bf2e6ff9c91a7834c1ba4fdd244a5e8aa4e537bd38aeae4b073aa.
//
// Solidity: event Unpaused(address account)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseUnpaused(log types.Log) (*TaikoL1ClientUnpaused, error) {
	event := new(TaikoL1ClientUnpaused)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Unpaused", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TaikoL1ClientUpgradedIterator is returned from FilterUpgraded and is used to iterate over the raw logs and unpacked data for Upgraded events raised by the TaikoL1Client contract.
type TaikoL1ClientUpgradedIterator struct {
	Event *TaikoL1ClientUpgraded // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *TaikoL1ClientUpgradedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TaikoL1ClientUpgraded)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(TaikoL1ClientUpgraded)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *TaikoL1ClientUpgradedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TaikoL1ClientUpgradedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TaikoL1ClientUpgraded represents a Upgraded event raised by the TaikoL1Client contract.
type TaikoL1ClientUpgraded struct {
	Implementation common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterUpgraded is a free log retrieval operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TaikoL1Client *TaikoL1ClientFilterer) FilterUpgraded(opts *bind.FilterOpts, implementation []common.Address) (*TaikoL1ClientUpgradedIterator, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TaikoL1Client.contract.FilterLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return &TaikoL1ClientUpgradedIterator{contract: _TaikoL1Client.contract, event: "Upgraded", logs: logs, sub: sub}, nil
}

// WatchUpgraded is a free log subscription operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TaikoL1Client *TaikoL1ClientFilterer) WatchUpgraded(opts *bind.WatchOpts, sink chan<- *TaikoL1ClientUpgraded, implementation []common.Address) (event.Subscription, error) {

	var implementationRule []interface{}
	for _, implementationItem := range implementation {
		implementationRule = append(implementationRule, implementationItem)
	}

	logs, sub, err := _TaikoL1Client.contract.WatchLogs(opts, "Upgraded", implementationRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TaikoL1ClientUpgraded)
				if err := _TaikoL1Client.contract.UnpackLog(event, "Upgraded", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseUpgraded is a log parse operation binding the contract event 0xbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b.
//
// Solidity: event Upgraded(address indexed implementation)
func (_TaikoL1Client *TaikoL1ClientFilterer) ParseUpgraded(log types.Log) (*TaikoL1ClientUpgraded, error) {
	event := new(TaikoL1ClientUpgraded)
	if err := _TaikoL1Client.contract.UnpackLog(event, "Upgraded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
