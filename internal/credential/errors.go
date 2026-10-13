/*
Copyright 2026 Red Hat, Inc..

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

package credential

import "errors"

var (
	ErrNoSecretSource         = errors.New("has no secret credential source")
	ErrReadingSecret          = errors.New("reading admin kubeconfig secret")
	ErrMissingKubeconfigKey   = errors.New("admin kubeconfig secret missing 'kubeconfig' key")
	ErrParsingKubeconfig      = errors.New("parsing kubeconfig from secret")
	ErrCertFilePathUsed       = errors.New("certificate-authority file path; use 'kubectl config view --flatten' to convert to inline data")
	ErrClientCertFilePathUsed = errors.New("client-certificate/key file paths; use 'kubectl config view --flatten' to convert to inline data")
	ErrTokenFileUsed          = errors.New("tokenFile which would read the hub pod's SA token; use a static token instead")
	ErrUnsupportedAuth        = errors.New("must contain a token or client certificate; exec-based auth (e.g. oc login) is not supported")
	ErrSerializingKubeconfig  = errors.New("serializing standing kubeconfig")
	ErrSettingOwnerRef        = errors.New("setting controller reference")
)
