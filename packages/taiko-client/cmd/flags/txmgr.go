package flags

import (
	"time"

	"github.com/urfave/cli/v2"
)

var (
	NumConfirmations = &cli.Uint64Flag{
		Name:     "tx.numConfirmations",
		Usage:    "Number of confirmations which we will wait after sending a transaction",
		Value:    0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_NUM_CONFIRMATIONS"},
	}
	SafeAbortNonceTooLowCount = &cli.Uint64Flag{
		Name: "tx.safeAbortNonceTooLowCount",
		Usage: "Number of ErrNonceTooLow observations required to give up on " +
			"a tx at a particular nonce without receiving confirmation",
		Value:    3,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_SAFE_ABORT_NONCE_TOO_LOW_COUNT"},
	}
	FeeLimitMultiplier = &cli.Uint64Flag{
		Name:     "tx.feeLimitMultiplier",
		Usage:    "The multiplier applied to fee suggestions to put a hard limit on fee increases",
		Value:    10,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_FEE_LIMIT_MULTIPLIER"},
	}
	FeeLimitThreshold = &cli.Float64Flag{
		Name: "tx.feeLimitThreshold",
		Usage: "The minimum threshold (in GWei) at which fee bumping starts to be capped. " +
			"Allows arbitrary fee bumps below this threshold.",
		Value:    100.0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_FEE_LIMIT_THRESHOLD"},
	}
	MinTipCap = &cli.Float64Flag{
		Name:     "tx.minTipCap",
		Usage:    "Enforces a minimum tip cap (in GWei) to use when determining tx fees. 1 GWei by default.",
		Value:    1.0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_MIN_TIP_CAP"},
	}
	MinBaseFee = &cli.Float64Flag{
		Name:     "tx.minBaseFee",
		Usage:    "Enforces a minimum base fee (in GWei) to assume when determining tx fees. 1 GWei by default.",
		Value:    1.0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_MIN_BASE_FEE"},
	}
	ResubmissionTimeout = &cli.DurationFlag{
		Name:     "tx.resubmissionTimeout",
		Usage:    "Duration we will wait before resubmitting a transaction to L1",
		Value:    48 * time.Second,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_RESUBMISSION_TIMEOUT"},
	}
	TxSendTimeout = &cli.DurationFlag{
		Name:     "tx.sendTimeout",
		Usage:    "Timeout for sending transactions. If 0 it is disabled.",
		Value:    0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_SEND_TIMEOUT"},
	}
	TxNotInMempoolTimeout = &cli.DurationFlag{
		Name:     "tx.notInMempoolTimeout",
		Usage:    "Timeout for aborting a tx send if the tx does not make it to the mempool.",
		Value:    2 * time.Minute,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_NOT_IN_MEMPOOL_TIMEOUT"},
	}
	ReceiptQueryInterval = &cli.DurationFlag{
		Name:     "tx.receiptQueryInterval",
		Usage:    "Frequency to poll for receipts",
		Value:    12 * time.Second,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_RECEIPT_QUERY_INTERVAL"},
	}
	TxGasLimit = &cli.Uint64Flag{
		Name:     "tx.gasLimit",
		Usage:    "Gas limit will be used for transactions (0 means using gas estimation)",
		Value:    0,
		Category: txmgrCategory,
		EnvVars:  []string{"TX_GAS_LIMIT"},
	}
)

var TxmgrFlags = []cli.Flag{
	NumConfirmations,
	SafeAbortNonceTooLowCount,
	FeeLimitMultiplier,
	FeeLimitThreshold,
	MinTipCap,
	MinBaseFee,
	ResubmissionTimeout,
	TxSendTimeout,
	TxNotInMempoolTimeout,
	ReceiptQueryInterval,
	TxGasLimit,
}
