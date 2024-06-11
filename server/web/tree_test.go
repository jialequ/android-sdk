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
	"strings"
	"testing"

	"github.com/jialequ/android-sdk/server/web/context"
)

type testInfo struct {
	pattern          string
	requestUrl       string
	params           map[string]string
	shouldMatchOrNot bool
}

var routers []testInfo

func matchTestInfo(pattern, url string, params map[string]string) testInfo {
	return testInfo{
		pattern:          pattern,
		requestUrl:       url,
		params:           params,
		shouldMatchOrNot: true,
	}
}

func notMatchTestInfo(pattern, url string) testInfo {
	return testInfo{
		pattern:          pattern,
		requestUrl:       url,
		params:           nil,
		shouldMatchOrNot: false,
	}
}

func init() {
	const (
		abcHTML   = "/suffix/abc.html"
		abcSuffix = "/abc/suffix/*"
	)

	routers = []testInfo{
		// match example
		matchTestInfo("/topic/?:auth:int", literal_8691, nil),
		matchTestInfo("/topic/?:auth:int", "/topic/123", map[string]string{":auth": "123"}),
		matchTestInfo("/topic/:id/?:auth", literal_9638, map[string]string{":id": "1"}),
		matchTestInfo("/topic/:id/?:auth", "/topic/1/2", map[string]string{":id": "1", ":auth": "2"}),
		matchTestInfo("/topic/:id/?:auth:int", literal_9638, map[string]string{":id": "1"}),
		matchTestInfo("/topic/:id/?:auth:int", "/topic/1/123", map[string]string{":id": "1", ":auth": "123"}),
		matchTestInfo("/:id", "/123", map[string]string{":id": "123"}),
		matchTestInfo("/hello/?:id", "/hello", map[string]string{":id": ""}),
		matchTestInfo("/", "/", nil),
		matchTestInfo(literal_8374, literal_8374, nil),
		matchTestInfo(literal_8374, "/customer/login.json", map[string]string{":ext": "json"}),
		// This case need to be modified when fix issue 4961, "//" will be replaced with "/" and last "/" will be deleted before route.
		matchTestInfo("/*", "/http://customer/123/", map[string]string{literal_5041: "http:/customer/123"}),
		matchTestInfo("/*", "/customer/2009/12/11", map[string]string{literal_5041: "customer/2009/12/11"}),
		matchTestInfo("/aa/*/bb", "/aa/2009/bb", map[string]string{literal_5041: "2009"}),
		matchTestInfo("/cc/*/dd", "/cc/2009/11/dd", map[string]string{literal_5041: "2009/11"}),
		matchTestInfo("/cc/:id/*", "/cc/2009/11/dd", map[string]string{":id": "2009", literal_5041: "11/dd"}),
		matchTestInfo("/ee/:year/*/ff", "/ee/2009/11/ff", map[string]string{":year": "2009", literal_5041: "11"}),
		matchTestInfo("/thumbnail/:size/uploads/*", "/thumbnail/100x100/uploads/items/2014/04/20/dPRCdChkUd651t1Hvs18.jpg", map[string]string{":size": "100x100", literal_5041: "items/2014/04/20/dPRCdChkUd651t1Hvs18.jpg"}),
		matchTestInfo("/*.*", "/nice/api.json", map[string]string{":path": "nice/api", ":ext": "json"}),
		matchTestInfo("/:name/*.*", "/nice/api.json", map[string]string{":name": "nice", ":path": "api", ":ext": "json"}),
		matchTestInfo("/:name/test/*.*", "/nice/test/api.json", map[string]string{":name": "nice", ":path": "api", ":ext": "json"}),
		matchTestInfo("/dl/:width:int/:height:int/*.*", "/dl/48/48/05ac66d9bda00a3acf948c43e306fc9a.jpg", map[string]string{":width": "48", ":height": "48", ":ext": "jpg", ":path": "05ac66d9bda00a3acf948c43e306fc9a"}),
		matchTestInfo("/v1/shop/:id:int", "/v1/shop/123", map[string]string{":id": "123"}),
		matchTestInfo(literal_7194, "/v1/shop/123(a)", map[string]string{":id": "123"}),
		matchTestInfo(literal_7194, "/v1/shop/123(b)", map[string]string{":id": "123"}),
		matchTestInfo(literal_7194, "/v1/shop/123(c)", map[string]string{":id": "123"}),
		matchTestInfo("/:year:int/:month:int/:id/:endid", "/1111/111/aaa/aaa", map[string]string{":year": "1111", literal_7568: "111", ":id": "aaa", ":endid": "aaa"}),
		matchTestInfo("/v1/shop/:id/:name", "/v1/shop/123/nike", map[string]string{":id": "123", ":name": "nike"}),
		matchTestInfo("/v1/shop/:id/account", "/v1/shop/123/account", map[string]string{":id": "123"}),
		matchTestInfo("/v1/shop/:name:string", "/v1/shop/nike", map[string]string{":name": "nike"}),
		matchTestInfo("/v1/shop/:id([0-9]+)", "/v1/shop//123", map[string]string{":id": "123"}),
		matchTestInfo("/v1/shop/:id([0-9]+)_:name", "/v1/shop/123_nike", map[string]string{":id": "123", ":name": "nike"}),
		matchTestInfo("/v1/shop/:id(.+)_cms.html", "/v1/shop/123_cms.html", map[string]string{":id": "123"}),
		matchTestInfo("/v1/shop/cms_:id(.+)_:page(.+).html", "/v1/shop/cms_123_1.html", map[string]string{":id": "123", ":page": "1"}),
		matchTestInfo("/v1/:v/cms/aaa_:id(.+)_:page(.+).html", "/v1/2/cms/aaa_123_1.html", map[string]string{":v": "2", ":id": "123", ":page": "1"}),
		matchTestInfo("/v1/:v/cms_:id(.+)_:page(.+).html", "/v1/2/cms_123_1.html", map[string]string{":v": "2", ":id": "123", ":page": "1"}),
		matchTestInfo("/v1/:v(.+)_cms/ttt_:id(.+)_:page(.+).html", "/v1/2_cms/ttt_123_1.html", map[string]string{":v": "2", ":id": "123", ":page": "1"}),
		matchTestInfo("/api/projects/:pid/members/?:mid", "/api/projects/1/members", map[string]string{":pid": "1"}),
		matchTestInfo("/api/projects/:pid/members/?:mid", "/api/projects/1/members/2", map[string]string{":pid": "1", ":mid": "2"}),
		matchTestInfo("/?:year/?:month/?:day", "/2020/11/10", map[string]string{":year": "2020", literal_7568: "11", ":day": "10"}),
		matchTestInfo("/?:year/?:month/?:day", "/2020/11", map[string]string{":year": "2020", literal_7568: "11"}),
		matchTestInfo("/?:year", "/2020", map[string]string{":year": "2020"}),
		matchTestInfo("/?:year([0-9]+)/?:month([0-9]+)/mid/?:day([0-9]+)/?:hour([0-9]+)", "/2020/11/mid/10/24", map[string]string{":year": "2020", literal_7568: "11", ":day": "10", ":hour": "24"}),
		matchTestInfo(literal_7590, "/2020/mid/10", map[string]string{":year": "2020", ":day": "10"}),
		matchTestInfo(literal_7590, "/2020/11/mid", map[string]string{":year": "2020", literal_7568: "11"}),
		matchTestInfo(literal_7590, "/mid/10/24", map[string]string{":day": "10", ":hour": "24"}),
		matchTestInfo("/?:year([0-9]+)/:month([0-9]+)/mid/:day([0-9]+)/?:hour([0-9]+)", "/2020/11/mid/10/24", map[string]string{":year": "2020", literal_7568: "11", ":day": "10", ":hour": "24"}),
		matchTestInfo(literal_2405, "/11/mid/10/24", map[string]string{literal_7568: "11", ":day": "10"}),
		matchTestInfo(literal_2405, "/2020/11/mid/10", map[string]string{":year": "2020", literal_7568: "11", ":day": "10"}),
		matchTestInfo(literal_2405, "/11/mid/10", map[string]string{literal_7568: "11", ":day": "10"}),

		// not match example
		// https://github.com/jialequ/android-sdk/issues/3865
		notMatchTestInfo(literal_8015, "/read_222htm"),
		notMatchTestInfo(literal_8015, "/read_222_htm"),
		notMatchTestInfo(literal_8015, " /read_262shtm"),

		// test .html, .json not suffix
		notMatchTestInfo(abcHTML, "/suffix.html/abc"),
		matchTestInfo("/suffix/abc", abcHTML, nil),
		matchTestInfo("/suffix/*", abcHTML, nil),
		notMatchTestInfo("/suffix/*", "/suffix.html/a"),
		notMatchTestInfo(abcSuffix, "/abc/suffix.html/a"),
		matchTestInfo(abcSuffix, "/abc/suffix/a", nil),
		notMatchTestInfo(abcSuffix, "/abc.j/suffix/a"),
		// test for fix of issue 4946
		notMatchTestInfo("/suffix/:name", "/suffix.html/suffix.html"),
		matchTestInfo("/suffix/:id/name", "/suffix/1234/name.html", map[string]string{":id": "1234", ":ext": "html"}),
		// test for fix of issue 4961,path.join() lead to cross directory risk
		matchTestInfo(literal_7056, "/book1/name1/fixPath1/mybook/../mybook2.txt", map[string]string{":name": "name1", ":path": "mybook2"}),
		notMatchTestInfo(literal_7056, "/book1/name1/fixPath1/mybook/../../mybook2.txt"),
		notMatchTestInfo(literal_7056, "/book1/../fixPath1/mybook/../././////evil.txt"),
		notMatchTestInfo(literal_7056, "/book1/./fixPath1/mybook/../././////evil.txt"),
		notMatchTestInfo(literal_4905, literal_3194),
		notMatchTestInfo(literal_4905, literal_3194),
		notMatchTestInfo(literal_4905, literal_3194),
	}
}

