// Copyright (c) 2026 the go-widgets/isoicons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package isoicons

import "embed"

// cloudnativeFS holds the vendored 128px PNG art of the "cloudnative" pack.
// The assets are unmodified isometric icons from
// github.com/fjudith/cloud-native-isometric-icons (Apache-2.0); see NOTICE and
// assets/cloudnative/LICENSE.
//
//go:embed assets/cloudnative/*.png
var cloudnativeFS embed.FS

// cnAsset reads one "cloudnative" PNG from the embedded filesystem.
func cnAsset(name string) ([]byte, error) {
	return cloudnativeFS.ReadFile("assets/cloudnative/" + name)
}

// cnConcept maps a stable pack-local id to its embedded PNG file name. A
// neutral concept sets file; a theme-aware concept leaves file empty and sets
// light and dark instead.
type cnConcept struct {
	id          string
	file        string
	light, dark string
}

// assetName returns the embedded file name to load for theme t.
func (c cnConcept) assetName(t Theme) string {
	if c.file != "" {
		return c.file
	}
	if t == ThemeDark {
		return c.dark
	}
	return c.light
}

// cloudNativeConcepts is the curated concept table of the "cloudnative" pack,
// in stable listing order. Its ids become "cloudnative/<id>". "container" and
// "package" ship theme-aware light/dark variants; the rest are neutral.
var cloudNativeConcepts = []cnConcept{
	// Kubernetes control plane.
	{id: "apiserver", file: "apiserver.png"},
	{id: "controller-manager", file: "controller-manager.png"},
	{id: "scheduler", file: "scheduler.png"},
	{id: "kubelet", file: "kubelet.png"},
	{id: "kube-proxy", file: "kube-proxy.png"},
	{id: "cloud-controller-manager", file: "cloud-controller-manager.png"},
	// Kubernetes infrastructure.
	{id: "etcd", file: "etcd.png"},
	{id: "master", file: "master.png"},
	{id: "node", file: "node.png"},
	// Kubernetes resources.
	{id: "pod", file: "pod.png"},
	{id: "deployment", file: "deployment.png"},
	{id: "statefulset", file: "statefulset.png"},
	{id: "daemonset", file: "daemonset.png"},
	{id: "job", file: "job.png"},
	{id: "cronjob", file: "cronjob.png"},
	{id: "configmap", file: "configmap.png"},
	{id: "secret", file: "secret.png"},
	{id: "namespace", file: "namespace.png"},
	{id: "ingress", file: "ingress.png"},
	{id: "service", file: "service.png"},
	{id: "persistent-volume", file: "persistent-volume.png"},
	{id: "persistent-volume-claim", file: "persistent-volume-claim.png"},
	// Compute / servers.
	{id: "server", file: "server.png"},
	{id: "server-rack", file: "server-rack.png"},
	{id: "virtual-machine", file: "virtual-machine.png"},
	{id: "micro-vm", file: "micro-vm.png"},
	{id: "storage-server", file: "storage-server.png"},
	{id: "cloud", file: "cloud.png"},
	// Networking.
	{id: "dns", file: "dns.png"},
	{id: "internet", file: "internet.png"},
	{id: "load-balancer", file: "load-balancer.png"},
	// Repositories / registries.
	{id: "code-repository", file: "code-repository.png"},
	{id: "container-registry", file: "container-registry.png"},
	// Storage.
	{id: "object-storage", file: "object-storage.png"},
	{id: "folder", file: "folder.png"},
	{id: "documents", file: "documents.png"},
	{id: "index", file: "index.png"},
	// Pipelines / software factory.
	{id: "deployment-pipeline", file: "deployment-pipeline.png"},
	{id: "integration-pipeline", file: "integration-pipeline.png"},
	{id: "drone", file: "drone.png"},
	// Data services.
	{id: "postgresql", file: "postgresql.png"},
	{id: "minio", file: "minio.png"},
	// Theme-aware concepts (light + dark variants).
	{id: "container", light: "container_light.png", dark: "container_dark.png"},
	{id: "package", light: "package_light.png", dark: "package_dark.png"},
}
