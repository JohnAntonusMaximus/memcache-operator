// +build !ignore_autogenerated

/*
Copyright The Kubernetes Authors.

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

// Code generated by conversion-gen. DO NOT EDIT.

package v1alpha1

import (
	config "k8s.io/apimachinery/pkg/apis/config"
	conversion "k8s.io/apimachinery/pkg/conversion"
	runtime "k8s.io/apimachinery/pkg/runtime"
)

func init() {
	localSchemeBuilder.Register(RegisterConversions)
}

// RegisterConversions adds conversion functions to the given scheme.
// Public to allow building arbitrary schemes.
func RegisterConversions(s *runtime.Scheme) error {
	if err := s.AddGeneratedConversionFunc((*ClientConnectionConfiguration)(nil), (*config.ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(a.(*ClientConnectionConfiguration), b.(*config.ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddGeneratedConversionFunc((*config.ClientConnectionConfiguration)(nil), (*ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(a.(*config.ClientConnectionConfiguration), b.(*ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*config.ClientConnectionConfiguration)(nil), (*ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(a.(*config.ClientConnectionConfiguration), b.(*ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	if err := s.AddConversionFunc((*ClientConnectionConfiguration)(nil), (*config.ClientConnectionConfiguration)(nil), func(a, b interface{}, scope conversion.Scope) error {
		return Convert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(a.(*ClientConnectionConfiguration), b.(*config.ClientConnectionConfiguration), scope)
	}); err != nil {
		return err
	}
	return nil
}

func autoConvert_v1alpha1_ClientConnectionConfiguration_To_config_ClientConnectionConfiguration(in *ClientConnectionConfiguration, out *config.ClientConnectionConfiguration, s conversion.Scope) error {
	out.Kubeconfig = in.Kubeconfig
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}

func autoConvert_config_ClientConnectionConfiguration_To_v1alpha1_ClientConnectionConfiguration(in *config.ClientConnectionConfiguration, out *ClientConnectionConfiguration, s conversion.Scope) error {
	out.Kubeconfig = in.Kubeconfig
	out.AcceptContentTypes = in.AcceptContentTypes
	out.ContentType = in.ContentType
	out.QPS = in.QPS
	out.Burst = in.Burst
	return nil
}
