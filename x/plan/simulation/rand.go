package simulation

import (
	"math/rand"

	"github.com/sentinel-official/hub/v0.5/x/plan/types"
)

func RandomPlan(r *rand.Rand, plans types.Plans) types.Plan {
	if len(plans) == 0 {
		return types.Plan{
			ID: 1,
		}
	}

	return plans[r.Intn(
		len(plans),
	)]
}
