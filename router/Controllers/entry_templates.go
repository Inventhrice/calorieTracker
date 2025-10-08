package controllers

import (
	"fmt"
	"net/http"

	models "example.com/m/v2/Models"
	"github.com/gin-gonic/gin"
)

func getAllTemplates(ctx *gin.Context) {
	var errmsg string
	if userID, err := helper_GetUserID(ctx); err != nil {
		errmsg = err.Error()
	} else {
		if templates, err := models.GetAllTemplates(userID); err != nil {
			errmsg = err.Error()
		} else {
			ctx.JSON(http.StatusOK, templates)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func getTemplate(ctx *gin.Context) {
	var errmsg string
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid template ID"
	} else {
		if userID, err := helper_GetUserID(ctx); err != nil {
			errmsg = err.Error()
		} else {
			if template, err := models.GetTemplate(userID, id); err != nil {
				errmsg = err.Error()
			} else {
				ctx.JSON(http.StatusOK, template)
				return
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func addTemplate(ctx *gin.Context) {
	var errmsg string
	var template models.EntryTemplate
	if err := ctx.BindJSON(&template); err != nil {
		errmsg = err.Error()
	} else {
		if template.UserID, err = helper_GetUserID(ctx); err != nil {
			errmsg = err.Error()
		} else {
			if newID, err := models.AddTemplate(template); err != nil {
				errmsg = err.Error()
			} else {
				ctx.JSON(http.StatusOK, gin.H{"message": "Template added successfully", "addedID": newID})
				return
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func deleteTemplate(ctx *gin.Context) {
	var errmsg string
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid template ID"
	} else {
		if err := models.DeleteTemplate(id); err != nil {
			errmsg = err.Error()
		} else {
			ctx.Status(http.StatusOK)
			return
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func updateTemplate(ctx *gin.Context) {
	var errmsg string
	if id, err := helper_getIntFromStr(ctx.Param("id")); err != nil {
		errmsg = "Invalid template ID"
	} else {
		var template models.EntryTemplate
		if err := ctx.BindJSON(&template); err != nil {
			errmsg = err.Error()
		} else {
			if id == template.ID {
				if err := models.UpdateTemplate(template); err != nil {
					errmsg = err.Error()
				} else {
					ctx.Status(http.StatusOK)
					return
				}
			} else {
				errmsg = fmt.Sprintf("IDs do not match, aborting request. id: %v, template.ID: %v", id, template.ID)
			}
		}
	}
	Helper_ctx400(ctx, errmsg)
}

func InitTemplateRoutes(group *gin.RouterGroup) {
	group.GET("/all", getAllTemplates)
	group.GET("/:id", getTemplate)
	group.PATCH("/:id", updateTemplate)
	group.POST("/", addTemplate)
	group.DELETE("/:id", deleteTemplate)
}
