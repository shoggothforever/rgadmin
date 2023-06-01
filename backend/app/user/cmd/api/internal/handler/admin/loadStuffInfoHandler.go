package admin

import (
	"net/http"

	"Table/app/user/cmd/api/internal/logic/admin"
	"Table/app/user/cmd/api/internal/svc"
	"github.com/zeromicro/go-zero/rest/httpx"
)

func LoadStuffInfoHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		l := admin.NewLoadStuffInfoLogic(r.Context(), svcCtx)
		err := l.LoadStuffInfo()
		if err != nil {
			httpx.ErrorCtx(r.Context(), w, err)
		} else {
			httpx.Ok(w)
		}
	}
}
