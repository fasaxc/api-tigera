// Copyright (c) 2021 Tigera, Inc. All rights reserved.

// Code generated by lister-gen. DO NOT EDIT.

package v3

import (
	v3 "github.com/projectcalico/api/pkg/apis/projectcalico/v3"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/labels"
	"k8s.io/client-go/tools/cache"
)

// StagedKubernetesNetworkPolicyLister helps list StagedKubernetesNetworkPolicies.
// All objects returned here must be treated as read-only.
type StagedKubernetesNetworkPolicyLister interface {
	// List lists all StagedKubernetesNetworkPolicies in the indexer.
	// Objects returned here must be treated as read-only.
	List(selector labels.Selector) (ret []*v3.StagedKubernetesNetworkPolicy, err error)
	// Get retrieves the StagedKubernetesNetworkPolicy from the index for a given name.
	// Objects returned here must be treated as read-only.
	Get(name string) (*v3.StagedKubernetesNetworkPolicy, error)
	StagedKubernetesNetworkPolicyListerExpansion
}

// stagedKubernetesNetworkPolicyLister implements the StagedKubernetesNetworkPolicyLister interface.
type stagedKubernetesNetworkPolicyLister struct {
	indexer cache.Indexer
}

// NewStagedKubernetesNetworkPolicyLister returns a new StagedKubernetesNetworkPolicyLister.
func NewStagedKubernetesNetworkPolicyLister(indexer cache.Indexer) StagedKubernetesNetworkPolicyLister {
	return &stagedKubernetesNetworkPolicyLister{indexer: indexer}
}

// List lists all StagedKubernetesNetworkPolicies in the indexer.
func (s *stagedKubernetesNetworkPolicyLister) List(selector labels.Selector) (ret []*v3.StagedKubernetesNetworkPolicy, err error) {
	err = cache.ListAll(s.indexer, selector, func(m interface{}) {
		ret = append(ret, m.(*v3.StagedKubernetesNetworkPolicy))
	})
	return ret, err
}

// Get retrieves the StagedKubernetesNetworkPolicy from the index for a given name.
func (s *stagedKubernetesNetworkPolicyLister) Get(name string) (*v3.StagedKubernetesNetworkPolicy, error) {
	obj, exists, err := s.indexer.GetByKey(name)
	if err != nil {
		return nil, err
	}
	if !exists {
		return nil, errors.NewNotFound(v3.Resource("stagedkubernetesnetworkpolicy"), name)
	}
	return obj.(*v3.StagedKubernetesNetworkPolicy), nil
}
