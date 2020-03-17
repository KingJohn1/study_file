// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package websever

import "errors"

type Args struct {
 A,B int
}

type Quotient struct{
	Quo,Rem int
}

var Addr ="127.0.0.1:333"
var Port=":333"

type Arith int
//Multiply
func (a *Arith)Multiply( args *Args,reply *int)error{
	*reply=args.A*args.B
	return nil
}

func (a *Arith)Divide(args *Args,quo*Quotient)error{
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}