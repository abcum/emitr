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

type Emitter struct {
	once map[string][]func()
	when map[string][]func()
}

func (e *Emitter) Once(name string, fn func()) {
	if e.once == nil {
		e.once = make(map[string][]func())
	}
	e.once[name] = append(e.once[name], fn)
}

func (e *Emitter) When(name string, fn func()) {
	if e.when == nil {
		e.when = make(map[string][]func())
	}
	e.when[name] = append(e.when[name], fn)
}

func (e *Emitter) Emit(name string) {

	if _, ok := e.when[name]; ok {
		for i := len(e.when[name]) - 1; i >= 0; i-- {
			e.when[name][i]()
		}
	}

	if _, ok := e.once[name]; ok {
		for i := len(e.once[name]) - 1; i >= 0; i-- {
			e.once[name][i]()
			e.once[name][i] = nil
			e.once[name] = e.once[name][:i]
		}
	}

}