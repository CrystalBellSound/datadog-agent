// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

// Package core implements the "core" bundle, providing services common to all
// agent flavors and binaries.
//
// The constituent components serve as utilities and are mostly independent of
// one another.  Other components should depend on any components they need.
//
// This bundle does not depend on any other bundles.

//go:build test

package core

import (
	"testing"

	"go.uber.org/fx"

	"github.com/DataDog/datadog-agent/comp/core/config"
	"github.com/DataDog/datadog-agent/comp/core/hostname/hostnameimpl"
	log "github.com/DataDog/datadog-agent/comp/core/log/def"
	logmock "github.com/DataDog/datadog-agent/comp/core/log/mock"
	"github.com/DataDog/datadog-agent/comp/core/sysprobeconfig/sysprobeconfigimpl"
	"github.com/DataDog/datadog-agent/comp/core/telemetry/telemetryimpl"
	"github.com/DataDog/datadog-agent/pkg/util/fxutil"
)

// team: agent-runtimes

// MakeMockBundle returns a core bundle with a customized set of fx.Option including sane defaults.
func MakeMockBundle(logParams, logger fx.Option) fxutil.BundleOptions {
	return fxutil.Bundle(
		fx.Provide(func(params BundleParams) config.Params { return params.ConfigParams }),
		config.MockModule(),
		logParams,
		logger,
		fx.Provide(func(params BundleParams) sysprobeconfigimpl.Params { return params.SysprobeConfigParams }),
		sysprobeconfigimpl.MockModule(),
		telemetryimpl.MockModule(),
		hostnameimpl.MockModule(),
	)
}

// MockBundle defines the mock fx options for this bundle.
func MockBundle() fxutil.BundleOptions {
	return MakeMockBundle(
		fx.Supply(log.Params{}),
		fx.Provide(func(t testing.TB) log.Component { return logmock.New(t) }),
	)
}
