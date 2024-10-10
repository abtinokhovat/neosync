package order

import (
	"time"
)

// Order represents a customer's order
type Order struct {
	ID         uint
	CreatedAt  time.Time
	UpdatedAt  time.Time
	CustomerID uint
	ProviderID uint
	Status     Status
}

// Status represents the status of an order
type Status uint

// Status constants
const (
	_                   = iota
	Pending      Status = 1
	InProgress   Status = 2
	ProviderSeen Status = 3
	PickedUp     Status = 4
	Delivered    Status = 5
)

// Validate checks if the provided status is within the valid range
func (s Status) Validate() bool {
	return Pending <= s && s <= Delivered
}

// String returns the English string representation of the status
func (s Status) String() string {
	if name, ok := englishNameMapping[s]; ok {
		return name
	}
	return "Unknown Status"
}

// FaString returns the Persian string representation of the status
func (s Status) FaString() string {
	if name, ok := persianNameMapping[s]; ok {
		return name
	}
	return "وضعیت ناشناخته"
}

// English name mapping for statuses
var englishNameMapping = map[Status]string{
	Pending:      "Pending",
	InProgress:   "In Progress",
	ProviderSeen: "Provider Seen",
	PickedUp:     "Picked Up",
	Delivered:    "Delivered",
}

// Persian name mapping for statuses
var persianNameMapping = map[Status]string{
	Pending:      "در انتظار",
	InProgress:   "در حال انجام",
	ProviderSeen: "پرووایدر مشاهده شد",
	PickedUp:     "تحویل گرفته شد",
	Delivered:    "تحویل داده شد",
}
