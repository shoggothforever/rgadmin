package admin

import (
	"Table/app/user/cmd/api/internal/svc"
	"Table/app/user/cmd/api/internal/types"
	"Table/app/user/cmd/api/model"
	"encoding/csv"
	"fmt"
	"github.com/minio/minio-go/v7"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/rest/httpx"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"time"
)

func AllwageHandler(svcCtx *svc.ServiceContext) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var req types.WageExcelReq
		if err := httpx.Parse(r, &req); err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "数据格式有误，无法正常解析"))
			return
		}
		if req.Month < 1 || req.Month > 12 || req.Year < 2000 || req.Year > 2023 {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(400, "输入日期信息无效"))
			return
		}
		filter := bson.D{{"year", req.Year}, {"month", req.Month}}
		cur, err := svcCtx.Mongo.Collection("wage").Find(r.Context(), filter)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
			return
		}
		defer cur.Close(r.Context())
		var tmpname = "temp.csv"
		var filename = fmt.Sprintf("wage-%d-%d.csv", req.Year, req.Month)
		var bucket = svcCtx.Config.Minio.Bucket
		file, err := os.Create(tmpname)
		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, fmt.Sprintf("无法创建临时 CSV 文件: %s", err.Error())))
			return
		}
		defer func() {
			err = file.Close()
			if err != nil {
				log.Println("无法关闭文件:", err)
			}
			err = os.Remove("temp.csv")
			if err != nil {
				log.Println("无法删除临时 CSV 文件:", err)
			}
		}()
		//file.WriteString("xEFxBBxBF")
		csvw := csv.NewWriter(file)
		csvw.Write([]string{"staffCode", "name", "wageBeforeTax", "tax", "actualWage"})
		var wage model.Wage
		for cur.Next(r.Context()) {
			cur.Decode(&wage)
			logx.Info("获取到薪水信息", wage)
			wbt := strconv.FormatFloat(wage.WageBeforeTax, 'f', 2, 64)
			tax := strconv.FormatFloat(wage.Tax, 'f', 2, 64)
			acw := strconv.FormatFloat(wage.ActualWage, 'f', 2, 64)
			csvw.Write([]string{wage.StaffCode, wage.Name, wbt, tax, acw})
		}
		csvw.Flush()
		_, err = svcCtx.Minio.FPutObject(r.Context(), bucket, filename, tmpname, minio.PutObjectOptions{
			ContentType: "application/csv",
		})
		//l := admin.NewAllwageLogic(r.Context(), svcCtx)
		//resp, err := l.Allwage(&req)

		if err != nil {
			httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
		} else {
			//w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
			//w.Header().Set("Content-Type", "application/octet-stream")
			reqParams := make(url.Values)
			reqParams.Set("response-content-disposition", fmt.Sprintf("attachment; filename=%s", filename))
			url, err := svcCtx.Minio.PresignedGetObject(r.Context(), bucket, filename, time.Duration(svcCtx.Config.Minio.LinkExpire)*time.Second, reqParams)
			if err != nil {
				httpx.OkJsonCtx(r.Context(), w, types.NewErrCodeMsg(500, err.Error()))
				return
			}
			resp := types.WageExcelResp{
				Response: types.Response{200, "成功生成Excel表格"},
				FileName: filename,
				FileType: "application/octet-stream",
				//DownloadUrl: fmt.Sprintf("%s/%s/%s", svcCtx.Minio.EndpointURL().String(), bucket, filename),
				DownloadUrl: url.String(),
			}
			httpx.OkJsonCtx(r.Context(), w, resp)
		}
	}
}
