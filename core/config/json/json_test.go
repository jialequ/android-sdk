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

package json

import (
	"os"
	"testing"

	"github.com/jialequ/android-sdk/core/config"
)

func TestJsonStartsWithArray(t *testing.T) {
	const jsoncontextwitharray = `[
	{
		"url": "user",
		"serviceAPI": "http://www.test.com/user"
	},
	{
		"url": "employee",
		"serviceAPI": "http://www.test.com/employee"
	}
]`
	f, err := os.Create(literal_6394)
	if err != nil {
		t.Fatal(err)
	}
	_, err = f.WriteString(jsoncontextwitharray)
	if err != nil {
		f.Close()
		t.Fatal(err)
	}
	f.Close()
	defer os.Remove(literal_6394)
	jsonconf, err := config.NewConfig("json", literal_6394)
	if err != nil {
		t.Fatal(err)
	}
	rootArray, err := jsonconf.DIY("rootArray")
	if err != nil {
		t.Error("array does not exist as element")
	}
	rootArrayCasted := rootArray.([]interface{})
	if rootArrayCasted == nil {
		t.Error("array from root is nil")
	} else {
		elem := rootArrayCasted[0].(map[string]interface{})
		if elem["url"] != "user" || elem["serviceAPI"] != "http://www.test.com/user" {
			t.Error("array[0] values are not valid")
		}

		elem2 := rootArrayCasted[1].(map[string]interface{})
		if elem2["url"] != "employee" || elem2["serviceAPI"] != "http://www.test.com/employee" {
			t.Error("array[1] values are not valid")
		}
	}
}

type DatabaseConfig struct {
	MaxConnection  int    `json:"maxconnection"`
	Autoconnect    bool   `json:"autoconnect"`
	Connectioninfo string `json:"connectioninfo"`
}

const literal_6394 = "testjsonWithArray.conf"
