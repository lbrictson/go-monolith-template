package web

import "embed"

// Assets represents the embedded files.
// You can add more files here by just extending this line, they will all be in the go executable
//
//go:embed static/css/*.css static/js/*.js static/img/*
var Assets embed.FS
