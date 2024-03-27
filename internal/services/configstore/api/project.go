// Copyright 2019 Sorint.lab
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied
// See the License for the specific language governing permissions and
// limitations under the License.

package api

import (
	"encoding/json"
	"net/http"
	"net/url"

	"github.com/gorilla/mux"
	"github.com/rs/zerolog"
	"github.com/sorintlab/errors"

	"agola.io/agola/internal/services/configstore/action"
	"agola.io/agola/internal/util"
	csapitypes "agola.io/agola/services/configstore/api/types"
	"agola.io/agola/services/configstore/types"
)

func projectResponse(project *types.Project, projectDynamicData *action.ProjectDynamicData) (*csapitypes.Project, error) {
	r, err := projectsResponse([]*types.Project{project}, map[string]*action.ProjectDynamicData{project.ID: projectDynamicData})
	if err != nil {
		return nil, errors.WithStack(err)
	}
	return r[0], nil
}

// TODO(sgotti) do these queries inside the action handler?
func projectsResponse(projects []*types.Project, projectsDynamicData map[string]*action.ProjectDynamicData) ([]*csapitypes.Project, error) {
	resProjects := make([]*csapitypes.Project, len(projects))

	for i, project := range projects {
		projectDynamicData := projectsDynamicData[project.ID]

		resProjects[i] = &csapitypes.Project{
			Project:          project,
			OwnerType:        projectDynamicData.OwnerType,
			OwnerID:          projectDynamicData.OwnerID,
			Path:             projectDynamicData.Path,
			ParentPath:       projectDynamicData.ParentPath,
			GlobalVisibility: projectDynamicData.GlobalVisibility,
		}
	}

	return resProjects, nil
}

type ProjectHandler struct {
	log zerolog.Logger
	ah  *action.ActionHandler
}

func NewProjectHandler(log zerolog.Logger, ah *action.ActionHandler) *ProjectHandler {
	return &ProjectHandler{log: log, ah: ah}
}

func (h *ProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	vars := mux.Vars(r)
	projectRef, err := url.PathUnescape(vars["projectref"])
	if err != nil {
		util.HTTPError(w, util.NewAPIErrorWrap(util.ErrBadRequest, err))
		return
	}

	res, err := h.ah.GetProject(ctx, projectRef)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	resProject, err := projectResponse(res.Project, res.ProjectDynamicData)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	if err := util.HTTPResponse(w, http.StatusOK, resProject); err != nil {
		h.log.Err(err).Send()
	}
}

type CreateProjectHandler struct {
	log zerolog.Logger
	ah  *action.ActionHandler
}

func NewCreateProjectHandler(log zerolog.Logger, ah *action.ActionHandler) *CreateProjectHandler {
	return &CreateProjectHandler{log: log, ah: ah}
}

func (h *CreateProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var req *csapitypes.CreateUpdateProjectRequest
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&req); err != nil {
		util.HTTPError(w, util.NewAPIErrorWrap(util.ErrBadRequest, err))
		return
	}

	areq := &action.CreateUpdateProjectRequest{
		Name:                        req.Name,
		Parent:                      req.Parent,
		Visibility:                  req.Visibility,
		RemoteRepositoryConfigType:  req.RemoteRepositoryConfigType,
		RemoteSourceID:              req.RemoteSourceID,
		LinkedAccountID:             req.LinkedAccountID,
		RepositoryID:                req.RepositoryID,
		RepositoryPath:              req.RepositoryPath,
		SSHPrivateKey:               req.SSHPrivateKey,
		SkipSSHHostKeyCheck:         req.SkipSSHHostKeyCheck,
		PassVarsToForkedPR:          req.PassVarsToForkedPR,
		DefaultBranch:               req.DefaultBranch,
		MembersCanPerformRunActions: req.MembersCanPerformRunActions,
	}

	res, err := h.ah.CreateProject(ctx, areq)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	resProject, err := projectResponse(res.Project, res.ProjectDynamicData)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	if err := util.HTTPResponse(w, http.StatusCreated, resProject); err != nil {
		h.log.Err(err).Send()
	}
}

type UpdateProjectHandler struct {
	log zerolog.Logger
	ah  *action.ActionHandler
}

func NewUpdateProjectHandler(log zerolog.Logger, ah *action.ActionHandler) *UpdateProjectHandler {
	return &UpdateProjectHandler{log: log, ah: ah}
}

func (h *UpdateProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	projectRef, err := url.PathUnescape(vars["projectref"])
	if err != nil {
		util.HTTPError(w, util.NewAPIErrorWrap(util.ErrBadRequest, err))
		return
	}

	var req *csapitypes.CreateUpdateProjectRequest
	d := json.NewDecoder(r.Body)
	if err := d.Decode(&req); err != nil {
		util.HTTPError(w, util.NewAPIErrorWrap(util.ErrBadRequest, err))
		return
	}

	areq := &action.CreateUpdateProjectRequest{
		Name:                        req.Name,
		Parent:                      req.Parent,
		Visibility:                  req.Visibility,
		RemoteRepositoryConfigType:  req.RemoteRepositoryConfigType,
		RemoteSourceID:              req.RemoteSourceID,
		LinkedAccountID:             req.LinkedAccountID,
		RepositoryID:                req.RepositoryID,
		RepositoryPath:              req.RepositoryPath,
		SSHPrivateKey:               req.SSHPrivateKey,
		SkipSSHHostKeyCheck:         req.SkipSSHHostKeyCheck,
		PassVarsToForkedPR:          req.PassVarsToForkedPR,
		DefaultBranch:               req.DefaultBranch,
		MembersCanPerformRunActions: req.MembersCanPerformRunActions,
	}

	res, err := h.ah.UpdateProject(ctx, projectRef, areq)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	resProject, err := projectResponse(res.Project, res.ProjectDynamicData)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
		return
	}

	if err := util.HTTPResponse(w, http.StatusCreated, resProject); err != nil {
		h.log.Err(err).Send()
	}
}

type DeleteProjectHandler struct {
	log zerolog.Logger
	ah  *action.ActionHandler
}

func NewDeleteProjectHandler(log zerolog.Logger, ah *action.ActionHandler) *DeleteProjectHandler {
	return &DeleteProjectHandler{log: log, ah: ah}
}

func (h *DeleteProjectHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	vars := mux.Vars(r)
	projectRef, err := url.PathUnescape(vars["projectref"])
	if err != nil {
		util.HTTPError(w, util.NewAPIErrorWrap(util.ErrBadRequest, err))
		return
	}

	err = h.ah.DeleteProject(ctx, projectRef)
	if util.HTTPError(w, err) {
		h.log.Err(err).Send()
	}
	if err := util.HTTPResponse(w, http.StatusNoContent, nil); err != nil {
		h.log.Err(err).Send()
	}
}

const (
	DefaultProjectsLimit = 10
	MaxProjectsLimit     = 20
)
