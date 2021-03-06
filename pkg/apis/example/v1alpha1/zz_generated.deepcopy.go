// +build !ignore_autogenerated

/*
This code is auto-generated by k8s.io/code-generator
*/
// This file was autogenerated by deepcopy-gen. Do not edit it manually!

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpleExample) DeepCopyInto(out *SimpleExample) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpleExample.
func (in *SimpleExample) DeepCopy() *SimpleExample {
	if in == nil {
		return nil
	}
	out := new(SimpleExample)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimpleExample) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpleExampleList) DeepCopyInto(out *SimpleExampleList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	out.ListMeta = in.ListMeta
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]SimpleExampleSpec, len(*in))
		copy(*out, *in)
	}
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpleExampleList.
func (in *SimpleExampleList) DeepCopy() *SimpleExampleList {
	if in == nil {
		return nil
	}
	out := new(SimpleExampleList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *SimpleExampleList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	} else {
		return nil
	}
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *SimpleExampleSpec) DeepCopyInto(out *SimpleExampleSpec) {
	*out = *in
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new SimpleExampleSpec.
func (in *SimpleExampleSpec) DeepCopy() *SimpleExampleSpec {
	if in == nil {
		return nil
	}
	out := new(SimpleExampleSpec)
	in.DeepCopyInto(out)
	return out
}
