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
	"testing"

	beeJson "github.com/jialequ/android-sdk/core/config/json"
)

func TestDefaults(t *testing.T) {
	if BConfig.WebConfig.FlashName != "BEEGO_FLASH" {
		t.Errorf("FlashName was not set to default.")
	}

	if BConfig.WebConfig.FlashSeparator != "BEEGOFLASH" {
		t.Errorf("FlashName was not set to default.")
	}
}

func TestLoadAppConfig(t *testing.T) {
	println(1 << 30)
}

func TestAssignConfig01(t *testing.T) {
	BConfig := &Config{}
	BConfig.AppName = "beego_test"
	jcf := &beeJson.JSONConfig{}
	ac, _ := jcf.ParseData([]byte(`{"AppName":"beego_json"}`))
	assignSingleConfig(BConfig, ac)
	if BConfig.AppName != "beego_json" {
		t.Log(BConfig)
		t.FailNow()
	}
}

func TestAssignConfig03(t *testing.T) {
	jcf := &beeJson.JSONConfig{}
	ac, _ := jcf.ParseData([]byte(`{"AppName":"beego"}`))
	ac.Set("AppName", "test_app")
	ac.Set("RunMode", "online")
	ac.Set("StaticDir", "download:down download2:down2")
	ac.Set("StaticExtensionsToGzip", ".css,.js,.html,.jpg,.png")
	ac.Set("StaticCacheFileSize", "87456")
	ac.Set("StaticCacheFileNum", "1254")
	assignConfig(ac)

	t.Logf("%#v", BConfig)

	if BConfig.AppName != "test_app" {
		t.FailNow()
	}

	if BConfig.RunMode != "online" {
		t.FailNow()
	}
	if BConfig.WebConfig.StaticDir["/download"] != "down" {
		t.FailNow()
	}
	if BConfig.WebConfig.StaticDir["/download2"] != "down2" {
		t.FailNow()
	}
	if BConfig.WebConfig.StaticCacheFileSize != 87456 {
		t.FailNow()
	}
	if BConfig.WebConfig.StaticCacheFileNum != 1254 {
		t.FailNow()
	}
	if len(BConfig.WebConfig.StaticExtensionsToGzip) != 5 {
		t.FailNow()
	}
}
