//go:build !ignore_autogenerated
// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by deepcopy-gen. DO NOT EDIT.

package v1alpha1

import (
	runtime "github.com/spotmaxtech/k8s-apimachinery-v0260/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectReview) DeepCopyInto(out *SelfSubjectReview) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Status.DeepCopyInto(&out.Status)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectReview.
func (in *SelfSubjectReview) DeepCopy() *SelfSubjectReview {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectReview)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SelfSubjectReview) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SelfSubjectReviewStatus) DeepCopyInto(out *SelfSubjectReviewStatus) {
	*out = *in
	in.UserInfo.DeepCopyInto(&out.UserInfo)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SelfSubjectReviewStatus.
func (in *SelfSubjectReviewStatus) DeepCopy() *SelfSubjectReviewStatus {
	if in == nil {
		return nil
	}
	out := new(SelfSubjectReviewStatus)
	in.DeepCopyInto(out)
	return out
}
