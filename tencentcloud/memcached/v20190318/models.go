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
    tcerr "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/errors"
    tchttp "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/http"
)

// Predefined struct for user
type DescribeInstancesRequestParams struct {
	// 枚举范围： AddTimeStamp, InstanceName, ProjectId
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 查找的关键字
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 子网ID列表
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// VIP列表
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// 0倒序，1正序，默认倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// VPC ID列表
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// 项目ID列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 实例列表的大小，参数默认值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

type DescribeInstancesRequest struct {
	*tchttp.BaseRequest
	
	// 枚举范围： AddTimeStamp, InstanceName, ProjectId
	OrderBy *string `json:"OrderBy,omitempty" name:"OrderBy"`

	// 查找的关键字
	SearchKeys []*string `json:"SearchKeys,omitempty" name:"SearchKeys"`

	// 子网ID列表
	UniqSubnetIds []*string `json:"UniqSubnetIds,omitempty" name:"UniqSubnetIds"`

	// VIP列表
	Vips []*string `json:"Vips,omitempty" name:"Vips"`

	// 0倒序，1正序，默认倒序
	OrderType *int64 `json:"OrderType,omitempty" name:"OrderType"`

	// 实例名称列表
	InstanceNames []*string `json:"InstanceNames,omitempty" name:"InstanceNames"`

	// VPC ID列表
	UniqVpcIds []*string `json:"UniqVpcIds,omitempty" name:"UniqVpcIds"`

	// 项目ID列表
	ProjectIds []*int64 `json:"ProjectIds,omitempty" name:"ProjectIds"`

	// 偏移量，取Limit整数倍
	Offset *uint64 `json:"Offset,omitempty" name:"Offset"`

	// 实例列表的大小，参数默认值100
	Limit *uint64 `json:"Limit,omitempty" name:"Limit"`

	// 实例ID列表
	InstanceIds []*string `json:"InstanceIds,omitempty" name:"InstanceIds"`
}

func (r *DescribeInstancesRequest) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesRequest) FromJsonString(s string) error {
	f := make(map[string]interface{})
	if err := json.Unmarshal([]byte(s), &f); err != nil {
		return err
	}
	delete(f, "OrderBy")
	delete(f, "SearchKeys")
	delete(f, "UniqSubnetIds")
	delete(f, "Vips")
	delete(f, "OrderType")
	delete(f, "InstanceNames")
	delete(f, "UniqVpcIds")
	delete(f, "ProjectIds")
	delete(f, "Offset")
	delete(f, "Limit")
	delete(f, "InstanceIds")
	if len(f) > 0 {
		return tcerr.NewTencentCloudSDKError("ClientError.BuildRequestError", "DescribeInstancesRequest has unknown keys!", "")
	}
	return json.Unmarshal([]byte(s), &r)
}

// Predefined struct for user
type DescribeInstancesResponseParams struct {
	// 实例详细信息列表
	InstanceList []*InstanceListInfo `json:"InstanceList,omitempty" name:"InstanceList"`

	// 实例数量
	TotalNum *int64 `json:"TotalNum,omitempty" name:"TotalNum"`

	// 唯一请求 ID，每次请求都会返回。定位问题时需要提供该次请求的 RequestId。
	RequestId *string `json:"RequestId,omitempty" name:"RequestId"`
}

type DescribeInstancesResponse struct {
	*tchttp.BaseResponse
	Response *DescribeInstancesResponseParams `json:"Response"`
}

func (r *DescribeInstancesResponse) ToJsonString() string {
    b, _ := json.Marshal(r)
    return string(b)
}

// FromJsonString It is highly **NOT** recommended to use this function
// because it has no param check, nor strict type check
func (r *DescribeInstancesResponse) FromJsonString(s string) error {
	return json.Unmarshal([]byte(s), &r)
}

type InstanceListInfo struct {
	// 实例修改时间
	ModTimeStamp *string `json:"ModTimeStamp,omitempty" name:"ModTimeStamp"`

	// 实例隔离时间
	IsolateTimeStamp *string `json:"IsolateTimeStamp,omitempty" name:"IsolateTimeStamp"`

	// 实例是否设置自动续费标识，1：设置自动续费；0：未设置自动续费
	AutoRenewFlag *int64 `json:"AutoRenewFlag,omitempty" name:"AutoRenewFlag"`

	// 仓库ID
	SetId *int64 `json:"SetId,omitempty" name:"SetId"`

	// 实例当前状态，0：发货中；1：运行中；2：创建失败；4：销毁中；5：隔离中；6：下线中
	Status *int64 `json:"Status,omitempty" name:"Status"`

	// 实例内置ID
	CmemId *int64 `json:"CmemId,omitempty" name:"CmemId"`

	// 实例关联的标签信息
	// 注意：此字段可能返回 null，表示取不到有效值。
	Tags []*TagInfo `json:"Tags,omitempty" name:"Tags"`

	// 实例ID
	InstanceId *string `json:"InstanceId,omitempty" name:"InstanceId"`

	// 地域id 1--广州 4--上海 5-- 香港 6--多伦多 7--上海金融 8--北京 9-- 新加坡 11--深圳金融 15--美西（硅谷）16--成都 17--德国 18--韩国 19--重庆 21--印度 22--美东（弗吉尼亚）23--泰国 24--俄罗斯 25--日本
	RegionId *int64 `json:"RegionId,omitempty" name:"RegionId"`

	// 实例描述信息
	InstanceDesc *string `json:"InstanceDesc,omitempty" name:"InstanceDesc"`

	// 过期策略
	Expire *int64 `json:"Expire,omitempty" name:"Expire"`

	// vpc网络下子网id 如：46315
	SubnetId *int64 `json:"SubnetId,omitempty" name:"SubnetId"`

	// 项目ID
	ProjectId *int64 `json:"ProjectId,omitempty" name:"ProjectId"`

	// 实例创建时间
	AddTimeStamp *string `json:"AddTimeStamp,omitempty" name:"AddTimeStamp"`

	// 区域ID
	ZoneId *int64 `json:"ZoneId,omitempty" name:"ZoneId"`

	// 计费模式：0-按量计费，1-包年包月
	PayMode *int64 `json:"PayMode,omitempty" name:"PayMode"`

	// vpc网络id 如：75101
	VpcId *int64 `json:"VpcId,omitempty" name:"VpcId"`

	// 实例名称
	InstanceName *string `json:"InstanceName,omitempty" name:"InstanceName"`

	// 实例截止时间
	DeadlineTimeStamp *string `json:"DeadlineTimeStamp,omitempty" name:"DeadlineTimeStamp"`

	// vpc网络id 如：vpc-fk33jsf43kgv
	UniqVpcId *string `json:"UniqVpcId,omitempty" name:"UniqVpcId"`

	// 实例vip
	Vip *string `json:"Vip,omitempty" name:"Vip"`

	// vpc网络下子网id 如：subnet-fd3j6l35mm0
	UniqSubnetId *string `json:"UniqSubnetId,omitempty" name:"UniqSubnetId"`

	// 用户AppID
	AppId *int64 `json:"AppId,omitempty" name:"AppId"`

	// 实例端口号
	Vport *int64 `json:"Vport,omitempty" name:"Vport"`
}

type TagInfo struct {
	// 标签键
	TagKey *string `json:"TagKey,omitempty" name:"TagKey"`

	// 标签值
	TagValue *string `json:"TagValue,omitempty" name:"TagValue"`
}