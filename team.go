package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	createTeam     = neteaseBaseURL + "/team/create.action"         // 创建群聊
	joinTeam       = neteaseBaseURL + "/team/add.action"            // 加入群聊（高级群）
	leaveTam       = neteaseBaseURL + "/team/leave.action"          // 主动退群 （高级群）
	updateTeamNick = neteaseBaseURL + "/team/updateTeamNick.action" // 修改群成员昵称 （高级群）
	removeTeam     = neteaseBaseURL + "/team/remove.action"         // 解散群 （高级群）
	kickTeam       = neteaseBaseURL + "/team/kick.action"           // 踢人出群 （高级群）
	updateTeam     = neteaseBaseURL + "/team/update.action"         // 修改群信息
	addManager     = neteaseBaseURL + "/team/addManager.action"     // 添加管理员
	removeManager  = neteaseBaseURL + "/team/removeManager.action"  // 移除管理员
	changeOwner    = neteaseBaseURL + "/team/changeOwner.action"    // 转移群主
	muteTlistAll   = neteaseBaseURL + "/team/muteTlistAll.action"   // 禁言群组
	muteTlist      = neteaseBaseURL + "/team/muteTlist.action"      // 禁言指定群成员
	muteTeam       = neteaseBaseURL + "/team/muteTeam.action"       // 设置群消息免打扰
	queryDetail    = neteaseBaseURL + "/team/queryDetail.action"    // 设置群消息免打扰
)

type CreateTeamRequest struct {
	Tname                   string  `json:"tname"`
	Owner                   string  `json:"owner"`
	Members                 string  `json:"members"`
	Announcement            *string `json:"announcement"`
	Intro                   *string `json:"intro"`
	Msg                     string  `json:"msg"`
	Magree                  *int    `json:"magree"`
	Joinmode                int     `json:"joinmode"`
	Custom                  *string `json:"custom"`
	Icon                    *string `json:"icon"`
	Beinvitemode            *int    `json:"beinvitemode"`
	Invitemode              *int    `json:"invitemode"`
	Uptinfomode             *int    `json:"uptinfomode"`
	Upcustommode            *int    `json:"upcustommode"`
	TeamMemberLimit         *int    `json:"teamMemberLimit"`
	IsNotifyCloseOnline     *int    `json:"isNotifyCloseOnline"`
	IsNotifyClosePersistent *int    `json:"isNotifyClosePersistent"`
	Attach                  *string `json:"attach"`
	Bid                     *string `json:"bid"`
}

type JoinTeamRequest struct {
	Tid     string  `json:"tid"`
	Owner   string  `json:"owner"`
	Members string  `json:"members"`
	Magree  *int    `json:"magree"`
	Msg     string  `json:"msg"`
	Attach  *string `json:"attach"`
}

type UpdateTeamRequest struct {
	Tid             string  `json:"tid"`
	Tname           *string `json:"tname"`
	Owner           string  `json:"owner"`
	Announcement    *string `json:"announcement"`
	Intro           *string `json:"intro"`
	Icon            *string `json:"icon"`
	Joinmode        *int    `json:"joinmode"`
	Beinvitemode    *int    `json:"beinvitemode"`
	Invitemode      *int    `json:"invitemode"`
	Uptinfomode     *int    `json:"uptinfomode"`
	Upcustommode    *int    `json:"upcustommode"`
	TeamMemberLimit *int    `json:"teamMemberLimit"`
	Custom          *string `json:"custom"`
	Attach          *string `json:"attach"`
	Bid             *string `json:"bid"`
}

type ManagerRequest struct {
	Tid     string  `json:"tid"`
	Owner   string  `json:"owner"`
	Members string  `json:"members"`
	Attach  *string `json:"attach"`
}

type ChangeOwnerRequest struct {
	Tid      string  `json:"tid"`
	Owner    string  `json:"owner"`
	Newowner string  `json:"newowner"`
	Leave    int     `json:"levae"`
	Attach   *string `json:"attach"`
}

