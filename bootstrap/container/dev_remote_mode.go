/*******************************************************************************
 * Copyright 2023 Intel Inc.
 *
 * Licensed under the Apache License, Version 2.0 (the "License"); you may not use this file except
 * in compliance with the License. You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software distributed under the License
 * is distributed on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express
 * or implied. See the License for the specific language governing permissions and limitations under
 * the License.
 *******************************************************************************/

package container

import (
	"github.com/agile-edgex/go-mod-bootstrap/v3/di"
)

type DevRemoteMode struct {
	InDevMode    bool
	InRemoteMode bool
}

// DevRemoteModeName contains the name of the DevRemoteMode struct in the DIC.
var DevRemoteModeName = di.TypeInstanceToName((*DevRemoteMode)(nil))

// DevRemoteModeFrom helper function queries the DIC and returns the Dev and Remotes mode flags.
func DevRemoteModeFrom(get di.Get) DevRemoteMode {
	devOrRemoteMode, ok := get(DevRemoteModeName).(*DevRemoteMode)
	if !ok {
		return DevRemoteMode{}
	}

	return *devOrRemoteMode
}
