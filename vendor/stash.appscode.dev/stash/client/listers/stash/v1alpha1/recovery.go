/*
Copyright 2019 The Stash Authors.

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

// Code generated by lister-gen. DO NOT EDIT.

package v1alpha1

import (
	v1alpha1 "stash.appscode.dev/stash/apis/stash/v1alpha1"

	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// RecoveryLister helps list Recoveries.
type RecoveryLister interface {
	// List lists all Recoveries in the indexer.
	List(selector labels.Selector) (ret []*v1alpha1.Recovery, err error)
	// Recoveries returns an object that can list and get Recoveries.
	Recoveries(namespace string) RecoveryNamespaceLister
	RecoveryListerExpansion
}

// recoveryLister implements the RecoveryLister interface.
type recoveryLister struct {
	indexer cache.Indexer
}

// NewRecoveryLister returns a new RecoveryLister.
func NewRecoveryLister(indexer cache.Indexer) RecoveryLister {
	return &recoveryLister{indexer: indexer}
}

// List lists all Recoveries in the indexer.
func (s *recoveryLister) List(selector labels.Selector) (ret []*v1alpha1.Recovery, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recovery))
	})
	return ret, err
}

// Recoveries returns an object that can list and get Recoveries.
func (s *recoveryLister) Recoveries(namespace string) RecoveryNamespaceLister {
	return recoveryNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// RecoveryNamespaceLister helps list and get Recoveries.
type RecoveryNamespaceLister interface {
	// List lists all Recoveries in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1alpha1.Recovery, err error)
	// Get retrieves the Recovery from the indexer for a given namespace and name.
	Get(name string) (*v1alpha1.Recovery, error)
	RecoveryNamespaceListerExpansion
}

// recoveryNamespaceLister implements the RecoveryNamespaceLister
// interface.
type recoveryNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all Recoveries in the indexer for a given namespace.
func (s recoveryNamespaceLister) List(selector labels.Selector) (ret []*v1alpha1.Recovery, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1alpha1.Recovery))
	})
	return ret, err
}

// Get retrieves the Recovery from the indexer for a given namespace and name.
func (s recoveryNamespaceLister) Get(name string) (*v1alpha1.Recovery, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1alpha1.Resource("recovery"), name)
	}
	return obj.(*v1alpha1.Recovery), nil
}
