/*
 * @Author: hsm
 * @Date: 2021-02-24 14:42:51
 * @Last Modified by: mikey.zhaopeng
 * @Last Modified time: 2021-02-24 14:44:42
 */
package jpath

import "testing"

func TestHeadHandler(t *testing.T) {

	var headHandlerStr = func(src string) string {
		return string(headHandler([]rune(src)))
	}

	if headHandlerStr("/123") != "/123 " {
		t.Error("error")
	}

	if headHandlerStr("//123") != "//123 " {
		t.Error("error")
	}

	if headHandlerStr("dasda") != "/dasda " {
		t.Error("error")
	}

	if headHandlerStr("  /123") != "/123 " {
		t.Error("error")
	}

	if headHandlerStr("  [@a='1']") != "/.[@a='1'] " {
		t.Error("error")
	}

	if headHandlerStr("  //123") != "//123 " {
		t.Error("error")
	}
}

func TestTarget(t *testing.T) {
	jexp1 := "/ad/bs[]"
	content := headHandler([]rune(jexp1))
	var i = 1
	o, ty1 := getTarget(content, &i)
	if o.GetString(content) != "ad" {
		t.Error()
	}

	if ty1 != nNextPath {
		t.Error(ty1)
	}

	jexp2 := "/bs[]"
	content2 := headHandler([]rune(jexp2))
	var i2 = 1
	o2, ty2 := getTarget(content2, &i2)
	if o2.GetString(content2) != "bs" {
		t.Error()
	}
	if ty2 != nIndexes {
		t.Error(ty2)
	}
}

func TestIndexes(t *testing.T) {

	for _, jexp1 := range []string{"/bs[]", "/bs[:]"} {

		content1 := headHandler([]rune(jexp1))
		var i1 = 1
		o1, ty1 := getTarget(content1, &i1)
		if o1.GetString(content1) != "bs" {
			t.Error()
		}

		if ty1 != nIndexes {
			t.Error(ty1)
		}

		idxs1, nt1 := getIndexes(content1, &i1)
		if nt1 != 0 {
			t.Error(nt1)
		}

		if idxs1.Start != 0 || idxs1.End != -1 {
			t.Error(idxs1)
		}
	}

	jexp1 := "/bs[1:]"
	content1 := headHandler([]rune(jexp1))
	var i1 = 1
	o1, ty1 := getTarget(content1, &i1)
	if o1.GetString(content1) != "bs" {
		t.Error()
	}

	if ty1 != nIndexes {
		t.Error(ty1)
	}

	idxs1, nt1 := getIndexes(content1, &i1)
	if nt1 != 0 {
		t.Error(nt1)
	}

	if idxs1.Start != 1 || idxs1.End != -1 {
		t.Error(idxs1)
	}

	jexp1 = "/bs[1:11]"
	content1 = headHandler([]rune(jexp1))
	i1 = 1
	o1, ty1 = getTarget(content1, &i1)
	if o1.GetString(content1) != "bs" {
		t.Error()
	}

	if ty1 != nIndexes {
		t.Error(ty1)
	}

	idxs1, nt1 = getIndexes(content1, &i1)
	if nt1 != 0 {
		t.Error(nt1)
	}

	if idxs1.Start != 1 || idxs1.End != 11 {
		t.Error(idxs1)
	}

	jexp1 = "/bs[1]"
	content1 = headHandler([]rune(jexp1))
	i1 = 1
	o1, ty1 = getTarget(content1, &i1)
	if o1.GetString(content1) != "bs" {
		t.Error()
	}

	if ty1 != nIndexes {
		t.Error(ty1)
	}

	idxs1, nt1 = getIndexes(content1, &i1)
	if nt1 != 0 {
		t.Error(nt1)
	}

	if idxs1.Start != 1 {
		t.Error(idxs1)
	}
}

func TestDepth(t *testing.T) {
	var jexp string
	var content []rune
	var i = 1
	var ty nexttype

	jexp = "/bs<>[1:11]"
	content = headHandler([]rune(jexp))

	o, ty := getTarget(content, &i)
	if o.GetString(content) != "bs" {
		t.Error()
	}

	if ty != nDepth {
		t.Error(ty)
	}

	if depth := getDepth(content, &i); depth != -1 {
		t.Error(depth)
	}

	i++
	idxs, nt := getIndexes(content, &i)
	if nt != 0 {
		t.Error(nt)
	}

	if idxs.Start != 1 && idxs.End != 11 {
		t.Error(idxs)
	}
}

func TestCondition(t *testing.T) {
	var jexp string
	var content []rune
	var i = 1
	var ty nexttype

	jexp = "/bs<>[1:11]()"
	content = headHandler([]rune(jexp))

	o, ty := getTarget(content, &i)
	if o.GetString(content) != "bs" {
		t.Error()
	}

	if ty != nDepth {
		t.Error(ty)
	}

	if depth := getDepth(content, &i); depth != -1 {
		t.Error(depth)
	}

	i++
	idxs, nt := getIndexes(content, &i)
	if nt != 0 {
		t.Error(nt)
	}
	if idxs.Start != 1 && idxs.End != 11 {
		t.Error(idxs)
	}

	i++
	if cond := getCondition(content, &i); cond != "" {
		t.Error(cond)
	}

	i = 1
	jexp = "/bs<>[1:11](se1)"
	content = headHandler([]rune(jexp))

	o, ty = getTarget(content, &i)
	if o.GetString(content) != "bs" {
		t.Error()
	}

	if ty != nDepth {
		t.Error(ty)
	}

	if depth := getDepth(content, &i); depth != -1 {
		t.Error(depth)
	}

	i++
	idxs, nt = getIndexes(content, &i)
	if nt != 0 {
		t.Error(nt)
	}

	if idxs.Start != 1 && idxs.End != 11 {
		t.Error(idxs)
	}

	i++
	if cond := getCondition(content, &i); cond != "se1" {
		t.Error(cond)
	}
}
