package proxy

import "github.com/Goplush/lachesisnode/m/hashgraph"

type AppProxy interface {
	SubmitCh() chan []byte
	CommitBlock(block hashgraph.Block) ([]byte, error)
}

type LachesisProxy interface {
	CommitCh() chan hashgraph.Block
	SubmitTx(tx []byte) error
}
