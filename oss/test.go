package service

//
//import (
//	"bytes"
//	"context"
//	"encoding/base64"
//	"encoding/csv"
//	"encoding/json"
//	"errors"
//	"fmt"
//	logger "git.100tal.com/wangxiao_go_lib/xesLogger"
//	"git.100tal.com/wangxiao_jiaoyan_go/library/sdk/analysesdk"
//	analyseproto "git.100tal.com/wangxiao_jiaoyan_go/library/sdk/analysesdk/proto"
//	jsoniter "github.com/json-iterator/go"
//	"net/url"
//	entity2 "saas-school-service/internal/biz/customer/domain/entity"
//	customerRepo "saas-school-service/internal/biz/customer/infra/repo"
//	"saas-school-service/internal/biz/event/domain/entity"
//	common2 "saas-school-service/internal/biz/event/infra/common"
//	repo2 "saas-school-service/internal/biz/product/infra/repo"
//	entity3 "saas-school-service/internal/biz/vipconfig/domain/entity"
//	"saas-school-service/internal/biz/vipconfig/infra/repo"
//	"saas-school-service/pkg/apicall"
//	"saas-school-service/pkg/basic"
//	"saas-school-service/pkg/common"
//	"saas-school-service/pkg/oss/upload"
//	proto "saas-school-service/pkg/proto/event"
//	"strconv"
//	"strings"
//	"sync"
//	"time"
//)
//
//func (e *eventSrv) GetPuvData(ctx context.Context, req *proto.PuvReq) (resp *proto.PuvResp, err error) {
//	puvResp, err := analysesdk.GetEventPuvData(ctx, &analyseproto.GetEventPUVDataReq{
//		EId:     req.EID,
//		Source:  int64(req.Source),
//		Channel: req.Channel,
//	})
//
//	if err != nil || puvResp == nil {
//		return nil, err
//	}
//
//	resp = &proto.PuvResp{
//		Pv: puvResp.PV,
//		Uv: puvResp.UV,
//	}
//	return
//}
//
//func (e *eventSrv) ExportData(ctx context.Context, req *proto.DataListReq) (err error) {
//	tag := "[eventSrv.ExportData]"
//	dataList, err := e.GetDataList(ctx, req)
//	if err != nil {
//		return
//	}
//
//	if len(dataList.List) == 0 {
//		return errors.New("数据不存在")
//	}
//
//	resp := make([][]string, 0, len(dataList.List)+1)
//	resp = append(resp, []string{"ID", "来源", "渠道", "用户名称", "用户id", "学部", "年级", "学科", "领取时间", "资料信息"})
//	for k := range dataList.List {
//		resp = append(resp, []string{
//			strconv.FormatInt(dataList.List[k].ID, 10),
//			dataList.List[k].SourceName,
//			strconv.FormatInt(dataList.List[k].Channel, 10),
//			dataList.List[k].UName,
//			strconv.FormatInt(dataList.List[k].UID, 10),
//			dataList.List[k].PeriodName,
//			dataList.List[k].GradeName,
//			dataList.List[k].SubjectName,
//			dataList.List[k].DrawTime,
//			func() (infos string) {
//				if len(dataList.List[k].RInfos) == 0 {
//					return
//				}
//				arrInfos := make([]string, 0, len(dataList.List[k].RInfos))
//				for _, j := range dataList.List[k].RInfos {
//					arrInfos = append(arrInfos, fmt.Sprintf("%s[%v]", j.RName, j.RID))
//				}
//
//				return strings.Join(arrInfos, "|")
//			}(),
//		})
//	}
//
//	adminUsersMap := apicall.GetAdminUserList(ctx, &apicall.AdminUserListReq{AdminIds: []int64{req.AdminId}})
//	if _, ok := adminUsersMap[req.AdminId]; !ok {
//		return errors.New("操作用户不存在")
//	}
//
//	dataBytes := new(bytes.Buffer)
//	dataBytes.WriteString("\xEF\xBB\xBF")
//	wr := csv.NewWriter(dataBytes)
//	for _, body := range resp {
//		if err = wr.Write(body); err != nil {
//			logger.Ex(ctx, tag, "write csv error, [err]:%+v, [req]:%+v", err, req)
//		}
//	}
//	wr.Flush()
//
//	var upUrl string
//	if upUrl, err = upload.NewOssUpload().PutObjectFromReader(ctx, &upload.PutObjectReq{
//		FileType:   upload.FileResource,
//		Reader:     dataBytes,
//		TargetPath: fmt.Sprintf("event/event_user_%v_%d.csv", adminUsersMap[req.AdminId].Workcode, time.Now().Unix()),
//	}, true); err != nil {
//		logger.Ex(ctx, tag, "upload.PutObject error, [err]:%+v, [req]:%+v", err, req)
//		return err
//	}
//
//	yachApi := apicall.NewYachApi()
//	msgStr, err := jsoniter.Marshal(map[string]interface{}{
//		"msgtype": "text",
//		"text": map[string]string{
//			"content": "你导出的教研云营销模板任务已经生成完成，下载点击: " + upUrl,
//		},
//	})
//	if err != nil {
//		logger.Ex(ctx, tag, "json encode msg error, [err]:%+v, [req]:%+v", err, req)
//		return err
//	}
//
//	_, err = yachApi.SendNotice(ctx, "workcode", adminUsersMap[req.AdminId].Workcode, msgStr)
//	return
//}
//
//func (e *eventSrv) GetDataList(ctx context.Context, req *proto.DataListReq) (resp *proto.DataListResp, err error) {
//	resp = new(proto.DataListResp)
//	eventInfo, err := e.eventRepo.GetEventInfo(ctx, &entity.EventSearch{
//		Id: req.EID,
//	})
//	if err != nil || eventInfo == nil {
//		return nil, err
//	}
//
//	resp.BaseInfo = &proto.BaseInfoItem{
//		Id:       eventInfo.Id,
//		Title:    eventInfo.Title,
//		ShowTime: time.Now().AddDate(0, 0, -1).Format("2006-01-02"),
//	}
//
//	total, list, err := e.eventRepo.GetList(ctx, &entity.EventUserSearch{
//		Eid:      req.EID,
//		Source:   req.Source,
//		Channel:  req.Channel,
//		Page:     req.Page,
//		PageSize: req.PageSize,
//	})
//	if err != nil {
//		return
//	}
//
//	resp.Total = total
//	if len(list) == 0 {
//		return
//	}
//
//	var (
//		wg             sync.WaitGroup
//		rids           = make([]string, 0, len(list))
//		userIds        = make([]int64, 0, len(list))
//		periodNameMap  = make(map[int64]string)
//		subjectNameMap = make(map[int64]string)
//		gradeNameMap   = make(map[int64]string)
//		rIDsMap        = make(map[int64][]string, len(list))
//		userInfoMap    = make(map[int64]*entity2.UserInfo, len(list))
//		rightInfoMap   = make(map[int64]*entity.EventRight, len(rids))
//		rightNameMap   = make(map[int64]struct {
//			Id   int64
//			Name string
//		}, len(rids))
//
//		getPSGMap = func() {
//			data, _ := basic.NewBasicConfig().GetBaseOptions(ctx)
//			for k := range data.Periods {
//				periodNameMap[data.Periods[k].ID] = data.Periods[k].Name
//			}
//
//			for k := range data.Subjects {
//				subjectNameMap[data.Subjects[k].ID] = data.Subjects[k].Name
//			}
//
//			for k := range data.Grades {
//				gradeNameMap[data.Grades[k].ID] = data.Grades[k].Name
//			}
//		}
//
//		getUserInfos = func() {
//			defer wg.Done()
//			userInfos, _ := customerRepo.GetUserRepo(customerRepo.GetUserOrgRepo()).BatchGetByUids(ctx, common.SliceInt2Unique(userIds))
//			if len(userInfos) == 0 {
//				return
//			}
//
//			for k := range userInfos {
//				userInfoMap[userInfos[k].Id] = userInfos[k]
//			}
//		}
//
//		getUserName = func(userId int64) (name string) {
//			if _, ok := userInfoMap[userId]; !ok {
//				return
//			}
//
//			return userInfoMap[userId].UName
//		}
//
//		getRightInfos = func() {
//			defer wg.Done()
//			rightInfos, _ := e.eventRepo.GetEventRights(ctx, &entity.EventRightSearch{
//				Eid: eventInfo.Id,
//				Rid: common.SliceString2Unique(rids),
//			})
//			if len(rightInfos) == 0 {
//				return
//			}
//			var (
//				vipIDs     = make([]int64, 0)
//				vipMap     = make(map[int64]string)
//				productIDs = make([]int64, 0)
//				productMap = make(map[int64]string)
//			)
//
//			for k := range rightInfos {
//				rightInfoMap[rightInfos[k].Id] = rightInfos[k]
//				if rightInfos[k].RType == common2.RTypeProduct {
//					productIDs = append(productIDs, rightInfos[k].RId)
//				}
//				if rightInfos[k].RType == common2.RTypeVip {
//					vipIDs = append(vipIDs, rightInfos[k].RId)
//				}
//			}
//			if len(vipIDs) > 0 {
//				vipInfos, _ := repo.GetVipConfigRepo().FindVipConfigs(ctx, &entity3.VipConfigSearch{
//					VipIds: vipIDs,
//					//Status: vipcfgsdk.StatusOn,
//				})
//				if len(vipInfos) > 0 {
//					for k := range vipInfos {
//						vipMap[vipInfos[k].Id] = vipInfos[k].Name
//					}
//				}
//			}
//
//			if len(productIDs) > 0 {
//				productInfos, _ := repo2.GetProductRepo().GetProductByIds(ctx, productIDs)
//				if len(productInfos) > 0 {
//					for k := range productInfos {
//						productMap[productInfos[k].Id] = productInfos[k].Name
//					}
//				}
//			}
//
//			for k := range rightInfos {
//				switch rightInfos[k].RType {
//				case common2.RTypeVip:
//					if _, ok := vipMap[rightInfos[k].RId]; !ok {
//						continue
//					}
//					rightNameMap[rightInfos[k].Id] = struct {
//						Id   int64
//						Name string
//					}{Id: rightInfos[k].RId, Name: vipMap[rightInfos[k].RId]}
//				case common2.RTypeProduct:
//					if _, ok := productMap[rightInfos[k].RId]; !ok {
//						continue
//					}
//					rightNameMap[rightInfos[k].Id] = struct {
//						Id   int64
//						Name string
//					}{Id: rightInfos[k].RId, Name: productMap[rightInfos[k].RId]}
//				default:
//					continue
//				}
//			}
//		}
//	)
//
//	for k := range list {
//		rid := strings.Split(list[k].Rid, "|")
//		tmpRids := make([]string, 0, len(rid))
//		for _, v := range rid {
//			if v == "" {
//				continue
//			}
//
//			tmpRids = append(tmpRids, v)
//		}
//
//		rIDsMap[list[k].Id] = tmpRids
//		rids = append(rids, tmpRids...)
//		userIds = append(userIds, list[k].UserId)
//	}
//
//	wg.Add(2)
//	go getUserInfos()
//	go getRightInfos()
//	wg.Wait()
//
//	getPSGMap()
//	for k := range list {
//		resp.List = append(resp.List, &proto.DataListItem{
//			ID:     list[k].Id,
//			Source: list[k].Source,
//			SourceName: func() (name string) {
//				if _, ok := common2.SourceMap[list[k].Source]; ok {
//					name = common2.SourceMap[list[k].Source]
//				}
//				return
//			}(),
//			Channel: list[k].Channel,
//			UID:     list[k].UserId,
//			UName:   getUserName(list[k].UserId),
//			Period:  list[k].PeriodId,
//			Grade:   list[k].GradeId,
//			Subject: list[k].SubjectId,
//			PeriodName: func() (name string) {
//				if _, ok := periodNameMap[list[k].PeriodId]; !ok {
//					return
//				}
//				return periodNameMap[list[k].PeriodId]
//			}(),
//			GradeName: func() (name string) {
//				if _, ok := gradeNameMap[list[k].GradeId]; !ok {
//					return
//				}
//				return gradeNameMap[list[k].GradeId]
//			}(),
//			SubjectName: func() (name string) {
//				if _, ok := subjectNameMap[list[k].SubjectId]; !ok {
//					return
//				}
//				return subjectNameMap[list[k].SubjectId]
//			}(),
//			DrawTime: list[k].CreatedAt.Format("2006-01-02 15:04:05"),
//			RInfos: func() (infos []*proto.RInfosItem) {
//				if _, ok := rIDsMap[list[k].Id]; !ok {
//					return
//				}
//
//				infos = make([]*proto.RInfosItem, 0, len(rIDsMap[list[k].Id]))
//				for _, v := range rIDsMap[list[k].Id] {
//					rid, _ := strconv.ParseInt(v, 0, 0)
//					if _, ok := rightNameMap[rid]; !ok {
//						continue
//					}
//					infos = append(infos, &proto.RInfosItem{
//						RID:   rightNameMap[rid].Id,
//						RName: rightNameMap[rid].Name,
//					})
//				}
//				return
//			}(),
//		})
//	}
//
//	return
//}
//
////yachAPI:
/////cmpts/msgchl/yach/notice/send
//
//// SendNotice 企业助手发送工作通知
//func (y *YachApi) SendNotice(ctx context.Context, userType, workCode string, msg []byte) (data *SendNoticeMsgId, err error) {
//	tag := "[apicall.yach.SendNotice]"
//	ticket, err := y.getTicket(ctx)
//	if ticket == "" {
//		return nil, err
//	}
//	base64msg := base64.StdEncoding.EncodeToString(msg)
//	endpoint := sendNoticeApi
//	params := map[string]string{
//		"ticket":      ticket,
//		"user_type":   userType,
//		"userid_list": workCode,
//		"message":     base64msg,
//	}
//	resp, err := request.PostForm(ctx, endpoint, params)
//	if err != nil {
//		logger.Ex(ctx, tag, "api call request failed, workcode:%s, err:%s", workCode, err)
//		return nil, err
//	}
//
//	var accountResp YachSendNoticeResp
//	json.Unmarshal(resp, &accountResp)
//	if accountResp.Errcode != 0 {
//		logger.Ex(ctx, tag, "api call unmarshal failed, workcode:%s, resp:%+v, err:%s", workCode, string(resp), err)
//		return nil, err
//	}
//	return &accountResp.Data, nil
//}
//
//func (o *ossUpload) PutObjectFromReader(ctx context.Context, req *PutObjectReq, isNeedCdn bool) (resUrl string, err error) {
//	tag := "[oss.ossUpload.PutObjectFromReader]"
//
//	if req.Reader == nil {
//		logger.Ex(ctx, tag, "req.Reader is nil error [req]:%+v", req)
//		return
//	}
//
//	var (
//		retryTimes = 0
//	)
//
//	for {
//		if retryTimes > RetryTimes {
//			logger.Ex(ctx, "o.bucket.PutObject retry.timeout, [retryTimes]:%d", RetryTimes)
//			break
//		}
//
//		if err = o.bucket.PutObject(req.TargetPath, req.Reader, ossSdk.ObjectACL(ossSdk.ACLPrivate)); err != nil {
//			logger.Ex(ctx, tag, "o.bucket.PutObject error, [err]:%+v, [req]:%+v", err, req)
//			retryTimes++
//			continue
//		}
//		resUrl, err = o.bucket.SignURL(req.TargetPath, ossSdk.HTTPGet, 3600)
//		if err != nil {
//			logger.Ex(ctx, tag, "o.bucket.SignURL error, [err]:%+v, [req]:%+v", err, req)
//		}
//		if isNeedCdn {
//			urlStruct, _ := url.Parse(resUrl)
//			path := urlStruct.RequestURI()
//			resUrl = fmt.Sprintf("%s%s", o.cdn, path)
//		}
//		return
//	}
//	return
//}
