package slackuserid

import (
	//	"fmt"

	"github.com/nlopes/slack"
)

func SlackIDtoUsername(slack_token string, userid string) string {
	api := slack.New(slack_token)
	//user, err := api.GetUserInfo(userid)
	user, _ := api.GetUserInfo(userid)

	/*if err != nil {
		fmt.Printf("%s\n", err)
		return
	}*/
	//fmt.Printf("ID: %s, Fullname: %s, Email: %s\n", user.ID, user.Profile.RealName, user.Profile.Email)
	return user.Profile.RealName
}
