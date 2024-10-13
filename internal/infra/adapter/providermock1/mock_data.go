package providermock1

import "time"

var mockData = []MockAdapter1ResponseItem{
	{
		TrackingCode: "abc",
		Status:       Delivered,
		UpdateAt:     time.Now().Add(5 * time.Minute),
	},
	{
		TrackingCode: "lls",
		Status:       InWarehouse,
		UpdateAt:     time.Now().Add(6 * time.Minute),
	},
	{
		TrackingCode: "asdf",
		Status:       Processing,
		UpdateAt:     time.Now().Add(7 * time.Minute),
	},
	{
		TrackingCode: "weori",
		Status:       Delivered,
		UpdateAt:     time.Now().Add(8 * time.Minute),
	},
	{
		TrackingCode: "awer",
		Status:       Processing,
		UpdateAt:     time.Now().Add(9 * time.Minute),
	},
	{
		TrackingCode: "dddc",
		Status:       Processing,
		UpdateAt:     time.Now().Add(10 * time.Minute),
	},
}
