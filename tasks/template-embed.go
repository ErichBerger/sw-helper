package tasks

import "embed"

//go:embed templates/storefront-js/*
var storefrontJSTemplates embed.FS

//go:embed templates/cms-element/*
var cmsElementTemplates embed.FS
