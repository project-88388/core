package params

import (
	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/terra-money/core/v2/app/config"
)

func RegisterDenomsConfig() error {

	err := sdk.RegisterDenom(config.MicroUSDDenom, sdk.NewDecWithPrec(1, 6))
	if err != nil {

	err := sdk.RegisterDenom(config.MicroLuna, sdk.NewDecWithPrec(1, 6))
	if err != nil {
		return err
	}
}
	
	return nil
}
