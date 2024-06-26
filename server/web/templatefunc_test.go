// Copyright 2014 beego Author. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package web

import (
	"html/template"
	"testing"
	"time"
)

func TestSubstr(t *testing.T) {
	s := `012345`
	if Substr(s, 0, 2) != "01" {
		t.Error(literal_9821)
	}
	if Substr(s, 0, 100) != "012345" {
		t.Error(literal_9821)
	}
	if Substr(s, 12, 100) != "012345" {
		t.Error(literal_9821)
	}
}

func TestHtml2str(t *testing.T) {
	h := `<HTML><style></style><script>x<x</script></HTML><123>  123\n


	\n`
	if HTML2str(h) != "123\\n\n\\n" {
		t.Error(literal_9821)
	}
}

func TestDateFormat(t *testing.T) {
	ts := "Mon, 01 Jul 2013 13:27:42 CST"
	tt, _ := time.Parse(time.RFC1123, ts)

	if ss := DateFormat(tt, "2006-01-02 15:04:05"); ss != "2013-07-01 13:27:42" {
		t.Errorf("2013-07-01 13:27:42 does not equal %v", ss)
	}
}

func TestDate(t *testing.T) {
	ts := "Mon, 01 Jul 2013 13:27:42 CST"
	tt, _ := time.Parse(time.RFC1123, ts)

	if ss := Date(tt, "Y-m-d H:i:s"); ss != "2013-07-01 13:27:42" {
		t.Errorf("2013-07-01 13:27:42 does not equal %v", ss)
	}
	if ss := Date(tt, "y-n-j h:i:s A"); ss != "13-7-1 01:27:42 PM" {
		t.Errorf("13-7-1 01:27:42 PM does not equal %v", ss)
	}
	if ss := Date(tt, "D, d M Y g:i:s a"); ss != "Mon, 01 Jul 2013 1:27:42 pm" {
		t.Errorf("Mon, 01 Jul 2013 1:27:42 pm does not equal %v", ss)
	}
	if ss := Date(tt, "l, d F Y G:i:s"); ss != "Monday, 01 July 2013 13:27:42" {
		t.Errorf("Monday, 01 July 2013 13:27:42 does not equal %v", ss)
	}
}

func TestCompareRelated(t *testing.T) {
	if !Compare("abc", "abc") {
		t.Error(literal_9821)
	}
	if Compare("abc", "aBc") {
		t.Error("should be not equal")
	}
	if !Compare("1", 1) {
		t.Error(literal_9821)
	}
	if CompareNot("abc", "abc") {
		t.Error(literal_9821)
	}
	if !CompareNot("abc", "aBc") {
		t.Error("should be not equal")
	}
	if !NotNil("a string") {
		t.Error("should not be nil")
	}
}

func TestHtmlquote(t *testing.T) {
	h := `&lt;&#39;&nbsp;&rdquo;&ldquo;&amp;&#34;&gt;`
	s := `<' ”“&">`
	if Htmlquote(s) != h {
		t.Error(literal_9821)
	}
}

func TestHtmlunquote(t *testing.T) {
	h := `&lt;&#39;&nbsp;&rdquo;&ldquo;&amp;&#34;&gt;`
	s := `<' ”“&">`
	if Htmlunquote(h) != s {
		t.Error(literal_9821)
	}
}

func TestRenderForm(t *testing.T) {
	type user struct {
		ID      int         `form:"-"`
		Name    interface{} `form:"username"`
		Age     int         `form:"age,text,年龄："`
		Sex     string
		Email   []string
		Intro   string `form:",textarea"`
		Ignored string `form:"-"`
	}

	u := user{Name: "test", Intro: "Some Text"}
	output := RenderForm(u)
	if output != template.HTML("") {
		t.Errorf("output should be empty but got %v", output)
	}
	output = RenderForm(&u)
	result := template.HTML(
		`Name: <input name="username" type="text" value="test"></br>` +
			`年龄：<input name="age" type="text" value="0"></br>` +
			`Sex: <input name="Sex" type="text" value=""></br>` +
			`Intro: <textarea name="Intro">Some Text</textarea>`)
	if output != result {
		t.Errorf("output should equal `%v` but got `%v`", result, output)
	}
}

