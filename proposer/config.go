package proposer

import (
	"crypto/ecdsa"
	"fmt"
	"math/big"
	"net/url"
	"strings"
	"time"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"

	"github.com/taikoxyz/taiko-client/cmd/flags"
	"github.com/taikoxyz/taiko-client/pkg/rpc"
)

// Config contains all configurations to initialize a Taiko proposer.
type Config struct {
	*rpc.ClientConfig
	AssignmentHookAddress               common.Address
	L1ProposerPrivKey                   *ecdsa.PrivateKey
	ExtraData                           string
	ProposeInterval                     *time.Duration
	LocalAddresses                      []common.Address
	LocalAddressesOnly                  bool
	ProposeEmptyBlocksInterval          *time.Duration
	MaxProposedTxListsPerEpoch          uint64
	ProposeBlockTxGasLimit              uint64
	ProposeBlockTxReplacementMultiplier uint64
	WaitReceiptTimeout                  time.Duration
	ProposeBlockTxGasTipCap             *big.Int
	ProverEndpoints                     []*url.URL
	OptimisticTierFee                   *big.Int
	SgxTierFee                          *big.Int
	PseZkevmTierFee                     *big.Int
	SgxAndPseZkevmTierFee               *big.Int
	TierFeePriceBump                    *big.Int
	MaxTierFeePriceBumps                uint64
	IncludeParentMetaHash               bool
}

// NewConfigFromCliContext initializes a Config instance from
// command line flags.
func NewConfigFromCliContext(c *cli.Context) (*Config, error) {
	l1ProposerPrivKey, err := crypto.ToECDSA(
		common.Hex2Bytes(c.String(flags.L1ProposerPrivKey.Name)),
	)
	if err != nil {
		return nil, fmt.Errorf("invalid L1 proposer private key: %w", err)
	}

	// Proposing configuration
	var proposingInterval *time.Duration
	if c.IsSet(flags.ProposeInterval.Name) {
		interval := c.Duration(flags.ProposeInterval.Name)
		proposingInterval = &interval
	}

	var proposeEmptyBlocksInterval *time.Duration
	if c.IsSet(flags.ProposeEmptyBlocksInterval.Name) {
		interval := c.Duration(flags.ProposeEmptyBlocksInterval.Name)
		proposeEmptyBlocksInterval = &interval
	}

	var localAddresses []common.Address
	if c.IsSet(flags.TxPoolLocals.Name) {
		for _, account := range strings.Split(c.String(flags.TxPoolLocals.Name), ",") {
			if trimmed := strings.TrimSpace(account); !common.IsHexAddress(trimmed) {
				return nil, fmt.Errorf("invalid account in --txpool.locals: %s", trimmed)
			}
			localAddresses = append(localAddresses, common.HexToAddress(account))
		}
	}

	proposeBlockTxReplacementMultiplier := c.Uint64(flags.ProposeBlockTxReplacementMultiplier.Name)
	if proposeBlockTxReplacementMultiplier == 0 {
		return nil, fmt.Errorf(
			"invalid --proposeBlockTxReplacementMultiplier value: %d",
			proposeBlockTxReplacementMultiplier,
		)
	}

	var proposeBlockTxGasTipCap *big.Int
	if c.IsSet(flags.ProposeBlockTxGasTipCap.Name) {
		proposeBlockTxGasTipCap = new(big.Int).SetUint64(c.Uint64(flags.ProposeBlockTxGasTipCap.Name))
	}

	var proverEndpoints []*url.URL
	for _, e := range strings.Split(c.String(flags.ProverEndpoints.Name), ",") {
		endpoint, err := url.Parse(e)
		if err != nil {
			return nil, err
		}
		proverEndpoints = append(proverEndpoints, endpoint)
	}

	return &Config{
		ClientConfig: &rpc.ClientConfig{
			L1Endpoint:        c.String(flags.L1WSEndpoint.Name),
			L2Endpoint:        c.String(flags.L2HTTPEndpoint.Name),
			TaikoL1Address:    common.HexToAddress(c.String(flags.TaikoL1Address.Name)),
			TaikoL2Address:    common.HexToAddress(c.String(flags.TaikoL2Address.Name)),
			TaikoTokenAddress: common.HexToAddress(c.String(flags.TaikoTokenAddress.Name)),
			RetryInterval:     c.Duration(flags.BackOffRetryInterval.Name),
			Timeout:           c.Duration(flags.RPCTimeout.Name),
		},
		AssignmentHookAddress:               common.HexToAddress(c.String(flags.ProposerAssignmentHookAddress.Name)),
		L1ProposerPrivKey:                   l1ProposerPrivKey,
		ExtraData:                           c.String(flags.ExtraData.Name),
		ProposeInterval:                     proposingInterval,
		LocalAddresses:                      localAddresses,
		LocalAddressesOnly:                  c.Bool(flags.TxPoolLocalsOnly.Name),
		ProposeEmptyBlocksInterval:          proposeEmptyBlocksInterval,
		MaxProposedTxListsPerEpoch:          c.Uint64(flags.MaxProposedTxListsPerEpoch.Name),
		ProposeBlockTxGasLimit:              c.Uint64(flags.ProposeBlockTxGasLimit.Name),
		ProposeBlockTxReplacementMultiplier: proposeBlockTxReplacementMultiplier,
		WaitReceiptTimeout:                  c.Duration(flags.WaitReceiptTimeout.Name),
		ProposeBlockTxGasTipCap:             proposeBlockTxGasTipCap,
		ProverEndpoints:                     proverEndpoints,
		OptimisticTierFee:                   new(big.Int).SetUint64(c.Uint64(flags.OptimisticTierFee.Name)),
		SgxTierFee:                          new(big.Int).SetUint64(c.Uint64(flags.SgxTierFee.Name)),
		PseZkevmTierFee:                     new(big.Int).SetUint64(c.Uint64(flags.PseZkevmTierFee.Name)),
		SgxAndPseZkevmTierFee:               new(big.Int).SetUint64(c.Uint64(flags.SgxAndPseZkevmTierFee.Name)),
		TierFeePriceBump:                    new(big.Int).SetUint64(c.Uint64(flags.TierFeePriceBump.Name)),
		MaxTierFeePriceBumps:                c.Uint64(flags.MaxTierFeePriceBumps.Name),
		IncludeParentMetaHash:               c.Bool(flags.ProposeBlockIncludeParentMetaHash.Name),
	}, nil
}
