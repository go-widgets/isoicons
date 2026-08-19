// Copyright (c) 2026 the go-widgets/isoicons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package isoicons

import (
	"errors"
	"strings"
	"testing"

	"github.com/go-gfx/gfx/raster"
	"github.com/go-widgets/toolkit"
)

// spriteOf resolves id, asserting it is a registered toolkit.IsoSpriteIcon with
// a non-nil sprite image, and returns that image.
func spriteOf(t *testing.T, reg *toolkit.IsoIconRegistry, id string) *raster.Image {
	t.Helper()
	icon, ok := reg.Resolve(id)
	if !ok {
		t.Fatalf("id %q did not resolve", id)
	}
	sprite, ok := icon.(toolkit.IsoSpriteIcon)
	if !ok {
		t.Fatalf("id %q resolved to %T, want toolkit.IsoSpriteIcon", id, icon)
	}
	if sprite.Img == nil {
		t.Fatalf("id %q has a nil sprite image", id)
	}
	return sprite.Img
}

// hasOpaquePixel reports whether img has at least one non-transparent pixel —
// proof the asset actually decoded/rasterized to visible art.
func hasOpaquePixel(img *raster.Image) bool {
	for i := 3; i < len(img.Pix); i += 4 {
		if img.Pix[i] != 0 {
			return true
		}
	}
	return false
}

// TestRegisterCloudNativeResolvesEveryConcept is the toothed assertion for the
// cloudnative pack: every curated id resolves to a sprite whose decoded image
// is 128px wide, has positive height, and carries visible (non-transparent)
// pixels.
func TestRegisterCloudNativeResolvesEveryConcept(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := RegisterCloudNative(reg); err != nil {
		t.Fatalf("RegisterCloudNative: %v", err)
	}
	ids := CloudNativeIDs()
	if len(ids) != len(cloudNativeConcepts) {
		t.Fatalf("CloudNativeIDs len = %d, want %d", len(ids), len(cloudNativeConcepts))
	}
	for _, id := range ids {
		if !strings.HasPrefix(id, cloudNativePrefix+"/") {
			t.Fatalf("id %q not namespaced by pack", id)
		}
		img := spriteOf(t, reg, id)
		if img.W != 128 {
			t.Fatalf("id %q width = %d, want 128", id, img.W)
		}
		if img.H <= 0 {
			t.Fatalf("id %q height = %d, want > 0", id, img.H)
		}
		if len(img.Pix) != 4*img.W*img.H {
			t.Fatalf("id %q Pix len = %d, want %d", id, len(img.Pix), 4*img.W*img.H)
		}
		if !hasOpaquePixel(img) {
			t.Fatalf("id %q decoded to a fully transparent image", id)
		}
	}
}

// TestCloudNativeThemeAwareVariantsDiffer proves the theme-aware concepts embed
// two genuinely different variants: the light and dark sprites for the same id
// must not be byte-identical.
func TestCloudNativeThemeAwareVariantsDiffer(t *testing.T) {
	light := toolkit.NewIsoIconRegistry()
	dark := toolkit.NewIsoIconRegistry()
	if err := RegisterCloudNative(light, ThemeLight); err != nil {
		t.Fatalf("RegisterCloudNative light: %v", err)
	}
	if err := RegisterCloudNative(dark, ThemeDark); err != nil {
		t.Fatalf("RegisterCloudNative dark: %v", err)
	}
	for _, id := range []string{"cloudnative/container", "cloudnative/package"} {
		lp := spriteOf(t, light, id).Pix
		dp := spriteOf(t, dark, id).Pix
		if len(lp) == len(dp) && string(lp) == string(dp) {
			t.Fatalf("id %q: light and dark variants are byte-identical", id)
		}
	}
	// A neutral concept resolves to the same art regardless of theme.
	lp := spriteOf(t, light, "cloudnative/pod").Pix
	dp := spriteOf(t, dark, "cloudnative/pod").Pix
	if string(lp) != string(dp) {
		t.Fatalf("neutral concept pod differs between themes")
	}
}

// TestRegisterAWSResolvesEveryConcept is the toothed assertion for the aws
// pack: every SVG rasterizes to a 128x128 sprite with visible pixels.
func TestRegisterAWSResolvesEveryConcept(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := RegisterAWS(reg); err != nil {
		t.Fatalf("RegisterAWS: %v", err)
	}
	ids := AWSIDs()
	if len(ids) != len(awsConcepts) {
		t.Fatalf("AWSIDs len = %d, want %d", len(ids), len(awsConcepts))
	}
	for _, id := range ids {
		if !strings.HasPrefix(id, awsPrefix+"/") {
			t.Fatalf("id %q not namespaced by pack", id)
		}
		img := spriteOf(t, reg, id)
		if img.W != awsSize || img.H != awsSize {
			t.Fatalf("id %q size = %dx%d, want %dx%d", id, img.W, img.H, awsSize, awsSize)
		}
		if !hasOpaquePixel(img) {
			t.Fatalf("id %q rasterized to a fully transparent image", id)
		}
	}
}

