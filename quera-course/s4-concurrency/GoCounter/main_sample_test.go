package main

import (
    "testing"
)

func TestMetrics_Sample(t *testing.T) {
    var m Metrics

    m.RecordView()
    m.RecordView()
    m.RecordView()

    m.RecordError()

    views, errors := m.GetStats()

    if views != 3 {
        t.Errorf("Expected 3 views, but got %d", views)
    }

    if errors != 1 {
        t.Errorf("Expected 1 error, but got %d", errors)
    }
}
