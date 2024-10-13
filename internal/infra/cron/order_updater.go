package cron

import (
	"context"
	"github.com/robfig/cron/v3"
	"log"
	"neosync/internal/domain/order"
	"neosync/internal/domain/provider"
	"time"
)

type OrderUpdaterConfig struct {
	CronExpression string        `koanf:"cron_expression"` // "0 0 * * *"
	TimeoutMinutes time.Duration `koanf:"timeout_minutes"` // 10*time.Minute
}

type OrderStatusUpdater struct {
	cfg             OrderUpdaterConfig
	orderService    *order.Service
	providerService *provider.Service
}

func NewOrderStatusUpdater(cfg OrderUpdaterConfig, orderService *order.Service, providerService *provider.Service) *OrderStatusUpdater {
	return &OrderStatusUpdater{
		cfg:             cfg,
		orderService:    orderService,
		providerService: providerService,
	}
}

// StartCronJob sets up a cron job to update order statuses daily
func (u *OrderStatusUpdater) StartCronJob() {
	c := cron.New()
	// Schedule the job to run every day at midnight
	_, err := c.AddFunc(u.cfg.CronExpression, func() {
		log.Println("Starting daily batch update for order statuses...")
		if err := u.updateOrderStatuses(); err != nil {
			log.Printf("Failed to update order statuses: %v", err)
		}
	})
	if err != nil {
		log.Fatalf("Failed to add cron job: %v", err)
	}

	c.Start()
	log.Println("Cron job started. Running daily at midnight.")
}

// updateOrderStatuses executes the batch update process
func (u *OrderStatusUpdater) updateOrderStatuses() error {
	ctx, cancel := context.WithTimeout(context.Background(), u.cfg.TimeoutMinutes*time.Minute)
	defer cancel()

	// Step 1: this can be optimised by checking the order list before and getting out the distinct list of orders
	// but for the lack of time I leave it here just like this :,)
	// even I've designed the input of the providerService.BatchFetchAll to accept limited amount of providers
	// I think the best place to implement is in the UpdateStatusMany this function should Accept an slice of orders and the repo method should return the distinct providers in a map inside the for.Next() loop
	// providers := make(map[uint]struct{})
	// for.Next(){
	// order := ...
	// providers[order.ProviderID] = struct{}
	//}
	// this just makes a map of providers in the most efficient way because we have to query over the orders table here
	providers, err := u.providerService.GetAll(ctx)
	if err != nil {
		return err
	}

	ids := make([]uint, 0)
	for _, value := range providers {
		ids = append(ids, value.ID)
	}

	// Step 2: Get statuses from external providers
	statusResponse, err := u.providerService.BatchFetchAll(ctx, provider.GetAllRequest{
		ProviderIDs: ids,
	})
	if err != nil {
		return err
	}

	// Step 3: Update statuses in the database
	err = u.orderService.UpdateStatusMany(ctx, order.UpdateStatusManyRequest{
		TrackingStatusMapping: statusResponse.Mapping,
	})

	return err
}
