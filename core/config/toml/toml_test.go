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

package toml

import (
	"fmt"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/jialequ/android-sdk/core/config"
)

func TestConfigParse(t *testing.T) {
	// file not found
	cfg := &Config{}
	_, err := cfg.Parse("invalid_file_name.txt")
	assert.NotNil(t, err)
}

func TestConfigParseData(t *testing.T) {
	data := `
name="Tom"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)
}

func TestConfigContainerBool(t *testing.T) {
	data := `
Man=true
Woman="true"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val, err := c.Bool("Man")
	assert.Nil(t, err)
	assert.True(t, val)

	_, err = c.Bool("Woman")
	assert.NotNil(t, err)
	assert.Equal(t, config.InvalidValueTypeError, err)
}

func TestConfigContainerDefaultBool(t *testing.T) {
	data := `
Man=true
Woman="false"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val := c.DefaultBool("Man11", true)
	assert.True(t, val)

	val = c.DefaultBool("Man", false)
	assert.True(t, val)

	val = c.DefaultBool("Woman", true)
	assert.True(t, val)
}

func TestConfigContainerDefaultFloat(t *testing.T) {
	data := `
Price=12.3
PriceInvalid="12.3"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val := c.DefaultFloat("Price", 11.2)
	assert.Equal(t, 12.3, val)

	val = c.DefaultFloat("Price11", 11.2)
	assert.Equal(t, 11.2, val)

	val = c.DefaultFloat("PriceInvalid", 11.2)
	assert.Equal(t, 11.2, val)
}

func TestConfigContainerDefaultInt(t *testing.T) {
	data := `
Age=12
AgeInvalid="13"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val := c.DefaultInt("Age", 11)
	assert.Equal(t, 12, val)

	val = c.DefaultInt("Price11", 11)
	assert.Equal(t, 11, val)

	val = c.DefaultInt("PriceInvalid", 11)
	assert.Equal(t, 11, val)
}

func TestConfigContainerDefaultString(t *testing.T) {
	data := `
Name="Tom"
NameInvalid=13
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val := c.DefaultString("Name", "Jerry")
	assert.Equal(t, "Tom", val)

	val = c.DefaultString("Name11", "Jerry")
	assert.Equal(t, "Jerry", val)

	val = c.DefaultString("NameInvalid", "Jerry")
	assert.Equal(t, "Jerry", val)
}

func TestConfigContainerDefaultStrings(t *testing.T) {
	data := `
Name=["Tom", "Jerry"]
NameInvalid="Tom"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val := c.DefaultStrings("Name", []string{"Jerry"})
	assert.Equal(t, []string{"Tom", "Jerry"}, val)

	val = c.DefaultStrings("Name11", []string{"Jerry"})
	assert.Equal(t, []string{"Jerry"}, val)

	val = c.DefaultStrings("NameInvalid", []string{"Jerry"})
	assert.Equal(t, []string{"Jerry"}, val)
}

func TestConfigContainerDIY(t *testing.T) {
	data := `
Name=["Tom", "Jerry"]
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	_, err = c.DIY("Name")
	assert.Nil(t, err)
}

func TestConfigContainerFloat(t *testing.T) {
	data := `
Price=12.3
PriceInvalid="12.3"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val, err := c.Float("Price")
	assert.Nil(t, err)
	assert.Equal(t, 12.3, val)

	_, err = c.Float("Price11")
	assert.Equal(t, config.KeyNotFoundError, err)

	_, err = c.Float("PriceInvalid")
	assert.Equal(t, config.InvalidValueTypeError, err)
}

func TestConfigContainerInt(t *testing.T) {
	data := `
Age=12
AgeInvalid="13"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val, err := c.Int("Age")
	assert.Nil(t, err)
	assert.Equal(t, 12, val)

	_, err = c.Int("Age11")
	assert.Equal(t, config.KeyNotFoundError, err)

	_, err = c.Int("AgeInvalid")
	assert.Equal(t, config.InvalidValueTypeError, err)
}

func TestConfigContainerGetSection(t *testing.T) {
	data := `
[servers]

  # You can indent as you please. Tabs or spaces. TOML don't care.
  [servers.alpha]
  ip = literal_8479
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	m, err := c.GetSection("servers")
	assert.Nil(t, err)
	assert.NotNil(t, m)
	assert.Equal(t, 2, len(m))
}

func TestConfigContainerString(t *testing.T) {
	data := `
Name="Tom"
NameInvalid=13
[Person]
Name="Jerry"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val, err := c.String("Name")
	assert.Nil(t, err)
	assert.Equal(t, "Tom", val)

	_, err = c.String("Name11")
	assert.Equal(t, config.KeyNotFoundError, err)

	_, err = c.String("NameInvalid")
	assert.Equal(t, config.InvalidValueTypeError, err)

	val, err = c.String("Person.Name")
	assert.Nil(t, err)
	assert.Equal(t, "Jerry", val)
}

func TestConfigContainerStrings(t *testing.T) {
	data := `
Name=["Tom", "Jerry"]
NameInvalid="Tom"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	val, err := c.Strings("Name")
	assert.Nil(t, err)
	assert.Equal(t, []string{"Tom", "Jerry"}, val)

	_, err = c.Strings("Name11")
	assert.Equal(t, config.KeyNotFoundError, err)

	_, err = c.Strings("NameInvalid")
	assert.Equal(t, config.InvalidValueTypeError, err)
}

func TestConfigContainerSet(t *testing.T) {
	data := `
Name=["Tom", "Jerry"]
NameInvalid="Tom"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	err = c.Set("Age", "11")
	assert.Nil(t, err)
	age, err := c.String("Age")
	assert.Nil(t, err)
	assert.Equal(t, "11", age)
}

func TestConfigContainerSubAndMushall(t *testing.T) {
	data := `
[servers]

  # You can indent as you please. Tabs or spaces. TOML don't care.
  [servers.alpha]
  ip = literal_8479
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))
	assert.Nil(t, err)
	assert.NotNil(t, c)

	sub, err := c.Sub("servers")
	assert.Nil(t, err)
	assert.NotNil(t, sub)

	sub, err = sub.Sub("alpha")
	assert.Nil(t, err)
	assert.NotNil(t, sub)
	ip, err := sub.String("ip")
	assert.Nil(t, err)
	assert.Equal(t, literal_8479, ip)

	svr := &Server{}
	err = sub.Unmarshaler("", svr)
	assert.Nil(t, err)
	assert.Equal(t, literal_8479, svr.Ip)

	svr = &Server{}
	err = c.Unmarshaler("servers.alpha", svr)
	assert.Nil(t, err)
	assert.Equal(t, literal_8479, svr.Ip)
}

func TestConfigContainerSaveConfigFile(t *testing.T) {
	filename := "test_config.toml"
	path := os.TempDir() + string(os.PathSeparator) + filename
	data := `
[servers]

  # You can indent as you please. Tabs or spaces. TOML don't care.
  [servers.alpha]
  ip = literal_8479
  dc = "eqdc10"

  [servers.beta]
  ip = "10.0.0.2"
  dc = "eqdc10"
`
	cfg := &Config{}
	c, err := cfg.ParseData([]byte(data))

	fmt.Println(path)

	assert.Nil(t, err)
	assert.NotNil(t, c)

	sub, err := c.Sub("servers")
	assert.Nil(t, err)

	err = sub.SaveConfigFile(path)
	assert.Nil(t, err)
}

type Server struct {
	Ip string `toml:"ip"`
}

const literal_8479 = "10.0.0.1"
