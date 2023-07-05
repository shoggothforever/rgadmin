package user

import (
	"Table/app/user/cmd/api/internal/types"
	"net/http"

	"Table/app/user/cmd/api/internal/logic/user"
	"Table/app/user/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func UsergetleaveHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := user.NewUsergetleaveLogic(r.Context(), svcCtx)
		resp, err := l.Usergetleave()
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(resp.Code, resp.Msg))
		} else {
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
