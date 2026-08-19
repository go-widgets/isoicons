// Copyright (c) 2026 the go-widgets/isoicons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package isoicons

import (
	"bytes"
	"fmt"
	"image"
	"image/png"

	"github.com/go-gfx/gfx/raster"
	"github.com/go-widgets/toolkit"
	"github.com/srwiley/oksvg"
	"github.com/srwiley/rasterx"
)

// Theme selects, for a concept that ships two variants, which one is
// registered. Concepts with a single neutral variant ignore it.
type Theme int

const (
	// ThemeLight registers the light-background variant (the default).
	ThemeLight Theme = iota
	// ThemeDark registers the dark-background variant.
	ThemeDark
)

// Pack id prefixes. Every registered id is "<prefix>/<concept>", so the two
// packs never collide with each other or with the toolkit's built-in icons.
const (
	cloudNativePrefix = "cloudnative"
	awsPrefix         = "aws"
)

// assetFunc reads one vendored asset's raw bytes by file name. It is the seam
// the Register* functions load through, so a test can drive the decode/error
// paths without a corrupt file on disk.
type assetFunc func(name string) ([]byte, error)

// decodePNG decodes an embedded PNG into a straight-alpha raster image.
func decodePNG(data []byte) (*raster.Image, error) {
	img, err := png.Decode(bytes.NewReader(data))
	if err != nil {
		return nil, fmt.Errorf("isoicons: decode png: %w", err)
	}
	return raster.FromImage(img), nil
}

// rasterizeSVG parses and rasterizes an embedded SVG into a w by h straight-
// alpha raster image using the pure-Go oksvg + rasterx rasterizer.
func rasterizeSVG(data []byte, w, h int) (*raster.Image, error) {
	icon, err := oksvg.ReadIconStream(bytes.NewReader(data), oksvg.IgnoreErrorMode)
	if err != nil {
		return nil, fmt.Errorf("isoicons: parse svg: %w", err)
	}
	icon.SetTarget(0, 0, float64(w), float64(h))
	rgba := image.NewRGBA(image.Rect(0, 0, w, h))
	dasher := rasterx.NewDasher(w, h, rasterx.NewScannerGV(w, h, rgba, rgba.Bounds()))
	icon.Draw(dasher, 1.0)
	return raster.FromImage(rgba), nil
}

// RegisterCloudNative registers the "cloudnative" pack into reg. The optional
// theme (default [ThemeLight]) selects the variant for the theme-aware
// concepts; every other concept ships a single neutral variant. It returns an
// error only if a vendored asset fails to decode.
func RegisterCloudNative(reg *toolkit.IsoIconRegistry, theme ...Theme) error {
	t := ThemeLight
	if len(theme) > 0 {
		t = theme[0]
	}
	return registerCloudNative(reg, t, cnAsset)
}

// registerCloudNative is the loader-injectable core of [RegisterCloudNative].
func registerCloudNative(reg *toolkit.IsoIconRegistry, t Theme, load assetFunc) error {
	for _, c := range cloudNativeConcepts {
		data, err := load(c.assetName(t))
		if err != nil {
			return err
		}
		img, err := decodePNG(data)
		if err != nil {
			return err
		}
		reg.Register(cloudNativePrefix+"/"+c.id, toolkit.IsoSpriteIcon{Img: img})
	}
	return nil
}

// RegisterAWS registers the "aws" pack into reg, rasterizing each vendored SVG
// at load time. It returns an error only if a vendored asset fails to parse.
func RegisterAWS(reg *toolkit.IsoIconRegistry) error {
	return registerAWS(reg, awsAsset)
}

// registerAWS is the loader-injectable core of [RegisterAWS].
func registerAWS(reg *toolkit.IsoIconRegistry, load assetFunc) error {
	for _, name := range awsConcepts {
		data, err := load(name + ".svg")
		if err != nil {
			return err
		}
		img, err := rasterizeSVG(data, awsSize, awsSize)
		if err != nil {
			return err
		}
		reg.Register(awsPrefix+"/"+name, toolkit.IsoSpriteIcon{Img: img})
	}
	return nil
}

// RegisterAll registers both packs into reg. The optional theme is forwarded to
// the "cloudnative" pack (the "aws" pack is fixed-colour).
func RegisterAll(reg *toolkit.IsoIconRegistry, theme ...Theme) error {
	t := ThemeLight
	if len(theme) > 0 {
		t = theme[0]
	}
	return registerAll(reg, t, cnAsset, awsAsset)
}

// registerAll is the loader-injectable core of [RegisterAll].
func registerAll(reg *toolkit.IsoIconRegistry, t Theme, cn, aws assetFunc) error {
	if err := registerCloudNative(reg, t, cn); err != nil {
		return err
	}
	return registerAWS(reg, aws)
}

// CloudNativeIDs returns the registered ids of the "cloudnative" pack, in a
// stable order — for building a palette or listing the installed library.
func CloudNativeIDs() []string {
	out := make([]string, len(cloudNativeConcepts))
	for i, c := range cloudNativeConcepts {
		out[i] = cloudNativePrefix + "/" + c.id
	}
	return out
}

// AWSIDs returns the registered ids of the "aws" pack, in a stable order.
func AWSIDs() []string {
	out := make([]string, len(awsConcepts))
	for i, name := range awsConcepts {
		out[i] = awsPrefix + "/" + name
	}
	return out
}