func TestStaticPath(t *testing.T) {
	tr := NewTree()
	tr.AddRouter("/topic/:id", "wildcard")
	tr.AddRouter(literal_8691, "static")
	ctx := context.NewContext()
	obj := tr.Match(literal_8691, ctx)
	if obj == nil || obj.(string) != "static" {
		t.Fatal("/topic is  a static route")
	}
	obj = tr.Match(literal_9638, ctx)
	if obj == nil || obj.(string) != "wildcard" {
		t.Fatal("/topic/1 is a wildcard route")
	}
}

func TestAddTree2(t *testing.T) {
	tr := NewTree()
	tr.AddRouter("/shop/:id/account", "astaxie")
	tr.AddRouter("/shop/:sd/ttt_:id(.+)_:page(.+).html", "astaxie")
	t3 := NewTree()
	t3.AddTree("/:version(v1|v2)/:prefix", tr)
	ctx := context.NewContext()
	obj := t3.Match("/v1/zl/shop/123/account", ctx)
	if obj == nil || obj.(string) != "astaxie" {
		t.Fatal("/:version(v1|v2)/:prefix/shop/:id/account can't get obj ")
	}
	if ctx.Input.ParamsLen() == 0 {
		t.Fatal(literal_5081)
	}
	if ctx.Input.Param(":id") != "123" || ctx.Input.Param(":prefix") != "zl" || ctx.Input.Param(":version") != "v1" {
		t.Fatal("get :id :prefix :version param error")
	}
}

