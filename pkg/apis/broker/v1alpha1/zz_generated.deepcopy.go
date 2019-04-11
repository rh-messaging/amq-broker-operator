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
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemis) DeepCopyInto(out *ActiveMQArtemis) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemis.
func (in *ActiveMQArtemis) DeepCopy() *ActiveMQArtemis {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemis) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddress) DeepCopyInto(out *ActiveMQArtemisAddress) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddress.
func (in *ActiveMQArtemisAddress) DeepCopy() *ActiveMQArtemisAddress {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddress)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisAddress) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressList) DeepCopyInto(out *ActiveMQArtemisAddressList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveMQArtemisAddress, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressList.
func (in *ActiveMQArtemisAddressList) DeepCopy() *ActiveMQArtemisAddressList {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisAddressList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressSpec) DeepCopyInto(out *ActiveMQArtemisAddressSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressSpec.
func (in *ActiveMQArtemisAddressSpec) DeepCopy() *ActiveMQArtemisAddressSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisAddressStatus) DeepCopyInto(out *ActiveMQArtemisAddressStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisAddressStatus.
func (in *ActiveMQArtemisAddressStatus) DeepCopy() *ActiveMQArtemisAddressStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisAddressStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisList) DeepCopyInto(out *ActiveMQArtemisList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ActiveMQArtemis, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisList.
func (in *ActiveMQArtemisList) DeepCopy() *ActiveMQArtemisList {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ActiveMQArtemisList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisSpec) DeepCopyInto(out *ActiveMQArtemisSpec) {
	*out = *in
	out.ClusterConfig = in.ClusterConfig
	out.SSLConfig = in.SSLConfig
	out.CommonConfig = in.CommonConfig
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisSpec.
func (in *ActiveMQArtemisSpec) DeepCopy() *ActiveMQArtemisSpec {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ActiveMQArtemisStatus) DeepCopyInto(out *ActiveMQArtemisStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ActiveMQArtemisStatus.
func (in *ActiveMQArtemisStatus) DeepCopy() *ActiveMQArtemisStatus {
	if in == nil {
		return nil
	}
	out := new(ActiveMQArtemisStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ClusterConfig) DeepCopyInto(out *ClusterConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ClusterConfig.
func (in *ClusterConfig) DeepCopy() *ClusterConfig {
	if in == nil {
		return nil
	}
	out := new(ClusterConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CommonConfig) DeepCopyInto(out *CommonConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CommonConfig.
func (in *CommonConfig) DeepCopy() *CommonConfig {
	if in == nil {
		return nil
	}
	out := new(CommonConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SSLConfig) DeepCopyInto(out *SSLConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SSLConfig.
func (in *SSLConfig) DeepCopy() *SSLConfig {
	if in == nil {
		return nil
	}
	out := new(SSLConfig)
	in.DeepCopyInto(out)
	return out
}
