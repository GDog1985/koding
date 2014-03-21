package followingfeed

import (
	"encoding/json"
	"errors"
	"fmt"
	"koding/tools/logger"
	"socialapi/models"
)

type Action func(*FollowingFeedController, *models.ChannelMessage) error

type FollowingFeedController struct {
	routes map[string]Action
	log    logger.Log
}

var HandlerNotFoundErr = errors.New("Handler Not Found")

func NewFollowingFeedController(log logger.Log) *FollowingFeedController {
	ffc := &FollowingFeedController{
		log: log,
	}

	routes := map[string]Action{
		"channel_message_created": (*FollowingFeedController).MessageSaved,
		"channel_message_update":  (*FollowingFeedController).MessageUpdated,
		"channel_message_deleted": (*FollowingFeedController).MessageDeleted,
	}

	ffc.routes = routes

	return ffc
}

func (f *FollowingFeedController) HandleEvent(event string, data []byte) error {
	f.log.Debug("New Event Recieved %s", event)
	handler, ok := f.routes[event]
	if !ok {
		return HandlerNotFoundErr
	}

	return handler(f, data)
}

func (f *FollowingFeedController) MessageSaved(data *models.ChannelMessage) error {
	a := models.NewAccount()
	a.Id = data.AccountId
	channelIds, err := a.FetchFollowerChannelIds()
	if err != nil {
		return err
	}
	fmt.Println(channelIds)
	return nil
}

func (f *FollowingFeedController) MessageUpdated(data *models.ChannelMessage) error {
	fmt.Println("update", data.InitialChannelId)

	return nil

}

func (f *FollowingFeedController) MessageDeleted(data *models.ChannelMessage) error {
	fmt.Println("delete", data.InitialChannelId)

	return nil
}

func mapMessage(data []byte) (*models.ChannelMessage, error) {
	cm := models.NewChannelMessage()
	if err := json.Unmarshal(data, cm); err != nil {
		return nil, err
	}

	return cm, nil
}

func isEligible(cm *models.ChannelMessage) (bool, error) {
	if cm.InitialChannelId == 0 {
		return false, nil
	}

	if cm.Type != models.ChannelMessage_TYPE_POST {
		return false, nil
	}

	return isChannelEligible(cm)
}

func isChannelEligible(cm *models.ChannelMessage) (bool, error) {
	c, err := fetchChannel(cm.InitialChannelId)
	if err != nil {
		return false, err
	}

	if c.Name != models.Channel_KODING_NAME {
		return false, nil
	}

	return true, nil
}

// todo add caching here
func fetchChannel(channelId int64) (*models.Channel, error) {
	c := models.NewChannel()
	c.Id = channelId
	// todo - fetch only name here
	if err := c.Fetch(); err != nil {
		return nil, err
	}

	return c, nil
}
