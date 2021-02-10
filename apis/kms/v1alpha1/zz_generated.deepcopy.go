// +build !ignore_autogenerated

/*
Copyright 2020 The Crossplane Authors.

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

package v1alpha1

import (
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AliasListEntry) DeepCopyInto(out *AliasListEntry) {
	*out = *in
	if in.AliasARN != nil {
		in, out := &in.AliasARN, &out.AliasARN
		*out = new(string)
		**out = **in
	}
	if in.TargetKeyID != nil {
		in, out := &in.TargetKeyID, &out.TargetKeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AliasListEntry.
func (in *AliasListEntry) DeepCopy() *AliasListEntry {
	if in == nil {
		return nil
	}
	out := new(AliasListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomKeyParameters) DeepCopyInto(out *CustomKeyParameters) {
	*out = *in
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.PendingWindowInDays != nil {
		in, out := &in.PendingWindowInDays, &out.PendingWindowInDays
		*out = new(int64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomKeyParameters.
func (in *CustomKeyParameters) DeepCopy() *CustomKeyParameters {
	if in == nil {
		return nil
	}
	out := new(CustomKeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CustomKeyStoresListEntry) DeepCopyInto(out *CustomKeyStoresListEntry) {
	*out = *in
	if in.CloudHsmClusterID != nil {
		in, out := &in.CloudHsmClusterID, &out.CloudHsmClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.CustomKeyStoreID != nil {
		in, out := &in.CustomKeyStoreID, &out.CustomKeyStoreID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CustomKeyStoresListEntry.
func (in *CustomKeyStoresListEntry) DeepCopy() *CustomKeyStoresListEntry {
	if in == nil {
		return nil
	}
	out := new(CustomKeyStoresListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *GrantListEntry) DeepCopyInto(out *GrantListEntry) {
	*out = *in
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new GrantListEntry.
func (in *GrantListEntry) DeepCopy() *GrantListEntry {
	if in == nil {
		return nil
	}
	out := new(GrantListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Key) DeepCopyInto(out *Key) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Key.
func (in *Key) DeepCopy() *Key {
	if in == nil {
		return nil
	}
	out := new(Key)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *Key) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyList) DeepCopyInto(out *KeyList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]Key, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyList.
func (in *KeyList) DeepCopy() *KeyList {
	if in == nil {
		return nil
	}
	out := new(KeyList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *KeyList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyListEntry) DeepCopyInto(out *KeyListEntry) {
	*out = *in
	if in.KeyARN != nil {
		in, out := &in.KeyARN, &out.KeyARN
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyListEntry.
func (in *KeyListEntry) DeepCopy() *KeyListEntry {
	if in == nil {
		return nil
	}
	out := new(KeyListEntry)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyMetadata) DeepCopyInto(out *KeyMetadata) {
	*out = *in
	if in.AWSAccountID != nil {
		in, out := &in.AWSAccountID, &out.AWSAccountID
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CloudHsmClusterID != nil {
		in, out := &in.CloudHsmClusterID, &out.CloudHsmClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.CustomKeyStoreID != nil {
		in, out := &in.CustomKeyStoreID, &out.CustomKeyStoreID
		*out = new(string)
		**out = **in
	}
	if in.CustomerMasterKeySpec != nil {
		in, out := &in.CustomerMasterKeySpec, &out.CustomerMasterKeySpec
		*out = new(string)
		**out = **in
	}
	if in.DeletionDate != nil {
		in, out := &in.DeletionDate, &out.DeletionDate
		*out = (*in).DeepCopy()
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionAlgorithms != nil {
		in, out := &in.EncryptionAlgorithms, &out.EncryptionAlgorithms
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExpirationModel != nil {
		in, out := &in.ExpirationModel, &out.ExpirationModel
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.KeyManager != nil {
		in, out := &in.KeyManager, &out.KeyManager
		*out = new(string)
		**out = **in
	}
	if in.KeyState != nil {
		in, out := &in.KeyState, &out.KeyState
		*out = new(string)
		**out = **in
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.SigningAlgorithms != nil {
		in, out := &in.SigningAlgorithms, &out.SigningAlgorithms
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ValidTo != nil {
		in, out := &in.ValidTo, &out.ValidTo
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyMetadata.
func (in *KeyMetadata) DeepCopy() *KeyMetadata {
	if in == nil {
		return nil
	}
	out := new(KeyMetadata)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyObservation) DeepCopyInto(out *KeyObservation) {
	*out = *in
	if in.AWSAccountID != nil {
		in, out := &in.AWSAccountID, &out.AWSAccountID
		*out = new(string)
		**out = **in
	}
	if in.ARN != nil {
		in, out := &in.ARN, &out.ARN
		*out = new(string)
		**out = **in
	}
	if in.CloudHsmClusterID != nil {
		in, out := &in.CloudHsmClusterID, &out.CloudHsmClusterID
		*out = new(string)
		**out = **in
	}
	if in.CreationDate != nil {
		in, out := &in.CreationDate, &out.CreationDate
		*out = (*in).DeepCopy()
	}
	if in.DeletionDate != nil {
		in, out := &in.DeletionDate, &out.DeletionDate
		*out = (*in).DeepCopy()
	}
	if in.Enabled != nil {
		in, out := &in.Enabled, &out.Enabled
		*out = new(bool)
		**out = **in
	}
	if in.EncryptionAlgorithms != nil {
		in, out := &in.EncryptionAlgorithms, &out.EncryptionAlgorithms
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ExpirationModel != nil {
		in, out := &in.ExpirationModel, &out.ExpirationModel
		*out = new(string)
		**out = **in
	}
	if in.KeyID != nil {
		in, out := &in.KeyID, &out.KeyID
		*out = new(string)
		**out = **in
	}
	if in.KeyManager != nil {
		in, out := &in.KeyManager, &out.KeyManager
		*out = new(string)
		**out = **in
	}
	if in.KeyState != nil {
		in, out := &in.KeyState, &out.KeyState
		*out = new(string)
		**out = **in
	}
	if in.SigningAlgorithms != nil {
		in, out := &in.SigningAlgorithms, &out.SigningAlgorithms
		*out = make([]*string, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(string)
				**out = **in
			}
		}
	}
	if in.ValidTo != nil {
		in, out := &in.ValidTo, &out.ValidTo
		*out = (*in).DeepCopy()
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyObservation.
func (in *KeyObservation) DeepCopy() *KeyObservation {
	if in == nil {
		return nil
	}
	out := new(KeyObservation)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyParameters) DeepCopyInto(out *KeyParameters) {
	*out = *in
	if in.BypassPolicyLockoutSafetyCheck != nil {
		in, out := &in.BypassPolicyLockoutSafetyCheck, &out.BypassPolicyLockoutSafetyCheck
		*out = new(bool)
		**out = **in
	}
	if in.CustomKeyStoreID != nil {
		in, out := &in.CustomKeyStoreID, &out.CustomKeyStoreID
		*out = new(string)
		**out = **in
	}
	if in.CustomerMasterKeySpec != nil {
		in, out := &in.CustomerMasterKeySpec, &out.CustomerMasterKeySpec
		*out = new(string)
		**out = **in
	}
	if in.Description != nil {
		in, out := &in.Description, &out.Description
		*out = new(string)
		**out = **in
	}
	if in.KeyUsage != nil {
		in, out := &in.KeyUsage, &out.KeyUsage
		*out = new(string)
		**out = **in
	}
	if in.Origin != nil {
		in, out := &in.Origin, &out.Origin
		*out = new(string)
		**out = **in
	}
	if in.Policy != nil {
		in, out := &in.Policy, &out.Policy
		*out = new(string)
		**out = **in
	}
	if in.Tags != nil {
		in, out := &in.Tags, &out.Tags
		*out = make([]*Tag, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(Tag)
				(*in).DeepCopyInto(*out)
			}
		}
	}
	in.CustomKeyParameters.DeepCopyInto(&out.CustomKeyParameters)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyParameters.
func (in *KeyParameters) DeepCopy() *KeyParameters {
	if in == nil {
		return nil
	}
	out := new(KeyParameters)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeySpec) DeepCopyInto(out *KeySpec) {
	*out = *in
	in.ResourceSpec.DeepCopyInto(&out.ResourceSpec)
	in.ForProvider.DeepCopyInto(&out.ForProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeySpec.
func (in *KeySpec) DeepCopy() *KeySpec {
	if in == nil {
		return nil
	}
	out := new(KeySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyStatus) DeepCopyInto(out *KeyStatus) {
	*out = *in
	in.ResourceStatus.DeepCopyInto(&out.ResourceStatus)
	in.AtProvider.DeepCopyInto(&out.AtProvider)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyStatus.
func (in *KeyStatus) DeepCopy() *KeyStatus {
	if in == nil {
		return nil
	}
	out := new(KeyStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *Tag) DeepCopyInto(out *Tag) {
	*out = *in
	if in.TagKey != nil {
		in, out := &in.TagKey, &out.TagKey
		*out = new(string)
		**out = **in
	}
	if in.TagValue != nil {
		in, out := &in.TagValue, &out.TagValue
		*out = new(string)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new Tag.
func (in *Tag) DeepCopy() *Tag {
	if in == nil {
		return nil
	}
	out := new(Tag)
	in.DeepCopyInto(out)
	return out
}
