package assetbundle

import (
	"github.com/aki017/assetbundle/body"
	"github.com/aki017/assetbundle/header"
)

// AssetBundle is Main struct
type AssetBundle struct {
	Header *header.Header
	Bodies []*body.Body
}