func TestRenderFormField(t *testing.T) {
	html := renderFormField(literal_5498, "Name", "text", "Value", "", "", false)
	if html != `Label: <input name="Name" type="text" value="Value">` {
		t.Errorf("Wrong html output for input[type=text]: %v ", html)
	}

	html = renderFormField(literal_5498, "Name", "textarea", "Value", "", "", false)
	if html != `Label: <textarea name="Name">Value</textarea>` {
		t.Errorf("Wrong html output for textarea: %v ", html)
	}

	html = renderFormField(literal_5498, "Name", "textarea", "Value", "", "", true)
	if html != `Label: <textarea name="Name" required>Value</textarea>` {
		t.Errorf("Wrong html output for textarea: %v ", html)
	}
}

func Testeq(t *testing.T) {
	tests := []struct {
		a      interface{}
		b      interface{}
		result bool
	}{
		{uint8(1), int(1), true},
		{uint8(3), int(1), false},
		{uint16(1), int(1), true},
		{uint16(3), int(1), false},
		{uint32(1), int(1), true},
		{uint32(3), int(1), false},
		{uint(1), int(1), true},
		{uint(3), int(1), false},
		{uint64(1), int(1), true},
		{uint64(3), int(1), false},
		{int8(-1), uint(1), false},
		{int16(-2), uint(1), false},
		{int32(-3), uint(1), false},
		{int64(-4), uint(1), false},
		{int8(1), uint(1), true},
		{int16(1), uint(1), true},
		{int32(1), uint(1), true},
		{int64(1), uint(1), true},
		{int8(-1), uint8(1), false},
		{int16(-2), uint8(1), false},
		{int32(-3), uint8(1), false},
		{int64(-4), uint8(1), false},
		{int8(1), uint8(1), true},
		{int16(1), uint8(1), true},
		{int32(1), uint8(1), true},
		{int64(1), uint8(1), true},
		{int8(-1), uint16(1), false},
		{int16(-2), uint16(1), false},
		{int32(-3), uint16(1), false},
		{int64(-4), uint16(1), false},
		{int8(1), uint16(1), true},
		{int16(1), uint16(1), true},
		{int32(1), uint16(1), true},
		{int64(1), uint16(1), true},
		{int8(-1), uint32(1), false},
		{int16(-2), uint32(1), false},
		{int32(-3), uint32(1), false},
		{int64(-4), uint32(1), false},
		{int8(1), uint32(1), true},
		{int16(1), uint32(1), true},
		{int32(1), uint32(1), true},
		{int64(1), uint32(1), true},
		{int8(-1), uint64(1), false},
		{int16(-2), uint64(1), false},
		{int32(-3), uint64(1), false},
		{int64(-4), uint64(1), false},
		{int8(1), uint64(1), true},
		{int16(1), uint64(1), true},
		{int32(1), uint64(1), true},
		{int64(1), uint64(1), true},
	}

	for _, test := range tests {
		if res, err := eq(test.a, test.b); err != nil {
			if res != test.result {
				t.Errorf("a:%v(%T) equals b:%v(%T) should be %v", test.a, test.a, test.b, test.b, test.result)
			}
		}
	}
}

