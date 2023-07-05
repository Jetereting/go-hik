package hik

import (
	"errors"
	"github.com/Jetereting/go-hik"
	"github.com/tidwall/gjson"
)

var HIK go_hik.HKConfig

// OrgAdd 添加组织
func OrgAdd(name, code, pCode string) error {
	_, e := HIK.HttpPost("/api/resource/v1/org/batch/add", []map[string]string{{
		"orgName":         name,
		"orgIndexCode":    code,
		"parentIndexCode": pCode,
	}})
	return e
}

// OrgUp 更新组织
func OrgUp(code, name string) error {
	_, e := HIK.HttpPost("/api/resource/v1/org/single/update", map[string]string{
		"orgName":      name,
		"orgIndexCode": code,
	})
	return e
}

// OrgDel 删除组织-需要下面没有人才能成功
func OrgDel(codes []string) error {
	// 删除权限
	personMap := []map[string]interface{}{
		{
			"indexCodes":     codes,
			"personDataType": "org",
		},
	}
	_, e := HIK.HttpPost("/api/acps/v1/auth_config/delete", map[string]interface{}{
		"personDatas": personMap,
	})
	if e != nil {
		return e
	}
	// 删除组织
	_, e = HIK.HttpPost("/api/resource/v1/org/batch/delete", map[string]interface{}{
		"indexCodes": codes,
	})
	return e
}

// UserAdd 添加人员
func UserAdd(name, code, cardNum, orgCode, faceData string) error {
	if cardNum == "" {
		cardNum = code
	}
	isNoFace := false
	if faceData == "" {
		isNoFace = true
		faceData = "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"
	}
	resp, e := HIK.HttpPost("/api/resource/v2/person/single/add", map[string]interface{}{
		"personName":      name,
		"personId":        code,
		"gender":          "0",
		"certificateType": "990",
		"certificateNo":   cardNum,
		"orgIndexCode":    orgCode,
		"faces":           []map[string]string{{"faceData": faceData}},
	})
	if e != nil {
		return e
	}
	if isNoFace {
		//删除人脸
		_, e = HIK.HttpPost("/api/resource/v1/face/single/delete", map[string]interface{}{
			"faceId": resp.Get("faceId").String(),
		})
	}
	return e
}

// UserDel 删除人员
func UserDel(codes []string) error {
	_, e := HIK.HttpPost("/api/resource/v1/person/batch/delete", map[string]interface{}{
		"personIds": codes,
	})
	return e
}

// UserFaceUp 修改人员人脸
func UserFaceUp(userCode, faceData string) error {
	resp, e := HIK.HttpPost("/api/resource/v2/person/advance/personList", map[string]interface{}{
		"pageNo":    1,
		"pageSize":  1,
		"personIds": userCode,
	})
	if e != nil {
		return e
	}
	if len(resp.Get("list").Array()) == 0 {
		return errors.New("未找到人员")
	}
	faceId := resp.Get("list").Array()[0].Get("personPhoto").Get("personPhotoIndexCode").Str
	if faceId == "" {
		//新增人脸
		_, e = HIK.HttpPost("/api/resource/v1/face/single/add", map[string]interface{}{
			"personId": userCode,
			"faceData": faceData,
		})
	} else {
		//修改人脸
		_, e = HIK.HttpPost("/api/resource/v1/face/single/update", map[string]interface{}{
			"faceId":   faceId,
			"faceData": faceData,
		})
	}
	return e
}

// PermissionOrgAddDoor 组织添加门禁设备权限
func PermissionOrgAddDoor(orgCodes, doors []string) error {
	var doorsMap []map[string]interface{}
	for _, v := range doors {
		doorsMap = append(doorsMap, map[string]interface{}{
			"resourceIndexCode": v,
			"resourceType":      "acsDevice",
			"channelNos":        []string{"1"},
		})
	}
	personMap := []map[string]interface{}{
		{
			"indexCodes":     orgCodes,
			"personDataType": "org",
		},
	}
	// 删除原有权限
	_, e := HIK.HttpPost("/api/acps/v1/auth_config/delete", map[string]interface{}{
		"personDatas": personMap,
	})
	if e != nil {
		return e
	}
	if len(doors) == 0 {
		return nil
	}
	// 添加权限配置
	reqBody := map[string]interface{}{
		"personDatas":   personMap,
		"resourceInfos": doorsMap,
	}
	_, e = HIK.HttpPost("/api/acps/v1/auth_config/add", reqBody)
	if e != nil {
		return e
	}
	// 快捷下载权限
	_, e = HIK.HttpPost("/api/acps/v1/authDownload/configuration/shortcut", map[string]interface{}{
		"taskType":      4,
		"resourceInfos": doorsMap,
	})
	return e
}

// PermissionProgress 权限任务进度
func PermissionProgress(taskId string) (gjson.Result, error) {
	return HIK.HttpPost("/api/acps/v1/authDownload/task/progress", map[string]interface{}{
		"taskId": taskId,
	})
}

// PermissionList 权限查询
func PermissionList(userCodes []string) {
	HIK.HttpPost("/api/acps/v1/auth_item/list/search", map[string]interface{}{
		"personIds": userCodes,
		"queryType": "acsDevice",
		"pageNo":    1,
		"pageSize":  100,
	})
}

func DoorList() (gjson.Result, error) {
	return HIK.HttpPost("/api/resource/v2/acsDevice/search", map[string]interface{}{
		"pageNo":    1,
		"pageSize":  1000,
		"orderBy":   "name",
		"orderType": "asc",
	})
}
