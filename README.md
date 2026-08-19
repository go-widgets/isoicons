# go-widgets/isoicons

[![CI](https://github.com/go-widgets/isoicons/actions/workflows/ci.yml/badge.svg)](https://github.com/go-widgets/isoicons/actions/workflows/ci.yml)
[![release](https://img.shields.io/github/v/release/go-widgets/isoicons?display_name=tag&sort=semver&color=0d9488)](https://github.com/go-widgets/isoicons/releases)
[![pkg.go.dev](https://img.shields.io/badge/pkg.go.dev-isoicons-007d9c?logo=go&logoColor=white)](https://pkg.go.dev/github.com/go-widgets/isoicons)
![coverage](https://img.shields.io/badge/coverage-100%25-1a7f37)
[![license](https://img.shields.io/badge/license-BSD--3--Clause-blue)](./LICENSE)

Reusable **isometric icon packs** for the
[go-widgets/toolkit](https://github.com/go-widgets/toolkit) isometric-diagram
widget. Two packs of ready-to-place isometric artwork are registered into a
`toolkit.IsoIconRegistry` as billboarded sprite icons — no drawing code, just
name a node's icon by id.

Pure Go, **CGO_ENABLED=0**, and **`GOOS=js GOARCH=wasm`-clean**: PNG art is
decoded with the standard library's `image/png`, and SVG art is rasterized with
the pure-Go [oksvg](https://github.com/srwiley/oksvg) +
[rasterx](https://github.com/srwiley/rasterx) rasterizer. 100% statement
coverage.

## Packs

| Pack          | id prefix       | concepts | source art                                                                                  | how it loads                          |
| ------------- | --------------- | -------- | ------------------------------------------------------------------------------------------- | ------------------------------------- |
| `cloudnative` | `cloudnative/…` | 44       | 128px PNG, [fjudith/cloud-native-isometric-icons](https://github.com/fjudith/cloud-native-isometric-icons) (Apache-2.0) | `//go:embed` + `image/png` decode     |
| `aws`         | `aws/…`         | 9        | SVG, [danieljoos/isometric-cloud-icons](https://github.com/danieljoos/isometric-cloud-icons) (CC0-1.0)                  | `//go:embed` + oksvg/rasterx raster   |

Every id is namespaced by its pack (`cloudnative/pod`, `aws/aws-lambda`) so the
packs never collide with each other or with the toolkit's built-in icons. Two
`cloudnative` concepts — `container` and `package` — ship **theme-aware**
light/dark variants selected by the `Theme` argument; all others are neutral.

## Install

```sh
go get github.com/go-widgets/isoicons
```

## Usage

```go
package main

import (
	"github.com/go-widgets/isoicons"
	"github.com/go-widgets/toolkit"
)

func main() {
	reg := toolkit.NewIsoIconRegistry()

	// Register both packs; light theme for the theme-aware concepts (default).
	if err := isoicons.RegisterAll(reg); err != nil {
		panic(err)
	}

	// Or a single pack, with an explicit theme:
	//   isoicons.RegisterCloudNative(reg, isoicons.ThemeDark)
	//   isoicons.RegisterAWS(reg)

	// A diagram node can now name any registered id:
	icon, ok := reg.Resolve("cloudnative/pod") // -> toolkit.IsoSpriteIcon, true
	_ = icon
	_ = ok

	// List what is installed, per pack:
	_ = isoicons.CloudNativeIDs() // []string{"cloudnative/apiserver", ...}
	_ = isoicons.AWSIDs()         // []string{"aws/aws-api-gateway", ...}
}
```

## API

```go
type Theme int
const (
	ThemeLight Theme = iota // default
	ThemeDark
)

// Register one pack, both packs, and list the ids of each.
func RegisterCloudNative(reg *toolkit.IsoIconRegistry, theme ...Theme) error
func RegisterAWS(reg *toolkit.IsoIconRegistry) error
func RegisterAll(reg *toolkit.IsoIconRegistry, theme ...Theme) error
func CloudNativeIDs() []string
func AWSIDs() []string
```

`theme` is variadic; omit it for `ThemeLight`. `RegisterAWS` takes no theme (the
AWS art is fixed-colour). Each `Register*` returns an error only if a vendored
asset fails to decode/parse — impossible with the shipped, tested assets, but
surfaced rather than swallowed.

### `cloudnative` concepts (44)

Kubernetes control plane: `apiserver`, `controller-manager`, `scheduler`,
`kubelet`, `kube-proxy`, `cloud-controller-manager`.
Kubernetes infrastructure: `etcd`, `master`, `node`.
Kubernetes resources: `pod`, `deployment`, `statefulset`, `daemonset`, `job`,
`cronjob`, `configmap`, `secret`, `namespace`, `ingress`, `service`,
`persistent-volume`, `persistent-volume-claim`.
Compute: `server`, `server-rack`, `virtual-machine`, `micro-vm`,
`storage-server`, `cloud`.
Networking: `dns`, `internet`, `load-balancer`.
Repositories: `code-repository`, `container-registry`.
Storage: `object-storage`, `folder`, `documents`, `index`.
Pipelines: `deployment-pipeline`, `integration-pipeline`, `drone`.
Data services: `postgresql`, `minio`.
Theme-aware (light/dark): `container`, `package`.

### `aws` concepts (9)

`aws-api-gateway`, `aws-cloudfront`, `aws-cognito-identity-pool`,
`aws-dynamodb`, `aws-lambda`, `aws-s3-bucket`, `black-box`,
`low-block-rounded-blue`, `webapp`.

## Licensing

This repository carries a **triple** licensing regime — the Go code and each
asset pack are licensed independently:

| Component                | License          | Holder                    |
| ------------------------ | ---------------- | ------------------------- |
| Go source code (this repo) | **BSD-3-Clause** | the go-widgets/isoicons authors |
| `assets/cloudnative/*.png` | **Apache-2.0**   | fjudith                   |
| `assets/aws/*.svg`         | **CC0-1.0**      | danieljoos (public domain) |

The assets are vendored **unmodified**. See [`NOTICE`](./NOTICE) for the exact
upstream repositories, commit SHAs, and raw URLs, and each
`assets/<pack>/LICENSE` for the full license text. CC0-1.0 requires no
attribution; the source is recorded only for traceability.

Total vendored asset weight: ~592 KB (540 KB PNG + 52 KB SVG).
