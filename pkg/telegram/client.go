package telegram

import (
	"bufio"
	"context"
	"errors"
	"fmt"
	"github.com/gotd/td/telegram"
	"github.com/gotd/td/telegram/auth"
	"github.com/gotd/td/tg"
	"go.uber.org/zap"
	"golang.org/x/crypto/ssh/terminal"
	"log"
	"os"
	"strings"
)

type Client struct {
	ApiId   int    `json:"api_id"`
	ApiHash string `json:"api_hash"`
	Options telegram.Options
	log     *zap.Logger
}

// New Constructor
func New(apiId int, apiHash string) *Client {
	log, err := zap.NewDevelopment()
	if err != nil {
		panic(err)
	}
	return &Client{
		ApiId:   apiId,
		ApiHash: apiHash,
		log:     log,
	}
}

// newClient get new telegram client instance
func (c *Client) newClient() *telegram.Client {
	client := telegram.NewClient(c.ApiId, c.ApiHash, c.Options)
	return client
}

// AuthSendCode Отправка кода подтверждения на номер
func (c *Client) AuthSendCode(phone string) (*tg.AuthSentCode, error) {
	// https://core.telegram.org/api/obtaining_api_id
	client := c.newClient()
	var response *tg.AuthSentCode
	var err error
	ctx := context.Background()
	if err := client.Run(ctx, func(ctx context.Context) error {
		// Now you can invoke MTProto RPC requests by calling the API.
		api := client.API()
		data := &tg.AuthSendCodeRequest{
			PhoneNumber: phone,
			APIHash:     c.ApiHash,
			APIID:       c.ApiId,
		}

		response, err = api.AuthSendCode(ctx, data)
		if err != nil {
			return err
		}

		log.Println("Telegram AuthSendCode: ", response)

		// Return to close client connection and free up resources.
		return nil
	}); err != nil {
		return nil, err
	}
	return response, nil
}

// SingIn Авторизаия
func (c *Client) SingIn(phone string, phoneHash string, code string) (tg.AuthAuthorizationClass, error) {
	// https://core.telegram.org/api/obtaining_api_id
	//client := c.newClient()
	//var (
	//	response tg.AuthAuthorizationClass
	//	err      error
	//)
	//if err := client.Run(context.Background(), func(ctx context.Context) error {
	//	// Now you can invoke MTProto RPC requests by calling the API.
	//	api := client.API()
	//	data := &tg.AuthSignInRequest{
	//		PhoneNumber:   phone,
	//		PhoneCode:     code,
	//		PhoneCodeHash: phoneHash,
	//	}
	//	response, err = api.AuthSignIn(ctx, data)
	//	if err != nil {
	//		return err
	//	}
	//
	//	log.Println("Telegram SingIn: ", response)
	//
	//	// Return to close client connection and free up resources.
	//	return nil
	//}); err != nil {
	//	return nil, err
	//}
	return nil, nil
}

// noSignUp can be embedded to prevent signing up.
type noSignUp struct{}

func (c noSignUp) SignUp(ctx context.Context) (auth.UserInfo, error) {
	return auth.UserInfo{}, errors.New("not implemented")
}

func (c noSignUp) AcceptTermsOfService(ctx context.Context, tos tg.HelpTermsOfService) error {
	return &auth.SignUpRequired{TermsOfService: tos}
}

// termAuth implements authentication via terminal.
type termAuth struct {
	noSignUp
	phone string
}

func (a termAuth) Phone(_ context.Context) (string, error) {
	return a.phone, nil
}

func (a termAuth) Password(_ context.Context) (string, error) {
	fmt.Print("Enter 2FA password: ")
	bytePwd, err := terminal.ReadPassword(0)
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(string(bytePwd)), nil
}

func (a termAuth) Code(_ context.Context, _ *tg.AuthSentCode) (string, error) {
	fmt.Print("Enter code: ")
	code, err := bufio.NewReader(os.Stdin).ReadString('\n')
	if err != nil {
		return "", err
	}
	return strings.TrimSpace(code), nil
}

func (c *Client) getClient() (*telegram.Client, error) {
	log := c.log
	client, err := telegram.ClientFromEnvironment(telegram.Options{
		Logger: log,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}

// Run runs f callback with context and logger, panics on error.
func (c *Client) Run(f func(ctx context.Context, client *telegram.Client) error) error {
	// No graceful shutdown.
	ctx := context.Background()
	client, err := c.getClient()
	if err != nil {
		return err
	}
	return f(ctx, client)
}

//func (c *Client) Run(f func(client *telegram.Client, fn func(ctx context.Context) error) error) error {
//	// No graceful shutdown.
//	ctx := context.Background()
//	client, err := c.getClient()
//	if err != nil {
//		return err
//	}
//	client.Run(ctx,fn(ctx))
//	return f(client, ))
//}

// NewFlow Авторизация с сохранением пользователя в сессию
func (c *Client) NewFlow(phone string) (user *tg.User, err error) {
	err = c.Run(func(ctx context.Context, client *telegram.Client) error {
		// Setting up authentication flow helper based on terminal auth.
		flow := auth.NewFlow(
			termAuth{phone: phone},
			auth.SendCodeOptions{},
		)
		return client.Run(ctx, func(ctx context.Context) error {
			if err := client.Auth().IfNecessary(ctx, flow); err != nil {
				return err
			}
			self, err := client.Self(ctx)
			if err != nil {
				return err
			}
			user = self
			return nil
		})
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (c *Client) LogOut() error {
	if err := c.Run(func(ctx context.Context, client *telegram.Client) error {
		client.Run(ctx, func(ctx context.Context) error {
			user, err := client.Self(ctx)
			if err != nil {
				return err
			}
			//logout, err := client.API().AuthLogOut(ctx)
			//fmt.Println("aaaaaaaaaaaaaaaaaaaaa", logout)
			//fmt.Println("bbbbbbbbbbbbbbbbbbb", err)
			fmt.Println("uuuuuuuuuuuuuuuuuu", user)

			return nil
		})
		return nil
	}); err != nil {
		return err
	}
	return nil
}

func (c *Client) Test() error {
	err := c.Run(func(ctx context.Context, client *telegram.Client) error {
		var excepted []int64
		chats, err := client.API().MessagesGetAllChats(ctx, excepted)
		if err != nil {
			return err
		}
		fmt.Println("aaaaaaaaaaaaaaa", chats)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (c *Client) GetSelf() (*tg.User, error) {
	var (
		user *tg.User
		err  error
	)
	err = c.Run(func(ctx context.Context, client *telegram.Client) error {
		client.Run(ctx, func(ctx context.Context) error {
			user, err = client.Self(ctx)
			if err != nil {
				return err
			}
			return nil
		})
		return nil
	})
	if err != nil {
		return nil, err
	}
	return user, nil
}
