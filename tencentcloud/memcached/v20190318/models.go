// Copyright (c) 2017-2018 THL A29 Limited, a Tencent company. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//    http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package v20190318

import (
    "encoding/json"

    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest

	// 实例ID组成的数组，数组下标从0开始
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds" list`

	// 实例名称组成的数组，数组下标从0开始
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames" list`

	// 实例列表的大小，参数默认值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 枚举范围： AddTimeStamp, InstanceName, ProjectId
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 0倒序，1正序，默认倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 项目ID组成的数组，数组下标从0开始
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds" list`

	// 搜索关键词：支持实例ID、实例名称、完整IP
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys" list`

	// 子网ID数组，数组下标从0开始，如：subnet-fdj24n34j2
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds" list`

	// 私有网络ID数组，数组下标从0开始，如果不传则默认选择基础网络，如：vpc-sad23jfdfk
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds" list`

	// 实例服务IP组成的数组，数组下标从0开始
	Vips []*string `json:"Vips,omitempty" name:"Vips" list`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesRequest) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *struct {

		// 实例详细信息列表
		InstanceList []*InstanceListInfo `json:"InstanceList,omitempty" name:"InstanceList" list`

		// 实例数量
		TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

		// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
		RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
	} `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

func (r *DescribeInstancesResponse) FromJsonString(s string) error {
    return json.Unmarshal([]byte(s), &r)
}

type InstanceListInfo struct {

	// 实例关联的标签信息
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags" list`

	// 实例创建时间
	AddTimeStamp *string `json:"AddTimeStamp,omitempty" name:"AddTimeStamp"`

	// 用户AppID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 实例是否设置自动续费标识，1：设置自动续费；0：未设置自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 实例内置ID
	CmemId *int64 `json:"CmemId,omitempty" name:"CmemId"`

	// 实例截止时间
	DeadlineTimeStamp *string `json:"DeadlineTimeStamp,omitempty" name:"DeadlineTimeStamp"`

	// 过期策略
	Expire *int64 `json:"Expire,omitempty" name:"Expire"`

	// 实例描述信息
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例隔离时间
	IsolateTimeStamp *string `json:"IsolateTimeStamp,omitempty" name:"IsolateTimeStamp"`

	// 实例修改时间
	ModTimeStamp *string `json:"ModTimeStamp,omitempty" name:"ModTimeStamp"`

	// 计费模式：0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 地域id 1--广州 4--上海 5-- 香港 6--多伦多 7--上海金融 8--北京 9-- 新加坡 11--深圳金融 15--美西（硅谷）16--成都 17--德国 18--韩国 19--重庆 21--印度 22--美东（弗吉尼亚）23--泰国 24--俄罗斯 25--日本
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 仓库ID
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// 实例当前状态，0：待初始化；1：实例在流程中；2：实例运行中；-2：实例已隔离；-3：实例待删除
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// vpc网络下子网id 如：46315
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// vpc网络下子网id 如：subnet-fd3j6l35mm0
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// vpc网络id 如：vpc-fk33jsf43kgv
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 实例vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vpc网络id 如：75101
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 实例端口号
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`
}

type TagInfo struct {

	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}
