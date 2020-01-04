package main

/*
#include <stdio.h>

void hi() {
    printf("hello world!\n");
	int a=1;
	printf("12123%d",a);
}
*/
import "C" //这里可看作封装的伪包C, 这条语句要紧挨着上面的注释块，不可在它俩之间间隔空行！

func main() {
	C.hi()
	print("Hi, vim-go")
}
