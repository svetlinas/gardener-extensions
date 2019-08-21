// +build !ignore_autogenerated

/*
Copyright (c) 2019 SAP SE or an SAP affiliate company. All rights reserved. This file is licensed under the Apache Software License, v. 2 except as noted otherwise in the LICENSE file

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	unsafe "unsafe"

	azure "github.com/gardener/gardener-extensions/controllers/provider-azure/pkg/apis/azure"
	corev1alpha1 "github.com/gardener/gardener/pkg/apis/core/v1alpha1"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*AvailabilitySet)(nil), (*azure.AvailabilitySet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_AvailabilitySet_To_azure_AvailabilitySet(a.(*AvailabilitySet), b.(*azure.AvailabilitySet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.AvailabilitySet)(nil), (*AvailabilitySet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_AvailabilitySet_To_v1alpha1_AvailabilitySet(a.(*azure.AvailabilitySet), b.(*AvailabilitySet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*CloudControllerManagerConfig)(nil), (*azure.CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_CloudControllerManagerConfig_To_azure_CloudControllerManagerConfig(a.(*CloudControllerManagerConfig), b.(*azure.CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.CloudControllerManagerConfig)(nil), (*CloudControllerManagerConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(a.(*azure.CloudControllerManagerConfig), b.(*CloudControllerManagerConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ControlPlaneConfig)(nil), (*azure.ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ControlPlaneConfig_To_azure_ControlPlaneConfig(a.(*ControlPlaneConfig), b.(*azure.ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.ControlPlaneConfig)(nil), (*ControlPlaneConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(a.(*azure.ControlPlaneConfig), b.(*ControlPlaneConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*DomainCount)(nil), (*azure.DomainCount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_DomainCount_To_azure_DomainCount(a.(*DomainCount), b.(*azure.DomainCount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.DomainCount)(nil), (*DomainCount)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_DomainCount_To_v1alpha1_DomainCount(a.(*azure.DomainCount), b.(*DomainCount), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureConfig)(nil), (*azure.InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureConfig_To_azure_InfrastructureConfig(a.(*InfrastructureConfig), b.(*azure.InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.InfrastructureConfig)(nil), (*InfrastructureConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(a.(*azure.InfrastructureConfig), b.(*InfrastructureConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*InfrastructureStatus)(nil), (*azure.InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_InfrastructureStatus_To_azure_InfrastructureStatus(a.(*InfrastructureStatus), b.(*azure.InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.InfrastructureStatus)(nil), (*InfrastructureStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(a.(*azure.InfrastructureStatus), b.(*InfrastructureStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImage)(nil), (*azure.MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImage_To_azure_MachineImage(a.(*MachineImage), b.(*azure.MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.MachineImage)(nil), (*MachineImage)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_MachineImage_To_v1alpha1_MachineImage(a.(*azure.MachineImage), b.(*MachineImage), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImageVersion)(nil), (*azure.MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImageVersion_To_azure_MachineImageVersion(a.(*MachineImageVersion), b.(*azure.MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.MachineImageVersion)(nil), (*MachineImageVersion)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_MachineImageVersion_To_v1alpha1_MachineImageVersion(a.(*azure.MachineImageVersion), b.(*MachineImageVersion), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*MachineImages)(nil), (*azure.MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_MachineImages_To_azure_MachineImages(a.(*MachineImages), b.(*azure.MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.MachineImages)(nil), (*MachineImages)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_MachineImages_To_v1alpha1_MachineImages(a.(*azure.MachineImages), b.(*MachineImages), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkConfig)(nil), (*azure.NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkConfig_To_azure_NetworkConfig(a.(*NetworkConfig), b.(*azure.NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.NetworkConfig)(nil), (*NetworkConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_NetworkConfig_To_v1alpha1_NetworkConfig(a.(*azure.NetworkConfig), b.(*NetworkConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*NetworkStatus)(nil), (*azure.NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_NetworkStatus_To_azure_NetworkStatus(a.(*NetworkStatus), b.(*azure.NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.NetworkStatus)(nil), (*NetworkStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_NetworkStatus_To_v1alpha1_NetworkStatus(a.(*azure.NetworkStatus), b.(*NetworkStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ProviderProfileConfig)(nil), (*azure.ProviderProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ProviderProfileConfig_To_azure_ProviderProfileConfig(a.(*ProviderProfileConfig), b.(*azure.ProviderProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.ProviderProfileConfig)(nil), (*ProviderProfileConfig)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_ProviderProfileConfig_To_v1alpha1_ProviderProfileConfig(a.(*azure.ProviderProfileConfig), b.(*ProviderProfileConfig), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*ResourceGroup)(nil), (*azure.ResourceGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ResourceGroup_To_azure_ResourceGroup(a.(*ResourceGroup), b.(*azure.ResourceGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.ResourceGroup)(nil), (*ResourceGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_ResourceGroup_To_v1alpha1_ResourceGroup(a.(*azure.ResourceGroup), b.(*ResourceGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*RouteTable)(nil), (*azure.RouteTable)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_RouteTable_To_azure_RouteTable(a.(*RouteTable), b.(*azure.RouteTable), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.RouteTable)(nil), (*RouteTable)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_RouteTable_To_v1alpha1_RouteTable(a.(*azure.RouteTable), b.(*RouteTable), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*SecurityGroup)(nil), (*azure.SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_SecurityGroup_To_azure_SecurityGroup(a.(*SecurityGroup), b.(*azure.SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.SecurityGroup)(nil), (*SecurityGroup)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_SecurityGroup_To_v1alpha1_SecurityGroup(a.(*azure.SecurityGroup), b.(*SecurityGroup), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*Subnet)(nil), (*azure.Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_Subnet_To_azure_Subnet(a.(*Subnet), b.(*azure.Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.Subnet)(nil), (*Subnet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_Subnet_To_v1alpha1_Subnet(a.(*azure.Subnet), b.(*Subnet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VNet)(nil), (*azure.VNet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VNet_To_azure_VNet(a.(*VNet), b.(*azure.VNet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.VNet)(nil), (*VNet)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_VNet_To_v1alpha1_VNet(a.(*azure.VNet), b.(*VNet), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*VNetStatus)(nil), (*azure.VNetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_VNetStatus_To_azure_VNetStatus(a.(*VNetStatus), b.(*azure.VNetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.VNetStatus)(nil), (*VNetStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_VNetStatus_To_v1alpha1_VNetStatus(a.(*azure.VNetStatus), b.(*VNetStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*WorkerStatus)(nil), (*azure.WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_WorkerStatus_To_azure_WorkerStatus(a.(*WorkerStatus), b.(*azure.WorkerStatus), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*azure.WorkerStatus)(nil), (*WorkerStatus)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_azure_WorkerStatus_To_v1alpha1_WorkerStatus(a.(*azure.WorkerStatus), b.(*WorkerStatus), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_AvailabilitySet_To_azure_AvailabilitySet(in *AvailabilitySet, out *azure.AvailabilitySet, s conversion.Scope) error {
	out.Purpose = azure.Purpose(in.Purpose)
	out.ID = in.ID
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_AvailabilitySet_To_azure_AvailabilitySet is an autogenerated conversion function.
func Convert_v1alpha1_AvailabilitySet_To_azure_AvailabilitySet(in *AvailabilitySet, out *azure.AvailabilitySet, s conversion.Scope) error {
	return autoConvert_v1alpha1_AvailabilitySet_To_azure_AvailabilitySet(in, out, s)
}

func autoConvert_azure_AvailabilitySet_To_v1alpha1_AvailabilitySet(in *azure.AvailabilitySet, out *AvailabilitySet, s conversion.Scope) error {
	out.Purpose = Purpose(in.Purpose)
	out.ID = in.ID
	out.Name = in.Name
	return nil
}

// Convert_azure_AvailabilitySet_To_v1alpha1_AvailabilitySet is an autogenerated conversion function.
func Convert_azure_AvailabilitySet_To_v1alpha1_AvailabilitySet(in *azure.AvailabilitySet, out *AvailabilitySet, s conversion.Scope) error {
	return autoConvert_azure_AvailabilitySet_To_v1alpha1_AvailabilitySet(in, out, s)
}

func autoConvert_v1alpha1_CloudControllerManagerConfig_To_azure_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *azure.CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_v1alpha1_CloudControllerManagerConfig_To_azure_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_v1alpha1_CloudControllerManagerConfig_To_azure_CloudControllerManagerConfig(in *CloudControllerManagerConfig, out *azure.CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_CloudControllerManagerConfig_To_azure_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_azure_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *azure.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	out.FeatureGates = *(*map[string]bool)(unsafe.Pointer(&in.FeatureGates))
	return nil
}

// Convert_azure_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig is an autogenerated conversion function.
func Convert_azure_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in *azure.CloudControllerManagerConfig, out *CloudControllerManagerConfig, s conversion.Scope) error {
	return autoConvert_azure_CloudControllerManagerConfig_To_v1alpha1_CloudControllerManagerConfig(in, out, s)
}

func autoConvert_v1alpha1_ControlPlaneConfig_To_azure_ControlPlaneConfig(in *ControlPlaneConfig, out *azure.ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*azure.CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_v1alpha1_ControlPlaneConfig_To_azure_ControlPlaneConfig is an autogenerated conversion function.
func Convert_v1alpha1_ControlPlaneConfig_To_azure_ControlPlaneConfig(in *ControlPlaneConfig, out *azure.ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ControlPlaneConfig_To_azure_ControlPlaneConfig(in, out, s)
}

func autoConvert_azure_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *azure.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	out.CloudControllerManager = (*CloudControllerManagerConfig)(unsafe.Pointer(in.CloudControllerManager))
	return nil
}

// Convert_azure_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig is an autogenerated conversion function.
func Convert_azure_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in *azure.ControlPlaneConfig, out *ControlPlaneConfig, s conversion.Scope) error {
	return autoConvert_azure_ControlPlaneConfig_To_v1alpha1_ControlPlaneConfig(in, out, s)
}

func autoConvert_v1alpha1_DomainCount_To_azure_DomainCount(in *DomainCount, out *azure.DomainCount, s conversion.Scope) error {
	out.Region = in.Region
	out.Count = in.Count
	return nil
}

// Convert_v1alpha1_DomainCount_To_azure_DomainCount is an autogenerated conversion function.
func Convert_v1alpha1_DomainCount_To_azure_DomainCount(in *DomainCount, out *azure.DomainCount, s conversion.Scope) error {
	return autoConvert_v1alpha1_DomainCount_To_azure_DomainCount(in, out, s)
}

func autoConvert_azure_DomainCount_To_v1alpha1_DomainCount(in *azure.DomainCount, out *DomainCount, s conversion.Scope) error {
	out.Region = in.Region
	out.Count = in.Count
	return nil
}

// Convert_azure_DomainCount_To_v1alpha1_DomainCount is an autogenerated conversion function.
func Convert_azure_DomainCount_To_v1alpha1_DomainCount(in *azure.DomainCount, out *DomainCount, s conversion.Scope) error {
	return autoConvert_azure_DomainCount_To_v1alpha1_DomainCount(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureConfig_To_azure_InfrastructureConfig(in *InfrastructureConfig, out *azure.InfrastructureConfig, s conversion.Scope) error {
	out.ResourceGroup = (*azure.ResourceGroup)(unsafe.Pointer(in.ResourceGroup))
	if err := Convert_v1alpha1_NetworkConfig_To_azure_NetworkConfig(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_v1alpha1_InfrastructureConfig_To_azure_InfrastructureConfig is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureConfig_To_azure_InfrastructureConfig(in *InfrastructureConfig, out *azure.InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureConfig_To_azure_InfrastructureConfig(in, out, s)
}

func autoConvert_azure_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *azure.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	out.ResourceGroup = (*ResourceGroup)(unsafe.Pointer(in.ResourceGroup))
	if err := Convert_azure_NetworkConfig_To_v1alpha1_NetworkConfig(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	return nil
}

// Convert_azure_InfrastructureConfig_To_v1alpha1_InfrastructureConfig is an autogenerated conversion function.
func Convert_azure_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in *azure.InfrastructureConfig, out *InfrastructureConfig, s conversion.Scope) error {
	return autoConvert_azure_InfrastructureConfig_To_v1alpha1_InfrastructureConfig(in, out, s)
}

func autoConvert_v1alpha1_InfrastructureStatus_To_azure_InfrastructureStatus(in *InfrastructureStatus, out *azure.InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_NetworkStatus_To_azure_NetworkStatus(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	if err := Convert_v1alpha1_ResourceGroup_To_azure_ResourceGroup(&in.ResourceGroup, &out.ResourceGroup, s); err != nil {
		return err
	}
	out.AvailabilitySets = *(*[]azure.AvailabilitySet)(unsafe.Pointer(&in.AvailabilitySets))
	out.RouteTables = *(*[]azure.RouteTable)(unsafe.Pointer(&in.RouteTables))
	out.SecurityGroups = *(*[]azure.SecurityGroup)(unsafe.Pointer(&in.SecurityGroups))
	return nil
}

// Convert_v1alpha1_InfrastructureStatus_To_azure_InfrastructureStatus is an autogenerated conversion function.
func Convert_v1alpha1_InfrastructureStatus_To_azure_InfrastructureStatus(in *InfrastructureStatus, out *azure.InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_InfrastructureStatus_To_azure_InfrastructureStatus(in, out, s)
}

func autoConvert_azure_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *azure.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	if err := Convert_azure_NetworkStatus_To_v1alpha1_NetworkStatus(&in.Networks, &out.Networks, s); err != nil {
		return err
	}
	if err := Convert_azure_ResourceGroup_To_v1alpha1_ResourceGroup(&in.ResourceGroup, &out.ResourceGroup, s); err != nil {
		return err
	}
	out.AvailabilitySets = *(*[]AvailabilitySet)(unsafe.Pointer(&in.AvailabilitySets))
	out.RouteTables = *(*[]RouteTable)(unsafe.Pointer(&in.RouteTables))
	out.SecurityGroups = *(*[]SecurityGroup)(unsafe.Pointer(&in.SecurityGroups))
	return nil
}

// Convert_azure_InfrastructureStatus_To_v1alpha1_InfrastructureStatus is an autogenerated conversion function.
func Convert_azure_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in *azure.InfrastructureStatus, out *InfrastructureStatus, s conversion.Scope) error {
	return autoConvert_azure_InfrastructureStatus_To_v1alpha1_InfrastructureStatus(in, out, s)
}

func autoConvert_v1alpha1_MachineImage_To_azure_MachineImage(in *MachineImage, out *azure.MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Publisher = in.Publisher
	out.Offer = in.Offer
	out.SKU = in.SKU
	return nil
}

// Convert_v1alpha1_MachineImage_To_azure_MachineImage is an autogenerated conversion function.
func Convert_v1alpha1_MachineImage_To_azure_MachineImage(in *MachineImage, out *azure.MachineImage, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImage_To_azure_MachineImage(in, out, s)
}

func autoConvert_azure_MachineImage_To_v1alpha1_MachineImage(in *azure.MachineImage, out *MachineImage, s conversion.Scope) error {
	out.Name = in.Name
	out.Version = in.Version
	out.Publisher = in.Publisher
	out.Offer = in.Offer
	out.SKU = in.SKU
	return nil
}

// Convert_azure_MachineImage_To_v1alpha1_MachineImage is an autogenerated conversion function.
func Convert_azure_MachineImage_To_v1alpha1_MachineImage(in *azure.MachineImage, out *MachineImage, s conversion.Scope) error {
	return autoConvert_azure_MachineImage_To_v1alpha1_MachineImage(in, out, s)
}

func autoConvert_v1alpha1_MachineImageVersion_To_azure_MachineImageVersion(in *MachineImageVersion, out *azure.MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.URN = in.URN
	return nil
}

// Convert_v1alpha1_MachineImageVersion_To_azure_MachineImageVersion is an autogenerated conversion function.
func Convert_v1alpha1_MachineImageVersion_To_azure_MachineImageVersion(in *MachineImageVersion, out *azure.MachineImageVersion, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImageVersion_To_azure_MachineImageVersion(in, out, s)
}

func autoConvert_azure_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *azure.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	out.Version = in.Version
	out.URN = in.URN
	return nil
}

// Convert_azure_MachineImageVersion_To_v1alpha1_MachineImageVersion is an autogenerated conversion function.
func Convert_azure_MachineImageVersion_To_v1alpha1_MachineImageVersion(in *azure.MachineImageVersion, out *MachineImageVersion, s conversion.Scope) error {
	return autoConvert_azure_MachineImageVersion_To_v1alpha1_MachineImageVersion(in, out, s)
}

func autoConvert_v1alpha1_MachineImages_To_azure_MachineImages(in *MachineImages, out *azure.MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]azure.MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_v1alpha1_MachineImages_To_azure_MachineImages is an autogenerated conversion function.
func Convert_v1alpha1_MachineImages_To_azure_MachineImages(in *MachineImages, out *azure.MachineImages, s conversion.Scope) error {
	return autoConvert_v1alpha1_MachineImages_To_azure_MachineImages(in, out, s)
}

func autoConvert_azure_MachineImages_To_v1alpha1_MachineImages(in *azure.MachineImages, out *MachineImages, s conversion.Scope) error {
	out.Name = in.Name
	out.Versions = *(*[]MachineImageVersion)(unsafe.Pointer(&in.Versions))
	return nil
}

// Convert_azure_MachineImages_To_v1alpha1_MachineImages is an autogenerated conversion function.
func Convert_azure_MachineImages_To_v1alpha1_MachineImages(in *azure.MachineImages, out *MachineImages, s conversion.Scope) error {
	return autoConvert_azure_MachineImages_To_v1alpha1_MachineImages(in, out, s)
}

func autoConvert_v1alpha1_NetworkConfig_To_azure_NetworkConfig(in *NetworkConfig, out *azure.NetworkConfig, s conversion.Scope) error {
	if err := Convert_v1alpha1_VNet_To_azure_VNet(&in.VNet, &out.VNet, s); err != nil {
		return err
	}
	out.Workers = corev1alpha1.CIDR(in.Workers)
	return nil
}

// Convert_v1alpha1_NetworkConfig_To_azure_NetworkConfig is an autogenerated conversion function.
func Convert_v1alpha1_NetworkConfig_To_azure_NetworkConfig(in *NetworkConfig, out *azure.NetworkConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkConfig_To_azure_NetworkConfig(in, out, s)
}

func autoConvert_azure_NetworkConfig_To_v1alpha1_NetworkConfig(in *azure.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	if err := Convert_azure_VNet_To_v1alpha1_VNet(&in.VNet, &out.VNet, s); err != nil {
		return err
	}
	out.Workers = corev1alpha1.CIDR(in.Workers)
	return nil
}

// Convert_azure_NetworkConfig_To_v1alpha1_NetworkConfig is an autogenerated conversion function.
func Convert_azure_NetworkConfig_To_v1alpha1_NetworkConfig(in *azure.NetworkConfig, out *NetworkConfig, s conversion.Scope) error {
	return autoConvert_azure_NetworkConfig_To_v1alpha1_NetworkConfig(in, out, s)
}

func autoConvert_v1alpha1_NetworkStatus_To_azure_NetworkStatus(in *NetworkStatus, out *azure.NetworkStatus, s conversion.Scope) error {
	if err := Convert_v1alpha1_VNetStatus_To_azure_VNetStatus(&in.VNet, &out.VNet, s); err != nil {
		return err
	}
	out.Subnets = *(*[]azure.Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_v1alpha1_NetworkStatus_To_azure_NetworkStatus is an autogenerated conversion function.
func Convert_v1alpha1_NetworkStatus_To_azure_NetworkStatus(in *NetworkStatus, out *azure.NetworkStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_NetworkStatus_To_azure_NetworkStatus(in, out, s)
}

func autoConvert_azure_NetworkStatus_To_v1alpha1_NetworkStatus(in *azure.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	if err := Convert_azure_VNetStatus_To_v1alpha1_VNetStatus(&in.VNet, &out.VNet, s); err != nil {
		return err
	}
	out.Subnets = *(*[]Subnet)(unsafe.Pointer(&in.Subnets))
	return nil
}

// Convert_azure_NetworkStatus_To_v1alpha1_NetworkStatus is an autogenerated conversion function.
func Convert_azure_NetworkStatus_To_v1alpha1_NetworkStatus(in *azure.NetworkStatus, out *NetworkStatus, s conversion.Scope) error {
	return autoConvert_azure_NetworkStatus_To_v1alpha1_NetworkStatus(in, out, s)
}

func autoConvert_v1alpha1_ProviderProfileConfig_To_azure_ProviderProfileConfig(in *ProviderProfileConfig, out *azure.ProviderProfileConfig, s conversion.Scope) error {
	out.CountUpdateDomains = *(*[]azure.DomainCount)(unsafe.Pointer(&in.CountUpdateDomains))
	out.CountFaultDomains = *(*[]azure.DomainCount)(unsafe.Pointer(&in.CountFaultDomains))
	out.MachineImages = *(*[]azure.MachineImages)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_ProviderProfileConfig_To_azure_ProviderProfileConfig is an autogenerated conversion function.
func Convert_v1alpha1_ProviderProfileConfig_To_azure_ProviderProfileConfig(in *ProviderProfileConfig, out *azure.ProviderProfileConfig, s conversion.Scope) error {
	return autoConvert_v1alpha1_ProviderProfileConfig_To_azure_ProviderProfileConfig(in, out, s)
}

func autoConvert_azure_ProviderProfileConfig_To_v1alpha1_ProviderProfileConfig(in *azure.ProviderProfileConfig, out *ProviderProfileConfig, s conversion.Scope) error {
	out.CountUpdateDomains = *(*[]DomainCount)(unsafe.Pointer(&in.CountUpdateDomains))
	out.CountFaultDomains = *(*[]DomainCount)(unsafe.Pointer(&in.CountFaultDomains))
	out.MachineImages = *(*[]MachineImages)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_azure_ProviderProfileConfig_To_v1alpha1_ProviderProfileConfig is an autogenerated conversion function.
func Convert_azure_ProviderProfileConfig_To_v1alpha1_ProviderProfileConfig(in *azure.ProviderProfileConfig, out *ProviderProfileConfig, s conversion.Scope) error {
	return autoConvert_azure_ProviderProfileConfig_To_v1alpha1_ProviderProfileConfig(in, out, s)
}

func autoConvert_v1alpha1_ResourceGroup_To_azure_ResourceGroup(in *ResourceGroup, out *azure.ResourceGroup, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_ResourceGroup_To_azure_ResourceGroup is an autogenerated conversion function.
func Convert_v1alpha1_ResourceGroup_To_azure_ResourceGroup(in *ResourceGroup, out *azure.ResourceGroup, s conversion.Scope) error {
	return autoConvert_v1alpha1_ResourceGroup_To_azure_ResourceGroup(in, out, s)
}

func autoConvert_azure_ResourceGroup_To_v1alpha1_ResourceGroup(in *azure.ResourceGroup, out *ResourceGroup, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_azure_ResourceGroup_To_v1alpha1_ResourceGroup is an autogenerated conversion function.
func Convert_azure_ResourceGroup_To_v1alpha1_ResourceGroup(in *azure.ResourceGroup, out *ResourceGroup, s conversion.Scope) error {
	return autoConvert_azure_ResourceGroup_To_v1alpha1_ResourceGroup(in, out, s)
}

func autoConvert_v1alpha1_RouteTable_To_azure_RouteTable(in *RouteTable, out *azure.RouteTable, s conversion.Scope) error {
	out.Purpose = azure.Purpose(in.Purpose)
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_RouteTable_To_azure_RouteTable is an autogenerated conversion function.
func Convert_v1alpha1_RouteTable_To_azure_RouteTable(in *RouteTable, out *azure.RouteTable, s conversion.Scope) error {
	return autoConvert_v1alpha1_RouteTable_To_azure_RouteTable(in, out, s)
}

func autoConvert_azure_RouteTable_To_v1alpha1_RouteTable(in *azure.RouteTable, out *RouteTable, s conversion.Scope) error {
	out.Purpose = Purpose(in.Purpose)
	out.Name = in.Name
	return nil
}

// Convert_azure_RouteTable_To_v1alpha1_RouteTable is an autogenerated conversion function.
func Convert_azure_RouteTable_To_v1alpha1_RouteTable(in *azure.RouteTable, out *RouteTable, s conversion.Scope) error {
	return autoConvert_azure_RouteTable_To_v1alpha1_RouteTable(in, out, s)
}

func autoConvert_v1alpha1_SecurityGroup_To_azure_SecurityGroup(in *SecurityGroup, out *azure.SecurityGroup, s conversion.Scope) error {
	out.Purpose = azure.Purpose(in.Purpose)
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_SecurityGroup_To_azure_SecurityGroup is an autogenerated conversion function.
func Convert_v1alpha1_SecurityGroup_To_azure_SecurityGroup(in *SecurityGroup, out *azure.SecurityGroup, s conversion.Scope) error {
	return autoConvert_v1alpha1_SecurityGroup_To_azure_SecurityGroup(in, out, s)
}

func autoConvert_azure_SecurityGroup_To_v1alpha1_SecurityGroup(in *azure.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	out.Purpose = Purpose(in.Purpose)
	out.Name = in.Name
	return nil
}

// Convert_azure_SecurityGroup_To_v1alpha1_SecurityGroup is an autogenerated conversion function.
func Convert_azure_SecurityGroup_To_v1alpha1_SecurityGroup(in *azure.SecurityGroup, out *SecurityGroup, s conversion.Scope) error {
	return autoConvert_azure_SecurityGroup_To_v1alpha1_SecurityGroup(in, out, s)
}

func autoConvert_v1alpha1_Subnet_To_azure_Subnet(in *Subnet, out *azure.Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.Purpose = azure.Purpose(in.Purpose)
	return nil
}

// Convert_v1alpha1_Subnet_To_azure_Subnet is an autogenerated conversion function.
func Convert_v1alpha1_Subnet_To_azure_Subnet(in *Subnet, out *azure.Subnet, s conversion.Scope) error {
	return autoConvert_v1alpha1_Subnet_To_azure_Subnet(in, out, s)
}

func autoConvert_azure_Subnet_To_v1alpha1_Subnet(in *azure.Subnet, out *Subnet, s conversion.Scope) error {
	out.Name = in.Name
	out.Purpose = Purpose(in.Purpose)
	return nil
}

// Convert_azure_Subnet_To_v1alpha1_Subnet is an autogenerated conversion function.
func Convert_azure_Subnet_To_v1alpha1_Subnet(in *azure.Subnet, out *Subnet, s conversion.Scope) error {
	return autoConvert_azure_Subnet_To_v1alpha1_Subnet(in, out, s)
}

func autoConvert_v1alpha1_VNet_To_azure_VNet(in *VNet, out *azure.VNet, s conversion.Scope) error {
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.CIDR = (*corev1alpha1.CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_v1alpha1_VNet_To_azure_VNet is an autogenerated conversion function.
func Convert_v1alpha1_VNet_To_azure_VNet(in *VNet, out *azure.VNet, s conversion.Scope) error {
	return autoConvert_v1alpha1_VNet_To_azure_VNet(in, out, s)
}

func autoConvert_azure_VNet_To_v1alpha1_VNet(in *azure.VNet, out *VNet, s conversion.Scope) error {
	out.Name = (*string)(unsafe.Pointer(in.Name))
	out.CIDR = (*corev1alpha1.CIDR)(unsafe.Pointer(in.CIDR))
	return nil
}

// Convert_azure_VNet_To_v1alpha1_VNet is an autogenerated conversion function.
func Convert_azure_VNet_To_v1alpha1_VNet(in *azure.VNet, out *VNet, s conversion.Scope) error {
	return autoConvert_azure_VNet_To_v1alpha1_VNet(in, out, s)
}

func autoConvert_v1alpha1_VNetStatus_To_azure_VNetStatus(in *VNetStatus, out *azure.VNetStatus, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_v1alpha1_VNetStatus_To_azure_VNetStatus is an autogenerated conversion function.
func Convert_v1alpha1_VNetStatus_To_azure_VNetStatus(in *VNetStatus, out *azure.VNetStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_VNetStatus_To_azure_VNetStatus(in, out, s)
}

func autoConvert_azure_VNetStatus_To_v1alpha1_VNetStatus(in *azure.VNetStatus, out *VNetStatus, s conversion.Scope) error {
	out.Name = in.Name
	return nil
}

// Convert_azure_VNetStatus_To_v1alpha1_VNetStatus is an autogenerated conversion function.
func Convert_azure_VNetStatus_To_v1alpha1_VNetStatus(in *azure.VNetStatus, out *VNetStatus, s conversion.Scope) error {
	return autoConvert_azure_VNetStatus_To_v1alpha1_VNetStatus(in, out, s)
}

func autoConvert_v1alpha1_WorkerStatus_To_azure_WorkerStatus(in *WorkerStatus, out *azure.WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]azure.MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_v1alpha1_WorkerStatus_To_azure_WorkerStatus is an autogenerated conversion function.
func Convert_v1alpha1_WorkerStatus_To_azure_WorkerStatus(in *WorkerStatus, out *azure.WorkerStatus, s conversion.Scope) error {
	return autoConvert_v1alpha1_WorkerStatus_To_azure_WorkerStatus(in, out, s)
}

func autoConvert_azure_WorkerStatus_To_v1alpha1_WorkerStatus(in *azure.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	out.MachineImages = *(*[]MachineImage)(unsafe.Pointer(&in.MachineImages))
	return nil
}

// Convert_azure_WorkerStatus_To_v1alpha1_WorkerStatus is an autogenerated conversion function.
func Convert_azure_WorkerStatus_To_v1alpha1_WorkerStatus(in *azure.WorkerStatus, out *WorkerStatus, s conversion.Scope) error {
	return autoConvert_azure_WorkerStatus_To_v1alpha1_WorkerStatus(in, out, s)
}
