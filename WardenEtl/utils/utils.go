package utils

import (
	"math/big"

	"github.com/shopspring/decimal"

	log "github.com/sirupsen/logrus"
)

func GetOnboardFeeToken(tokenPrice *big.Int, onboardFeeDollars int64) *big.Int {
	feeToken := decimal.NewFromBigInt(
		big.NewInt(onboardFeeDollars), 8).Div(
		decimal.NewFromBigInt(tokenPrice, 0)).Mul(
		decimal.NewFromBigInt(big.NewInt(1), 18))
	return feeToken.BigInt()
}

func GetOffboardFeeToken(etherPrice, tokenPrice, gasPrice *big.Int, offboardFeeDollars int64, gas uint64) *big.Int {
	gasFee := decimal.NewFromBigInt(gasPrice, -18).Mul(
		decimal.NewFromInt(int64(gas)),
	).Mul(
		decimal.NewFromBigInt(etherPrice, -8),
	)

	log.Infof("Offboard gasfee[%v]\n", gasFee)
	totalFee := decimal.NewFromInt(offboardFeeDollars).Add(gasFee)

	feeToken := totalFee.Div(
		decimal.NewFromBigInt(tokenPrice, -8),
	).Mul(
		decimal.NewFromBigInt(big.NewInt(1), 18),
	)

	return feeToken.BigInt()
}
