// Copyright (c) liuhao. 2012-2050. All rights reserved.
package _83_最低票价

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_递归(t *testing.T) {
	var days = []int{1, 4, 6, 7, 8, 20}
	var costs = []int{2, 7, 15}
	assert.Equal(t, mincostTickets(days, costs), 11)

}

func Test_递归1(t *testing.T) {

	var days = []int{1,2,3,4,6,8,9,10,13,14,16,17,19,21,24,26,27,28,29}
	var costs = []int{3,14,50}
	assert.Equal(t, mincostTickets(days, costs), 50)
	fmt.Println(dp,cheap7th,cheap30th)

}
