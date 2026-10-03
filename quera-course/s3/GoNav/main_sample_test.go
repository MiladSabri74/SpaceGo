package main

import (
    "reflect"
    "testing"
)

func TestCalculateWeight_Sample(t *testing.T) {
    // Normal setup: 1 truck with 2 packages
    truck := &DeliveryTruck{
        Cargo: []*Package{
            {Weight: 10},
            {Weight: 15},
        },
    }

    // Act
    weight := CalculateWeight(truck)

    // Assert
    if weight != 25 {
        t.Errorf("Expected weight to be 25, got %d", weight)
    }
}

func TestAssignDriver_Sample(t *testing.T) {
    // Normal setup: a driverless truck and a valid driver
    truck := &DeliveryTruck{}
    driver := &Driver{Name: "Sam", LicenseGrade: 2}

    // Act
    success := AssignDriver(truck, driver)

    // Assert
    if !success {
        t.Errorf("Expected driver assignment to succeed")
    }
    if truck.Driver == nil || truck.Driver.Name != "Sam" {
        t.Errorf("Expected driver to be correctly assigned to the truck")
    }
}

func TestLoadPackage_Sample(t *testing.T) {
    // Normal setup: Truck with capacity and driver, valid package
    truck := &DeliveryTruck{
        Vehicle: Vehicle{MaxCapacity: 100},
        Driver:  &Driver{Name: "Alex"},
        Cargo:   []*Package{},
    }
    pkg := &Package{TrackingCode: "PKG-123", Weight: 20}

    // Act
    success := LoadPackage(truck, pkg)

    // Assert
    if !success {
        t.Errorf("Expected package to be loaded successfully")
    }
    if len(truck.Cargo) != 1 || truck.Cargo[0].TrackingCode != "PKG-123" {
        t.Errorf("Expected package to be inside the truck's cargo")
    }
}

func TestTransferPackage_Sample(t *testing.T) {
    // Normal setup: Two trucks in the same zone, receiver has driver and capacity
    zone := Zone{City: "Chicago", Code: 606}
    pkg := &Package{TrackingCode: "X-100", Weight: 10, IsFragile: false}

    sender := &DeliveryTruck{
        CurrentZone: zone,
        Cargo:       []*Package{pkg},
    }
    receiver := &DeliveryTruck{
        CurrentZone: zone,
        Vehicle:     Vehicle{MaxCapacity: 50},
        Driver:      &Driver{Name: "Taylor", LicenseGrade: 2},
        Cargo:       []*Package{},
    }

    // Act
    success := TransferPackage(sender, receiver, "X-100")

    // Assert
    if !success {
        t.Errorf("Expected package transfer to succeed")
    }
    if len(sender.Cargo) != 0 {
        t.Errorf("Expected sender to have empty cargo after transfer")
    }
    if len(receiver.Cargo) != 1 || receiver.Cargo[0].TrackingCode != "X-100" {
        t.Errorf("Expected receiver to have the transferred package")
    }
}

func TestFindTrucksByDestination_Sample(t *testing.T) {
    // Normal setup: Fleet of trucks, finding the one going to Miami
    miami := Zone{City: "Miami", Code: 305}
    dallas := Zone{City: "Dallas", Code: 214}

    fleet := []*DeliveryTruck{
        {
            Vehicle: Vehicle{LicensePlate: "TRK-MIAMI"},
            Cargo:   []*Package{{Destination: miami}},
        },
        {
            Vehicle: Vehicle{LicensePlate: "TRK-DALLAS"},
            Cargo:   []*Package{{Destination: dallas}},
        },
    }

    // Act
    results := FindTrucksByDestination(fleet, miami)

    // Assert
    expected := []string{"TRK-MIAMI"}
    if !reflect.DeepEqual(results, expected) {
        t.Errorf("Expected %v, got %v", expected, results)
    }
}
