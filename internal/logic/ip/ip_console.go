/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package ip

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/dao"
	"bamboo-service/internal/model/do"
	"context"
	"encoding/csv"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/glog"
	"github.com/gogf/gf/v2/os/gtime"
	"os"
	"sync"
	"sync/atomic"
)

var mu sync.Mutex

// IPv4FileUpload
//
// # 上传IPv4数据库
//
// 上传IPv4数据库，用于上传IPv4数据库操作；
// 该接口将会对上传的文件进行解码，解码成功后将会将文件写入到 upload/ip_location/database_location_ipv4.scv 文件中；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - file			文件(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sIP) IPv4FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv4FileUpload | 上传IPv4数据库接口")
	// 检查原文件是否存在
	err = gfile.Remove("upload/ip_location/database_location_ipv4.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件删除失败")
	}
	file.Filename = "database_location_ipv4.scv"
	_, err = file.Save("upload/ip_location/")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件写入失败")
	}
	err = os.Chmod("upload/ip_location/database_location_ipv4.scv", 0755)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件权限修改失败")
	}
	return nil
}

// IPv6FileUpload
//
// # 上传IPv6数据库
//
// 上传IPv6数据库，用于上传IPv6数据库操作；
// 该接口将会对上传的文件进行解码，解码成功后将会将文件写入到 upload/ip_location/database_location_ipv6.scv 文件中；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - file			文件(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sIP) IPv6FileUpload(ctx context.Context, file *ghttp.UploadFile) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv6FileUpload | 上传IPv6数据库接口")
	// 检查原文件是否存在
	err = gfile.Remove("upload/ip_location/database_location_ipv6.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件删除失败")
	}
	file.Filename = "database_location_ipv6.scv"
	_, err = file.Save("upload/ip_location/")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件写入失败")
	}
	err = os.Chmod("upload/ip_location/database_location_ipv6.scv", 0755)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件权限修改失败")
	}
	return nil
}

// IPv4FileImport
//
// # 导入IPv4数据库
//
// 导入IPv4数据库，用于导入IPv4数据库操作；
// 该接口将会从 upload/ip_location/database_location_ipv4.scv 文件中导入数据到数据库中；
// 该接口将会清空原有的数据；
//
// # 参数
//   - ctx			上下文(context.Context)
//
// # 返回
//   - err			错误信息(error)
func (s *sIP) IPv4FileImport(ctx context.Context) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv4FileImport | 导入IPv4数据库接口")
	// 事务前期统计
	mu.Lock()
	atomic.StoreInt32(&constant.ImportIPv4CtxCancel, 1)
	mu.Unlock()
	// 事务时间统计
	g.Log().Notice(ctx, "[SERV] 开始导入 ipv4 数据库")
	startTime := gtime.Now().TimestampMilli()
	// 检查文件是否存在
	if !gfile.Exists("upload/ip_location/database_location_ipv4.scv") {
		return berror.NewError(bcode.ServerInternalError, "文件不存在")
	}
	getFile, err := gfile.Open("upload/ip_location/database_location_ipv4.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件打开失败")
	}
	csvGetAll, err := csv.NewReader(getFile).ReadAll()
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件读取失败")
	}
	// 导入总体数目进行计算
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_4_upload_total"}).Update(do.Info{Value: len(csvGetAll)})
	if err != nil {
		g.Log().Warningf(ctx, "[SERV] 导入 ipv4 数据库总数更新失败 [%v]", err)
	}
	// 事务操作
	err = dao.LocationIpv4.Transaction(ctx, func(ctx context.Context, tx gdb.TX) error {
		// 清空原有数据
		_, err := tx.Ctx(ctx).Exec("TRUNCATE TABLE fy_location_ipv4;")
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库清空失败")
		}
		getIntervalTime := gtime.TimestampMilli()
		// 导入数据
		for i, getCSV := range csvGetAll {
			if i%1000 == 0 && g.Log().GetLevel() > glog.LEVEL_NONE {
				g.Log().Debugf(ctx, "[SERV] 导入 ipv4 数据库 [%v, %v]", i, len(csvGetAll))
			}
			if i%10000 == 0 && g.Log().GetLevel() <= glog.LEVEL_DEBU {
				g.Log().Infof(ctx, "[SERV] 导入 ipv4 数据库 [%v, %v]", i, len(csvGetAll))
			}
			// 对时间进行计算
			if i%100000 == 0 {
				// 统计十万次周期时间
				_, infoErr := dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_4_period_time"}).
					Update(do.Info{Value: gtime.Now().TimestampMilli() - getIntervalTime})
				getIntervalTime = gtime.TimestampMilli()
				if infoErr != nil {
					g.Log().Warningf(ctx, "[SERV] 导入 ipv4 数据库时间更新失败 [%v]", infoErr)
				}
			}
			mu.Lock()
			getValue := atomic.LoadInt32(&constant.ImportIPv4CtxCancel)
			mu.Unlock()
			if getValue == 0 {
				g.Log().Warningf(ctx, "[SERV] 导入 ipv4 数据库取消")
				return berror.NewError(bcode.OperationFailed, "导入取消")
			}
			_, err := tx.Ctx(ctx).Insert("fy_location_ipv4", do.LocationIpv4{
				IpForm:      getCSV[0],
				IpTo:        getCSV[1],
				CountryCode: getCSV[2],
				CountryName: getCSV[3],
				RegionName:  getCSV[4],
				CityName:    getCSV[5],
				Latitude:    getCSV[6],
				Longitude:   getCSV[7],
				ZipCode:     getCSV[8],
				TimeZone:    getCSV[9],
			})
			if err != nil {
				g.Log().Debug(ctx, getCSV)
				return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入失败")
			}
		}
		return nil
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入失败")
	}
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_4_import_time"}).Update(do.Info{Value: gtime.Now()})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入时间更新失败")
	}
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_4_import_spending"}).
		Update(do.Info{Value: gtime.Now().TimestampMilli() - startTime})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入时间更新失败")
	}
	mu.Lock()
	atomic.StoreInt32(&constant.ImportIPv4CtxCancel, 0)
	mu.Unlock()
	g.Log().Notice(ctx, "[SERV] 导入 ipv4 数据库完成")
	return nil
}

