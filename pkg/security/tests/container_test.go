// Unless explicitly stated otherwise all files in this repository are licensed
// under the Apache License Version 2.0.
// This product includes software developed at Datadog (https://www.datadoghq.com/).
// Copyright 2016-present Datadog, Inc.

//go:build linux && functionaltests

// Package tests holds tests related files
package tests

import (
	"os/exec"
	"testing"
	"time"

	"github.com/DataDog/datadog-agent/pkg/security/ebpf/kernel"
	"github.com/DataDog/datadog-agent/pkg/security/secl/containerutils"
	"github.com/DataDog/datadog-agent/pkg/security/secl/model"
	"github.com/DataDog/datadog-agent/pkg/security/secl/rules"
	"github.com/stretchr/testify/assert"
)

func TestContainerCreatedAt(t *testing.T) {
	SkipIfNotAvailable(t)

	checkKernelCompatibility(t, "OpenSUSE 15.3 kernel", func(kv *kernel.Version) bool {
		// because of some strange btrfs subvolume error
		return kv.IsOpenSUSELeap15_3Kernel()
	})

	ruleDefs := []*rules.RuleDefinition{
		{
			ID:         "test_container_created_at",
			Expression: `container.id != "" && container.created_at < 3s && open.file.path == "{{.Root}}/test-open"`,
		},
		{
			ID:         "test_container_created_at_delay",
			Expression: `container.id != "" && container.created_at > 3s && open.file.path == "{{.Root}}/test-open-delay"`,
		},
	}
	test, err := newTestModule(t, nil, ruleDefs)
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-open")
	if err != nil {
		t.Fatal(err)
	}

	testFileDelay, _, err := test.Path("test-open-delay")
	if err != nil {
		t.Fatal(err)
	}

	dockerWrapper, err := newDockerCmdWrapper(test.Root(), test.Root(), "ubuntu", "")
	if err != nil {
		t.Skip("Skipping created time in containers tests: Docker not available")
		return
	}
	defer dockerWrapper.stop()

	dockerWrapper.Run(t, "container-created-at", func(t *testing.T, _ wrapperType, cmdFunc func(cmd string, args []string, envs []string) *exec.Cmd) {
		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFile}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_created_at")
			assertFieldEqual(t, event, "open.file.path", testFile)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")

			test.validateOpenSchema(t, event)
		})
	})

	dockerWrapper.Run(t, "container-created-at-delay", func(t *testing.T, _ wrapperType, cmdFunc func(cmd string, args []string, envs []string) *exec.Cmd) {
		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFileDelay}, nil) // shouldn't trigger an event
			if err := cmd.Run(); err != nil {
				return err
			}
			time.Sleep(3 * time.Second)
			cmd = cmdFunc("touch", []string{testFileDelay}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_created_at_delay")
			assertFieldEqual(t, event, "open.file.path", testFileDelay)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")
			assert.Equal(t, event.CGroupContext.CGroupFlags, containerutils.CGroupFlags(containerutils.CGroupManagerDocker))
			createdAtNano, _ := event.GetFieldValue("container.created_at")
			createdAt := time.Unix(0, int64(createdAtNano.(int)))
			assert.True(t, time.Since(createdAt) > 3*time.Second)

			test.validateOpenSchema(t, event)
		})
	})
}

func TestContainerFlagsDocker(t *testing.T) {
	SkipIfNotAvailable(t)

	ruleDefs := []*rules.RuleDefinition{
		{
			ID:         "test_container_flags",
			Expression: `container.runtime == "docker" && container.id != "" && open.file.path == "{{.Root}}/test-open" && cgroup.id =~ "*docker*"`,
		},
	}
	test, err := newTestModule(t, nil, ruleDefs)
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-open")
	if err != nil {
		t.Fatal(err)
	}

	dockerWrapper, err := newDockerCmdWrapper(test.Root(), test.Root(), "ubuntu", "")
	if err != nil {
		t.Skipf("Skipping container test: Docker not available (%s)", err.Error())
		return
	}
	defer dockerWrapper.stop()

	dockerWrapper.Run(t, "container-runtime", func(t *testing.T, _ wrapperType, cmdFunc func(cmd string, args []string, envs []string) *exec.Cmd) {
		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFile}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_flags")
			assertFieldEqual(t, event, "open.file.path", testFile)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")
			assertFieldEqual(t, event, "container.runtime", "docker")
			assert.Equal(t, containerutils.CGroupFlags(containerutils.CGroupManagerDocker), event.CGroupContext.CGroupFlags)

			test.validateOpenSchema(t, event)
		})
	})
}

