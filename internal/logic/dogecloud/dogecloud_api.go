/*
 * ------------------------------------------------------------
 * Copyright (c) 2016-NOW 锋楪技术 All Rights Reserved. 版权所有
 * 开源协议请遵循 MIT 开源协议，参考代码中的 LICENSE 部分
 * ------------------------------------------------------------
 * 代码若需进行商用请务必联系我，同意后方可使用。在使用部分请注明出处
 * 作者：锋楪技术（筱锋xiao_lfeng）
 * ------------------------------------------------------------
 */

package dogecloud

import (
	"bamboo-service/internal/constant"
	"bamboo-service/internal/model/rdo"
	"bamboo-service/internal/service"
	"bytes"
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"
	"github.com/bamboo-services/bamboo-utils/bcode"
	"github.com/bamboo-services/bamboo-utils/berror"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/gclient"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/util/gconv"
	"net/url"
)

// Api
//
// # 多吉云 API 请求
//
// 通过该接口可以请求多吉云的 API 接口；
// 该接口会自动处理签名等信息；
// 该接口会自动处理 JSON 数据；
// 该接口会自动处理请求头等信息；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - apiPath		API 路径(string)
//   - data			请求数据(map[string]interface{})
//   - jsonMode		是否为 JSON 模式(bool)
//
// # 返回
//   - ret			返回数据(map[string]interface{})
//   - err			错误信息(error)
func (s *sDogeCloud) Api(
	ctx context.Context,
	apiPath string,
	data map[string]interface{},
	jsonMode bool,
) (ref map[string]interface{}, err error) {
	g.Log().Notice(ctx, "[SERV] doge-cloud.Api | 多吉云 API 请求")
	var body, mime string
	if jsonMode {
		_body, err := json.Marshal(data)
		if err != nil {
			return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "Json 数据处理失败")
		}
		body = string(_body)
		mime = "application/json"
	} else {
		values := url.Values{}
		for k, v := range data {
			values.Set(k, v.(string))
		}
		body = values.Encode()
		mime = "application/x-www-form-urlencoded"
	}

	hmacObj := hmac.New(sha1.New, []byte(constant.DogeCloudSecretKey))
	hmacObj.Write([]byte(apiPath + "\n" + body))

	client := g.Client()
	client.SetHeader("Content-Type", mime)
	client.SetHeader(
		"Authorization",
		"TOKEN "+constant.DogeCloudAccessKey+":"+hex.EncodeToString(hmacObj.Sum(nil)),
	)
	response, err := client.Post(ctx, "https://api.dogecloud.com"+apiPath, body)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "请求失败")
	}
	defer func(response *gclient.Response) {
		_ = response.Close()
	}(response)

	// 数据读取操作
	err = json.Unmarshal(response.ReadAll(), &ref)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "Json 数据解析失败")
	}
	return ref, nil
}

// GetAccessTokenApi
//
// # 获取多吉云存储 Token 权限
//
// 该借口操作不建议单独直接进行调用，在 task 模块进行定时循环调用；
// 目的为获取权限 Token，用于后续的操作；
//
// # 参数
//   - ctx			上下文(context.Context)
//
// # 返回
//   - bucket		多吉云存储 Token 信息(*rdo.DogeCloudBucketRDO)
//   - err			错误信息(error)
func (s *sDogeCloud) GetAccessTokenApi(ctx context.Context) (bucket *rdo.DogeCloudBucketRDO, err error) {
	g.Log().Notice(ctx, "[SERV] doge-cloud.GetAccessTokenApi | 获取多吉云存储 Token 权限")
	data := gconv.Map(g.Map{
		"channel": "OSS_FULL",
		"scopes":  "bamboo-service",
	})
	ref, err := service.DogeCloud().Api(ctx, "/auth/tmp_token.json", data, true)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取 Token 失败")
	}
	getJson := gjson.New(ref)
	// 数据处理
	bucketRDO := new(rdo.DogeCloudBucketRDO)
	bucketRDO.Name = getJson.Get("data.Buckets.0.name").String()
	bucketRDO.Bucket = getJson.Get("data.Buckets.0.s3Bucket").String()
	bucketRDO.Endpoint = getJson.Get("data.Buckets.0.s3Endpoint").String()
	bucketRDO.EndpointHost = getJson.Get("data.Buckets.0.s3EndpointHost").String()
	bucketRDO.AccessKeyID = getJson.Get("data.Credentials.accessKeyId").String()
	bucketRDO.SecretAccessKey = getJson.Get("data.Credentials.secretAccessKey").String()
	bucketRDO.SessionToken = getJson.Get("data.Credentials.sessionToken").String()
	bucketRDO.ExpiredAt = gtime.NewFromTime(getJson.Get("data.ExpiredAt").Time())
	// 数据返回
	return bucketRDO, nil
}

// GetToken
//
// # 获取多吉云存储 Token 权限
//
// 该接口为从缓存调用令牌信息；
// 用于获取多吉云存储 Token 权限；
//
// # 参数
//   - ctx			上下文(context.Context)
//
// # 返回
//   - bucket		多吉云存储 Token 信息(*rdo.DogeCloudBucketRDO)
//   - err			错误信息(error)
func (s *sDogeCloud) GetToken(ctx context.Context) (bucket *rdo.DogeCloudBucketRDO, err error) {
	g.Log().Notice(ctx, "[SERV] doge-cloud.GetToken | 获取多吉云存储 Token 权限")
	// 从缓存读取
	hGet, err := g.Redis().HGetAll(ctx, "global:dc:bucket")
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取 Token 失败")
	}
	err = hGet.Scan(&bucket)
	if err != nil {
		return nil, berror.NewErrorHasError(bcode.ServerInternalError, err, "获取 Token 失败")
	}
	return bucket, nil
}

// UploadData
//
// # 上传数据
//
// 该接口用于上传数据到多吉云存储；
//
// # 参数
//   - ctx			上下文(context.Context)
//   - path			路径(string)
//   - fileName		文件名(string)
//   - body			数据(*io.Reader)
//
// # 返回
//   - err			错误信息(error)
func (s *sDogeCloud) UploadData(ctx context.Context, path, fileName string, body []byte) (err error) {
	g.Log().Notice(ctx, "[SERV] doge-cloud.UploadData | 上传数据")
	bucket, err := s.GetToken(ctx)
	if err != nil {
		return err
	}
	// 上传数据
	s3Config := &aws.Config{
		Credentials: credentials.NewStaticCredentials(bucket.AccessKeyID, bucket.SecretAccessKey, bucket.SessionToken),
		Region:      aws.String("automatic"),
		Endpoint:    aws.String(bucket.Endpoint),
	}
	newSession, err := session.NewSession(s3Config)
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "创建 Session 失败")
	}
	client := s3manager.NewUploaderWithClient(s3.New(newSession))
	_, err = client.Upload(&s3manager.UploadInput{
		Bucket: aws.String(bucket.Bucket),
		Key:    aws.String(path + fileName),
		Body:   bytes.NewReader(body),
	})
	if err != nil {
		return berror.NewErrorHasError(bcode.ServerInternalError, err, "上传失败")
	}
	return nil
}