// IPv6FileImport
//
// # 导入IPv6数据库
//
// 导入IPv6数据库，用于导入IPv6数据库操作；
// 该接口将会从 upload/ip_location/database_location_ipv6.scv 文件中导入数据到数据库中；
// 该接口将会清空原有的数据；
//
// # 参数
//   - ctx			上下文(context.Context)
//
// # 返回
//   - err			错误信息(error)
func (s *sIP) IPv6FileImport(ctx context.Context) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.IPv6FileImport | 导入IPv6数据库接口")
	// 事务前期统计
	mu.Lock()
	atomic.StoreInt32(&constant.ImportIPv6CtxCancel, 1)
	mu.Unlock()
	// 事务时间统计
	g.Log().Notice(ctx, "[SERV] 开始导入 ipv6 数据库")
	startTime := gtime.Now().TimestampMilli()
	// 检查文件是否存在
	if !gfile.Exists("upload/ip_location/database_location_ipv6.scv") {
		return berror.NewError(bcode.ServerInternalError, "文件不存在")
	}
	getFile, err := gfile.Open("upload/ip_location/database_location_ipv6.scv")
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件打开失败")
	}
	csvGetAll, err := csv.NewReader(getFile).ReadAll()
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "文件读取失败")
	}
	// 导入总体数目进行计算
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_6_upload_total"}).Update(do.Info{Value: len(csvGetAll)})
	if err != nil {
		g.Log().Warningf(ctx, "[SERV] 导入 ipv6 数据库总数更新失败 [%v]", err)
	}
	// 事务操作
	err = dao.LocationIpv6.Transaction(ctx, func(_ context.Context, tx gdb.TX) error {
		// 清空原有数据
		_, err := tx.Ctx(ctx).Exec("TRUNCATE TABLE fy_location_ipv6;")
		if err != nil {
			return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库清空失败")
		}
		// 导入数据
		for i, getCSV := range csvGetAll {
			if i%1000 == 0 {
				g.Log().Debugf(ctx, "[SERV] 导入 ipv6 数据库 [%v, %v]", i, len(csvGetAll))
			}
			if i%10000 == 0 {
				g.Log().Infof(ctx, "[SERV] 导入 ipv6 数据库 [%v, %v]", i, len(csvGetAll))
			}
			if i%100000 == 0 {
				g.Log().Noticef(ctx, "[SERV] 导入 ipv6 数据库 [%v, %v]", i, len(csvGetAll))
			}
			mu.Lock()
			getValue := atomic.LoadInt32(&constant.ImportIPv6CtxCancel)
			mu.Unlock()
			if getValue == 0 {
				g.Log().Warningf(ctx, "[SERV] 导入 ipv6 数据库取消")
				return berror.NewError(bcode.OperationFailed, "导入取消")
			}
			_, err = dao.LocationIpv6.Ctx(ctx).Where(do.LocationIpv4{
				IpForm:      getCSV[0],
				IpTo:        getCSV[1],
				CountryCode: getCSV[2],
				CountryName: getCSV[3],
				RegionName:  getCSV[4],
				CityName:    getCSV[5],
				Latitude:    getCSV[6],
				Longitude:   getCSV[7],
				ZipCode:     getCSV[8],
				TimeZone:    getCSV[9],
			}).Insert()
			if err != nil {
				g.Log().Debug(ctx, getCSV)
				return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入失败")
			}
		}
		return nil
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入失败")
	}
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_6_import_time"}).Update(do.Info{Value: gtime.Now()})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入时间更新失败")
	}
	_, err = dao.Info.Ctx(ctx).Where(do.Info{Key: "ip_6_import_spending"}).
		Update(do.Info{Value: gtime.Now().TimestampMilli() - startTime})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "数据库导入时间更新失败")
	}
	mu.Lock()
	atomic.StoreInt32(&constant.ImportIPv6CtxCancel, 0)
	mu.Unlock()
	g.Log().Notice(ctx, "[SERV] 导入 ipv6 数据库完成")
	return nil
}

// CancelImportIP
//
// # 取消导入IPv4数据库
//
// 取消导入IPv4数据库，用于取消导入IPv4数据库操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - targetCtx	目标上下文(string)
//
// # 返回
//   - err			错误信息(error)
func (s *sIP) CancelImportIP(ctx context.Context, ipType bool) (err error) {
	g.Log().Notice(ctx, "[SERV] ip.CancelImportIPv4 | 取消导入IPv4数据库接口")
	// 取消 CTX
	mu.Lock()
	if ipType {
		if atomic.LoadInt32(&constant.ImportIPv6CtxCancel) == 0 {
			return berror.NewError(bcode.OperationFailed, "没有正在导入的内容")
		}
		atomic.StoreInt32(&constant.ImportIPv6CtxCancel, 0)
	} else {
		g.Log().Debug(ctx, atomic.LoadInt32(&constant.ImportIPv4CtxCancel))
		if atomic.LoadInt32(&constant.ImportIPv4CtxCancel) == 0 {
			return berror.NewError(bcode.OperationFailed, "没有正在导入的内容")
		}
		atomic.StoreInt32(&constant.ImportIPv4CtxCancel, 0)
	}
	mu.Unlock()
	return nil
}
