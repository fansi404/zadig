/*
Copyright 2021 The KodeRover Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/koderover/zadig/pkg/tool/errors"

	"github.com/koderover/zadig/pkg/microservice/aslan/core/common/service/fs"
	"github.com/koderover/zadig/pkg/microservice/aslan/core/templatestore/service"
	templateservice "github.com/koderover/zadig/pkg/microservice/aslan/core/templatestore/service"
	internalhandler "github.com/koderover/zadig/pkg/shared/handler"
)

type addChartArgs struct {
	*fs.DownloadFromSourceArgs

	Name string `json:"name"`
}

type updateChartTemplateVariablesReq struct {
	Variables []*templateservice.Variable `json:"variables"`
}

func GetChartTemplate(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.Err = service.GetChartTemplate(c.Param("name"), ctx.Logger)
}

func GetTemplateVariables(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.Err = service.GetChartTemplateVariables(c.Param("name"), ctx.Logger)
}

func ListFiles(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	// TODO: support to return multiple files in a bulk
	ctx.Resp, ctx.Err = service.GetFileContentForTemplate(c.Param("name"), c.Query("filePath"), c.Query("fileName"), ctx.Logger)
}

func GetChartTemplateReferences(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()
	return
}

func ListChartTemplates(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Resp, ctx.Err = service.ListChartTemplates(ctx.Logger)
}

func AddChartTemplate(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := &addChartArgs{}
	if err := c.ShouldBindJSON(args); err != nil {
		ctx.Err = err
		return
	}

	ctx.Err = service.AddChartTemplate(args.Name, args.DownloadFromSourceArgs, ctx.Logger)
}

func UpdateChartTemplate(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := &addChartArgs{}
	if err := c.ShouldBindJSON(args); err != nil {
		ctx.Err = err
		return
	}

	ctx.Err = service.UpdateChartTemplate(c.Param("name"), args.DownloadFromSourceArgs, ctx.Logger)
}

func UpdateChartTemplateVariables(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	args := make([]*templateservice.Variable, 0)
	if err := c.ShouldBindJSON(&args); err != nil {
		ctx.Err = errors.ErrInvalidParam.AddErr(err)
		return
	}

	ctx.Err = service.UpdateChartTemplateVariables(c.Param("name"), args, ctx.Logger)
}

func RemoveChartTemplate(c *gin.Context) {
	ctx := internalhandler.NewContext(c)
	defer func() { internalhandler.JSONResponse(c, ctx) }()

	ctx.Err = service.RemoveChartTemplate(c.Param("name"), ctx.Logger)
}
