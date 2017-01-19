// Copyright © 2016 Abcum Ltd
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package emitr

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMain(t *testing.T) {

	e := &Emitter{}

	Convey("Emitter created successfully", t, func() {
		So(e, ShouldNotBeNil)
	})

	Convey("Can call emit with no callbacks", t, func() {
		e.Emit("none")
		So(e, ShouldNotBeNil)
		e.Emit("test")
		So(e, ShouldNotBeNil)
	})

	Convey("Can add a callback using `once`", t, func() {
		e.Once("test", func() {})
		So(len(e.once["test"]), ShouldEqual, 1)
		e.Once("test", func() {})
		So(len(e.once["test"]), ShouldEqual, 2)
	})

	Convey("Can add a callback using `when`", t, func() {
		e.When("test", func() {})
		So(len(e.when["test"]), ShouldEqual, 1)
		e.When("test", func() {})
		So(len(e.when["test"]), ShouldEqual, 2)
	})

	Convey("Can emit no events using `emit`", t, func() {
		e.Emit("none")
		So(len(e.once["test"]), ShouldEqual, 2)
		So(len(e.when["test"]), ShouldEqual, 2)
	})

	Convey("Can emit all events using `emit` for 1st time", t, func() {
		e.Emit("test")
		So(len(e.once["test"]), ShouldEqual, 0)
		So(len(e.when["test"]), ShouldEqual, 2)
	})

	Convey("Can emit all events using `emit` for 2nd time", t, func() {
		e.Emit("test")
		So(len(e.once["test"]), ShouldEqual, 0)
		So(len(e.when["test"]), ShouldEqual, 2)
	})

}
