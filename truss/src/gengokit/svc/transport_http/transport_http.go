package transport_http

import (
	"github.com/metaverse/truss/gengokit"
	"github.com/metaverse/truss/gengokit/svc/transport_http/templates"
	"github.com/pkg/errors"
	"io"
)

const TransportHttpPath = "svc/transport_http/transport_http.gotemplate"

// NewTransportHttp returns a new NewTransportHttpRender
func NewTransportHttp() *TransportHttp {
	return &TransportHttp{}
}

type TransportHttp struct {
	prev io.Reader
}

// Load loads the previous version of the middleware file.
func (p *TransportHttp) Load(prev io.Reader) {
	p.prev = prev
}

// Render creates the middlewares.go file. With no previous version it renders
// the templates, if there was a previous version loaded in, it passes that
// through.
func (p *TransportHttp) Render(path string, data *gengokit.Data) (io.Reader, error) {
	if path != TransportHttpPath {
		return nil, errors.Errorf("cannot render unknown file: %q", path)
	}
	if p.prev != nil {
		return p.prev, nil
	}
	return data.ApplyTemplate(templates.TransportHttp, "TransportHttp")
}
