package processor

import "github.com/ElrondNetwork/elrond-go/dataRetriever"

// ArgTxInterceptorProcessor is the argument for the interceptor processor used for transaction-like data
type ArgTxInterceptorProcessor struct {
	ShardedDataCache dataRetriever.ShardedDataCacherNotifier
}