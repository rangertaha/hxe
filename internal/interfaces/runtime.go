/*
Copyright © 2025 Rangertaha <rangertaha@gmail.com>

Licensed under the MIT License (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    https://opensource.org/licenses/MIT

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package interfaces

import "time"

type Status interface {
	State() int32
	Uptime() time.Duration
	Message() string
	Progress() int32
}

type Runner interface {
	Id() string
	Init() error
	Start() error
	Stop() error
	Fill() error
	Test() error
	Train() error
	Status() Status
}

type Plugin interface {
	Init() error
	Stdin() error
	Stdcmd() error
	Stdout() error
	Stderr() error
}
