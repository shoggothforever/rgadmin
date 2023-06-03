package admin

import (
	"net/http"

	"Table/app/user/cmd/api/internal/logic/admin"
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoadInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.LoadInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "数据格式有误，无法正常解析"))
			return
		}

		l := admin.NewLoadInfoLogic(r.Context(), svcCtx)
		resp, err := l.LoadInfo(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(resp.Code, resp.Msg))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
