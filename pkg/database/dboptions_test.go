/*
Copyright 2021 CodeNotary, Inc. All rights reserved.

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

package database

import (
	"testing"

	"github.com/codenotary/immudb/embedded/store"
	"github.com/stretchr/testify/require"
)

func TestDefaultOptions(t *testing.T) {
	op := DefaultOption()
	if op.GetDbName() != "db_name" {
		t.Errorf("default sysdb name not what expected")
	}
	if op.GetDbRootPath() != DefaultOption().dbRootPath {
		t.Errorf("default db rootpath not what expected")
	}
	if op.GetCorruptionChecker() {
		t.Errorf("default corruption checker not what expected")
	}

	DbName := "Charles_Aznavour"
	rootpath := "rootpath"
	storeOpts := store.DefaultOptions()

	replicaOpts := &ReplicationOptions{}
	replicaOpts.AsReplica(true).
		WithSrcDatabase("defaultdb").
		WithSrcAddress("127.0.0.1").
		WithSrcPort(3322).
		WithFollowerUsr("immudb").
		WithFollowerPwd("immdub")

	op = DefaultOption().
		WithDbName(DbName).
		WithDbRootPath(rootpath).
		WithCorruptionChecker(true).
		WithStoreOptions(storeOpts).
		WithReplicationOptions(replicaOpts)

	if op.GetDbName() != DbName {
		t.Errorf("db name not set correctly , expected %s got %s", DbName, op.GetDbName())
	}
	if op.GetDbRootPath() != rootpath {
		t.Errorf("rootpath not set correctly , expected %s got %s", rootpath, op.GetDbRootPath())
	}

	if !op.GetCorruptionChecker() {
		t.Errorf("corruption checker not set correctly , expected %v got %v", true, op.GetCorruptionChecker())
	}

	require.Equal(t, storeOpts, op.storeOpts)
	require.Equal(t, replicaOpts, op.GetReplicationOptions())
}
