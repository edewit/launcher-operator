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
func (in *GitConfig) DeepCopyInto(out *GitConfig) {
	*out = *in
	out.Token = in.Token
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GitConfig.
func (in *GitConfig) DeepCopy() *GitConfig {
	if in == nil {
		return nil
	}
	out := new(GitConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Launcher) DeepCopyInto(out *Launcher) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Launcher.
func (in *Launcher) DeepCopy() *Launcher {
	if in == nil {
		return nil
	}
	out := new(Launcher)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Launcher) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LauncherList) DeepCopyInto(out *LauncherList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Launcher, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LauncherList.
func (in *LauncherList) DeepCopy() *LauncherList {
	if in == nil {
		return nil
	}
	out := new(LauncherList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *LauncherList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LauncherSpec) DeepCopyInto(out *LauncherSpec) {
	*out = *in
	out.Git = in.Git
	out.OpenShift = in.OpenShift
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LauncherSpec.
func (in *LauncherSpec) DeepCopy() *LauncherSpec {
	if in == nil {
		return nil
	}
	out := new(LauncherSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *LauncherStatus) DeepCopyInto(out *LauncherStatus) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new LauncherStatus.
func (in *LauncherStatus) DeepCopy() *LauncherStatus {
	if in == nil {
		return nil
	}
	out := new(LauncherStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OpenShiftConfig) DeepCopyInto(out *OpenShiftConfig) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OpenShiftConfig.
func (in *OpenShiftConfig) DeepCopy() *OpenShiftConfig {
	if in == nil {
		return nil
	}
	out := new(OpenShiftConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SecretKeyRef) DeepCopyInto(out *SecretKeyRef) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SecretKeyRef.
func (in *SecretKeyRef) DeepCopy() *SecretKeyRef {
	if in == nil {
		return nil
	}
	out := new(SecretKeyRef)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SensitiveValue) DeepCopyInto(out *SensitiveValue) {
	*out = *in
	out.ValueFrom = in.ValueFrom
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SensitiveValue.
func (in *SensitiveValue) DeepCopy() *SensitiveValue {
	if in == nil {
		return nil
	}
	out := new(SensitiveValue)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ValueFrom) DeepCopyInto(out *ValueFrom) {
	*out = *in
	out.SecretKeyRef = in.SecretKeyRef
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ValueFrom.
func (in *ValueFrom) DeepCopy() *ValueFrom {
	if in == nil {
		return nil
	}
	out := new(ValueFrom)
	in.DeepCopyInto(out)
	return out
}
