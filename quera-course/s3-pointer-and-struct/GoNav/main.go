package main

type Zone struct {
    // TODO: Implement Zone
}

type Package struct {
    // TODO: Implement Package
}

type Driver struct {
    // TODO: Implement Driver
}

type Vehicle struct {
    // TODO: Implement Vehicle
}

type DeliveryTruck struct {
    // TODO: Implement DeliveryTruck
}

func CalculateWeight(truck *DeliveryTruck) int {
    // TODO: Implement CalculateWeight
}

func AssignDriver(truck *DeliveryTruck, d *Driver) bool {
    // TODO: Implement AssignDriver
}

func TransferPackage(sender, receiver *DeliveryTruck, trackingCode string) bool {
    // TODO: Implement TransferPackage
}

func FindTrucksByDestination(fleet []*DeliveryTruck, targetZone Zone) []string {
    // TODO: Implement FindTrucksByDestination
}
