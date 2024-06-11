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
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

//
// The unregroute_test.go contains tests for the unregister route
// functionality, that allows overriding route paths in children project
// that embed parent routers.
//

const (
	contentRootOriginal      = "ok-original-root"
	contentLevel1Original    = "ok-original-level1"
	contentLevel2Original    = "ok-original-level2"
	contentRootReplacement   = "ok-replacement-root"
	contentLevel1Replacement = "ok-replacement-level1"
	contentLevel2Replacement = "ok-replacement-level2"
)

// TestPreUnregController will supply content for the original routes,
// before unregistration
type TestPreUnregController struct {
	Controller
}

func (tc *TestPreUnregController) GetFixedRoot() {
	tc.Ctx.Output.Body([]byte(contentRootOriginal))
}

func (tc *TestPreUnregController) GetFixedLevel1() {
	tc.Ctx.Output.Body([]byte(contentLevel1Original))
}

func (tc *TestPreUnregController) GetFixedLevel2() {
	tc.Ctx.Output.Body([]byte(contentLevel2Original))
}

// TestPostUnregController will supply content for the overriding routes,
// after the original ones are unregistered.
type TestPostUnregController struct {
	Controller
}

func (tc *TestPostUnregController) GetFixedRoot() {
	tc.Ctx.Output.Body([]byte(contentRootReplacement))
}

func (tc *TestPostUnregController) GetFixedLevel1() {
	tc.Ctx.Output.Body([]byte(contentLevel1Replacement))
}

func (tc *TestPostUnregController) GetFixedLevel2() {
	tc.Ctx.Output.Body([]byte(contentLevel2Replacement))
}

// TestUnregisterFixedRouteRoot replaces just the root fixed route path.
// In this case, for a path like literal_5074 or literal_4926, those actions
// should remain intact, and continue to serve the original content.
func TestUnregisterFixedRouteRoot(t *testing.T) {
	method := "GET"

	handler := NewControllerRegister()
	handler.Add("/", &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_2806))
	handler.Add(literal_4926, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_4678))
	handler.Add(literal_5074, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_0637))

	// Test original root
	testHelperFnContentCheck(t, handler, "Test original root",
		method, "/", contentRootOriginal)

	// Test original level 1
	testHelperFnContentCheck(t, handler, "Test original level 1",
		method, literal_4926, contentLevel1Original)

	// Test original level 2
	testHelperFnContentCheck(t, handler, "Test original level 2",
		method, literal_5074, contentLevel2Original)

	// Remove only the root path
	findAndRemoveSingleTree(handler.routers[method])

	// Replace the root path TestPreUnregController action with the action from
	// TestPostUnregController
	handler.Add("/", &TestPostUnregController{}, WithRouterMethods(&TestPostUnregController{}, literal_2806))

	// Test replacement root (expect change)
	testHelperFnContentCheck(t, handler, "Test replacement root (expect change)", method, "/", contentRootReplacement)

	// Test level 1 (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test level 1 (expect no change from the original)", method, literal_4926, contentLevel1Original)

	// Test level 2 (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test level 2 (expect no change from the original)", method, literal_5074, contentLevel2Original)
}

// TestUnregisterFixedRouteLevel1 replaces just the literal_4926 fixed route path.
// In this case, for a path like literal_5074 or "/", those actions
// should remain intact, and continue to serve the original content.
func TestUnregisterFixedRouteLevel1(t *testing.T) {
	method := "GET"

	handler := NewControllerRegister()
	handler.Add("/", &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_2806))
	handler.Add(literal_4926, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_4678))
	handler.Add(literal_5074, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_0637))

	// Test original root
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original root",
		method, "/", contentRootOriginal)

	// Test original level 1
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original level 1",
		method, literal_4926, contentLevel1Original)

	// Test original level 2
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original level 2",
		method, literal_5074, contentLevel2Original)

	// Remove only the level1 path
	subPaths := splitPath(literal_4926)
	if handler.routers[method].prefix == strings.Trim(literal_4926, "/ ") {
		findAndRemoveSingleTree(handler.routers[method])
	} else {
		findAndRemoveTree(subPaths, handler.routers[method], method)
	}

	// Replace the "level1" path TestPreUnregController action with the action from
	// TestPostUnregController
	handler.Add(literal_4926, &TestPostUnregController{}, WithRouterMethods(&TestPostUnregController{}, literal_4678))

	// Test replacement root (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test replacement root (expect no change from the original)", method, "/", contentRootOriginal)

	// Test level 1 (expect change)
	testHelperFnContentCheck(t, handler, "Test level 1 (expect change)", method, literal_4926, contentLevel1Replacement)

	// Test level 2 (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test level 2 (expect no change from the original)", method, literal_5074, contentLevel2Original)
}

// TestUnregisterFixedRouteLevel2 unregisters just the literal_5074 fixed
// route path. In this case, for a path like literal_4926 or "/", those actions
// should remain intact, and continue to serve the original content.
func TestUnregisterFixedRouteLevel2(t *testing.T) {
	method := "GET"

	handler := NewControllerRegister()
	handler.Add("/", &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_2806))
	handler.Add(literal_4926, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_4678))
	handler.Add(literal_5074, &TestPreUnregController{}, WithRouterMethods(&TestPreUnregController{}, literal_0637))

	// Test original root
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original root",
		method, "/", contentRootOriginal)

	// Test original level 1
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original level 1",
		method, literal_4926, contentLevel1Original)

	// Test original level 2
	testHelperFnContentCheck(t, handler,
		"TestUnregisterFixedRouteLevel1.Test original level 2",
		method, literal_5074, contentLevel2Original)

	// Remove only the level2 path
	subPaths := splitPath(literal_5074)
	if handler.routers[method].prefix == strings.Trim(literal_5074, "/ ") {
		findAndRemoveSingleTree(handler.routers[method])
	} else {
		findAndRemoveTree(subPaths, handler.routers[method], method)
	}

	// Replace the literal_5074 path TestPreUnregController action with the action from
	// TestPostUnregController
	handler.Add(literal_5074, &TestPostUnregController{}, WithRouterMethods(&TestPostUnregController{}, literal_0637))

	// Test replacement root (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test replacement root (expect no change from the original)", method, "/", contentRootOriginal)

	// Test level 1 (expect no change from the original)
	testHelperFnContentCheck(t, handler, "Test level 1 (expect no change from the original)", method, literal_4926, contentLevel1Original)

	// Test level 2 (expect change)
	testHelperFnContentCheck(t, handler, "Test level 2 (expect change)", method, literal_5074, contentLevel2Replacement)
}

func testHelperFnContentCheck(t *testing.T, handler *ControllerRegister,
	testName, method, path, expectedBodyContent string,
) {
	r, err := http.NewRequest(method, path, nil)
	if err != nil {
		t.Errorf("httpRecorderBodyTest NewRequest error: %v", err)
		return
	}
	w := httptest.NewRecorder()
	handler.ServeHTTP(w, r)
	body := w.Body.String()
	if body != expectedBodyContent {
		t.Errorf("%s: expected [%s], got [%s];", testName, expectedBodyContent, body)
	}
}

const literal_2806 = "get:GetFixedRoot"

const literal_4926 = "/level1"

const literal_4678 = "get:GetFixedLevel1"

const literal_5074 = "/level1/level2"

const literal_0637 = "get:GetFixedLevel2"
