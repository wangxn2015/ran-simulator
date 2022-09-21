// SPDX-FileCopyrightText: 2020-present Open Networking Foundation <info@opennetworking.org>
//
// SPDX-License-Identifier: Apache-2.0

package main

import (
	"github.com/wangxn2015/helmit/pkg/registry"
	"github.com/wangxn2015/helmit/pkg/test"
	"github.com/wangxn2015/ran-simulator/tests/e2t"
	"github.com/wangxn2015/ran-simulator/tests/ransim"
)

func main() {
	registry.RegisterTestSuite("e2t", &e2t.TestSuite{})
	registry.RegisterTestSuite("ransim", &ransim.TestSuite{})
	test.Main()
}
