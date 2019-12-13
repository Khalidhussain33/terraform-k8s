// +build !ignore_autogenerated

// This file was autogenerated by openapi-gen. Do not edit it manually!

package v1alpha1

import (
	spec "github.com/go-openapi/spec"
	common "k8s.io/kube-openapi/pkg/common"
)

func GetOpenAPIDefinitions(ref common.ReferenceCallback) map[string]common.OpenAPIDefinition {
	return map[string]common.OpenAPIDefinition{
		"./pkg/apis/app/v1alpha1.Workspace":       schema_pkg_apis_app_v1alpha1_Workspace(ref),
		"./pkg/apis/app/v1alpha1.WorkspaceSpec":   schema_pkg_apis_app_v1alpha1_WorkspaceSpec(ref),
		"./pkg/apis/app/v1alpha1.WorkspaceStatus": schema_pkg_apis_app_v1alpha1_WorkspaceStatus(ref),
	}
}

func schema_pkg_apis_app_v1alpha1_Workspace(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "Workspace is the Schema for the workspaces API",
				Type:        []string{"object"},
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
							Ref: ref("./pkg/apis/app/v1alpha1.WorkspaceSpec"),
						},
					},
					"status": {
						SchemaProps: spec.SchemaProps{
							Ref: ref("./pkg/apis/app/v1alpha1.WorkspaceStatus"),
						},
					},
				},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.WorkspaceSpec", "./pkg/apis/app/v1alpha1.WorkspaceStatus", "k8s.io/apimachinery/pkg/apis/meta/v1.ObjectMeta"},
	}
}

func schema_pkg_apis_app_v1alpha1_WorkspaceSpec(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceSpec defines the desired state of Workspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"organization": {
						SchemaProps: spec.SchemaProps{
							Description: "Terraform Cloud organization",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"module": {
						SchemaProps: spec.SchemaProps{
							Description: "Module source and version to use",
							Ref:         ref("./pkg/apis/app/v1alpha1.Module"),
						},
					},
					"variables": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Variables as inputs to module",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/app/v1alpha1.Variable"),
									},
								},
							},
						},
					},
					"secretsMountPath": {
						SchemaProps: spec.SchemaProps{
							Description: "File path within operator pod to load workspace secrets",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"outputs": {
						VendorExtensible: spec.VendorExtensible{
							Extensions: spec.Extensions{
								"x-kubernetes-list-type": "set",
							},
						},
						SchemaProps: spec.SchemaProps{
							Description: "Outputs denote outputs wanted",
							Type:        []string{"array"},
							Items: &spec.SchemaOrArray{
								Schema: &spec.Schema{
									SchemaProps: spec.SchemaProps{
										Ref: ref("./pkg/apis/app/v1alpha1.Output"),
									},
								},
							},
						},
					},
				},
				Required: []string{"organization", "module", "secretsMountPath"},
			},
		},
		Dependencies: []string{
			"./pkg/apis/app/v1alpha1.Module", "./pkg/apis/app/v1alpha1.Output", "./pkg/apis/app/v1alpha1.Variable"},
	}
}

func schema_pkg_apis_app_v1alpha1_WorkspaceStatus(ref common.ReferenceCallback) common.OpenAPIDefinition {
	return common.OpenAPIDefinition{
		Schema: spec.Schema{
			SchemaProps: spec.SchemaProps{
				Description: "WorkspaceStatus defines the observed state of Workspace",
				Type:        []string{"object"},
				Properties: map[string]spec.Schema{
					"runStatus": {
						SchemaProps: spec.SchemaProps{
							Description: "Run Status gets the run status",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"workspaceID": {
						SchemaProps: spec.SchemaProps{
							Description: "Workspace ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"runID": {
						SchemaProps: spec.SchemaProps{
							Description: "Run ID",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"configHash": {
						SchemaProps: spec.SchemaProps{
							Description: "Configuration hash",
							Type:        []string{"string"},
							Format:      "",
						},
					},
					"stateDownloadURL": {
						SchemaProps: spec.SchemaProps{
							Description: "State Version download URL",
							Type:        []string{"string"},
							Format:      "",
						},
					},
				},
				Required: []string{"runStatus", "workspaceID", "runID", "configHash", "stateDownloadURL"},
			},
		},
	}
}
