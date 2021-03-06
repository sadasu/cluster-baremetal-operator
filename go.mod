module github.com/openshift/cluster-baremetal-operator

go 1.15

require (
	github.com/ghodss/yaml v1.0.0
	github.com/go-bindata/go-bindata v3.1.2+incompatible
	github.com/golangci/golangci-lint v1.32.0
	github.com/google/go-cmp v0.5.2
	github.com/metal3-io/baremetal-operator v0.0.0-00010101000000-000000000000
	github.com/openshift/api v0.0.0-20201214114959-164a2fb63b5f
	github.com/openshift/client-go v0.0.0-20201020074620-f8fd44879f7c
	github.com/openshift/library-go v0.0.0-20201203122949-352bc2d14339
	github.com/pkg/errors v0.9.1
	github.com/stretchr/stew v0.0.0-20130812190256-80ef0842b48b
	github.com/stretchr/testify v1.6.1
	golang.org/x/crypto v0.0.0-20201221181555-eec23a3978ad
	k8s.io/api v0.20.2
	k8s.io/apimachinery v0.20.2
	k8s.io/client-go v0.20.2
	k8s.io/klog/v2 v2.4.0
	k8s.io/utils v0.0.0-20210111153108-fddb29f9d009
	sigs.k8s.io/controller-runtime v0.8.3
	sigs.k8s.io/controller-tools v0.4.1
	sigs.k8s.io/kustomize/kustomize/v3 v3.9.4
)

replace github.com/metal3-io/baremetal-operator => github.com/openshift/baremetal-operator v0.0.0-20210527161605-4e331bfd4b1d
