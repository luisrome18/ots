package main

import (
	"crypto/sha512"
	"encoding/base64"
	"path"
	"text/template"
)

var tplFuncs = template.FuncMap{
	"SRIHash": assetSRIHash,
}

func assetSRIHash(assetName string) string {
	data := MustAsset(path.Join("frontend", assetName))

	h := sha512.New384()
	h.Write(data)
	sum := h.Sum(nil)

	return "sha384-" + base64.StdEncoding.EncodeToString(sum)
}
