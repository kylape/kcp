/*
Copyright 2021 The KCP Authors.

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

package helper

import (
	"testing"

	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"

	tenancyv1alpha1 "github.com/kcp-dev/kcp/pkg/apis/tenancy/v1alpha1"
)

func TestEncodeLogicalClusterName(t *testing.T) {
	for _, testCase := range []struct {
		name        string
		input       *tenancyv1alpha1.Workspace
		expected    string
		expectedErr bool
	}{
		{
			name: "organization workspace",
			input: &tenancyv1alpha1.Workspace{
				ObjectMeta: metav1.ObjectMeta{
					ClusterName: "admin",
					Name:        "organization",
				},
			},
			expected: "admin_organization",
		},
		{
			name: "normal workspace",
			input: &tenancyv1alpha1.Workspace{
				ObjectMeta: metav1.ObjectMeta{
					ClusterName: "admin_organization",
					Name:        "workspace",
				},
			},
			expected: "organization_workspace",
		},
		{
			name: "organization workspace in wrong root cluster",
			input: &tenancyv1alpha1.Workspace{
				ObjectMeta: metav1.ObjectMeta{
					ClusterName: "somethingwrong",
					Name:        "organization",
				},
			},
			expectedErr: true,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			actual, actualErr := EncodeLogicalClusterName(testCase.input)
			if actualErr != nil && !testCase.expectedErr {
				t.Errorf("%s: expected no error, got %v", testCase.name, actualErr)
			}
			if actualErr == nil && testCase.expectedErr {
				t.Errorf("%s: expected error, got none", testCase.name)
			}
			if actual != testCase.expected {
				t.Errorf("%s: got incorrect logical cluster name, expected %s got %s", testCase.name, testCase.expected, actual)
			}
		})
	}
}

func TestParseLogicalClusterName(t *testing.T) {
	for _, testCase := range []struct {
		name         string
		input        string
		expectedOrg  string
		expectedName string
		expectedErr  bool
	}{
		{
			name:         "valid name",
			input:        "organization_workspace",
			expectedOrg:  "organization",
			expectedName: "workspace",
		},
		{
			name:        "invalid name",
			input:       "whoathere",
			expectedErr: true,
		},
	} {
		t.Run(testCase.name, func(t *testing.T) {
			actualOrg, actualName, actualErr := ParseLogicalClusterName(testCase.input)
			if actualErr != nil && !testCase.expectedErr {
				t.Errorf("%s: expected no error, got %v", testCase.name, actualErr)
			}
			if actualErr == nil && testCase.expectedErr {
				t.Errorf("%s: expected error, got none", testCase.name)
			}
			if actualOrg != testCase.expectedOrg {
				t.Errorf("%s: got incorrect logical cluster name, expected %s got %s", testCase.name, testCase.expectedOrg, actualOrg)
			}
			if actualName != testCase.expectedName {
				t.Errorf("%s: got incorrect logical cluster name, expected %s got %s", testCase.name, testCase.expectedName, actualName)
			}
		})
	}
}
