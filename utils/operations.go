package utils

import (
	"math"

	"github.com/kaseat/pManager/models"
	"github.com/oleiade/lane"
)

// GetSum returns blance for given operations
func GetSum(operations []models.Operation) float64 {
	sum := int64(0)
	for _, opertion := range operations {
		amount := int64(opertion.Price*1e6) * opertion.Volume
		switch opertion.OperationType {
		case models.PayIn, models.Sell:
			sum += amount
		default:
			sum -= amount
		}
	}
	return float64(sum) / 1e6
}

// GetAverage returns average price of given operations
func GetAverage(ops []models.Operation) float64 {
	d := lane.NewDeque()
	for _, op := range ops {
		if op.OperationType == models.Buy {
			d.Append(op)
		} else {
			for {
				if d.Empty() {
					break
				}
				o := d.Shift().(models.Operation)
				if o.Volume-op.Volume <= 0 {
					op.Volume -= o.Volume
				} else {
					o.Volume -= op.Volume
					d.Prepend(o)
					break
				}
			}
		}
	}

	cost, vol := float64(0), float64(0)
	for {
		if d.Empty() {
			break
		}
		op := d.Pop().(models.Operation)
		cost += op.Price * float64(op.Volume)
		vol += float64(op.Volume)
	}

	result := float64(0)
	if vol != 0 {
		result = math.Round(cost/vol*1e6) / 1e6
	}
	return result
}