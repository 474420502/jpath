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
	t.Error(ty1)

	jexp2 := "/bs[]"
	content2 := headHandler([]rune(jexp2))
	var i2 = 1
	o2, ty2 := getTarget(content2, &i2)
	if o2.GetString(content2) != "bs" {
		t.Error()
	}
	t.Error(ty2)
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

	if idxs1.Start != 11 {
		t.Error(idxs1)
	}
}
