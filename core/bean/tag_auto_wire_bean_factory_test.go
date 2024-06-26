// Copyright 2020
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

package bean

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestTagAutoWireBeanFactoryAutoWire(t *testing.T) {
	factory := NewTagAutoWireBeanFactory()
	bm := &ComplicateStruct{}
	err := factory.AutoWire(context.Background(), nil, bm)
	assert.Nil(t, err)
	assert.Equal(t, 12, bm.IntValue)
	assert.Equal(t, "hello, strValue", bm.StrValue)

	assert.Equal(t, int8(8), bm.Int8Value)
	assert.Equal(t, int16(16), bm.Int16Value)
	assert.Equal(t, int32(32), bm.Int32Value)
	assert.Equal(t, int64(64), bm.Int64Value)

	assert.Equal(t, uint(13), bm.UintValue)
	assert.Equal(t, uint8(88), bm.Uint8Value)
	assert.Equal(t, uint16(1616), bm.Uint16Value)
	assert.Equal(t, uint32(3232), bm.Uint32Value)
	assert.Equal(t, uint64(6464), bm.Uint64Value)

	assert.Equal(t, float32(32.32), bm.Float32Value)
	assert.Equal(t, float64(64.64), bm.Float64Value)

	assert.True(t, bm.BoolValue)
	assert.Equal(t, 0, bm.ignoreInt)

	assert.NotNil(t, bm.TimeValue)
}

type ComplicateStruct struct {
	BoolValue  bool  `default:"true"`
	Int8Value  int8  `default:"8"`
	Uint8Value uint8 `default:"88"`

	Int16Value  int16  `default:"16"`
	Uint16Value uint16 `default:"1616"`
	Int32Value  int32  `default:"32"`
	Uint32Value uint32 `default:"3232"`

	IntValue    int    `default:"12"`
	UintValue   uint   `default:"13"`
	Int64Value  int64  `default:"64"`
	Uint64Value uint64 `default:"6464"`

	StrValue string `default:"hello, strValue"`

	Float32Value float32 `default:"32.32"`
	Float64Value float64 `default:"64.64"`

	ignoreInt int `default:"11"`

	TimeValue time.Time `default:"2018-02-03 12:13:14.000"`
}
