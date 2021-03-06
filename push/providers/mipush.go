package providers

import (
	"crypto/rand"
	"fmt"
	"io"
	"math/big"
	"net/http"
	"net/url"
	"notify-api/push"
	"os"
	"strconv"
	"strings"
	"time"
)

const APIURL = "https://api.xmpush.xiaomi.com/v2/message/user_account"

type MiPushProvider struct {
	MiPushSecret string
	MiPushClient *http.Client
}

func (p *MiPushProvider) Init() error {
	p.MiPushClient = &http.Client{}
	return nil
}

func (p *MiPushProvider) Send(msg *push.Message) error {
	n, _ := rand.Int(rand.Reader, big.NewInt(1000000))
	notifyID := n.Int64()

	// Build MIPush request
	title := msg.Title
	content := msg.Content
	long := msg.Long
	msgID := msg.ID

	intentUriFormat := "intent:#Intent;launchFlags=0x14000000;component=top.learningman.push/.TranslucentActivity;S.user_id=%s;S.long=%s;S.msg_id=%s;S.title=%s;S.created_at=%s;S.content=%s;end"
	intentUri := fmt.Sprintf(intentUriFormat,
		url.PathEscape(msg.UserID),
		url.PathEscape(long),
		url.PathEscape(msgID),
		url.PathEscape(title),
		url.PathEscape(msg.CreatedAt.Format(time.RFC3339)),
		url.PathEscape(content))

	postData := url.Values{
		"user_account":            {msg.UserID},
		"payload":                 {long},
		"restricted_package_name": {"top.learningman.push"},
		"pass_through":            {"0"},
		"title":                   {title},
		"description":             {content},
		"notify_id":               {strconv.Itoa(int(notifyID))},
		"extra.id":                {msgID},
		"extra.notify_effect":     {"2"}, // https://dev.mi.com/console/doc/detail?pId=1278#_3_2
		"extra.intent_uri":        {intentUri},
	}.Encode()

	req, err := http.NewRequest(
		http.MethodPost,
		APIURL,
		strings.NewReader(postData))

	req.Header.Set("Authorization", fmt.Sprintf("key=%s", p.MiPushSecret))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	resp, err := p.MiPushClient.Do(req)

	if err != nil {
		return err
	}

	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(resp.Body)

	return nil
}

func (p *MiPushProvider) ChannelName() string {
	return "MiPush"
}

func (p *MiPushProvider) Check() error {
	MiPushSecret := os.Getenv("MiPushSecret")
	if MiPushSecret == "" {
		return fmt.Errorf("MiPushSecret is not set")
	} else {
		p.MiPushSecret = MiPushSecret
		return nil
	}
}
