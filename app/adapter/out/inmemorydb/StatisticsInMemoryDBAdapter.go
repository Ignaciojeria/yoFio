package inmemorydb

import (
	"sync"
	"yofio/app/domain"
)

type StatisticsInMemoryDBAdapter struct {
	mu sync.Mutex
}

var inMemoryStatistics domain.Statistics = domain.Statistics{}

func (a *StatisticsInMemoryDBAdapter) LoadStatisticsPortOUT() domain.Statistics {
	return inMemoryStatistics
}

func (a *StatisticsInMemoryDBAdapter) UpdateStatisticsPortOUT(statistics domain.Statistics) error {
	a.mu.Lock()
	inMemoryStatistics = statistics
	a.mu.Unlock()
	return nil
}
