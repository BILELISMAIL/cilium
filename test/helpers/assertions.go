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

package helpers

import (
	"github.com/onsi/gomega"
)

// ExpectCiliumReady asserts that cilium status is ready
func ExpectCiliumReady(vm *SSHMeta) {
	err := vm.WaitUntilReady(100)
	gomega.ExpectWithOffset(1, err).To(gomega.BeNil(), "Cilium-agent cannot be started")

	vm.NetworkCreate(CiliumDockerNetwork, "")
	res := vm.NetworkGet(CiliumDockerNetwork)
	gomega.ExpectWithOffset(1, res).To(CMDSuccess(), "Cilium docker network is not created")

	res = vm.SetPolicyEnforcement(PolicyEnforcementDefault)
	gomega.ExpectWithOffset(1, res).To(CMDSuccess(), "Cannot set policy enforcement default")
}
