package twitch

import (
	bus "github.com/autonomouskoi/core-tinygo"
)

func ListProfiles(req *ListProfilesRequest) (*ListProfilesResponse, error) {
	msg := &bus.BusMessage{
		Topic: BusTopics_TWITCH_REQUEST.String(),
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_LIST_PROFILES_REQ),
	}
	msg.Message, _ = req.MarshalVT()
	reply, err := bus.WaitForReply(msg, 1000)
	if err != nil {
		return nil, err
	}
	resp := &ListProfilesResponse{}
	err = resp.UnmarshalVT(reply.GetMessage())
	return resp, err
}

func SendChat(req *TwitchChatRequestSendRequest) error {
	msg := &bus.BusMessage{
		Topic: BusTopics_TWITCH_CHAT_REQUEST.String(),
		Type:  int32(MessageTypeTwitchChatRequest_TWITCH_CHAT_REQUEST_TYPE_SEND_REQ),
	}
	var err error
	msg.Message, err = req.MarshalVT()
	if err != nil {
		return err
	}
	return bus.Send(msg)
}
