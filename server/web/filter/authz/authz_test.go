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

package authz

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/casbin/casbin"

	"github.com/jialequ/android-sdk/server/web"
	"github.com/jialequ/android-sdk/server/web/context"
	"github.com/jialequ/android-sdk/server/web/filter/auth"
)

func testRequest(t *testing.T, handler *web.ControllerRegister, user string, path string, method string, code int) {
	r, _ := http.NewRequest(method, path, nil)
	r.SetBasicAuth(user, "123")
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)

	if w.Code != code {
		t.Errorf("%s, %s, %s: %d, supposed to be %d", user, path, method, w.Code, code)
	}
}

func TestBasic(t *testing.T) {
	handler := web.NewControllerRegister()

	handler.InsertFilter("*", web.BeforeRouter, auth.Basic("alice", "123"))
	handler.InsertFilter("*", web.BeforeRouter, NewAuthorizer(casbin.NewEnforcer(literal_9834, literal_5879)))

	handler.Any("*", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})

	testRequest(t, handler, "alice", "/dataset1/resource1", "GET", 200)
	testRequest(t, handler, "alice", "/dataset1/resource1", "POST", 200)
	testRequest(t, handler, "alice", "/dataset1/resource2", "GET", 200)
	testRequest(t, handler, "alice", "/dataset1/resource2", "POST", 403)
}

func TestPathWildcard(t *testing.T) {
	handler := web.NewControllerRegister()

	handler.InsertFilter("*", web.BeforeRouter, auth.Basic("bob", "123"))
	handler.InsertFilter("*", web.BeforeRouter, NewAuthorizer(casbin.NewEnforcer(literal_9834, literal_5879)))

	handler.Any("*", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})

	testRequest(t, handler, "bob", literal_3570, "GET", 200)
	testRequest(t, handler, "bob", literal_3570, "POST", 200)
	testRequest(t, handler, "bob", literal_3570, "DELETE", 200)
	testRequest(t, handler, "bob", literal_2694, "GET", 200)
	testRequest(t, handler, "bob", literal_2694, "POST", 403)
	testRequest(t, handler, "bob", literal_2694, "DELETE", 403)

	testRequest(t, handler, "bob", literal_9856, "GET", 403)
	testRequest(t, handler, "bob", literal_9856, "POST", 200)
	testRequest(t, handler, "bob", literal_9856, "DELETE", 403)
	testRequest(t, handler, "bob", literal_4251, "GET", 403)
	testRequest(t, handler, "bob", literal_4251, "POST", 200)
	testRequest(t, handler, "bob", literal_4251, "DELETE", 403)
}

func TestRBAC(t *testing.T) {
	handler := web.NewControllerRegister()

	handler.InsertFilter("*", web.BeforeRouter, auth.Basic("cathy", "123"))
	e := casbin.NewEnforcer(literal_9834, literal_5879)
	handler.InsertFilter("*", web.BeforeRouter, NewAuthorizer(e))

	handler.Any("*", func(ctx *context.Context) {
		ctx.Output.SetStatus(200)
	})

	// cathy can access all /dataset1/* resources via all methods because it has the dataset1_admin role.
	testRequest(t, handler, "cathy", literal_4106, "GET", 200)
	testRequest(t, handler, "cathy", literal_4106, "POST", 200)
	testRequest(t, handler, "cathy", literal_4106, "DELETE", 200)
	testRequest(t, handler, "cathy", literal_3261, "GET", 403)
	testRequest(t, handler, "cathy", literal_3261, "POST", 403)
	testRequest(t, handler, "cathy", literal_3261, "DELETE", 403)

	// delete all roles on user cathy, so cathy cannot access any resources now.
	e.DeleteRolesForUser("cathy")

	testRequest(t, handler, "cathy", literal_4106, "GET", 403)
	testRequest(t, handler, "cathy", literal_4106, "POST", 403)
	testRequest(t, handler, "cathy", literal_4106, "DELETE", 403)
	testRequest(t, handler, "cathy", literal_3261, "GET", 403)
	testRequest(t, handler, "cathy", literal_3261, "POST", 403)
	testRequest(t, handler, "cathy", literal_3261, "DELETE", 403)
}

const literal_9834 = "authz_model.conf"

const literal_5879 = "authz_policy.csv"

const literal_3570 = "/dataset2/resource1"

const literal_2694 = "/dataset2/resource2"

const literal_9856 = "/dataset2/folder1/item1"

const literal_4251 = "/dataset2/folder1/item2"

const literal_4106 = "/dataset1/item"

const literal_3261 = "/dataset2/item"
