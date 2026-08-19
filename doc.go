// Copyright (c) 2026 the go-widgets/isoicons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

// Package isoicons provides reusable isometric icon packs for the
// github.com/go-widgets/toolkit isometric-diagram widget.
//
// Two packs are shipped, each vendored from a permissively licensed upstream
// and registered into a toolkit.IsoIconRegistry as billboarded sprite icons
// (toolkit.IsoSpriteIcon):
//
//   - "cloudnative": ~45 cloud-native concepts (Kubernetes control-plane and
//     resources, servers, networking, storage, pipelines, …) decoded from
//     embedded 128px PNG art. Two concepts ("container", "package") ship
//     theme-aware light/dark variants selected by the Theme argument.
//   - "aws": the nine isometric AWS/cloud SVGs, rasterized at load time with
//     the pure-Go github.com/srwiley/oksvg + github.com/srwiley/rasterx
//     rasterizer (CGO-free, wasm-clean).
//
// Every id is namespaced by its pack ("cloudnative/pod", "aws/aws-lambda") so
// the two packs never collide with each other or with the toolkit's built-in
// icons. All code in this repository is BSD-3-Clause; the vendored assets keep
// their upstream licenses (assets/cloudnative is Apache-2.0, assets/aws is
// CC0-1.0) — see the NOTICE file and each assets/*/LICENSE.
//
// The package is pure Go (CGO_ENABLED=0) and compiles for GOOS=js
// GOARCH=wasm: the PNG decoder is image/png from the standard library and the
// SVG rasterizer is pure Go.
package isoicons
