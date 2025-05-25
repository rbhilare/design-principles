package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// Violating Single Responsibility Principle
type Resources struct {
	// Deployment related attributes
	DeploymentTypeMeta   metav1.TypeMeta   `json:",inline"`
	DeploymentObjectMeta metav1.ObjectMeta `json:"deploymentmetadata,omitempty"`
	DeploymentSpec       DeploymentSpec    `json:"deploymentspec,omitempty"`
	DeploymentStatus     DeploymentStatus  `json:"deploymentstatus,omitempty"`

	// Pod related attributes
	PodTypeMeta   metav1.TypeMeta   `json:",inline"`
	PodObjectMeta metav1.ObjectMeta `json:"podmetadata,omitempty"`
	PodSpec       PodSpec           `json:"podspec,omitempty"`
	PodStatus     PodStatus         `json:"podstatus,omitempty"`
}

type DeploymentSpec struct {
	LabelSelector map[string]string `json:"labelSelector,omitempty"`
}

type DeploymentStatus struct {
	LastRestartTime string `json:"lastRestartTime,omitempty"`
}

type PodSpec struct {
	LabelSelector map[string]string `json:"labelSelector,omitempty"`
}

type PodStatus struct {
	LastRestartTime string `json:"lastRestartTime,omitempty"`
}
