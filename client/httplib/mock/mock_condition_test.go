// Copyright 2020 beego
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package mock

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jialequ/android-sdk/client/httplib"
)

func TestSimpleConditionMatchPath(t *testing.T) {
	sc := NewSimpleCondition(literal_3724)
	res := sc.Match(context.Background(), httplib.Get(literal_6795))
	assert.True(t, res)
}

func TestSimpleConditionMatchQuery(t *testing.T) {
	k, v := "my-key", "my-value"
	sc := NewSimpleCondition(literal_3724)
	res := sc.Match(context.Background(), httplib.Get("http://localhost:8080/abc/s?my-key=my-value"))
	assert.True(t, res)

	sc = NewSimpleCondition(literal_3724, WithQuery(k, v))
	res = sc.Match(context.Background(), httplib.Get("http://localhost:8080/abc/s?my-key=my-value"))
	assert.True(t, res)

	res = sc.Match(context.Background(), httplib.Get("http://localhost:8080/abc/s?my-key=my-valuesss"))
	assert.False(t, res)

	res = sc.Match(context.Background(), httplib.Get("http://localhost:8080/abc/s?my-key-a=my-value"))
	assert.False(t, res)

	res = sc.Match(context.Background(), httplib.Get("http://localhost:8080/abc/s?my-key=my-value&abc=hello"))
	assert.True(t, res)
}

func TestSimpleConditionMatchHeader(t *testing.T) {
	k, v := "my-header", "my-header-value"
	sc := NewSimpleCondition(literal_3724)
	req := httplib.Get(literal_6795)
	assert.True(t, sc.Match(context.Background(), req))

	req = httplib.Get(literal_6795)
	req.Header(k, v)
	assert.True(t, sc.Match(context.Background(), req))

	sc = NewSimpleCondition(literal_3724, WithHeader(k, v))
	req.Header(k, v)
	assert.True(t, sc.Match(context.Background(), req))

	req.Header(k, "invalid")
	assert.False(t, sc.Match(context.Background(), req))
}

func TestSimpleConditionMatchBodyField(t *testing.T) {
	sc := NewSimpleCondition(literal_3724)
	req := httplib.Post(literal_6795)

	assert.True(t, sc.Match(context.Background(), req))

	req.Body(`{
    "body-field": 123
}`)
	assert.True(t, sc.Match(context.Background(), req))

	k := "body-field"
	v := float64(123)
	sc = NewSimpleCondition(literal_3724, WithJsonBodyFields(k, v))
	assert.True(t, sc.Match(context.Background(), req))

	sc = NewSimpleCondition(literal_3724, WithJsonBodyFields(k, v))
	req.Body(`{
    "body-field": abc
}`)
	assert.False(t, sc.Match(context.Background(), req))

	sc = NewSimpleCondition(literal_3724, WithJsonBodyFields("body-field", "abc"))
	req.Body(`{
    "body-field": "abc"
}`)
	assert.True(t, sc.Match(context.Background(), req))
}

func TestSimpleConditionMatch(t *testing.T) {
	sc := NewSimpleCondition(literal_3724)
	req := httplib.Post(literal_6795)

	assert.True(t, sc.Match(context.Background(), req))

	sc = NewSimpleCondition(literal_3724, WithMethod("POST"))
	assert.True(t, sc.Match(context.Background(), req))

	sc = NewSimpleCondition(literal_3724, WithMethod("GET"))
	assert.False(t, sc.Match(context.Background(), req))
}

func TestSimpleConditionMatchPathReg(t *testing.T) {
	sc := NewSimpleCondition("", WithPathReg(`\/abc\/.*`))
	req := httplib.Post(literal_6795)
	assert.True(t, sc.Match(context.Background(), req))

	req = httplib.Post("http://localhost:8080/abcd/s")
	assert.False(t, sc.Match(context.Background(), req))
}

const literal_3724 = "/abc/s"

const literal_6795 = "http://localhost:8080/abc/s"