type MuteTlistAllRequest struct {
	Tid      string  `json:"tid"`
	Owner    string  `json:"owner"`
	Mute     *bool   `json:"mute"`
	MuteType *int    `json:"mute_type"`
	Attach   *string `json:"attach"`
}

type MuteTlistRequest struct {
	Tid    string  `json:"tid"`
	Owner  string  `json:"owner"`
	Accid  string  `json:"accid"`
	Mute   int     `json:"mute"`
	Attach *string `json:"attach"`
}

type KickRequest struct {
	Tid     string  `json:"tid"`
	Owner   string  `json:"owner"`
	Member  *string `json:"member"`
	Members *string `json:"members"`
	Attach  *string `json:"attach"`
}

type UpdateTeamNickRequest struct {
	Tid    string  `json:"tid"`
	Owner  string  `json:"owner"`
	Accid  string  `json:"accid"`
	Nick   *string `json:"nick"`
	Custom *string `json:"custom"`
}

type LeaveRequest struct {
	Tid    string  `json:"tid"`
	Accid  string  `json:"accid"`
	Attach *string `json:"attach"`
}

type MuteTeamRequest struct {
	Tid   string `json:"tid"`
	Accid string `json:"accid"`
	Ope   int    `json:"ope"`
}

type RemoveRequest struct {
	Tid    string  `json:"tid"`
	Owner  string  `json:"owner"`
	Attach *string `json:"attach"`
}

type QueryDetailRequest struct {
	Tid string `json:"tid"`
}

/** 拉人入群
 * @param tid 群id
 * @param owner 邀请人的用户账号accid
 * @param members 被邀请入群的用户列表
 * @param msg 邀请发送的文字
 */
// func (c *ImClient) JoinTeam(tid, owner, members, msg string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "owner": owner, "members": members, "msg": msg}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(joinTeam)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// func (c *ImClient) LeaveTeam(tid, accid string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "accid": accid}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(leaveTam)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// func (c *ImClient) UpdateTeamNick(tid, owner, accid, nick string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "owner": owner, "accid": accid, "nick": nick}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(updateTeamNick)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// func (c *ImClient) RemoveTeam(tid, owner string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "owner": owner}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(removeTeam)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// func (c *ImClient) KickBatchTeam(tid, owner, members, attach string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "owner": owner, "members": members}
// 	if len(attach) > 0 {
// 		param["attach"] = attach
// 	}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(kickTeam)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// func (c *ImClient) KickTeam(tid, owner, member, attach string) (map[string]*json.RawMessage, error) {
// 	param := map[string]string{"tid": tid, "owner": owner, "member": member}
// 	if len(attach) > 0 {
// 		param["attach"] = attach
// 	}
// 	client := c.client.R()
// 	c.setCommonHead(client)
// 	client.SetFormData(param)

// 	resp, err := client.Post(kickTeam)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var jsonRes map[string]*json.RawMessage
// 	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
// 	if err != nil {
// 		return nil, err
// 	}

// 	var code int
// 	err = json.Unmarshal(*jsonRes["code"], &code)
// 	if err != nil {
// 		return nil, err
// 	}

// 	if code != 200 {
// 		var msg string
// 		json.Unmarshal(*jsonRes["desc"], &msg)
// 		return nil, errors.New(msg)
// 	}

// 	return jsonRes, nil
// }

