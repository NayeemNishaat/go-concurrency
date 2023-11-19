package lib

import "web/model"

func (cfg *Config) GetInvoice(u *model.User, plan *model.Plan) (string, error) {
	// TODO: Perform complex calculation
	return plan.PlanAmountFormatted, nil
}
