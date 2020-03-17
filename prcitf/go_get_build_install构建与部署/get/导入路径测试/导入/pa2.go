// Copyright (c) Huawei Technologies Co., Ltd. 2012-2019. All rights reserved.
package 导入

import aa "prcitf/go_get_build_install构建与部署/get/导入路径测试/被导入"

//todo  修改package的import，使用import的路径无法成功导入，是不是理解有误；
func a(){
	b:=&aa.A{

	}
}
