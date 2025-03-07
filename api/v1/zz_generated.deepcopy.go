// +build !ignore_autogenerated

/*
Copyright 2020 Red Hat OpenShift Container Storage.

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
	"github.com/noobaa/noobaa-operator/v2/pkg/apis/noobaa/v1alpha1"
	conditionsv1 "github.com/openshift/custom-resource-status/conditions/v1"
	ceph_rook_iov1 "github.com/rook/rook/pkg/apis/ceph.rook.io/v1"
	rook_iov1 "github.com/rook/rook/pkg/apis/rook.io/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ArbiterSpec) DeepCopyInto(out *ArbiterSpec) {
	*out = *in
	if in.ArbiterMonPVCTemplate != nil {
		in, out := &in.ArbiterMonPVCTemplate, &out.ArbiterMonPVCTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ArbiterSpec.
func (in *ArbiterSpec) DeepCopy() *ArbiterSpec {
	if in == nil {
		return nil
	}
	out := new(ArbiterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ComponentImageStatus) DeepCopyInto(out *ComponentImageStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ComponentImageStatus.
func (in *ComponentImageStatus) DeepCopy() *ComponentImageStatus {
	if in == nil {
		return nil
	}
	out := new(ComponentImageStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *EncryptionSpec) DeepCopyInto(out *EncryptionSpec) {
	*out = *in
	out.KeyManagementService = in.KeyManagementService
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new EncryptionSpec.
func (in *EncryptionSpec) DeepCopy() *EncryptionSpec {
	if in == nil {
		return nil
	}
	out := new(EncryptionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ExternalStorageClusterSpec) DeepCopyInto(out *ExternalStorageClusterSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ExternalStorageClusterSpec.
func (in *ExternalStorageClusterSpec) DeepCopy() *ExternalStorageClusterSpec {
	if in == nil {
		return nil
	}
	out := new(ExternalStorageClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ImagesStatus) DeepCopyInto(out *ImagesStatus) {
	*out = *in
	if in.Ceph != nil {
		in, out := &in.Ceph, &out.Ceph
		*out = new(ComponentImageStatus)
		**out = **in
	}
	if in.NooBaaCore != nil {
		in, out := &in.NooBaaCore, &out.NooBaaCore
		*out = new(ComponentImageStatus)
		**out = **in
	}
	if in.NooBaaDB != nil {
		in, out := &in.NooBaaDB, &out.NooBaaDB
		*out = new(ComponentImageStatus)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ImagesStatus.
func (in *ImagesStatus) DeepCopy() *ImagesStatus {
	if in == nil {
		return nil
	}
	out := new(ImagesStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *KeyManagementServiceSpec) DeepCopyInto(out *KeyManagementServiceSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new KeyManagementServiceSpec.
func (in *KeyManagementServiceSpec) DeepCopy() *KeyManagementServiceSpec {
	if in == nil {
		return nil
	}
	out := new(KeyManagementServiceSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephBlockPools) DeepCopyInto(out *ManageCephBlockPools) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephBlockPools.
func (in *ManageCephBlockPools) DeepCopy() *ManageCephBlockPools {
	if in == nil {
		return nil
	}
	out := new(ManageCephBlockPools)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephConfig) DeepCopyInto(out *ManageCephConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephConfig.
func (in *ManageCephConfig) DeepCopy() *ManageCephConfig {
	if in == nil {
		return nil
	}
	out := new(ManageCephConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephDashboard) DeepCopyInto(out *ManageCephDashboard) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephDashboard.
func (in *ManageCephDashboard) DeepCopy() *ManageCephDashboard {
	if in == nil {
		return nil
	}
	out := new(ManageCephDashboard)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephFilesystems) DeepCopyInto(out *ManageCephFilesystems) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephFilesystems.
func (in *ManageCephFilesystems) DeepCopy() *ManageCephFilesystems {
	if in == nil {
		return nil
	}
	out := new(ManageCephFilesystems)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephObjectStoreUsers) DeepCopyInto(out *ManageCephObjectStoreUsers) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephObjectStoreUsers.
func (in *ManageCephObjectStoreUsers) DeepCopy() *ManageCephObjectStoreUsers {
	if in == nil {
		return nil
	}
	out := new(ManageCephObjectStoreUsers)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManageCephObjectStores) DeepCopyInto(out *ManageCephObjectStores) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManageCephObjectStores.
func (in *ManageCephObjectStores) DeepCopy() *ManageCephObjectStores {
	if in == nil {
		return nil
	}
	out := new(ManageCephObjectStores)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ManagedResourcesSpec) DeepCopyInto(out *ManagedResourcesSpec) {
	*out = *in
	out.CephConfig = in.CephConfig
	out.CephDashboard = in.CephDashboard
	out.CephBlockPools = in.CephBlockPools
	out.CephFilesystems = in.CephFilesystems
	out.CephObjectStores = in.CephObjectStores
	out.CephObjectStoreUsers = in.CephObjectStoreUsers
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ManagedResourcesSpec.
func (in *ManagedResourcesSpec) DeepCopy() *ManagedResourcesSpec {
	if in == nil {
		return nil
	}
	out := new(ManagedResourcesSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MirroringSpec) DeepCopyInto(out *MirroringSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MirroringSpec.
func (in *MirroringSpec) DeepCopy() *MirroringSpec {
	if in == nil {
		return nil
	}
	out := new(MirroringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MonitoringSpec) DeepCopyInto(out *MonitoringSpec) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MonitoringSpec.
func (in *MonitoringSpec) DeepCopy() *MonitoringSpec {
	if in == nil {
		return nil
	}
	out := new(MonitoringSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *MultiCloudGatewaySpec) DeepCopyInto(out *MultiCloudGatewaySpec) {
	*out = *in
	if in.Endpoints != nil {
		in, out := &in.Endpoints, &out.Endpoints
		*out = new(v1alpha1.EndpointsSpec)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new MultiCloudGatewaySpec.
func (in *MultiCloudGatewaySpec) DeepCopy() *MultiCloudGatewaySpec {
	if in == nil {
		return nil
	}
	out := new(MultiCloudGatewaySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *NodeTopologyMap) DeepCopyInto(out *NodeTopologyMap) {
	*out = *in
	if in.Labels != nil {
		in, out := &in.Labels, &out.Labels
		*out = make(map[string]TopologyLabelValues, len(*in))
		for key, val := range *in {
			var outVal []string
			if val == nil {
				(*out)[key] = nil
			} else {
				in, out := &val, &outVal
				*out = make(TopologyLabelValues, len(*in))
				copy(*out, *in)
			}
			(*out)[key] = outVal
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new NodeTopologyMap.
func (in *NodeTopologyMap) DeepCopy() *NodeTopologyMap {
	if in == nil {
		return nil
	}
	out := new(NodeTopologyMap)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitialization) DeepCopyInto(out *OCSInitialization) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitialization.
func (in *OCSInitialization) DeepCopy() *OCSInitialization {
	if in == nil {
		return nil
	}
	out := new(OCSInitialization)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCSInitialization) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationList) DeepCopyInto(out *OCSInitializationList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OCSInitialization, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationList.
func (in *OCSInitializationList) DeepCopy() *OCSInitializationList {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OCSInitializationList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationSpec) DeepCopyInto(out *OCSInitializationSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationSpec.
func (in *OCSInitializationSpec) DeepCopy() *OCSInitializationSpec {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OCSInitializationStatus) DeepCopyInto(out *OCSInitializationStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OCSInitializationStatus.
func (in *OCSInitializationStatus) DeepCopy() *OCSInitializationStatus {
	if in == nil {
		return nil
	}
	out := new(OCSInitializationStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageCluster) DeepCopyInto(out *StorageCluster) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageCluster.
func (in *StorageCluster) DeepCopy() *StorageCluster {
	if in == nil {
		return nil
	}
	out := new(StorageCluster)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageCluster) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterList) DeepCopyInto(out *StorageClusterList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]StorageCluster, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterList.
func (in *StorageClusterList) DeepCopy() *StorageClusterList {
	if in == nil {
		return nil
	}
	out := new(StorageClusterList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *StorageClusterList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterSpec) DeepCopyInto(out *StorageClusterSpec) {
	*out = *in
	if in.LabelSelector != nil {
		in, out := &in.LabelSelector, &out.LabelSelector
		*out = new(metav1.LabelSelector)
		(*in).DeepCopyInto(*out)
	}
	out.ExternalStorage = in.ExternalStorage
	if in.Placement != nil {
		in, out := &in.Placement, &out.Placement
		*out = make(rook_iov1.PlacementSpec, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	if in.Resources != nil {
		in, out := &in.Resources, &out.Resources
		*out = make(map[string]corev1.ResourceRequirements, len(*in))
		for key, val := range *in {
			(*out)[key] = *val.DeepCopy()
		}
	}
	out.Encryption = in.Encryption
	if in.StorageDeviceSets != nil {
		in, out := &in.StorageDeviceSets, &out.StorageDeviceSets
		*out = make([]StorageDeviceSet, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.MonPVCTemplate != nil {
		in, out := &in.MonPVCTemplate, &out.MonPVCTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.MultiCloudGateway != nil {
		in, out := &in.MultiCloudGateway, &out.MultiCloudGateway
		*out = new(MultiCloudGatewaySpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Monitoring != nil {
		in, out := &in.Monitoring, &out.Monitoring
		*out = new(MonitoringSpec)
		(*in).DeepCopyInto(*out)
	}
	if in.Network != nil {
		in, out := &in.Network, &out.Network
		*out = new(ceph_rook_iov1.NetworkSpec)
		(*in).DeepCopyInto(*out)
	}
	out.ManagedResources = in.ManagedResources
	if in.NodeTopologies != nil {
		in, out := &in.NodeTopologies, &out.NodeTopologies
		*out = new(NodeTopologyMap)
		(*in).DeepCopyInto(*out)
	}
	in.Arbiter.DeepCopyInto(&out.Arbiter)
	out.Mirroring = in.Mirroring
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterSpec.
func (in *StorageClusterSpec) DeepCopy() *StorageClusterSpec {
	if in == nil {
		return nil
	}
	out := new(StorageClusterSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageClusterStatus) DeepCopyInto(out *StorageClusterStatus) {
	*out = *in
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]conditionsv1.Condition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.RelatedObjects != nil {
		in, out := &in.RelatedObjects, &out.RelatedObjects
		*out = make([]corev1.ObjectReference, len(*in))
		copy(*out, *in)
	}
	if in.NodeTopologies != nil {
		in, out := &in.NodeTopologies, &out.NodeTopologies
		*out = new(NodeTopologyMap)
		(*in).DeepCopyInto(*out)
	}
	if in.FailureDomainValues != nil {
		in, out := &in.FailureDomainValues, &out.FailureDomainValues
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	in.Images.DeepCopyInto(&out.Images)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageClusterStatus.
func (in *StorageClusterStatus) DeepCopy() *StorageClusterStatus {
	if in == nil {
		return nil
	}
	out := new(StorageClusterStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDeviceSet) DeepCopyInto(out *StorageDeviceSet) {
	*out = *in
	in.Resources.DeepCopyInto(&out.Resources)
	in.PreparePlacement.DeepCopyInto(&out.PreparePlacement)
	in.Placement.DeepCopyInto(&out.Placement)
	out.Config = in.Config
	in.DataPVCTemplate.DeepCopyInto(&out.DataPVCTemplate)
	if in.MetadataPVCTemplate != nil {
		in, out := &in.MetadataPVCTemplate, &out.MetadataPVCTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
	if in.WalPVCTemplate != nil {
		in, out := &in.WalPVCTemplate, &out.WalPVCTemplate
		*out = new(corev1.PersistentVolumeClaim)
		(*in).DeepCopyInto(*out)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDeviceSet.
func (in *StorageDeviceSet) DeepCopy() *StorageDeviceSet {
	if in == nil {
		return nil
	}
	out := new(StorageDeviceSet)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *StorageDeviceSetConfig) DeepCopyInto(out *StorageDeviceSetConfig) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new StorageDeviceSetConfig.
func (in *StorageDeviceSetConfig) DeepCopy() *StorageDeviceSetConfig {
	if in == nil {
		return nil
	}
	out := new(StorageDeviceSetConfig)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in TopologyLabelValues) DeepCopyInto(out *TopologyLabelValues) {
	{
		in := &in
		*out = make(TopologyLabelValues, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new TopologyLabelValues.
func (in TopologyLabelValues) DeepCopy() TopologyLabelValues {
	if in == nil {
		return nil
	}
	out := new(TopologyLabelValues)
	in.DeepCopyInto(out)
	return *out
}
