/*
Copyright The Flagger Authors.

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

// Code generated by client-gen. DO NOT EDIT.

package fake

import (
	appmeshv1beta1 "github.com/weaveworks/flagger/pkg/apis/appmesh/v1beta1"
	flaggerv1alpha1 "github.com/weaveworks/flagger/pkg/apis/flagger/v1alpha1"
	flaggerv1alpha3 "github.com/weaveworks/flagger/pkg/apis/flagger/v1alpha3"
	gloov1 "github.com/weaveworks/flagger/pkg/apis/gloo/v1"
	networkingv1alpha3 "github.com/weaveworks/flagger/pkg/apis/istio/v1alpha3"
	projectcontourv1 "github.com/weaveworks/flagger/pkg/apis/projectcontour/v1"
	splitv1alpha1 "github.com/weaveworks/flagger/pkg/apis/smi/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	schema "k8s.io/apimachinery/pkg/runtime/schema"
	serializer "k8s.io/apimachinery/pkg/runtime/serializer"
	utilruntime "k8s.io/apimachinery/pkg/util/runtime"
)

var scheme = runtime.NewScheme()
var codecs = serializer.NewCodecFactory(scheme)
var parameterCodec = runtime.NewParameterCodec(scheme)
var localSchemeBuilder = runtime.SchemeBuilder{
	appmeshv1beta1.AddToScheme,
	flaggerv1alpha3.AddToScheme,
	flaggerv1alpha1.AddToScheme,
	gloov1.AddToScheme,
	networkingv1alpha3.AddToScheme,
	projectcontourv1.AddToScheme,
	splitv1alpha1.AddToScheme,
}

// AddToScheme adds all types of this clientset into the given scheme. This allows composition
// of clientsets, like in:
//
//   import (
//     "k8s.io/client-go/kubernetes"
//     clientsetscheme "k8s.io/client-go/kubernetes/scheme"
//     aggregatorclientsetscheme "k8s.io/kube-aggregator/pkg/client/clientset_generated/clientset/scheme"
//   )
//
//   kclientset, _ := kubernetes.NewForConfig(c)
//   _ = aggregatorclientsetscheme.AddToScheme(clientsetscheme.Scheme)
//
// After this, RawExtensions in Kubernetes types will serialize kube-aggregator types
// correctly.
var AddToScheme = localSchemeBuilder.AddToScheme

func init() {
	v1.AddToGroupVersion(scheme, schema.GroupVersion{Version: "v1"})
	utilruntime.Must(AddToScheme(scheme))
}
