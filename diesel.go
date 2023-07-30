package main

import (
"fmt"
"math/rand"
"time"
)

// DieselLevel is a struct that represents the level of diesel in a tank.
type DieselLevel struct {
CurrentLevel int
MinimumLevel int
}

// Supplier is a struct that represents a diesel supplier.
type Supplier struct {
Name string
Price float64
}

// Suppliers is a slice of Supplier.
type Suppliers []Supplier

// GetQuotes gets quotes from three different suppliers for the specified diesel level.
func GetQuotes(level DieselLevel) Suppliers {
suppliers := Suppliers{
Supplier{"Supplier 1", 1.00},
Supplier{"Supplier 2", 1.10},
Supplier{"Supplier 3", 1.20},
}

// Randomly select three suppliers.
selectedSuppliers := make([]Supplier, 3)
for i := 0; i < 3; i++ {
    selectedSupplier := suppliers[rand.Intn(len(suppliers))]
    selectedSuppliers[i] = selectedSupplier
}

return selectedSuppliers
}

// Main is the entry point for the application.
func main() {
// Create a DieselLevel struct.
level := DieselLevel{
CurrentLevel: 100,
MinimumLevel: 50,
}

// Get quotes from three different suppliers.
quotes := GetQuotes(level)

// Print the quotes.
for _, quote := range quotes {
    fmt.Println(quote.Name, quote.Price)
}

// If the current level of diesel is below the minimum level, send an email notification to the manager.
if level.CurrentLevel < level.MinimumLevel {
    fmt.Println("Sending email notification to manager.")
