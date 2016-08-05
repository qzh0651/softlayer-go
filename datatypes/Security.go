/**
 * Copyright 2016 IBM Corp.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *    http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package datatypes

type Security_Certificate struct {
	Entity

	AssociatedServiceCount            *int                                                                    `json:"associatedServiceCount,omitempty"`
	Certificate                       *string                                                                 `json:"certificate,omitempty"`
	CertificateSigningRequest         *string                                                                 `json:"certificateSigningRequest,omitempty"`
	CommonName                        *string                                                                 `json:"commonName,omitempty"`
	CreateDate                        *Time                                                                   `json:"createDate,omitempty"`
	Id                                *int                                                                    `json:"id,omitempty"`
	IntermediateCertificate           *string                                                                 `json:"intermediateCertificate,omitempty"`
	KeySize                           *int                                                                    `json:"keySize,omitempty"`
	LoadBalancerVirtualIpAddressCount *uint                                                                   `json:"loadBalancerVirtualIpAddressCount,omitempty"`
	LoadBalancerVirtualIpAddresses    []Network_Application_Delivery_Controller_LoadBalancer_VirtualIpAddress `json:"loadBalancerVirtualIpAddresses,omitempty"`
	ModifyDate                        *Time                                                                   `json:"modifyDate,omitempty"`
	Notes                             *string                                                                 `json:"notes,omitempty"`
	OrganizationName                  *string                                                                 `json:"organizationName,omitempty"`
	PrivateKey                        *string                                                                 `json:"privateKey,omitempty"`
	ValidityBegin                     *Time                                                                   `json:"validityBegin,omitempty"`
	ValidityDays                      *int                                                                    `json:"validityDays,omitempty"`
	ValidityEnd                       *Time                                                                   `json:"validityEnd,omitempty"`
}

type Security_Certificate_Entry struct {
	Entity

	CertificateId    *int    `json:"certificateId,omitempty"`
	CommonName       *string `json:"commonName,omitempty"`
	KeySize          *int    `json:"keySize,omitempty"`
	OrganizationName *string `json:"organizationName,omitempty"`
	ValidityBegin    *Time   `json:"validityBegin,omitempty"`
	ValidityDays     *int    `json:"validityDays,omitempty"`
	ValidityEnd      *Time   `json:"validityEnd,omitempty"`
}

type Security_Certificate_Request struct {
	Entity

	Account                      *Account                             `json:"account,omitempty"`
	AccountId                    *int                                 `json:"accountId,omitempty"`
	ApproverEmailAddress         *string                              `json:"approverEmailAddress,omitempty"`
	CertificateAuthorityName     *string                              `json:"certificateAuthorityName,omitempty"`
	CertificateSigningRequest    *string                              `json:"certificateSigningRequest,omitempty"`
	CommonName                   *string                              `json:"commonName,omitempty"`
	CreateDate                   *Time                                `json:"createDate,omitempty"`
	EffectiveDate                *Time                                `json:"effectiveDate,omitempty"`
	ExpirationDate               *Time                                `json:"expirationDate,omitempty"`
	Id                           *int                                 `json:"id,omitempty"`
	ModifyDate                   *Time                                `json:"modifyDate,omitempty"`
	Order                        *Billing_Order                       `json:"order,omitempty"`
	OrderItem                    *Billing_Order_Item                  `json:"orderItem,omitempty"`
	Status                       *Security_Certificate_Request_Status `json:"status,omitempty"`
	StatusId                     *int                                 `json:"statusId,omitempty"`
	TechnicalContactEmailAddress *string                              `json:"technicalContactEmailAddress,omitempty"`
}

type Security_Certificate_Request_ServerType struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
	Value       *int    `json:"value,omitempty"`
}

type Security_Certificate_Request_Status struct {
	Entity

	Description *string `json:"description,omitempty"`
	Id          *int    `json:"id,omitempty"`
	Name        *string `json:"name,omitempty"`
}

type Security_Directory_Service_Host_Xref_Hardware struct {
	Entity

	Host *Hardware `json:"host,omitempty"`
}

type Security_SecureTransportCipher struct {
	Entity

	KeyName *string `json:"keyName,omitempty"`
}

type Security_SecureTransportProtocol struct {
	Entity

	KeyName                         *string                          `json:"keyName,omitempty"`
	SupportedSecureTransportCiphers []Security_SecureTransportCipher `json:"supportedSecureTransportCiphers,omitempty"`
}

type Security_Ssh_Key struct {
	Entity

	Account                       *Account                                    `json:"account,omitempty"`
	BlockDeviceTemplateGroupCount *uint                                       `json:"blockDeviceTemplateGroupCount,omitempty"`
	BlockDeviceTemplateGroups     []Virtual_Guest_Block_Device_Template_Group `json:"blockDeviceTemplateGroups,omitempty"`
	CreateDate                    *Time                                       `json:"createDate,omitempty"`
	Fingerprint                   *string                                     `json:"fingerprint,omitempty"`
	Id                            *int                                        `json:"id,omitempty"`
	Key                           *string                                     `json:"key,omitempty"`
	Label                         *string                                     `json:"label,omitempty"`
	ModifyDate                    *Time                                       `json:"modifyDate,omitempty"`
	Notes                         *string                                     `json:"notes,omitempty"`
	SoftwarePasswordCount         *uint                                       `json:"softwarePasswordCount,omitempty"`
	SoftwarePasswords             []Software_Component_Password               `json:"softwarePasswords,omitempty"`
}
