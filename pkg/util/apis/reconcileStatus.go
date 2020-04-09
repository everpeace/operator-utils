/*
Copyright 2019 Red Hat, Inc.

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

package apis

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// ReconcileStatus represent the status of the last reconcile cycle. It's used to communicate success or failer and the error message
// +k8s:openapi-gen=true
type ReconcileStatus struct {

	// Status can be "Success" or "Failure"
	// +kubebuilder:validation:Enum=Success;Failure
	Status ResourceConditionType `json:"status,omitempty"`

	// LastUpdate this is the time of the last when the status was last updated
	LastUpdate metav1.Time `json:"lastUpdate,omitempty"`

	// Reason a custom message describing the status or the error
	// +kubebuilder:validation:Optional
	Reason string `json:"reason,omitempty"`
}

// ResourceConditionType can be "Enforcing" of "Failure" depending on whether the LockedResourceReconciler is currently able to enforce the resource
type ResourceConditionType string

const (
	// Enforcing means that the patch has been succesfully reconciled and it's being enforced
	Success ResourceConditionType = "Success"

	// Failure means that the patch has not been successfully reconciled and we cannot guarntee that it's being enforced
	Failure ResourceConditionType = "Failure"
)

// ReconcileStatusAware represnt a CRD type that has been enabled with ReconcileStatus, it can then benefit of a series of utility methods.
type ReconcileStatusAware interface {
	GetReconcileStatus() ReconcileStatus
	SetReconcileStatus(reconcileStatus ReconcileStatus)
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ReconcileStatus) DeepCopyInto(out *ReconcileStatus) {
	*out = *in
	in.LastUpdate.DeepCopyInto(&out.LastUpdate)
	return
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ReconcileStatus.
func (in *ReconcileStatus) DeepCopy() *ReconcileStatus {
	if in == nil {
		return nil
	}
	out := new(ReconcileStatus)
	in.DeepCopyInto(out)
	return out
}
