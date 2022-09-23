package chain

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

	log.Infof("Offboard gasfee[%v]", gasFee)
	totalFee := decimal.NewFromInt(offboardFeeDollars).Add(gasFee)

	feeToken := totalFee.Div(
		decimal.NewFromBigInt(tokenPrice, -8),
	).Mul(
		decimal.NewFromBigInt(big.NewInt(1), 18),
	)

	return feeToken.BigInt()
}

func ToWei(iamount interface{}, decimals int) *big.Int {
	amount := decimal.NewFromFloat(0)
	switch v := iamount.(type) {
	case string:
		amount, _ = decimal.NewFromString(v)
	case float64:
		amount = decimal.NewFromFloat(v)
	case int64:
		amount = decimal.NewFromFloat(float64(v))
	case decimal.Decimal:
		amount = v
	case *decimal.Decimal:
		amount = *v
	}

	mul := decimal.NewFromFloat(float64(10)).Pow(decimal.NewFromFloat(float64(decimals)))
	result := amount.Mul(mul)

	wei := new(big.Int)
	wei.SetString(result.String(), 10)

	return wei
}
