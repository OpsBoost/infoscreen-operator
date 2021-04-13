// +build !ignore_autogenerated

/*
Copyright 2021 OpsBoost Crew <info@opsboost.dev>.

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

// Code generated by controller-gen. DO NOT EDIT.

package v1

import (
	corev1 "k8s.io/api/core/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CredentialsSpec) DeepCopyInto(out *CredentialsSpec) {
	*out = *in
	if in.SecretRef != nil {
		in, out := &in.SecretRef, &out.SecretRef
		*out = new(corev1.SecretEnvSource)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CredentialsSpec.
func (in *CredentialsSpec) DeepCopy() *CredentialsSpec {
	if in == nil {
		return nil
	}
	out := new(CredentialsSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Firefox) DeepCopyInto(out *Firefox) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Firefox.
func (in *Firefox) DeepCopy() *Firefox {
	if in == nil {
		return nil
	}
	out := new(Firefox)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Firefox) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirefoxList) DeepCopyInto(out *FirefoxList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Firefox, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirefoxList.
func (in *FirefoxList) DeepCopy() *FirefoxList {
	if in == nil {
		return nil
	}
	out := new(FirefoxList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *FirefoxList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirefoxSpec) DeepCopyInto(out *FirefoxSpec) {
	*out = *in
	in.SessionSpec.DeepCopyInto(&out.SessionSpec)
	if in.Credentials != nil {
		in, out := &in.Credentials, &out.Credentials
		*out = new(CredentialsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirefoxSpec.
func (in *FirefoxSpec) DeepCopy() *FirefoxSpec {
	if in == nil {
		return nil
	}
	out := new(FirefoxSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *FirefoxStatus) DeepCopyInto(out *FirefoxStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new FirefoxStatus.
func (in *FirefoxStatus) DeepCopy() *FirefoxStatus {
	if in == nil {
		return nil
	}
	out := new(FirefoxStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionClusterRefSpec) DeepCopyInto(out *SessionClusterRefSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionClusterRefSpec.
func (in *SessionClusterRefSpec) DeepCopy() *SessionClusterRefSpec {
	if in == nil {
		return nil
	}
	out := new(SessionClusterRefSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionResolutionSpec) DeepCopyInto(out *SessionResolutionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionResolutionSpec.
func (in *SessionResolutionSpec) DeepCopy() *SessionResolutionSpec {
	if in == nil {
		return nil
	}
	out := new(SessionResolutionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SessionSpec) DeepCopyInto(out *SessionSpec) {
	*out = *in
	if in.Resolution != nil {
		in, out := &in.Resolution, &out.Resolution
		*out = new(SessionResolutionSpec)
		**out = **in
	}
	if in.ClusterRef != nil {
		in, out := &in.ClusterRef, &out.ClusterRef
		*out = new(SessionClusterRefSpec)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SessionSpec.
func (in *SessionSpec) DeepCopy() *SessionSpec {
	if in == nil {
		return nil
	}
	out := new(SessionSpec)
	in.DeepCopyInto(out)
	return out
}
