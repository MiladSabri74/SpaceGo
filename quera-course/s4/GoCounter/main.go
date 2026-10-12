package main

import (
	"sync/atomic"
)

type Metrics struct {
	viewStat  atomic.Int64
	errorStat atomic.Int64
}

func (m *Metrics) RecordView() {
	// TODO: Atomically increment the view counter by 1
	m.viewStat.Add(1)
}

func (m *Metrics) RecordError() {
	// TODO: Atomically increment the error counter by 1
	m.errorStat.Add(1)
}

func (m *Metrics) GetStats() (int64, int64) {
	return m.viewStat.Load(), m.errorStat.Load()
}
