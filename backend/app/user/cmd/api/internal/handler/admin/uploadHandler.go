package admin

import (
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"encoding/csv"
	"fmt"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"mime/multipart"
	"net/http"
)

func UploadHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.UploadInfoReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "数据格式有误，无法正常解析"))
			return
		}
		file, _, err := r.FormFile("file")
		//file, fileheader, err := r.FormFile("file")
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
				return
			}
		}(file)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		csvRead := csv.NewReader(file)
		record, _ := csvRead.ReadAll()
		for k, v := range record {
			if k == 10 {
				break
			}
			httpx.OkJsonCtx(r.Context(), w, fmt.Sprintf("csvRead读取到了信息,%v ,%v", k, v))
			logx.Info("csvRead读取到了信息", k, v)
		}
		//var user model.User
		//var wage model.Wage
		//err = svcCtx.Mongo.Collection("employee").FindOne(r.Context(), fileter).Decode(&user)
		//if err != nil {
		//	resp := &types.UploadInfoResp{Response: types.NewResponse(401, err.Error())}
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//	return
		//}
		////logx.WithContext(l.ctx).Info("查询到用户信息:", user)
		//var base = model.GetBasewage(user.Role)
		//base += base*req.WorkTime/1200 + model.ElseFee
		//tax := model.CalTax(base)
		//wage := model.Wage{
		//	StaffCode:     user.StaffCode,
		//	Name:          user.Name,
		//	Year:          time.Now().Year(),
		//	Month:         int(time.Now().Month()),
		//	WorkTime:      req.WorkTime,
		//	WageBeforeTax: base,
		//	Tax:           tax,
		//	ActualWage:    base - tax,
		//}
		////logx.WithContext(l.ctx).Info("计算到薪水信息:", wage)
		//_, err = svcCtx.Mongo.Collection("wage").InsertOne(r.Context(), &wage)
		//if err != nil {
		//	resp := &types.UploadInfoResp{Response: types.NewResponse(401, err.Error())}
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//	return
		//}
		//l := admin.NewUploadLogic(r.Context(), svcCtx)
		//resp, err := l.Upload(&req)
		//err = dbminio.Upload2Minio(r.Context(), svcCtx.Minio, "rgadmin", file, fileheader)
		//if err != nil {
		//	httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "存储文件失败"+err.Error()))
		//} else {
		//	//_, err = svcCtx.Redis.Set(r.Context(), fileheader.Filename, 1, 86400*30).Result()
		//	if err != nil {
		//		httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
		//		return
		//	}
		//	resp.Response = types.NewResponse(200, "存储文件成功")
		//	httpx.OkJsonCtx(r.Context(), w, resp)
		//}

	}
}
