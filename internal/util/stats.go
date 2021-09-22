package util

import (
	"os"
	"sync"
	"sync/atomic"
	"time"

	log "github.com/sirupsen/logrus"
)

type StatsType int

const (
	Consume StatsType = iota
	Produce
	ProduceLatency
)

type Stats struct {
	Consume       uint64
	ConsumePerSec uint64
	Produce       uint64
	ProducePerSec uint64

	ProduceLatency      uint64
	ProduceLatencyCount uint64
	ProduceLatencyAvg   uint64
	ProduceLatencyMutex sync.RWMutex
}

func NewStats() *Stats {
	return &Stats{}
}

func (s *Stats) PerSec() {
	lastConsume := uint64(0)
	lastProduce := uint64(0)
	ticker := time.NewTicker(1 * time.Second)
	for {
		select {
		case <-ticker.C:
			s.ConsumePerSec = s.Consume - lastConsume
			s.ProducePerSec = s.Produce - lastProduce

			s.ProduceLatencyMutex.Lock()
			latencyCount := s.ProduceLatencyCount
			latencySum := s.ProduceLatency
			s.ProduceLatencyCount = 0
			s.ProduceLatency = 0
			s.ProduceLatencyMutex.Unlock()

			if latencyCount > 0 {
				s.ProduceLatencyAvg = latencySum / latencyCount
			}

			lastConsume = s.Consume
			lastProduce = s.Produce

		}
	}
}

func (s *Stats) Inc(statsType StatsType) {
	s.IncBy(statsType, 1)
}

func (s *Stats) IncBy(statsType StatsType, value uint64) {
	switch statsType {
	case Consume:
		atomic.AddUint64(&s.Consume, value)
	case Produce:
		atomic.AddUint64(&s.Produce, value)
	case ProduceLatency:
		s.ProduceLatencyMutex.Lock()
		s.ProduceLatencyCount += 1
		s.ProduceLatency += value
		s.ProduceLatencyMutex.Unlock()
	}
}

func (s *Stats) Monitor(interval time.Duration) {
	log.Infof("Stats display interval: %s", interval)
	ticker := time.NewTicker(interval)
	log.SetOutput(os.Stdout)
	for {
		select {
		case <-ticker.C:
			log.Infof("Consumed events:[%d]", s.Consume)
			log.Infof("Produced events: [%d]", s.Produce)
		}
	}
}
