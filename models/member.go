package models

import (
	"github.com/astaxie/beego/validation"
	"regexp"
)

type Member struct {
	Id int `gorm:"primary_key,AUTO_INCREMENT" json:"id"`
	Nickname string `json:"nickname"`
	HeadImg string `json:"head_img"`
	MemberFrom int `json:"member_from"`
	MemberLink string `json:"member_link"`
	IsUse int `json:"is_use"`
	CreateTime int `json:"create_time"`
	UpdateTime int `json:"update_time"`
}

//添加用户
func AddMember(m Member) int{
	if IsExistMember(m.MemberLink){
		return GetMemberInfoByLink(m.MemberLink)
	}
	db.Create(m)
	return m.Id
}

//用户是否存在 true-存在 false-不存在
func IsExistMember(link string) bool {
	count := 0
	db.Model(Member{}).Where("member_link = ?",link).Count(&count)

	return count > 0
}

//获取用户id
func GetMemberInfoByLink(link string) int{
	m := Member{}
	db.Where("member_link = ?",link).First(&m)

	return m.Id
}

//获取未被使用的用户链接
func GetNotUseMemberLink() string {
	member := Member{}
	db.Where("is_use = 0").First(&member)
	member.IsUse = 1
	db.Save(&member)
	return member.MemberLink
}

//设置用户状态为已使用
func SetMemberIsUse(link string){
	db.Model(&Member{}).Where("member_link = ?",link).Update("is_use",1)
}

//验证添加用户数据
func (m *Member)CheckAddMemberData(v *validation.Validation){
	v.Required(m.Nickname,"nick_name").Message("请输入用户昵称")
	v.MinSize(m.Nickname,4,"nick_name").Message("请输入4-15位字符的用户昵称")
	v.MaxSize(m.Nickname,15,"nick_name").Message("请输入4-15位字符的用户昵称")
	v.Required(m.MemberLink,"member_link").Message("请输入用户个性域名")
	v.Match(m.MemberLink,regexp.MustCompile("/^[1-9A-Za-z][0-9A-Za-z_]{3,19}$/"),"member_link").Message("请输入4-20位字符的性域名")
	v.Required(m.HeadImg,"head_img").Message("请上传用户头像")
}

func GetUseMemberLink() []Member {
	member := []Member{}
	db.Select("id,member_link").Where("is_use = 1").Find(&member)
	return member
}

func SetMemberIsNotUse(Id int){
	db.Model(&Member{}).Where("id = ?",Id).Update("is_use",0)
}

func GetMemberList() []Member {
	member := []Member{}
	db.Select("id,member_link").Where("is_use = 1 AND id >= 2000 AND id <= 5000").Find(&member)
	return member
}