// TestRegisterAll registers both packs and checks the combined id set.
func TestRegisterAll(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := RegisterAll(reg); err != nil {
		t.Fatalf("RegisterAll: %v", err)
	}
	if err := RegisterAll(reg, ThemeDark); err != nil {
		t.Fatalf("RegisterAll dark: %v", err)
	}
	want := len(cloudNativeConcepts) + len(awsConcepts)
	if got := len(reg.IDs()); got != want {
		t.Fatalf("registry has %d ids, want %d", got, want)
	}
	// Spot-check one id from each pack still resolves as a sprite.
	spriteOf(t, reg, "cloudnative/pod")
	spriteOf(t, reg, "aws/aws-lambda")
}

// TestIDsAreStableAndDistinct checks the two id lists share no member and each
// is internally unique.
func TestIDsAreStableAndDistinct(t *testing.T) {
	seen := map[string]bool{}
	for _, id := range append(CloudNativeIDs(), AWSIDs()...) {
		if seen[id] {
			t.Fatalf("duplicate id %q", id)
		}
		seen[id] = true
	}
}

// --- error branches ---------------------------------------------------------

var errBoom = errors.New("boom")

func failLoad(string) ([]byte, error)    { return nil, errBoom }
func corruptLoad(string) ([]byte, error) { return []byte("not a valid asset"), nil }

func TestDecodePNGError(t *testing.T) {
	if _, err := decodePNG([]byte("not a png")); err == nil {
		t.Fatal("decodePNG(garbage) = nil error, want error")
	}
}

func TestDecodePNGSuccess(t *testing.T) {
	data, err := cnAsset("pod.png")
	if err != nil {
		t.Fatalf("cnAsset: %v", err)
	}
	img, err := decodePNG(data)
	if err != nil {
		t.Fatalf("decodePNG: %v", err)
	}
	if img.W != 128 {
		t.Fatalf("width = %d, want 128", img.W)
	}
}

func TestRasterizeSVGError(t *testing.T) {
	if _, err := rasterizeSVG([]byte("<svg><unclosed"), 16, 16); err == nil {
		t.Fatal("rasterizeSVG(malformed) = nil error, want error")
	}
}

func TestRegisterCloudNativeLoadError(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := registerCloudNative(reg, ThemeLight, failLoad); !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestRegisterCloudNativeDecodeError(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := registerCloudNative(reg, ThemeLight, corruptLoad); err == nil {
		t.Fatal("registerCloudNative(corrupt) = nil error, want error")
	}
}

func TestRegisterAWSLoadError(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := registerAWS(reg, failLoad); !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestRegisterAWSParseError(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	bad := func(string) ([]byte, error) { return []byte("<svg><unclosed"), nil }
	if err := registerAWS(reg, bad); err == nil {
		t.Fatal("registerAWS(malformed) = nil error, want error")
	}
}

func TestRegisterAllPropagatesCloudNativeError(t *testing.T) {
	reg := toolkit.NewIsoIconRegistry()
	if err := registerAll(reg, ThemeLight, failLoad, awsAsset); !errors.Is(err, errBoom) {
		t.Fatalf("err = %v, want errBoom", err)
	}
}

func TestAssetNameThemeSelection(t *testing.T) {
	neutral := cnConcept{id: "x", file: "x.png"}
	if got := neutral.assetName(ThemeDark); got != "x.png" {
		t.Fatalf("neutral assetName = %q, want x.png", got)
	}
	themed := cnConcept{id: "y", light: "y_light.png", dark: "y_dark.png"}
	if got := themed.assetName(ThemeLight); got != "y_light.png" {
		t.Fatalf("themed light assetName = %q, want y_light.png", got)
	}
	if got := themed.assetName(ThemeDark); got != "y_dark.png" {
		t.Fatalf("themed dark assetName = %q, want y_dark.png", got)
	}
}

func TestAssetReadersMissingFile(t *testing.T) {
	if _, err := cnAsset("does-not-exist.png"); err == nil {
		t.Fatal("cnAsset(missing) = nil error, want error")
	}
	if _, err := awsAsset("does-not-exist.svg"); err == nil {
		t.Fatal("awsAsset(missing) = nil error, want error")
	}
}
