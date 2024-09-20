package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	addFriend = neteaseBaseURL + "/friend/add.action"
)

/**添加好友
 * accid 发起者accid
 * faccid 接受者accid
 * friend_type 	1 直接加好友（无需对方同意）；2 请求加好友（需要对方同意）；3 同意加好友；4 拒绝加好友
 * msg 加好友对应的请求消息
 * serverex	 服务器端扩展字段 此字段 client 端只读，server 端读写
**/
func (c *ImClient) AddFriend(accid, faccid string, friend_type int, msg, serverex *string) (map[string]*json.RawMessage, error) {
	param := map[string]string{"accid": accid}
	param["faccid"] = faccid
	param["type"] = strconv.Itoa(friend_type)
	if msg != nil {
		param["msg"] = *msg
	}
	if serverex != nil {
		param["serverex"] = *serverex
	}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(addFriend)
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