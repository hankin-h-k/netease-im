package netease

import (
	"encoding/json"
	"errors"

	jsoniter "github.com/json-iterator/go"
)

const (
	joinTeam       = neteaseBaseURL + "/team/add.action"           // 加入群聊（高级群）
	leaveTam       = neteaseBaseURL + "team/leave.action"          // 主动退群 （高级群）
	updateTeamNick = neteaseBaseURL + "team/updateTeamNick.action" // 修改群成员昵称 （高级群）
	removeTeam     = neteaseBaseURL + "team/remove.action"         // 解散群 （高级群）
	kickTeam       = neteaseBaseURL + "team/kick.action"           // 踢人出群 （高级群）

)

func (c *ImClient) JoinTeam(tid, owner, members, msg string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "owner": owner, "members": members, "msg": msg}
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

func (c *ImClient) LeaveTeam(tid, accid string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "accid": accid}
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

func (c *ImClient) UpdateTeamNick(tid, owner, accid, nick string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "owner": owner, "accid": accid, "nick": nick}
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

func (c *ImClient) RemoveTeam(tid, owner string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "owner": owner}
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

func (c *ImClient) KickBatchTeam(tid, owner, members, attach string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "owner": owner, "members": members, "attach": attach}
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

func (c *ImClient) KickTeam(tid, owner, member, attach string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"tid": tid, "owner": owner, "member": member, "attach": attach}
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