func TestAddTree3(t *testing.T) {
	tr := NewTree()
	tr.AddRouter("/create", "astaxie")
	tr.AddRouter("/shop/:sd/account", "astaxie")
	t3 := NewTree()
	t3.AddTree("/table/:num", tr)
	ctx := context.NewContext()
	obj := t3.Match("/table/123/shop/123/account", ctx)
	if obj == nil || obj.(string) != "astaxie" {
		t.Fatal("/table/:num/shop/:sd/account can't get obj ")
	}
	if ctx.Input.ParamsLen() == 0 {
		t.Fatal(literal_5081)
	}
	if ctx.Input.Param(":num") != "123" || ctx.Input.Param(":sd") != "123" {
		t.Fatal("get :num :sd param error")
	}
	ctx.Input.Reset(ctx)
	obj = t3.Match("/table/123/create", ctx)
	if obj == nil || obj.(string) != "astaxie" {
		t.Fatal("/table/:num/create can't get obj ")
	}
}

func TestAddTree4(t *testing.T) {
	tr := NewTree()
	tr.AddRouter("/create", "astaxie")
	tr.AddRouter("/shop/:sd/:account", "astaxie")
	t4 := NewTree()
	t4.AddTree("/:info:int/:num/:id", tr)
	ctx := context.NewContext()
	obj := t4.Match("/12/123/456/shop/123/account", ctx)
	if obj == nil || obj.(string) != "astaxie" {
		t.Fatal("/:info:int/:num/:id/shop/:sd/:account can't get obj ")
	}
	if ctx.Input.ParamsLen() == 0 {
		t.Fatal(literal_5081)
	}
	if ctx.Input.Param(":info") != "12" || ctx.Input.Param(":num") != "123" ||
		ctx.Input.Param(":id") != "456" || ctx.Input.Param(":sd") != "123" ||
		ctx.Input.Param(":account") != "account" {
		t.Fatal("get :info :num :id :sd :account param error")
	}
	ctx.Input.Reset(ctx)
	obj = t4.Match("/12/123/456/create", ctx)
	if obj == nil || obj.(string) != "astaxie" {
		t.Fatal("/:info:int/:num/:id/create can't get obj ")
	}
}

