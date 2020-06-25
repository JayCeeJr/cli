// Copyright (c) 2020 Target Brands, Inc. All rights reserved.
//
// Use of this source code is governed by the LICENSE file in this repository.

package action

import (
	"flag"
	"net/http/httptest"
	"testing"

	"github.com/go-vela/mock/server"

	"github.com/urfave/cli/v2"
)

func TestAction_SecretView(t *testing.T) {
	// setup test server
	s := httptest.NewServer(server.FakeHandler())

	// setup flags
	authSet := flag.NewFlagSet("test", 0)
	authSet.String("addr", s.URL, "doc")
	authSet.String("token", "superSecretToken", "doc")

	fullSet := flag.NewFlagSet("test", 0)
	fullSet.String("addr", s.URL, "doc")
	fullSet.String("token", "superSecretToken", "doc")
	fullSet.String("engine", "native", "doc")
	fullSet.String("type", "repo", "doc")
	fullSet.String("org", "github", "doc")
	fullSet.String("repo", "octocat", "doc")
	fullSet.String("name", "foo", "doc")
	fullSet.String("output", "json", "doc")

	// setup tests
	tests := []struct {
		failure bool
		set     *flag.FlagSet
	}{
		{
			failure: false,
			set:     fullSet,
		},
		{
			failure: true,
			set:     authSet,
		},
		{
			failure: true,
			set:     flag.NewFlagSet("test", 0),
		},
	}

	// run tests
	for _, test := range tests {
		err := secretView(cli.NewContext(nil, test.set, nil))

		if test.failure {
			if err == nil {
				t.Errorf("secretView should have returned err")
			}

			continue
		}

		if err != nil {
			t.Errorf("secretView returned err: %v", err)
		}
	}
}