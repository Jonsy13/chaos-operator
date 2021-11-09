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

// Code generated by informer-gen. DO NOT EDIT.

package v1alpha1

import (
	time "time"

	litmuschaosv1alpha1 "github.com/Jonsy13/chaos-operator/pkg/apis/litmuschaos/v1alpha1"
	versioned "github.com/Jonsy13/chaos-operator/pkg/client/clientset/versioned"
	internalinterfaces "github.com/Jonsy13/chaos-operator/pkg/client/informers/externalversions/internalinterfaces"
	v1alpha1 "github.com/Jonsy13/chaos-operator/pkg/client/listers/litmuschaos/v1alpha1"
	v1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	runtime "k8s.io/apimachinery/pkg/runtime"
	watch "k8s.io/apimachinery/pkg/watch"
	cache "k8s.io/client-go/tools/cache"
)

// ChaosResultInformer provides access to a shared informer and lister for
// ChaosResults.
type ChaosResultInformer interface {
	Informer() cache.SharedIndexInformer
	Lister() v1alpha1.ChaosResultLister
}

type chaosResultInformer struct {
	factory          internalinterfaces.SharedInformerFactory
	tweakListOptions internalinterfaces.TweakListOptionsFunc
	namespace        string
}

// NewChaosResultInformer constructs a new informer for ChaosResult type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewChaosResultInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers) cache.SharedIndexInformer {
	return NewFilteredChaosResultInformer(client, namespace, resyncPeriod, indexers, nil)
}

// NewFilteredChaosResultInformer constructs a new informer for ChaosResult type.
// Always prefer using an informer factory to get a shared informer instead of getting an independent
// one. This reduces memory footprint and number of connections to the server.
func NewFilteredChaosResultInformer(client versioned.Interface, namespace string, resyncPeriod time.Duration, indexers cache.Indexers, tweakListOptions internalinterfaces.TweakListOptionsFunc) cache.SharedIndexInformer {
	return cache.NewSharedIndexInformer(
		&cache.ListWatch{
			ListFunc: func(options v1.ListOptions) (runtime.Object, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LitmuschaosV1alpha1().ChaosResults(namespace).List(options)
			},
			WatchFunc: func(options v1.ListOptions) (watch.Interface, error) {
				if tweakListOptions != nil {
					tweakListOptions(&options)
				}
				return client.LitmuschaosV1alpha1().ChaosResults(namespace).Watch(options)
			},
		},
		&litmuschaosv1alpha1.ChaosResult{},
		resyncPeriod,
		indexers,
	)
}

func (f *chaosResultInformer) defaultInformer(client versioned.Interface, resyncPeriod time.Duration) cache.SharedIndexInformer {
	return NewFilteredChaosResultInformer(client, f.namespace, resyncPeriod, cache.Indexers{cache.NamespaceIndex: cache.MetaNamespaceIndexFunc}, f.tweakListOptions)
}

func (f *chaosResultInformer) Informer() cache.SharedIndexInformer {
	return f.factory.InformerFor(&litmuschaosv1alpha1.ChaosResult{}, f.defaultInformer)
}

func (f *chaosResultInformer) Lister() v1alpha1.ChaosResultLister {
	return v1alpha1.NewChaosResultLister(f.Informer().GetIndexer())
}
