package twitch

import bus "github.com/autonomouskoi/core-tinygo"

func GetUser(req *GetUserRequest) (*GetUserResponse, error) {
	msg := &bus.BusMessage{
		Topic: BusTopics_TWITCH_REQUEST.String(),
		Type:  int32(MessageTypeRequest_TYPE_REQUEST_GET_USER_REQ),
	}
	if bus.MarshalMessage(msg, req); msg.Error != nil {
		return nil, msg.Error
	}

	reply, err := bus.WaitForReply(msg, 5000)
	if err != nil {
		return nil, err
	}
	if reply.Error != nil {
		return nil, reply.Error
	}

	resp := &GetUserResponse{}
	if err := bus.UnmarshalMessage(reply, resp); err != nil {
		return nil, err
	}
	return resp, nil
}
