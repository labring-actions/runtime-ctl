/*
Copyright 2023 cuisongliu@qq.com.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package cri

import v1 "github.com/labring-actions/runtime-ctl/types/v1"

func CRIRuntime(cri string, component v1.RuntimeDefaultComponent) (string, string) {
	switch cri {
	case "docker":
		return "runc", component.Runc
	case "crio":
		return "crun", component.Crun
	case "containerd":
		return "runc", component.Runc
	default:
		return "runc", component.Runc
	}
}