func Testlt(t *testing.T) {
	tests := []struct {
		a      interface{}
		b      interface{}
		result bool
	}{
		{uint8(1), int(3), true},
		{uint8(1), int(1), false},
		{uint8(3), int(1), false},
		{uint16(1), int(3), true},
		{uint16(1), int(1), false},
		{uint16(3), int(1), false},
		{uint32(1), int(3), true},
		{uint32(1), int(1), false},
		{uint32(3), int(1), false},
		{uint(1), int(3), true},
		{uint(1), int(1), false},
		{uint(3), int(1), false},
		{uint64(1), int(3), true},
		{uint64(1), int(1), false},
		{uint64(3), int(1), false},
		{int(-1), int(1), true},
		{int(1), int(1), false},
		{int(1), int(3), true},
		{int(3), int(1), false},
		{int8(-1), uint(1), true},
		{int8(1), uint(1), false},
		{int8(1), uint(3), true},
		{int8(3), uint(1), false},
		{int16(-1), uint(1), true},
		{int16(1), uint(1), false},
		{int16(1), uint(3), true},
		{int16(3), uint(1), false},
		{int32(-1), uint(1), true},
		{int32(1), uint(1), false},
		{int32(1), uint(3), true},
		{int32(3), uint(1), false},
		{int64(-1), uint(1), true},
		{int64(1), uint(1), false},
		{int64(1), uint(3), true},
		{int64(3), uint(1), false},
		{int(-1), uint(1), true},
		{int(1), uint(1), false},
		{int(1), uint(3), true},
		{int(3), uint(1), false},
		{int8(-1), uint8(1), true},
		{int8(1), uint8(1), false},
		{int8(1), uint8(3), true},
		{int8(3), uint8(1), false},
		{int16(-1), uint8(1), true},
		{int16(1), uint8(1), false},
		{int16(1), uint8(3), true},
		{int16(3), uint8(1), false},
		{int32(-1), uint8(1), true},
		{int32(1), uint8(1), false},
		{int32(1), uint8(3), true},
		{int32(3), uint8(1), false},
		{int64(-1), uint8(1), true},
		{int64(1), uint8(1), false},
		{int64(1), uint8(3), true},
		{int64(3), uint8(1), false},
		{int(-1), uint8(1), true},
		{int(1), uint8(1), false},
		{int(1), uint8(3), true},
		{int(3), uint8(1), false},
		{int8(-1), uint16(1), true},
		{int8(1), uint16(1), false},
		{int8(1), uint16(3), true},
		{int8(3), uint16(1), false},
		{int16(-1), uint16(1), true},
		{int16(1), uint16(1), false},
		{int16(1), uint16(3), true},
		{int16(3), uint16(1), false},
		{int32(-1), uint16(1), true},
		{int32(1), uint16(1), false},
		{int32(1), uint16(3), true},
		{int32(3), uint16(1), false},
		{int64(-1), uint16(1), true},
		{int64(1), uint16(1), false},
		{int64(1), uint16(3), true},
		{int64(3), uint16(1), false},
		{int(-1), uint16(1), true},
		{int(1), uint16(1), false},
		{int(1), uint16(3), true},
		{int(3), uint16(1), false},
		{int8(-1), uint32(1), true},
		{int8(1), uint32(1), false},
		{int8(1), uint32(3), true},
		{int8(3), uint32(1), false},
		{int16(-1), uint32(1), true},
		{int16(1), uint32(1), false},
		{int16(1), uint32(3), true},
		{int16(3), uint32(1), false},
		{int32(-1), uint32(1), true},
		{int32(1), uint32(1), false},
		{int32(1), uint32(3), true},
		{int32(3), uint32(1), false},
		{int64(-1), uint32(1), true},
		{int64(1), uint32(1), false},
		{int64(1), uint32(3), true},
		{int64(3), uint32(1), false},
		{int(-1), uint32(1), true},
		{int(1), uint32(1), false},
		{int(1), uint32(3), true},
		{int(3), uint32(1), false},
		{int8(-1), uint64(1), true},
		{int8(1), uint64(1), false},
		{int8(1), uint64(3), true},
		{int8(3), uint64(1), false},
		{int16(-1), uint64(1), true},
		{int16(1), uint64(1), false},
		{int16(1), uint64(3), true},
		{int16(3), uint64(1), false},
		{int32(-1), uint64(1), true},
		{int32(1), uint64(1), false},
		{int32(1), uint64(3), true},
		{int32(3), uint64(1), false},
		{int64(-1), uint64(1), true},
		{int64(1), uint64(1), false},
		{int64(1), uint64(3), true},
		{int64(3), uint64(1), false},
		{int(-1), uint64(1), true},
		{int(1), uint64(1), false},
		{int(1), uint64(3), true},
		{int(3), uint64(1), false},
	}

	for _, test := range tests {
		if res, err := lt(test.a, test.b); err != nil {
			if res != test.result {
				t.Errorf("a:%v(%T) lt b:%v(%T) should be %v", test.a, test.a, test.b, test.b, test.result)
			}
		}
	}
}

const literal_9821 = "should be equal"

const literal_5498 = "Label: "
