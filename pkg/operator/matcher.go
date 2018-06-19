// Copyright 2018 ReactiveOps
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package operator

import (
	"fmt"
	"reflect"

	rbacv1 "k8s.io/api/rbac/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

func crbMatches(existingCRB *rbacv1.ClusterRoleBinding, requestedCRB *rbacv1.ClusterRoleBinding) bool {
	if !metaMatches(&existingCRB.ObjectMeta, &requestedCRB.ObjectMeta) {
		return false
	}
	if !subjectsMatch(&existingCRB.Subjects, &requestedCRB.Subjects) {
		return false
	}

	if !roleRefMatches(&existingCRB.RoleRef, &requestedCRB.RoleRef) {
		return false
	}

	return true
}

func rbMatches(existingRB *rbacv1.RoleBinding, requestedRB *rbacv1.RoleBinding) bool {
	if !metaMatches(&existingRB.ObjectMeta, &requestedRB.ObjectMeta) {
		return false
	}
	if !subjectsMatch(&existingRB.Subjects, &requestedRB.Subjects) {
		return false
	}

	if !roleRefMatches(&existingRB.RoleRef, &requestedRB.RoleRef) {
		return false
	}

	return true
}

func metaMatches(existingMeta *metav1.ObjectMeta, requestedMeta *metav1.ObjectMeta) bool {
	if existingMeta.Name != requestedMeta.Name {
		return false
	}

	if existingMeta.Namespace != requestedMeta.Namespace {
		return false
	}

	if !reflect.DeepEqual(existingMeta.OwnerReferences, requestedMeta.OwnerReferences) {
		fmt.Printf("Owner References did not match: %v != %v", existingMeta.OwnerReferences, requestedMeta.OwnerReferences)
		return false
	}

	return true
}

func subjectsMatch(existingSubjects *[]rbacv1.Subject, requestedSubjects *[]rbacv1.Subject) bool {
	rSubjects := *requestedSubjects
	eSubjects := *existingSubjects

	if len(eSubjects) != len(rSubjects) {
		return false
	}

	for index, existingSubject := range eSubjects {
		if !subjectMatches(&existingSubject, &rSubjects[index]) {
			return false
		}
	}

	return true
}

func subjectMatches(existingSubject *rbacv1.Subject, requestedSubject *rbacv1.Subject) bool {
	if existingSubject.Kind != requestedSubject.Kind {
		return false
	}

	if existingSubject.Name != requestedSubject.Name {
		return false
	}

	if existingSubject.Namespace != requestedSubject.Namespace {
		return false
	}

	return true
}

func roleRefMatches(existingRoleRef *rbacv1.RoleRef, requestedRoleRef *rbacv1.RoleRef) bool {
	if existingRoleRef.Kind != requestedRoleRef.Kind {
		return false
	}

	if existingRoleRef.Name != requestedRoleRef.Name {
		return false
	}

	return true
}
