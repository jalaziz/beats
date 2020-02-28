// Copyright Elasticsearch B.V. and/or licensed to Elasticsearch B.V. under one
// or more contributor license agreements. Licensed under the Elastic License;
// you may not use this file except in compliance with the Elastic License.

// Code generated by beats/dev-tools/cmd/asset/asset.go - DO NOT EDIT.

package statsd

import (
	"github.com/elastic/beats/libbeat/asset"
)

func init() {
	if err := asset.SetFields("metricbeat", "statsd", asset.ModuleFieldsPri, AssetStatsd); err != nil {
		panic(err)
	}
}

// AssetStatsd returns asset data.
// This is the base64 encoded gzipped contents of module/statsd.
func AssetStatsd() string {
	return "eJyUkTtqxTAQRXut4qLGYHAWoCKbyAKMIo2NEv2QxoV3H/wL4uFXeMp7z4gzaMAvrQqVNVcrAHbsSUF+7YEUgKVqisvsUlT4FABwlAjJLp4EUMiTrqQwawFMjrytaicHRB2oeX8bXvPGlrTkM2lX2rWu/zBpidz9N9d2+v4hw018BOPR+hTn+24MOmcX5xOUGykb9Obaa86rdyEq9U62fyo6+aTflK+m/TPNQFycaS3bX/oLAAD//wYwi2Y="
}
