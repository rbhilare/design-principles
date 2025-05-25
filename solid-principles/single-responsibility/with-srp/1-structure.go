package main

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// With Single Responsibility Principle

// Deployment related attributes
type Deployment struct {
	TypeMeta   metav1.TypeMeta   `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec       DeploymentSpec    `json:"spec,omitempty"`
	Status     DeploymentStatus  `json:"status,omitempty"`
}

// Pod related attributes
type Pod struct {
	TypeMeta   metav1.TypeMeta   `json:",inline"`
	ObjectMeta metav1.ObjectMeta `json:"metadata,omitempty"`
	Spec       PodSpec           `json:"spec,omitempty"`
	Status     PodStatus         `json:"status,omitempty"`
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
