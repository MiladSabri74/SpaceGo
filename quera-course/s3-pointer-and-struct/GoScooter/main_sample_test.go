package main

import (
    "testing"
)

func TestNewScooter_Sample(t *testing.T) {
    s := NewScooter(20, 80)

    if s.speed != 20 || s.battery != 80 {
        t.Errorf("NewScooter did not initialize properly. Got speed: %d, battery: %d", s.speed, s.battery)
    }
}

func TestGetSpeed_Sample(t *testing.T) {
    s := &Scooter{speed: 15}

    if got := GetSpeed(s); got != 15 {
        t.Errorf("GetSpeed() = %d; want 15", got)
    }
}

func TestGetBattery_Sample(t *testing.T) {
    s := &Scooter{battery: 60}

    if got := GetBattery(s); got != 60 {
        t.Errorf("GetBattery() = %d; want 60", got)
    }
}

func TestChargeScooter_Sample(t *testing.T) {
    s := &Scooter{battery: 40}

    // Charging for 20 minutes adds 10 battery
    ChargeScooter(s, 20)

    if s.battery != 50 {
        t.Errorf("ChargeScooter() resulted in battery %d; want 50", s.battery)
    }
}

func TestTryRide_Sample(t *testing.T) {
    s := &Scooter{speed: 10, battery: 100}

    // Distance of 20 takes 2 units of time (20/10), requires 10 battery (20/2)
    result := TryRide(s, 20)

    if result != "2.00" {
        t.Errorf("TryRide() returned %q; want \"2.00\"", result)
    }

    if s.battery != 90 {
        t.Errorf("TryRide() left battery at %d; want 90", s.battery)
    }
}
