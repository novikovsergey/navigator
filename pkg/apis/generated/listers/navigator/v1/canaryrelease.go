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

// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	v1 "github.com/avito-tech/navigator/pkg/apis/navigator/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// CanaryReleaseLister helps list CanaryReleases.
type CanaryReleaseLister interface {
	// List lists all CanaryReleases in the indexer.
	List(selector labels.Selector) (ret []*v1.CanaryRelease, err error)
	// CanaryReleases returns an object that can list and get CanaryReleases.
	CanaryReleases(namespace string) CanaryReleaseNamespaceLister
	CanaryReleaseListerExpansion
}

// canaryReleaseLister implements the CanaryReleaseLister interface.
type canaryReleaseLister struct {
	indexer cache.Indexer
}

// NewCanaryReleaseLister returns a new CanaryReleaseLister.
func NewCanaryReleaseLister(indexer cache.Indexer) CanaryReleaseLister {
	return &canaryReleaseLister{indexer: indexer}
}

// List lists all CanaryReleases in the indexer.
func (s *canaryReleaseLister) List(selector labels.Selector) (ret []*v1.CanaryRelease, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CanaryRelease))
	})
	return ret, err
}

// CanaryReleases returns an object that can list and get CanaryReleases.
func (s *canaryReleaseLister) CanaryReleases(namespace string) CanaryReleaseNamespaceLister {
	return canaryReleaseNamespaceLister{indexer: s.indexer, namespace: namespace}
}

// CanaryReleaseNamespaceLister helps list and get CanaryReleases.
type CanaryReleaseNamespaceLister interface {
	// List lists all CanaryReleases in the indexer for a given namespace.
	List(selector labels.Selector) (ret []*v1.CanaryRelease, err error)
	// Get retrieves the CanaryRelease from the indexer for a given namespace and name.
	Get(name string) (*v1.CanaryRelease, error)
	CanaryReleaseNamespaceListerExpansion
}

// canaryReleaseNamespaceLister implements the CanaryReleaseNamespaceLister
// interface.
type canaryReleaseNamespaceLister struct {
	indexer   cache.Indexer
	namespace string
}

// List lists all CanaryReleases in the indexer for a given namespace.
func (s canaryReleaseNamespaceLister) List(selector labels.Selector) (ret []*v1.CanaryRelease, err error) {
	err = cache.ListAllByNamespace(s.indexer, s.namespace, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.CanaryRelease))
	})
	return ret, err
}

// Get retrieves the CanaryRelease from the indexer for a given namespace and name.
func (s canaryReleaseNamespaceLister) Get(name string) (*v1.CanaryRelease, error) {
	obj, exists, err := s.indexer.GetByKey(s.namespace + "/" + name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("canaryrelease"), name)
	}
	return obj.(*v1.CanaryRelease), nil
}