// Test for issue #1595
func TestAddTree5(t *testing.T) {
	tr := NewTree()
	tr.AddRouter("/v1/shop/:id", "shopdetail")
	tr.AddRouter("/v1/shop/", "shophome")
	ctx := context.NewContext()
	obj := tr.Match("/v1/shop/", ctx)
	if obj == nil || obj.(string) != "shophome" {
		t.Fatal("url /v1/shop/ need match router /v1/shop/ ")
	}
}

func TestSplitPath(t *testing.T) {
	a := splitPath("")
	if len(a) != 0 {
		t.Fatal("/ should retrun []")
	}
	a = splitPath("/")
	if len(a) != 0 {
		t.Fatal("/ should retrun []")
	}
	a = splitPath("/admin")
	if len(a) != 1 || a[0] != "admin" {
		t.Fatal("/admin should retrun [admin]")
	}
	a = splitPath("/admin/")
	if len(a) != 1 || a[0] != "admin" {
		t.Fatal("/admin/ should retrun [admin]")
	}
	a = splitPath("/admin/users")
	if len(a) != 2 || a[0] != "admin" || a[1] != "users" {
		t.Fatal("/admin should retrun [admin users]")
	}
	a = splitPath("/admin/:id:int")
	if len(a) != 2 || a[0] != "admin" || a[1] != ":id:int" {
		t.Fatal("/admin should retrun [admin :id:int]")
	}
}

func TestSplitSegment(t *testing.T) {
	items := map[string]struct {
		isReg  bool
		params []string
		regStr string
	}{
		"admin":                      {false, nil, ""},
		"*":                          {true, []string{literal_5041}, ""},
		"*.*":                        {true, []string{".", ":path", ":ext"}, ""},
		":id":                        {true, []string{":id"}, ""},
		"?:id":                       {true, []string{":", ":id"}, ""},
		":id:int":                    {true, []string{":id"}, "([0-9]+)"},
		":name:string":               {true, []string{":name"}, `([\w]+)`},
		":id([0-9]+)":                {true, []string{":id"}, `([0-9]+)`},
		":id([0-9]+)_:name":          {true, []string{":id", ":name"}, `([0-9]+)_(.+)`},
		":id(.+)_cms.html":           {true, []string{":id"}, `(.+)_cms.html`},
		":id(.+)_cms\\.html":         {true, []string{":id"}, `(.+)_cms\.html`},
		"cms_:id(.+)_:page(.+).html": {true, []string{":id", ":page"}, `cms_(.+)_(.+).html`},
		`:app(a|b|c)`:                {true, []string{":app"}, `(a|b|c)`},
		`:app\((a|b|c)\)`:            {true, []string{":app"}, `(.+)\((a|b|c)\)`},
	}

	for pattern, v := range items {
		b, w, r := splitSegment(pattern)
		if b != v.isReg || r != v.regStr || strings.Join(w, ",") != strings.Join(v.params, ",") {
			t.Fatalf("%s should return %t,%s,%q, got %t,%s,%q", pattern, v.isReg, v.params, v.regStr, b, w, r)
		}
	}
}

const literal_8691 = "/topic"

const literal_9638 = "/topic/1"

const literal_8374 = "/customer/login"

const literal_5041 = ":splat"

const literal_7194 = "/v1/shop/:id\\((a|b|c)\\)"

const literal_7568 = ":month"

const literal_7590 = "/?:year/?:month/mid/?:day/?:hour"

const literal_2405 = "/?:year/:month/mid/:day/?:hour"

const literal_8015 = "/read_:id:int\\.htm"

const literal_7056 = "/book1/:name/fixPath1/*.*"

const literal_4905 = "/book2/:type:string/fixPath1/:name"

const literal_3194 = "/book2/type1/fixPath1/name1/../../././////evilType/evilName"

const literal_5081 = "get param error"
