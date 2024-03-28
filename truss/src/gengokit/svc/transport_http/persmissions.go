package transport_http

import (
	"github.com/metaverse/truss/gengokit"
	"github.com/metaverse/truss/gengokit/svc/transport_http/templates"
	"github.com/pkg/errors"
	"io"
)

const PermissionsHttpPath = "svc/transport_http/permissions.gotemplate"

func NewPermissions() *Permissions {
	return &Permissions{}
}

type Permissions struct {
	prev io.Reader
}

// Load loads the previous version of the middleware file.
func (p *Permissions) Load(prev io.Reader) {
	p.prev = prev
}

// Render creates the middlewares.go file. With no previous version it renders
// the templates, if there was a previous version loaded in, it passes that
// through.
func (p *Permissions) Render(path string, data *gengokit.Data) (io.Reader, error) {
	if path != PermissionsHttpPath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}
	if p.prev != nil {
		return p.prev, nil
	}
	return data.ApplyTemplate(templates.Permissions, "Permissions")
}
