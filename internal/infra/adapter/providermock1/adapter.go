package providermock1

import (
	"context"
	"neosync/internal/domain/order"
	"neosync/internal/domain/provider"
	"time"
)

type Config struct {
	// url would be read from here or any adapter configs, also the config package can be easily expanded too reading the configs from various sources like a hashicorp/vault or a db
	URL string `koanf:"url"`
}

type Status uint

const (
	_           Status = iota
	InWarehouse Status = 1
	Processing  Status = 2
	Delivered   Status = 3
)

func (s Status) OrderStatus() uint {
	return uint(statusOrderStatusMapping[s])
}

var statusOrderStatusMapping = map[Status]order.Status{
	InWarehouse: order.PickedUp,
	Processing:  order.PickedUp,
	Delivered:   order.Delivered,
}

type MockAdapter1 struct {
}

type MockAdapter1ResponseItem struct {
	TrackingCode string
	Status       Status
	UpdateAt     time.Time
}

func New() *MockAdapter1 {
	return &MockAdapter1{}
}

func (a *MockAdapter1) Name() string {
	return "mock-provider-1"
}

func (a *MockAdapter1) GetAll(_ context.Context) (map[string]provider.AdapterResponseItem, error) {
	orders := make(map[string]provider.AdapterResponseItem)
	for _, record := range mockData {
		orders[record.TrackingCode] = provider.AdapterResponseItem{
			Status:    record.Status.OrderStatus(),
			UpdatedAt: record.UpdateAt,
		}
	}
	return orders, nil
}
