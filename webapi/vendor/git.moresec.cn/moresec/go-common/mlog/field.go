/*
@Time : 2020-10-27 10:40
@Author : gaoxl@moresec.cn
@Description:
@Software: GoLand
*/
package mlog

//模块名
func FieldMod(modName string) Field {
	return String("module", modName)
}

//服务地址
func FieldAddr(addr string) Field {
	return String("addr", addr)
}
