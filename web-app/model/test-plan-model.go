package model

import (
	"fmt"
	"time"
)

// Plan is the type for subscription plans
type TestPlan struct {
	ID                  int
	PlanName            string
	PlanAmount          int
	PlanAmountFormatted string
	CreatedAt           time.Time
	UpdatedAt           time.Time
}

func (p *TestPlan) GetAll() ([]*Plan, error) {
	var plans []*Plan

	plan := Plan{
		ID:         1,
		PlanName:   "abc",
		PlanAmount: 100,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	plans = append(plans, &plan)

	return plans, nil
}

// GetOne returns one plan by id
func (p *TestPlan) GetOne(id int) (*Plan, error) {
	plan := Plan{
		ID:         1,
		PlanName:   "abc",
		PlanAmount: 100,
		CreatedAt:  time.Now(),
		UpdatedAt:  time.Now(),
	}

	return &plan, nil
}

// SubscribeUserToPlan subscribes a user to one plan by insert
// values into user_plans table
func (p *TestPlan) SubscribeUserToPlan(user User, plan Plan) error {
	return nil
}

// AmountForDisplay formats the price we have in the DB as a currency string
func (p *TestPlan) AmountForDisplay() string {
	amount := float64(p.PlanAmount) / 100.0
	return fmt.Sprintf("$%.2f", amount)
}
