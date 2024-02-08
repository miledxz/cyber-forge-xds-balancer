package processor

import (
	"math/rand"

	"github.com/envoyproxy/go-control-plane/pkg/cache"
	"github.com/sirupsen/logrus"
)

type Processor struct {
	cache  cache.SnapshotCache
	nodeID string

	snapshotVersion int64

	logrus.FieldLogger

	
}

func NewProcessor(cache cache.SnapshotCache, nodeID string, log logrus.FieldLogger) *Processor {
	return &Processor{
		cache:           cache,
		nodeID:          nodeID,
		snapshotVersion: rand.Int63n(1000),
		FieldLogger:     log,
	}
}