// 踢人出群
func (c *ImClient) KickTeam(kickRequest KickRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": kickRequest.Tid, "owner": kickRequest.Owner}

	if kickRequest.Member != nil {
		param["member"] = *kickRequest.Member
	}

	if kickRequest.Members != nil {
		param["members"] = *kickRequest.Members
	}

	if kickRequest.Attach != nil {
		param["attach"] = *kickRequest.Attach
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(kickTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 设置群消息免打扰
func (c *ImClient) MuteTeam(muteTeamRequest MuteTeamRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": muteTeamRequest.Tid, "accid": muteTeamRequest.Accid, "ope": strconv.Itoa(muteTeamRequest.Ope)}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 拉人入群
func (c *ImClient) JoinTeam(joinTeamRequest JoinTeamRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": joinTeamRequest.Tid, "owner": joinTeamRequest.Owner, "members": joinTeamRequest.Members, "msg": joinTeamRequest.Msg}
	if joinTeamRequest.Magree != nil {
		param["magree"] = strconv.Itoa(*joinTeamRequest.Magree)
	}
	if joinTeamRequest.Attach != nil {
		param["attach"] = *joinTeamRequest.Attach
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(joinTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 添加管理员
func (c *ImClient) AddManager(managerRequest ManagerRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": managerRequest.Tid, "owner": managerRequest.Owner, "members": managerRequest.Members}
	if managerRequest.Attach != nil {
		param["attach"] = *managerRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(addManager)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 移除管理员
func (c *ImClient) RemoveManager(managerRequest ManagerRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": managerRequest.Tid, "owner": managerRequest.Owner, "members": managerRequest.Members}
	if managerRequest.Attach != nil {
		param["attach"] = *managerRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(removeManager)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 转移群主
func (c *ImClient) ChangeOwner(changeOwnerRequest ChangeOwnerRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": changeOwnerRequest.Tid, "owner": changeOwnerRequest.Owner, "newowner": changeOwnerRequest.Newowner, "leave": strconv.Itoa(changeOwnerRequest.Leave)}
	if changeOwnerRequest.Attach != nil {
		param["attach"] = *changeOwnerRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(changeOwner)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 禁言群组
func (c *ImClient) MuteTlistAll(muteTlistAllRequest MuteTlistAllRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": muteTlistAllRequest.Tid, "owner": muteTlistAllRequest.Owner}
	if muteTlistAllRequest.Mute != nil {
		if *muteTlistAllRequest.Mute {
			param["mute"] = "true"
		} else {
			param["mute"] = "false"
		}
	}
	if muteTlistAllRequest.MuteType != nil {
		param["muteType"] = strconv.Itoa(*muteTlistAllRequest.MuteType)
	}
	if muteTlistAllRequest.Attach != nil {
		param["attach"] = *muteTlistAllRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteTlistAll)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 禁言指定群成员
func (c *ImClient) MuteTlist(muteTlistRequest MuteTlistRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": muteTlistRequest.Tid, "owner": muteTlistRequest.Owner, "accid": muteTlistRequest.Accid, "mute": strconv.Itoa(muteTlistRequest.Mute)}

	if muteTlistRequest.Attach != nil {
		param["attach"] = *muteTlistRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(muteTlist)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 解散群聊
func (c *ImClient) RemoveTeam(removeRequest RemoveRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": removeRequest.Tid, "owner": removeRequest.Owner}

	if removeRequest.Attach != nil {
		param["attach"] = *removeRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(removeTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 创建群聊
func (c *ImClient) CreateTeam(createTeamRequest CreateTeamRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tname": createTeamRequest.Tname, "owner": createTeamRequest.Owner, "members": createTeamRequest.Members, "msg": createTeamRequest.Msg, "joinmode": strconv.Itoa(createTeamRequest.Joinmode)}

	if createTeamRequest.Announcement != nil {
		param["announcement"] = *createTeamRequest.Announcement
	}
	if createTeamRequest.Magree != nil {
		param["magree"] = strconv.Itoa(*createTeamRequest.Magree)
	}
	if createTeamRequest.Custom != nil {
		param["custom"] = *createTeamRequest.Custom
	}
	if createTeamRequest.Icon != nil {
		param["icon"] = *createTeamRequest.Icon
	}
	if createTeamRequest.Beinvitemode != nil {
		param["beinvitemode"] = strconv.Itoa(*createTeamRequest.Beinvitemode)
	}
	if createTeamRequest.Invitemode != nil {
		param["invitemode"] = strconv.Itoa(*createTeamRequest.Invitemode)
	}
	if createTeamRequest.Uptinfomode != nil {
		param["uptinfomode"] = strconv.Itoa(*createTeamRequest.Uptinfomode)
	}
	if createTeamRequest.Upcustommode != nil {
		param["upcustommode"] = strconv.Itoa(*createTeamRequest.Upcustommode)
	}
	if createTeamRequest.TeamMemberLimit != nil {
		param["teamMemberLimit"] = strconv.Itoa(*createTeamRequest.TeamMemberLimit)
	}
	if createTeamRequest.IsNotifyCloseOnline != nil {
		param["isNotifyCloseOnline"] = strconv.Itoa(*createTeamRequest.IsNotifyCloseOnline)
	}
	if createTeamRequest.IsNotifyClosePersistent != nil {
		param["isNotifyClosePersistent"] = strconv.Itoa(*createTeamRequest.IsNotifyClosePersistent)
	}
	if createTeamRequest.Attach != nil {
		param["attach"] = *createTeamRequest.Attach
	}
	if createTeamRequest.Bid != nil {
		param["bid"] = *createTeamRequest.Bid
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(createTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 主动退出群聊
func (c *ImClient) LeaveTeam(leaveRequest LeaveRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": leaveRequest.Tid, "accid": leaveRequest.Accid}

	if leaveRequest.Attach != nil {
		param["attach"] = *leaveRequest.Attach
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(leaveTam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 修改群信息
func (c *ImClient) UpdateTeam(updateTeamRequest UpdateTeamRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": updateTeamRequest.Tid, "owner": updateTeamRequest.Owner}
	if updateTeamRequest.Tname != nil {
		param["tname"] = *updateTeamRequest.Tname
	}
	if updateTeamRequest.Announcement != nil {
		param["announcement"] = *updateTeamRequest.Announcement
	}
	if updateTeamRequest.Intro != nil {
		param["intro"] = *updateTeamRequest.Intro
	}
	if updateTeamRequest.Icon != nil {
		param["icon"] = *updateTeamRequest.Icon
	}
	if updateTeamRequest.Joinmode != nil {
		param["joinmode"] = strconv.Itoa(*updateTeamRequest.Joinmode)
	}
	if updateTeamRequest.Beinvitemode != nil {
		param["beinvitemode"] = strconv.Itoa(*updateTeamRequest.Beinvitemode)
	}
	if updateTeamRequest.Invitemode != nil {
		param["invitemode"] = strconv.Itoa(*updateTeamRequest.Invitemode)
	}
	if updateTeamRequest.Uptinfomode != nil {
		param["uptinfomode"] = strconv.Itoa(*updateTeamRequest.Uptinfomode)
	}
	if updateTeamRequest.Upcustommode != nil {
		param["upcustommode"] = strconv.Itoa(*updateTeamRequest.Upcustommode)
	}
	if updateTeamRequest.TeamMemberLimit != nil {
		param["teamMemberLimit"] = strconv.Itoa(*updateTeamRequest.TeamMemberLimit)
	}
	if updateTeamRequest.Custom != nil {
		param["custom"] = *updateTeamRequest.Custom
	}
	if updateTeamRequest.Attach != nil {
		param["attach"] = *updateTeamRequest.Attach
	}
	if updateTeamRequest.Bid != nil {
		param["bid"] = *updateTeamRequest.Bid
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateTeam)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 修改群成员昵称
func (c *ImClient) UpdateTeamNick(updateTeamNickRequest UpdateTeamNickRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": updateTeamNickRequest.Tid, "owner": updateTeamNickRequest.Owner, "accid": updateTeamNickRequest.Accid}

	if updateTeamNickRequest.Nick != nil {
		param["nick"] = *updateTeamNickRequest.Nick
	}

	if updateTeamNickRequest.Custom != nil {
		param["custom"] = *updateTeamNickRequest.Custom
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(updateTeamNick)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}

// 获取群详细信息
func (c *ImClient) QueryDetail(queryDetailRequest QueryDetailRequest) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": queryDetailRequest.Tid}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(queryDetail)
	if err != nil {
		return nil, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return nil, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return nil, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return nil, errors.New(msg)
	}

	return jsonRes, nil
}
