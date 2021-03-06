package types

import (
	appsv1 "k8s.io/api/apps/v1"
	corev1 "k8s.io/api/core/v1"
	networkingv1 "k8s.io/api/networking/v1"
)

type ClusterState struct {
	Namespaces      []*corev1.Namespace
	Pods            []*corev1.Pod
	Services        []*corev1.Service
	ReplicaSets     []*appsv1.ReplicaSet
	Deployments     []*appsv1.Deployment
	NetworkPolicies []*networkingv1.NetworkPolicy
}

type Pod struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

type PodRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type PodIsolation struct {
	Pod               PodRef `json:"pod"`
	IsIngressIsolated bool   `json:"isIngressIsolated"`
	IsEgressIsolated  bool   `json:"isEgressIsolated"`
}

type NetworkPolicy struct {
	Name      string            `json:"name"`
	Namespace string            `json:"namespace"`
	Labels    map[string]string `json:"labels"`
}

type AllowedRoute struct {
	SourcePod       PodRef          `json:"sourcePod"`
	EgressPolicies  []NetworkPolicy `json:"egressPolicies"`
	TargetPod       PodRef          `json:"targetPod"`
	IngressPolicies []NetworkPolicy `json:"ingressPolicies"`
	Ports           []int32         `json:"ports"`
}

type Service struct {
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	TargetPods []PodRef `json:"targetPods"`
}

type ReplicaSet struct {
	Name       string   `json:"name"`
	Namespace  string   `json:"namespace"`
	TargetPods []PodRef `json:"targetPods"`
}

type ReplicaSetRef struct {
	Name      string `json:"name"`
	Namespace string `json:"namespace"`
}

type Deployment struct {
	Name              string          `json:"name"`
	Namespace         string          `json:"namespace"`
	TargetReplicaSets []ReplicaSetRef `json:"targetReplicaSets"`
}

type AnalysisResult struct {
	Pods          []*Pod          `json:"pods"`
	PodIsolations []*PodIsolation `json:"podIsolations"`
	AllowedRoutes []*AllowedRoute `json:"allowedRoutes"`
	Services      []*Service      `json:"services"`
	ReplicaSets   []*ReplicaSet   `json:"replicaSets"`
	Deployments   []*Deployment   `json:"deployments"`
}