func TestContainerFlagsPodman(t *testing.T) {
	SkipIfNotAvailable(t)

	ruleDefs := []*rules.RuleDefinition{
		{
			ID:         "test_container_flags",
			Expression: `container.runtime == "podman" && container.id != "" && open.file.path == "{{.Root}}/test-open" && cgroup.id =~ "*libpod*"`,
		},
	}
	test, err := newTestModule(t, nil, ruleDefs)
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-open")
	if err != nil {
		t.Fatal(err)
	}

	podmanWrapper, err := newDockerCmdWrapper(test.Root(), test.Root(), "ubuntu", string(podmanWrapperType))
	if err != nil {
		t.Skip("Skipping created time in containers tests: podman not available")
		return
	}
	defer podmanWrapper.stop()

	podmanWrapper.Run(t, "container-runtime", func(t *testing.T, _ wrapperType, cmdFunc func(cmd string, args []string, envs []string) *exec.Cmd) {
		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFile}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_flags")
			assertFieldEqual(t, event, "open.file.path", testFile)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")
			assertFieldEqual(t, event, "container.runtime", "podman")
			assert.Equal(t, containerutils.CGroupFlags(containerutils.CGroupManagerPodman), event.CGroupContext.CGroupFlags)

			test.validateOpenSchema(t, event)
		})
	})
}

func TestContainerVariables(t *testing.T) {
	SkipIfNotAvailable(t)

	ruleDefs := []*rules.RuleDefinition{
		{
			ID:         "test_container_set_variable",
			Expression: `container.id != "" && open.file.path == "{{.Root}}/test-open"`,
			Actions: []*rules.ActionDefinition{
				{
					Set: &rules.SetDefinition{
						Scope: "container",
						Value: 1,
						Name:  "foo",
					},
				},
			},
		},
		{
			ID:         "test_container_check_variable",
			Expression: `container.id != "" && open.file.path == "{{.Root}}/test-open2" && ${container.foo} == 1`,
		},
	}
	test, err := newTestModule(t, nil, ruleDefs)
	if err != nil {
		t.Fatal(err)
	}
	defer test.Close()

	testFile, _, err := test.Path("test-open")
	if err != nil {
		t.Fatal(err)
	}

	testFile2, _, err := test.Path("test-open2")
	if err != nil {
		t.Fatal(err)
	}

	dockerWrapper, err := newDockerCmdWrapper(test.Root(), test.Root(), "ubuntu", "")
	if err != nil {
		t.Skip("Skipping created time in containers tests: Docker not available")
		return
	}
	defer dockerWrapper.stop()

	dockerWrapper.Run(t, "container-variables", func(t *testing.T, _ wrapperType, cmdFunc func(cmd string, args []string, envs []string) *exec.Cmd) {
		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFile}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_set_variable")
			assertFieldEqual(t, event, "open.file.path", testFile)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")

			test.validateOpenSchema(t, event)
		})

		test.WaitSignal(t, func() error {
			cmd := cmdFunc("touch", []string{testFile2}, nil)
			return cmd.Run()
		}, func(event *model.Event, rule *rules.Rule) {
			assertTriggeredRule(t, rule, "test_container_check_variable")
			assertFieldEqual(t, event, "open.file.path", testFile2)
			assertFieldNotEmpty(t, event, "container.id", "container id shouldn't be empty")

			test.validateOpenSchema(t, event)
		})
	})
}
