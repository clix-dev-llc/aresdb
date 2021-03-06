//  Copyright (c) 2017-2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package common

import (
	"github.com/m3db/m3/src/cluster/placement"
	"github.com/uber/aresdb/controller/models"
	"github.com/uber/aresdb/metastore/common"
)

// TableSchemaMutator mutates table metadata
type TableSchemaMutator interface {
	ListTables(namespace string) ([]string, error)
	GetTable(namespace, name string) (*common.Table, error)
	CreateTable(namespace string, table *common.Table, force bool) error
	DeleteTable(namespace, name string) error
	UpdateTable(namespace string, table common.Table, force bool) error
	GetHash(namespace string) (string, error)
}

// SubscriberMutator defines rw operations
// only read ops needed for now, creating and removing subscriber should be done by each subscriber process
type SubscriberMutator interface {
	// GetSubscriber returns a subscriber
	GetSubscriber(namespace, subscriberName string) (models.Subscriber, error)
	// GetSubscribers returns a list of subscribers
	GetSubscribers(namespace string) ([]models.Subscriber, error)
	// GetHash returns hash of all subscribers
	GetHash(namespace string) (string, error)
}

// NamespaceMutator mutates table metadata
type NamespaceMutator interface {
	CreateNamespace(namespace string) error
	ListNamespaces() ([]string, error)
}

// JobMutator defines rw operations
type JobMutator interface {
	GetJob(namespace, name string) (job models.JobConfig, err error)
	GetJobs(namespace string) (job []models.JobConfig, err error)
	DeleteJob(namespace, name string) error
	UpdateJob(namespace string, job models.JobConfig) error
	AddJob(namespace string, job models.JobConfig) error
	GetHash(namespace string) (string, error)
}

// IngestionAssignmentMutator defines rw operations
type IngestionAssignmentMutator interface {
	GetIngestionAssignment(namespace, name string) (IngestionAssignment models.IngestionAssignment, err error)
	GetIngestionAssignments(namespace string) (IngestionAssignment []models.IngestionAssignment, err error)
	DeleteIngestionAssignment(namespace, name string) error
	UpdateIngestionAssignment(namespace string, IngestionAssignment models.IngestionAssignment) error
	AddIngestionAssignment(namespace string, IngestionAssignment models.IngestionAssignment) error
	GetHash(namespace, subscriber string) (string, error)
}

// EnumReader reads enum cases
type EnumReader interface {
	// GetEnumCases get all enum cases for the given table column
	GetEnumCases(namespace, table, column string) ([]string, error)
}

// EnumMutator defines EnumMutator interface
type EnumMutator interface {
	EnumReader
	// ExtendEnumCases try to extend new enum cases to given column
	ExtendEnumCases(namespace, table, column string, enumCases []string) ([]int, error)
}

// MembershipMutator defines membership rw operations
type MembershipMutator interface {
	// Join registers an instance to a namespace
	Join(namespace string, instance models.Instance) error
	// GetInstance returns an instance
	GetInstance(namespace, instanceName string) (models.Instance, error)
	// GetInstances returns a list of instances in a namespace
	GetInstances(namespace string) ([]models.Instance, error)
	// Leave removes an instance
	Leave(namespace, instanceName string) error
	// GetHash returns hash of all instances
	GetHash(namespace string) (string, error)
}

// PlacementServiceBuilder defines mutator interface for placement objects
type PlacementMutator interface {
	BuildInitialPlacement(namespace string, numShards int, numReplica int, instances []placement.Instance) (placement.Placement, error)
	GetCurrentPlacement(namespace string) (placement.Placement, error)
	AddInstance(namespace string, instances []placement.Instance) (placement.Placement, error)
	ReplaceInstance(namespace string, leavingInstances []string, newInstances []placement.Instance) (placement.Placement, error)
	RemoveInstance(namespace string, leavingInstances []string) (placement.Placement, error)
	MarkNamespaceAvailable(namespace string) (placement.Placement, error)
	MarkInstanceAvailable(namespace string, instance string) (placement.Placement, error)
	MarkShardsAvailable(namespace string, instance string, shards []uint32) (placement.Placement, error)
}
