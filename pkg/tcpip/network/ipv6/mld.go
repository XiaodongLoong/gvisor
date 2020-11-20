// Copyright 2020 The gVisor Authors.
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

package ipv6

import (
	"gvisor.dev/gvisor/pkg/tcpip"
	"gvisor.dev/gvisor/pkg/tcpip/header"
)

// mldState is the per-interface MLD state.
//
// mldState.init MUST be called to initialize the MLD state.
type mldState struct {
	// The IPv6 endpoint this mldState is for.
	ep *endpoint
}

// init sets up an mldState struct, and is required to be called before using
// a new mldState.
func (mld *mldState) init(ep *endpoint) {
	mld.ep = ep
}

func (mld *mldState) handleMulticastListenerQuery(mldHdr header.MLD) {
	// TODO(gvisor.dev/issue/4861): Handle MLD Query messages.
}

func (mld *mldState) handleMulticastListenerReport(mldHdr header.MLD) {
	// TODO(gvisor.dev/issue/4861): Handle MLD Report messages.
}

// joinGroup handles joining a new group and sending and scheduling the required
// messages.
//
// If the group is already joined, returns tcpip.ErrDuplicateAddress.
func (mld *mldState) joinGroup(groupAddress tcpip.Address) *tcpip.Error {
	// TODO(gvisor.dev/issue/4861): Handle joining a multicast group.
	return nil
}

// leaveGroup handles removing the group from the membership map, cancels any
// delay timers associated with that group, and sends the Done message, if
// required.
//
// If the group is not joined, this function will do nothing.
func (mld *mldState) leaveGroup(groupAddress tcpip.Address) {
	// TODO(gvisor.dev/issue/4861): Handle leaving a multicast group.
}

// leaveAllGroups leaves all groups that were joined.
func (mld *mldState) leaveAllGroups() {
	// TODO(gvisor.dev/issue/4861): Handle leaving all multicast groups.
}
