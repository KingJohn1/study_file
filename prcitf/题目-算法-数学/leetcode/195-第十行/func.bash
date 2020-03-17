#!/usr/bin/env bash

#本题如果先head再tail的话会导致【假如文件不到10行依然会显示某一个错误的行】。所以先tail -n +10选出从第10行开始的所有行（保证了假如文件不到10行时，不应该有任何输出），然后再用head -1输出筛选出的内容的第1行。
#
#tail -n +10 file.txt | head -1
#
#作者：gao-si-huang-bu
#链接：https://leetcode-cn.com/problems/tenth-line/solution/tailtop-by-gao-si-huang-bu/
#来源：力扣（LeetCode）
#著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
tail -n +10 file.txt | head -1

sed -n '10p' file.txt
awk 'NR==10' file.txt
head -10 file.txt | tail -1

cat `file.txt`|echo

demoFun(){
    echo "这是我的第一个 shell 函数!"
}
echo "-----函数开始执行-----"
demoFun
echo "-----函数执行完毕-----"