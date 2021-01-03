package wechat

// -- Basic Common API --

// BasicAccessToken BasicAccessToken
func (c *Client) BasicAccessToken(accessToken IAccessToken) *BasicAccessToken {
	return NewBasicAccessToken(c, accessToken)
}

// BasicMessage BasicMessage
func (c *Client) BasicMessage(accessToken IAccessToken, message IBasicMessage) *BasicMessage {
	return NewBasicMessage(c, accessToken, message)
}

// -- OfficeAccount API --

// MiniProgramAccessToken Miniprogram Auth
func (c *Client) OfficeAccountAccessToken() *OfficeAccountAccessToken {
	return NewOfficeAccountAccessToken(c)
}

// -- Miniprogram API --

// MiniProgramAuth Miniprogram Auth
func (c *Client) MiniProgramAuth() *MiniProgramAuth {
	return NewMiniProgramAuth(c)
}

// MiniProgramAccessToken Miniprogram Auth
func (c *Client) MiniProgramAccessToken() *MiniProgramAccessToken {
	return NewMiniProgramAccessToken(c)
}

// MiniProgramPaid Miniprogram Auth
func (c *Client) MiniProgramPaid() *MiniProgramPaid {
	return NewMiniProgramPaid(c)
}

// MiniProgramActivityMessageCreate MiniProgramActivityMessageCreate
func (c *Client) MiniProgramActivityMessageCreate() *MiniProgramActivityMessageCreate {
	return NewMiniProgramActivityMessageCreate(c)
}

// MiniProgramActivityMessageUpdate MiniProgramActivityMessageUpdate
func (c *Client) MiniProgramActivityMessageUpdate() *MiniProgramActivityMessageUpdate {
	return NewMiniProgramActivityMessageUpdate(c)
}

// MiniProgramAppCodeGet MiniProgramAppCodeGet
func (c *Client) MiniProgramAppCodeGet() *MiniProgramAppCodeGet {
	return NewMiniProgramAppCodeGet(c)
}

// MiniProgramAppCodeGetUnlimit MiniProgramAppCodeGetUnlimit
func (c *Client) MiniProgramAppCodeGetUnlimit() *MiniProgramAppCodeGetUnlimit {
	return NewMiniProgramAppCodeGetUnlimit(c)
}

// MiniProgramAppCodeCreate MiniProgramAppCodeCreate
func (c *Client) MiniProgramAppCodeCreate() *MiniProgramAppCodeCreate {
	return NewMiniProgramAppCodeCreate(c)
}

// MiniProgramSecImg MiniProgramSecImg
func (c *Client) MiniProgramSecImg() *MiniProgramSecImg {
	return NewMiniProgramSecImg(c)
}

// MiniProgramSecMsg MiniProgramSecMsg
func (c *Client) MiniProgramSecMsg() *MiniProgramSecMsg {
	return NewMiniProgramSecMsg(c)
}

// -- Work API --

// WorkAccessToken WorkAccessToken
func (c *Client) WorkAccessToken() *WorkAccessToken {
	return NewWorkAccessToken(c)
}
