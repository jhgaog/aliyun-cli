package ecs

//Licensed under the Apache License, Version 2.0 (the "License");
//you may not use this file except in compliance with the License.
//You may obtain a copy of the License at
//
//http://www.apache.org/licenses/LICENSE-2.0
//
//Unless required by applicable law or agreed to in writing, software
//distributed under the License is distributed on an "AS IS" BASIS,
//WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
//See the License for the specific language governing permissions and
//limitations under the License.
//
// Code generated by Alibaba Cloud SDK Code Generator.
// Changes may cause incorrect behavior and will be lost if the code is regenerated.

// Instance is a nested struct in ecs response
type Instance struct {
	Hostname                        string                               `json:"Hostname" xml:"Hostname"`
	ImageId                         string                               `json:"ImageId" xml:"ImageId"`
	InstanceType                    string                               `json:"InstanceType" xml:"InstanceType"`
	AutoReleaseTime                 string                               `json:"AutoReleaseTime" xml:"AutoReleaseTime"`
	LastInvokedTime                 string                               `json:"LastInvokedTime" xml:"LastInvokedTime"`
	OsType                          string                               `json:"OsType" xml:"OsType"`
	DeviceAvailable                 bool                                 `json:"DeviceAvailable" xml:"DeviceAvailable"`
	InstanceNetworkType             string                               `json:"InstanceNetworkType" xml:"InstanceNetworkType"`
	RegistrationTime                string                               `json:"RegistrationTime" xml:"RegistrationTime"`
	LocalStorageAmount              int                                  `json:"LocalStorageAmount" xml:"LocalStorageAmount"`
	NetworkType                     string                               `json:"NetworkType" xml:"NetworkType"`
	IntranetIp                      string                               `json:"IntranetIp" xml:"IntranetIp"`
	IsSpot                          bool                                 `json:"IsSpot" xml:"IsSpot"`
	InstanceChargeType              string                               `json:"InstanceChargeType" xml:"InstanceChargeType"`
	MachineId                       string                               `json:"MachineId" xml:"MachineId"`
	PrivatePoolOptionsId            string                               `json:"PrivatePoolOptionsId" xml:"PrivatePoolOptionsId"`
	ClusterId                       string                               `json:"ClusterId" xml:"ClusterId"`
	InstanceName                    string                               `json:"InstanceName" xml:"InstanceName"`
	PrivatePoolOptionsMatchCriteria string                               `json:"PrivatePoolOptionsMatchCriteria" xml:"PrivatePoolOptionsMatchCriteria"`
	DeploymentSetGroupNo            int                                  `json:"DeploymentSetGroupNo" xml:"DeploymentSetGroupNo"`
	CreditSpecification             string                               `json:"CreditSpecification" xml:"CreditSpecification"`
	GPUAmount                       int                                  `json:"GPUAmount" xml:"GPUAmount"`
	Connected                       bool                                 `json:"Connected" xml:"Connected"`
	InvocationCount                 int64                                `json:"InvocationCount" xml:"InvocationCount"`
	StartTime                       string                               `json:"StartTime" xml:"StartTime"`
	ZoneId                          string                               `json:"ZoneId" xml:"ZoneId"`
	InternetChargeType              string                               `json:"InternetChargeType" xml:"InternetChargeType"`
	InternetMaxBandwidthIn          int                                  `json:"InternetMaxBandwidthIn" xml:"InternetMaxBandwidthIn"`
	HostName                        string                               `json:"HostName" xml:"HostName"`
	Status                          string                               `json:"Status" xml:"Status"`
	CPU                             int                                  `json:"CPU" xml:"CPU"`
	Cpu                             int                                  `json:"Cpu" xml:"Cpu"`
	ISP                             string                               `json:"ISP" xml:"ISP"`
	OsVersion                       string                               `json:"OsVersion" xml:"OsVersion"`
	SpotPriceLimit                  float64                              `json:"SpotPriceLimit" xml:"SpotPriceLimit"`
	OSName                          string                               `json:"OSName" xml:"OSName"`
	OSNameEn                        string                               `json:"OSNameEn" xml:"OSNameEn"`
	SerialNumber                    string                               `json:"SerialNumber" xml:"SerialNumber"`
	RegionId                        string                               `json:"RegionId" xml:"RegionId"`
	IoOptimized                     bool                                 `json:"IoOptimized" xml:"IoOptimized"`
	InternetMaxBandwidthOut         int                                  `json:"InternetMaxBandwidthOut" xml:"InternetMaxBandwidthOut"`
	ResourceGroupId                 string                               `json:"ResourceGroupId" xml:"ResourceGroupId"`
	ActivationId                    string                               `json:"ActivationId" xml:"ActivationId"`
	InstanceTypeFamily              string                               `json:"InstanceTypeFamily" xml:"InstanceTypeFamily"`
	InstanceId                      string                               `json:"InstanceId" xml:"InstanceId"`
	DeploymentSetId                 string                               `json:"DeploymentSetId" xml:"DeploymentSetId"`
	GPUSpec                         string                               `json:"GPUSpec" xml:"GPUSpec"`
	Description                     string                               `json:"Description" xml:"Description"`
	Recyclable                      bool                                 `json:"Recyclable" xml:"Recyclable"`
	SaleCycle                       string                               `json:"SaleCycle" xml:"SaleCycle"`
	ExpiredTime                     string                               `json:"ExpiredTime" xml:"ExpiredTime"`
	OSType                          string                               `json:"OSType" xml:"OSType"`
	InternetIp                      string                               `json:"InternetIp" xml:"InternetIp"`
	Memory                          int                                  `json:"Memory" xml:"Memory"`
	CreationTime                    string                               `json:"CreationTime" xml:"CreationTime"`
	AgentVersion                    string                               `json:"AgentVersion" xml:"AgentVersion"`
	KeyPairName                     string                               `json:"KeyPairName" xml:"KeyPairName"`
	HpcClusterId                    string                               `json:"HpcClusterId" xml:"HpcClusterId"`
	LocalStorageCapacity            int64                                `json:"LocalStorageCapacity" xml:"LocalStorageCapacity"`
	VlanId                          string                               `json:"VlanId" xml:"VlanId"`
	StoppedMode                     string                               `json:"StoppedMode" xml:"StoppedMode"`
	SpotStrategy                    string                               `json:"SpotStrategy" xml:"SpotStrategy"`
	SpotDuration                    int                                  `json:"SpotDuration" xml:"SpotDuration"`
	DeletionProtection              bool                                 `json:"DeletionProtection" xml:"DeletionProtection"`
	SecurityGroupIds                SecurityGroupIdsInDescribeInstances  `json:"SecurityGroupIds" xml:"SecurityGroupIds"`
	InnerIpAddress                  InnerIpAddressInDescribeInstances    `json:"InnerIpAddress" xml:"InnerIpAddress"`
	PublicIpAddress                 PublicIpAddressInDescribeInstances   `json:"PublicIpAddress" xml:"PublicIpAddress"`
	RdmaIpAddress                   RdmaIpAddress                        `json:"RdmaIpAddress" xml:"RdmaIpAddress"`
	MetadataOptions                 MetadataOptions                      `json:"MetadataOptions" xml:"MetadataOptions"`
	DedicatedHostAttribute          DedicatedHostAttribute               `json:"DedicatedHostAttribute" xml:"DedicatedHostAttribute"`
	EipAddress                      EipAddressInDescribeInstances        `json:"EipAddress" xml:"EipAddress"`
	HibernationOptions              HibernationOptions                   `json:"HibernationOptions" xml:"HibernationOptions"`
	CpuOptions                      CpuOptions                           `json:"CpuOptions" xml:"CpuOptions"`
	EcsCapacityReservationAttr      EcsCapacityReservationAttr           `json:"EcsCapacityReservationAttr" xml:"EcsCapacityReservationAttr"`
	DedicatedInstanceAttribute      DedicatedInstanceAttribute           `json:"DedicatedInstanceAttribute" xml:"DedicatedInstanceAttribute"`
	VpcAttributes                   VpcAttributes                        `json:"VpcAttributes" xml:"VpcAttributes"`
	OperationLocks                  OperationLocksInDescribeInstances    `json:"OperationLocks" xml:"OperationLocks"`
	Tags                            TagsInDescribeInstances              `json:"Tags" xml:"Tags"`
	NetworkInterfaces               NetworkInterfacesInDescribeInstances `json:"NetworkInterfaces" xml:"NetworkInterfaces"`
}
