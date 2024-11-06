package netease

import (
	"encoding/json"
	"errors"
	"strconv"

	jsoniter "github.com/json-iterator/go"
)

const (
	UpdateImUserPoint = neteaseBaseURL + "/user/updateUinfo.action"
	GetImUserPoint    = neteaseBaseURL + "/user/getUinfos.action"
)

func (c *ImClient) GetImUser(accIds []string) (map[string]*json.RawMessage, error) {
	accidsStr, _ := json.Marshal(accIds)

	param := map[string]string{"accids": string(accidsStr)}
	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(GetImUserPoint)
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

	res := map[string]*json.RawMessage{}
	err = jsoniter.Unmarshal(*jsonRes["uinfos"], &res)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (c *ImClient) UpdateImUser(u *ImUser) (int, error) {
	param := map[string]string{"accid": u.ID}

	if len(u.Name) > 0 {
		param["name"] = u.Name
	}
	if len(u.IconURL) > 0 {
		param["icon"] = u.IconURL
	}
	if len(u.Sign) > 0 {
		param["sign"] = u.Sign
	}
	if len(u.Email) > 0 {
		param["email"] = u.Email
	}
	if len(u.Birthday) > 0 {
		param["birth"] = u.Birthday
	}
	if len(u.Mobile) > 0 {
		param["mobile"] = u.Mobile
	}
	if u.Gender == 1 || u.Gender == 2 {
		param["gender"] = strconv.Itoa(u.Gender)
	}
	if len(u.Extension) > 0 {
		param["ex"] = u.Extension
	}

	client := c.client.R()
	c.setCommonHead(client)
	client.SetFormData(param)

	resp, err := client.Post(UpdateImUserPoint)
	if err != nil {
		return 0, err
	}

	var jsonRes map[string]*json.RawMessage
	err = jsoniter.Unmarshal(resp.Body(), &jsonRes)
	if err != nil {
		return 0, err
	}

	var code int
	err = json.Unmarshal(*jsonRes["code"], &code)
	if err != nil {
		return code, err
	}

	if code != 200 {
		var msg string
		json.Unmarshal(*jsonRes["desc"], &msg)
		return code, errors.New(msg)
	}

	return code, nil
}
