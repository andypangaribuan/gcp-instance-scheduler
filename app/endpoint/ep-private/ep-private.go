package ep_private

import (
	"github.com/andypangaribuan/vision-go/core/api"
	"github.com/andypangaribuan/vision-go/vis"
	"net/http"
	"time"
)


/* ============================================
	Created by andy pangaribuan on 2021/05/20
	Copyright andypangaribuan. All rights reserved.
   ============================================ */
func Time(c api.Context) error {
	return c.ResponseStr(http.StatusOK, vis.Convert.TimeToStrFull(time.Now()))
}