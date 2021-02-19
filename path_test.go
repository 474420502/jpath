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
	jexp1 := "/ad/bs[@a='1']"
	content := headHandler([]rune(jexp1))
	var i = 1
	o := getTarget(content, &i)
	if o.GetString(content) != "ad" {
		t.Error()
	}

	jexp2 := "/bs[@a='1']"
	content2 := headHandler([]rune(jexp2))
	var i2 = 1
	o2 := getTarget(content2, &i2)
	if o2.GetString(content2) != "bs" {
		t.Error()
	}

}
