package admin

import (
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
	"github.com/xuri/excelize/v2"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.mongodb.org/mongo-driver/bson"
	"io"
	"mime/multipart"
	"net/http"
	"os"
	"strconv"
	"time"
)

func UploadWorkTimeHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		file, fileheader, err := r.FormFile("file")
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		//file, fileheader, err := r.FormFile("file")
		defer func(file multipart.File) {
			err := file.Close()
			if err != nil {
				httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
				return
			}
		}(file)
		dst, err := os.Create(fileheader.Filename)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		defer func() {
			dst.Close()
			os.Remove(fileheader.Filename)
		}()
		_, err = io.Copy(dst, file)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		f, err := excelize.OpenFile(fileheader.Filename)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		rows, err := f.GetRows("Sheet1")
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		filter := bson.D{{"staffcode", 0}}
		var wage model.Wage
		var user model.User
		for k, row := range rows {
			if k == 0 {
				continue
			}
			filter = bson.D{{"staffcode", row[0]}}
			err := svcCtx.Mongo.Collection("employee").FindOne(r.Context(), filter).Decode(&user)
			if err != nil {
				logx.Info(err.Error())
				continue
			}
			worktime, err := strconv.ParseFloat(row[1], 64)
			if err != nil {
				logx.Info(err.Error())
				continue
			}
			var base = model.GetBasewage(user.Role)
			base += base*worktime/1200 + model.ElseFee
			tax := model.CalTax(base)
			wage = model.Wage{
				StaffCode:     user.StaffCode,
				Name:          user.Name,
				Year:          time.Now().Year(),
				Month:         int(time.Now().Month()),
				WorkTime:      worktime,
				WageBeforeTax: base,
				Tax:           tax,
				ActualWage:    base - tax,
			}
			_, err = svcCtx.Mongo.Collection("wage").InsertOne(r.Context(), &wage)
			if err != nil {
				httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
				return
			}
		}
		httpx.OkJsonCtx(r.Context(), w, types.UploadInfoResp{Response: types.Response{
			Code: 200,
			Msg:  "上传用户工资表信息成功",
		}})
	}
}
