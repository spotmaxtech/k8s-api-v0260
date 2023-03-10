/*
Copyright 2022 The Kubernetes Authors.

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

package v1alpha1

import (
	v1 "github.com/spotmaxtech/k8s-api-v0260/core/v1"
	metav1 "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/apis/meta/v1"
)

// +genclient
// +genclient:nonNamespaced
// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.25

// ClusterCIDR represents a single configuration for per-Node Pod CIDR
// allocations when the MultiCIDRRangeAllocator is enabled (see the config for
// kube-controller-manager).  A cluster may have any number of ClusterCIDR
// resources, all of which will be considered when allocating a CIDR for a
// Node.  A ClusterCIDR is eligible to be used for a given Node when the node
// selector matches the node in question and has free CIDRs to allocate.  In
// case of multiple matching ClusterCIDR resources, the allocator will attempt
// to break ties using internal heuristics, but any ClusterCIDR whose node
// selector matches the Node may be used.
type ClusterCIDR struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ObjectMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Spec is the desired state of the ClusterCIDR.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
	// +optional
	Spec ClusterCIDRSpec `json:"spec,omitempty" protobuf:"bytes,2,opt,name=spec"`
}

// ClusterCIDRSpec defines the desired state of ClusterCIDR.
type ClusterCIDRSpec struct {
	// NodeSelector defines which nodes the config is applicable to.
	// An empty or nil NodeSelector selects all nodes.
	// This field is immutable.
	// +optional
	NodeSelector *v1.NodeSelector `json:"nodeSelector,omitempty" protobuf:"bytes,1,opt,name=nodeSelector"`

	// PerNodeHostBits defines the number of host bits to be configured per node.
	// A subnet mask determines how much of the address is used for network bits
	// and host bits. For example an IPv4 address of 192.168.0.0/24, splits the
	// address into 24 bits for the network portion and 8 bits for the host portion.
	// To allocate 256 IPs, set this field to 8 (a /24 mask for IPv4 or a /120 for IPv6).
	// Minimum value is 4 (16 IPs).
	// This field is immutable.
	// +required
	PerNodeHostBits int32 `json:"perNodeHostBits" protobuf:"varint,2,opt,name=perNodeHostBits"`

	// IPv4 defines an IPv4 IP block in CIDR notation(e.g. "10.0.0.0/8").
	// At least one of IPv4 and IPv6 must be specified.
	// This field is immutable.
	// +optional
	IPv4 string `json:"ipv4" protobuf:"bytes,3,opt,name=ipv4"`

	// IPv6 defines an IPv6 IP block in CIDR notation(e.g. "2001:db8::/64").
	// At least one of IPv4 and IPv6 must be specified.
	// This field is immutable.
	// +optional
	IPv6 string `json:"ipv6" protobuf:"bytes,4,opt,name=ipv6"`
}

// +k8s:deepcopy-gen:interfaces=k8s.io/apimachinery/pkg/runtime.Object
// +k8s:prerelease-lifecycle-gen:introduced=1.25

// ClusterCIDRList contains a list of ClusterCIDR.
type ClusterCIDRList struct {
	metav1.TypeMeta `json:",inline"`
	// Standard object's metadata.
	// More info: https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata
	// +optional
	metav1.ListMeta `json:"metadata,omitempty" protobuf:"bytes,1,opt,name=metadata"`

	// Items is the list of ClusterCIDRs.
	Items []ClusterCIDR `json:"items" protobuf:"bytes,2,rep,name=items"`
}
