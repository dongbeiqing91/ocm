// Code generated by lister-gen. DO NOT EDIT.

package v1

import (
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
	v1 "open-cluster-management.io/api/work/v1"
)

// AppliedManifestWorkLister helps list AppliedManifestWorks.
// All objects returned here must be treated as read-only.
type AppliedManifestWorkLister interface {
	// List lists all AppliedManifestWorks in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v1.AppliedManifestWork, err error)
	// Get retrieves the AppliedManifestWork from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v1.AppliedManifestWork, error)
	AppliedManifestWorkListerExpansion
}

// appliedManifestWorkLister implements the AppliedManifestWorkLister interface.
type appliedManifestWorkLister struct {
	indexer cache.Indexer
}

// NewAppliedManifestWorkLister returns a new AppliedManifestWorkLister.
func NewAppliedManifestWorkLister(indexer cache.Indexer) AppliedManifestWorkLister {
	return &appliedManifestWorkLister{indexer: indexer}
}

// List lists all AppliedManifestWorks in the indexer.
func (s *appliedManifestWorkLister) List(selector labels.Selector) (ret []*v1.AppliedManifestWork, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v1.AppliedManifestWork))
	})
	return ret, err
}

// Get retrieves the AppliedManifestWork from the index for a given name.
func (s *appliedManifestWorkLister) Get(name string) (*v1.AppliedManifestWork, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v1.Resource("appliedmanifestwork"), name)
	}
	return obj.(*v1.AppliedManifestWork), nil
}