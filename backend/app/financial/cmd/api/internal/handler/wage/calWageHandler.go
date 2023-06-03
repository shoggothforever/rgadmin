package wage

import (
	"net/http"

	"Table/app/financial/cmd/api/internal/logic/wage"
	"Table/app/financial/cmd/api/internal/svc"
	"Table/app/financial/cmd/api/internal/types"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func CalWageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.CalReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "数据格式有误，无法正常解析"))
			return
		}

		l := wage.NewCalWageLogic(r.Context(), svcCtx)
		resp, err := l.CalWage(&req)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(resp.Code, resp.Msg))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
