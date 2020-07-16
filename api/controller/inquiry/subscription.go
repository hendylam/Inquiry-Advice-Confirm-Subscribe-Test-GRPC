package inquiry

import (
	"net/http"

	"tapera.integrasi/api/util/httpx"
	srv "tapera.integrasi/service/inquiry"
)

func (c *Controller) Inquiry(w http.ResponseWriter, r *http.Request) {
	ext := httpx.New(w, r)
	var param srv.InquiryConfirmSubData
	if err := ext.BindJSON(&param); err != nil {
		ext.JSONerr(http.StatusBadRequest, "Invalid request payload")
		return
	}
	var result *srv.InquiryConfirmSubResponse
	result, _ = c.srv.InquiryConfirmSubs(r.Context(), &param)
	checker := result.Data
	if checker.NavDate == "" {
		rslt := c.srv.ConvertStringtoJson(result)
		ext.JSON(http.StatusOK, rslt)
	} else {
		ext.JSON(http.StatusOK, result)
	}
}
