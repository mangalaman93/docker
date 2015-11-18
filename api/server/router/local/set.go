package local

import (
	"fmt"
	"net/http"

	"github.com/docker/docker/api/server/httputils"
	"github.com/docker/docker/api/types"
	"github.com/docker/docker/runconfig"
	"golang.org/x/net/context"
)

func (s *router) postContainersSet(ctx context.Context, w http.ResponseWriter, r *http.Request, vars map[string]string) error {
	if vars == nil {
		return fmt.Errorf("Missing parameter")
	}
	if err := httputils.ParseForm(r); err != nil {
		return err
	}
	if err := httputils.CheckForJSON(r); err != nil {
		return err
	}

	name := vars["name"]
	_, hostConfig, err := runconfig.DecodeContainerConfig(r.Body)
	if err != nil {
		return err
	}

	warnings, err := s.daemon.ContainerSet(name, hostConfig)
	if err != nil {
		return err
	}

	return httputils.WriteJSON(w, http.StatusOK, &types.ContainerSetResponse{
		Warnings: warnings,
	})
}
