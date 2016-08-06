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

type Legal_RegulatedWorkload struct {
	Entity

	Account        *Account                      `json:"account,omitempty"`
	AccountId      *int                          `json:"accountId,omitempty"`
	EnabledFlag    *bool                         `json:"enabledFlag,omitempty"`
	Id             *int                          `json:"id,omitempty"`
	Type           *Legal_RegulatedWorkload_Type `json:"type,omitempty"`
	WorkloadTypeId *int                          `json:"workloadTypeId,omitempty"`
}

type Legal_RegulatedWorkload_Type struct {
	Entity

	Id      *int    `json:"id,omitempty"`
	KeyName *string `json:"keyName,omitempty"`
	Name    *string `json:"name,omitempty"`
}