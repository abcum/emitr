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

import "sync"

// Emitter represents an event emitter.
type Emitter struct {
	lock sync.Mutex
	once map[string][]func()
	when map[string][]func()
}

// Once adds a callback which is called only once when an
// event is emitted, and then removed immediately after.
func (e *Emitter) Once(name string, fn func()) {

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.once == nil {
		e.once = make(map[string][]func())
	}

	e.once[name] = append(e.once[name], fn)

}

// When adds a callback which is called each and every
// time that an event is emitted.
func (e *Emitter) When(name string, fn func()) {

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.when == nil {
		e.when = make(map[string][]func())
	}

	e.when[name] = append(e.when[name], fn)

}

// Emit runs callbacks for the specified event name,
// and ensures that one-off events, are removed when
// the event has finished emitting.
func (e *Emitter) Emit(name string) {

	e.lock.Lock()
	defer e.lock.Unlock()

	if e.when != nil {
		if _, ok := e.when[name]; ok {
			for i := len(e.when[name]) - 1; i >= 0; i-- {
				e.when[name][i]()
			}
		}
	}

	if e.once != nil {
		if _, ok := e.once[name]; ok {
			for i := len(e.once[name]) - 1; i >= 0; i-- {
				e.once[name][i]()
				e.once[name][i] = nil
				e.once[name] = e.once[name][:i]
			}
		}
	}

}
