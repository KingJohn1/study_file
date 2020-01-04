// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package main

import "testing"

func TestAdd3(t *testing.T) {

	r := Add(1, 2)
	println(t.Name())
	if r !=3 {

		t.Errorf("Add(1, 2) failed. Got %d, expected 3.", r)
	}

}