/*
Copyright 2018 The Kubernetes Authors.

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

package controller_test

import (
	"fmt"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/tsungming/controller-runtime/pkg/client"
	"github.com/tsungming/controller-runtime/pkg/controller"
	"github.com/tsungming/controller-runtime/pkg/manager"
	"github.com/tsungming/controller-runtime/pkg/reconcile"
	"github.com/tsungming/controller-runtime/pkg/runtime/inject"
)

var _ = Describe("controller.Controller", func() {
	var stop chan struct{}

	rec := reconcile.Func(func(reconcile.Request) (reconcile.Result, error) {
		return reconcile.Result{}, nil
	})
	BeforeEach(func() {
		stop = make(chan struct{})
	})

	AfterEach(func() {
		close(stop)
	})

	Describe("New", func() {
		It("should return an error if Name is not Specified", func(done Done) {
			m, err := manager.New(cfg, manager.Options{})
			Expect(err).NotTo(HaveOccurred())
			c, err := controller.New("", m, controller.Options{Reconcile: rec})
			Expect(c).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("must specify Name for Controller"))

			close(done)
		})

		It("should return an error if Reconcile is not Specified", func(done Done) {
			m, err := manager.New(cfg, manager.Options{})
			Expect(err).NotTo(HaveOccurred())

			c, err := controller.New("foo", m, controller.Options{})
			Expect(c).To(BeNil())
			Expect(err.Error()).To(ContainSubstring("must specify Reconcile"))

			close(done)
		})

		It("NewController should return an error if injecting Reconcile fails", func(done Done) {
			m, err := manager.New(cfg, manager.Options{})
			Expect(err).NotTo(HaveOccurred())

			c, err := controller.New("foo", m, controller.Options{Reconcile: &failRec{}})
			Expect(c).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err.Error()).To(ContainSubstring("expected error"))

			close(done)
		})
	})
})

var _ reconcile.Reconcile = &failRec{}
var _ inject.Client = &failRec{}

type failRec struct{}

func (*failRec) Reconcile(reconcile.Request) (reconcile.Result, error) {
	return reconcile.Result{}, nil
}

func (*failRec) InjectClient(client.Client) error {
	return fmt.Errorf("expected error")
}
