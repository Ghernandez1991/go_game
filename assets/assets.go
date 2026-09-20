package assets

import "embed"

//go:embed images/* sounds/*
var FS embed.FS
