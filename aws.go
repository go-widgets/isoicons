// Copyright (c) 2026 the go-widgets/isoicons authors. All rights reserved.
// Use of this source code is governed by a BSD-3-Clause license that can be
// found in the LICENSE file at the root of this repository.

package isoicons

import "embed"

// awsFS holds the vendored SVG art of the "aws" pack. The assets are
// unmodified isometric icons from github.com/danieljoos/isometric-cloud-icons,
// released under CC0-1.0 (public domain); see assets/aws/LICENSE.
//
//go:embed assets/aws/*.svg
var awsFS embed.FS

// awsAsset reads one "aws" SVG from the embedded filesystem.
func awsAsset(name string) ([]byte, error) {
	return awsFS.ReadFile("assets/aws/" + name)
}

// awsSize is the square pixel size each "aws" SVG is rasterized to.
const awsSize = 128

// awsConcepts is the "aws" pack's concept list, in stable listing order. Each
// name is both the embedded SVG stem and, prefixed, the registered id
// "aws/<name>".
var awsConcepts = []string{
	"aws-api-gateway",
	"aws-cloudfront",
	"aws-cognito-identity-pool",
	"aws-dynamodb",
	"aws-lambda",
	"aws-s3-bucket",
	"black-box",
	"low-block-rounded-blue",
	"webapp",
}
