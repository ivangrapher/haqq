package v1_3_0

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/cosmos/cosmos-sdk/types/module"
	distrkeeper "github.com/cosmos/cosmos-sdk/x/distribution/keeper"
	distrtypes "github.com/cosmos/cosmos-sdk/x/distribution/types"
	upgradetypes "github.com/cosmos/cosmos-sdk/x/upgrade/types"
)

// CreateUpgradeHandler creates an SDK upgrade handler for v1.0.2
func CreateUpgradeHandler(
	mm *module.Manager,
	configurator module.Configurator,
	dk distrkeeper.Keeper,
) upgradetypes.UpgradeHandler {
	return func(ctx sdk.Context, _ upgradetypes.Plan, vm module.VersionMap) (module.VersionMap, error) {
		logger := ctx.Logger()
		logger.Debug("run migration v1.3.0")

		// Refs:
		// - https://docs.cosmos.network/master/building-modules/upgrade.html#registering-migrations
		// - https://docs.cosmos.network/master/migrations/chain-upgrade-guide-044.html#chain-upgrade

		MigrateDistributionParams(ctx, dk)

		return mm.RunMigrations(ctx, configurator, vm)
	}
}

func MigrateDistributionParams(ctx sdk.Context, dk distrkeeper.Keeper) error {
	newDistrParams := distrtypes.Params{
		CommunityTax:        sdk.NewDecWithPrec(10, 3),
		BaseProposerReward:  sdk.NewDecWithPrec(1, 3),
		BonusProposerReward: sdk.NewDecWithPrec(5, 3),
		WithdrawAddrEnabled: true,
	}

	dk.SetParams(ctx, newDistrParams)

	return nil
}
