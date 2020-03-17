// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package bigint

import (
	"fmt"
	"math/big"
	"testing"
)

func Test_Rat(t *testing.T) {
	rat := big.NewRat(1, 3)
	fmt.Println(rat)
	rat.Neg(rat)
	fmt.Println(rat)
	fmt.Println(rat.Mul(rat, big.NewRat(3, 1)).Num())
}

func Test_Int(t *testing.T) {

	newInt:= new(big.Int)
	newInt,ok:= newInt.SetString("122332123123123123123123123123123123123123123211323",0)
	fmt.Println(newInt,ok)
	//newInt=newInt.And(newInt,big.NewInt(1))
	newInt.Mul(newInt, newInt)
	fmt.Println(newInt.Int64(),newInt.String())
}