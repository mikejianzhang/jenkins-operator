// +build !ignore_autogenerated

// Code generated by openapi-gen. DO NOT EDIT.

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha2

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Jenkins":       schema_pkg_apis_jenkins_v1alpha2_Jenkins(ref),
		"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsSpec":   schema_pkg_apis_jenkins_v1alpha2_JenkinsSpec(ref),
		"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsStatus": schema_pkg_apis_jenkins_v1alpha2_JenkinsStatus(ref),
	}
}

func schema_pkg_apis_jenkins_v1alpha2_Jenkins(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Jenkins is the Schema for the jenkins API",
				Properties: map[string]spec.Schema{
					"kind": {
						SchemaProps: spec.SchemaProps{
							Description: "Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#types-kinds",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"apiVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: https://git.k8s.io/community/contributors/devel/api-conventions.md#resources",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"metadata": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"),
						},
					},
					"spec": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsSpec", "github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_jenkins_v1alpha2_JenkinsSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsSpec defines the desired state of Jenkins",
				Properties: map[string]spec.Schema{
					"master": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL SPEC FIELDS - desired state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file",
							Ref:         ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsMaster"),
						},
					},
					"seedJobs": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.SeedJob"),
									},
								},
							},
						},
					},
					"service": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Service"),
						},
					},
					"slaveService": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Service"),
						},
					},
					"backup": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Backup"),
						},
					},
					"restore": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Restore"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Backup", "github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.JenkinsMaster", "github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Restore", "github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.SeedJob", "github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Service"},
	}
}

func schema_pkg_apis_jenkins_v1alpha2_JenkinsStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "JenkinsStatus defines the observed state of Jenkins",
				Properties: map[string]spec.Schema{
					"operatorVersion": {
						SchemaProps: spec.SchemaProps{
							Description: "INSERT ADDITIONAL STATUS FIELD - define observed state of cluster Important: Run \"operator-sdk generate k8s\" to regenerate code after modifying this file",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"provisionStartTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"baseConfigurationCompletedTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"userConfigurationCompletedTime": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("k8s.io/apimachinery/pkg/apis/meta/v1.Time"),
						},
					},
					"builds": {
						SchemaProps: spec.SchemaProps{
							Type: []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Build"),
									},
								},
							},
						},
					},
					"restoredBackup": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"lastBackup": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
					"pendingBackup": {
						SchemaProps: spec.SchemaProps{
							Type:   []string{"integer"},
							Format: "int64",
						},
					},
				},
			},
		},
		Dependencies: []string{
			"github.com/mikejianzhang/jenkins-operator/pkg/apis/jenkins/v1alpha2.Build", "k8s.io/apimachinery/pkg/apis/meta/v1.Time"},
	}
}
