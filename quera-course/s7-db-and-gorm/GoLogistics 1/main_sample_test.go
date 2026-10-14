package main

import (
    "testing"
)

func TestSampleMigration(t *testing.T) {
    testDB := GetDB("file::memory:?cache=shared")
    if testDB == nil {
        t.Fatalf("Failed to open test DB")
    }

    err := AutoMigrateModels(testDB)
    if err != nil {
        t.Fatalf("AutoMigrateModels returned error: %v", err)
    }
}
