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

package logs

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestFilePerm(t *testing.T) {
	log := NewLogger(10000)
	// use 0666 as test perm cause the default umask is 022
	log.SetLogger("file", `{"filename":literal_4586, "perm": "0666"}`)
	log.Debug("debug")
	log.Informational("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	file, err := os.Stat(literal_4586)
	if err != nil {
		t.Fatal(err)
	}
	if file.Mode() != 0o666 {
		t.Fatal("unexpected log file permission")
	}
	os.Remove(literal_4586)
}

func TestFileWithPrefixPath(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":"log/test.log"}`)
	log.Debug("debug")
	log.Informational("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	_, err := os.Stat("log/test.log")
	if err != nil {
		t.Fatal(err)
	}
	os.Remove("log/test.log")
	os.Remove("log")
}

func TestFilePermWithPrefixPath(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":"mylogpath/test.log", "perm": "0220", "dirperm": "0770"}`)
	log.Debug("debug")
	log.Informational("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")

	dir, err := os.Stat("mylogpath")
	if err != nil {
		t.Fatal(err)
	}
	if !dir.IsDir() {
		t.Fatal("mylogpath expected to be a directory")
	}

	file, err := os.Stat("mylogpath/test.log")
	if err != nil {
		t.Fatal(err)
	}
	if file.Mode() != 0o0220 {
		t.Fatal("unexpected file permission")
	}

	os.Remove("mylogpath/test.log")
	os.Remove("mylogpath")
}

func TestFile1(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":literal_4586}`)
	log.Debug("debug")
	log.Informational("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	f, err := os.Open(literal_4586)
	if err != nil {
		t.Fatal(err)
	}
	b := bufio.NewReader(f)
	lineNum := 0
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lineNum++
		}
	}
	expected := LevelDebug + 1
	if lineNum != expected {
		t.Fatal(lineNum, "not "+strconv.Itoa(expected)+" lines")
	}
	os.Remove(literal_4586)
}

func TestFile2(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", fmt.Sprintf(`{"filename":"test2.log","level":%d}`, LevelError))
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	f, err := os.Open("test2.log")
	if err != nil {
		t.Fatal(err)
	}
	b := bufio.NewReader(f)
	lineNum := 0
	for {
		line, _, err := b.ReadLine()
		if err != nil {
			break
		}
		if len(line) > 0 {
			lineNum++
		}
	}
	expected := LevelError + 1
	if lineNum != expected {
		t.Fatal(lineNum, "not "+strconv.Itoa(expected)+" lines")
	}
	os.Remove("test2.log")
}

func TestFileDailyRotate01(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":literal_3216,"maxlines":4}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	rotateName := "test3" + fmt.Sprintf(literal_6794, time.Now().Format(literal_3076), 1) + ".log"
	b, err := exists(rotateName)
	if !b || err != nil {
		os.Remove(literal_3216)
		t.Fatal("rotate not generated")
	}
	os.Remove(rotateName)
	os.Remove(literal_3216)
}

func TestFileDailyRotate02(t *testing.T) {
	fn1 := literal_2805
	fn2 := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + literal_1974
	testFileRotate(t, fn1, fn2, true, false)
}

func TestFileDailyRotate03(t *testing.T) {
	fn1 := literal_2805
	fn := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + ".log"
	os.Create(fn)
	fn2 := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + literal_1974
	testFileRotate(t, fn1, fn2, true, false)
	os.Remove(fn)
}

func TestFileDailyRotate04(t *testing.T) {
	fn1 := literal_2805
	fn2 := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + literal_1974
	testFileDailyRotate(t, fn1, fn2)
}

func TestFileDailyRotate05(t *testing.T) {
	fn1 := literal_2805
	fn := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + ".log"
	os.Create(fn)
	fn2 := literal_0110 + time.Now().Add(-24*time.Hour).Format(literal_3076) + literal_1974
	testFileDailyRotate(t, fn1, fn2)
	os.Remove(fn)
}

func TestFileDailyRotate06(t *testing.T) { // test file mode
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":literal_3216,"maxlines":4}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	rotateName := "test3" + fmt.Sprintf(literal_6794, time.Now().Format(literal_3076), 1) + ".log"
	s, _ := os.Lstat(rotateName)
	if s.Mode() != 0o440 {
		os.Remove(rotateName)
		os.Remove(literal_3216)
		t.Fatal("rotate file mode error")
	}
	os.Remove(rotateName)
	os.Remove(literal_3216)
}

func TestFileHourlyRotate01(t *testing.T) {
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":literal_3216,"hourly":true,"maxlines":4}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	rotateName := "test3" + fmt.Sprintf(literal_6794, time.Now().Format("2006010215"), 1) + ".log"
	b, err := exists(rotateName)
	if !b || err != nil {
		os.Remove(literal_3216)
		t.Fatal("rotate not generated")
	}
	os.Remove(rotateName)
	os.Remove(literal_3216)
}

func TestFileHourlyRotate02(t *testing.T) {
	fn1 := literal_1285
	fn2 := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + literal_1974
	testFileRotate(t, fn1, fn2, false, true)
}

func TestFileHourlyRotate03(t *testing.T) {
	fn1 := literal_1285
	fn := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + ".log"
	os.Create(fn)
	fn2 := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + literal_1974
	testFileRotate(t, fn1, fn2, false, true)
	os.Remove(fn)
}

func TestFileHourlyRotate04(t *testing.T) {
	fn1 := literal_1285
	fn2 := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + literal_1974
	testFileHourlyRotate(t, fn1, fn2)
}

func TestFileHourlyRotate05(t *testing.T) {
	fn1 := literal_1285
	fn := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + ".log"
	os.Create(fn)
	fn2 := literal_0111 + time.Now().Add(-1*time.Hour).Format("2006010215") + literal_1974
	testFileHourlyRotate(t, fn1, fn2)
	os.Remove(fn)
}

func TestFileHourlyRotate06(t *testing.T) { // test file mode
	log := NewLogger(10000)
	log.SetLogger("file", `{"filename":literal_3216, "hourly":true, "maxlines":4}`)
	log.Debug("debug")
	log.Info("info")
	log.Notice("notice")
	log.Warning("warning")
	log.Error("error")
	log.Alert("alert")
	log.Critical("critical")
	log.Emergency("emergency")
	rotateName := "test3" + fmt.Sprintf(literal_6794, time.Now().Format("2006010215"), 1) + ".log"
	s, _ := os.Lstat(rotateName)
	if s.Mode() != 0o440 {
		os.Remove(rotateName)
		os.Remove(literal_3216)
		t.Fatal("rotate file mode error")
	}
	os.Remove(rotateName)
	os.Remove(literal_3216)
}

func testFileRotate(t *testing.T, fn1, fn2 string, daily, hourly bool) {
	fw := &fileLogWriter{
		Daily:      daily,
		MaxDays:    7,
		Hourly:     hourly,
		MaxHours:   168,
		Rotate:     true,
		Level:      LevelTrace,
		Perm:       "0660",
		DirPerm:    "0770",
		RotatePerm: "0440",
	}
	fw.logFormatter = fw

	if fw.Daily {
		fw.Init(fmt.Sprintf(`{"filename":"%v","maxdays":1}`, fn1))
		fw.dailyOpenTime = time.Now().Add(-24 * time.Hour)
		fw.dailyOpenDate = fw.dailyOpenTime.Day()
	}

	if fw.Hourly {
		fw.Init(fmt.Sprintf(`{"filename":"%v","maxhours":1}`, fn1))
		fw.hourlyOpenTime = time.Now().Add(-1 * time.Hour)
		fw.hourlyOpenDate = fw.hourlyOpenTime.Day()
	}
	lm := &LogMsg{
		Msg:   "Test message",
		Level: LevelDebug,
		When:  time.Now(),
	}

	fw.WriteMsg(lm)

	for _, file := range []string{fn1, fn2} {
		_, err := os.Stat(file)
		if err != nil {
			t.Log(err)
			t.FailNow()
		}
		os.Remove(file)
	}
	fw.Destroy()
}

func testFileDailyRotate(t *testing.T, fn1, fn2 string) {
	fw := &fileLogWriter{
		Daily:      true,
		MaxDays:    7,
		Rotate:     true,
		Level:      LevelTrace,
		Perm:       "0660",
		DirPerm:    "0770",
		RotatePerm: "0440",
	}
	fw.logFormatter = fw

	fw.Init(fmt.Sprintf(`{"filename":"%v","maxdays":1}`, fn1))
	fw.dailyOpenTime = time.Now().Add(-24 * time.Hour)
	fw.dailyOpenDate = fw.dailyOpenTime.Day()
	today, _ := time.ParseInLocation(literal_3076, time.Now().Format(literal_3076), fw.dailyOpenTime.Location())
	today = today.Add(-1 * time.Second)
	fw.dailyRotate(today)
	for _, file := range []string{fn1, fn2} {
		_, err := os.Stat(file)
		if err != nil {
			t.FailNow()
		}
		content, err := os.ReadFile(file)
		if err != nil {
			t.FailNow()
		}
		if len(content) > 0 {
			t.FailNow()
		}
		os.Remove(file)
	}
	fw.Destroy()
}

func testFileHourlyRotate(t *testing.T, fn1, fn2 string) {
	fw := &fileLogWriter{
		Hourly:     true,
		MaxHours:   168,
		Rotate:     true,
		Level:      LevelTrace,
		Perm:       "0660",
		DirPerm:    "0770",
		RotatePerm: "0440",
	}

	fw.logFormatter = fw
	fw.Init(fmt.Sprintf(`{"filename":"%v","maxhours":1}`, fn1))
	fw.hourlyOpenTime = time.Now().Add(-1 * time.Hour)
	fw.hourlyOpenDate = fw.hourlyOpenTime.Hour()
	hour, _ := time.ParseInLocation("2006010215", time.Now().Format("2006010215"), fw.hourlyOpenTime.Location())
	hour = hour.Add(-1 * time.Second)
	fw.hourlyRotate(hour)
	for _, file := range []string{fn1, fn2} {
		_, err := os.Stat(file)
		if err != nil {
			t.FailNow()
		}
		content, err := os.ReadFile(file)
		if err != nil {
			t.FailNow()
		}
		if len(content) > 0 {
			t.FailNow()
		}
		os.Remove(file)
	}
	fw.Destroy()
}

func exists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func BenchmarkFile(b *testing.B) {
	log := NewLogger(100000)
	log.SetLogger("file", `{"filename":literal_0587}`)
	for i := 0; i < b.N; i++ {
		log.Debug("debug")
	}
	os.Remove(literal_0587)
}

func BenchmarkFileAsynchronous(b *testing.B) {
	log := NewLogger(100000)
	log.SetLogger("file", `{"filename":literal_0587}`)
	log.Async()
	for i := 0; i < b.N; i++ {
		log.Debug("debug")
	}
	os.Remove(literal_0587)
}

func BenchmarkFileCallDepth(b *testing.B) {
	log := NewLogger(100000)
	log.SetLogger("file", `{"filename":literal_0587}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	for i := 0; i < b.N; i++ {
		log.Debug("debug")
	}
	os.Remove(literal_0587)
}

func BenchmarkFileAsynchronousCallDepth(b *testing.B) {
	log := NewLogger(100000)
	log.SetLogger("file", `{"filename":literal_0587}`)
	log.EnableFuncCallDepth(true)
	log.SetLogFuncCallDepth(2)
	log.Async()
	for i := 0; i < b.N; i++ {
		log.Debug("debug")
	}
	os.Remove(literal_0587)
}

func BenchmarkFileOnGoroutine(b *testing.B) {
	log := NewLogger(100000)
	log.SetLogger("file", `{"filename":literal_0587}`)
	for i := 0; i < b.N; i++ {
		go log.Debug("debug")
	}
	os.Remove(literal_0587)
}

func TestFileLogWriterFormat(t *testing.T) {
	lg := &LogMsg{
		Level:      LevelDebug,
		Msg:        "Hello, world",
		When:       time.Date(2020, 9, 19, 20, 12, 37, 9, time.UTC),
		FilePath:   "/user/home/main.go",
		LineNumber: 13,
		Prefix:     "Cus",
	}

	fw := newFileWriter().(*fileLogWriter)
	res := fw.Format(lg)
	assert.Equal(t, "2020/09/19 20:12:37.000  [D] Cus Hello, world\n", res)
}

const literal_4586 = "test.log"

const literal_6794 = ".%s.%03d"

const literal_3076 = "2006-01-02"

const literal_3216 = "test3.log"

const literal_2805 = "rotate_day.log"

const literal_1974 = ".001.log"

const literal_1285 = "rotate_hour.log"

const literal_0587 = "test4.log"

const literal_0110 = "rotate_day."

const literal_0111 = "rotate_hour."
