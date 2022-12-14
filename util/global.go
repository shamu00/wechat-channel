package util

import (
	"context"
	"log"
	"math/rand"
	"os"
	"time"
)

var (
	GlobalConfigFetcher   IConfigurationFetcher
	GlobalChatGptClient   IChatClient
	GlobalWechatApiId     string
	GlobalWechatApiSecret string
)

const (
	defaultEndpoint = "https://config-center.azconfig.io"
)

func MustInitGlobal() {
	rand.Seed(time.Now().UnixNano())
	var credential = os.Getenv(AzureConfigCenterCredential)
	var secret = os.Getenv(AzureConfigCenterSecret)
	log.Printf("[Info]credential length:%d, secret length:%d", len(credential), len(secret))

	GlobalConfigFetcher = NewAzureConfigurationFetcher(defaultEndpoint, credential, secret)

	chatApiKey, err := GlobalConfigFetcher.GetString(context.Background(), KeyChatGptOpenKey)
	if err != nil {
		log.Fatalf("[Fatal]Couldn't fetch KeyChatOpenApiKey, err:%v", err)
	}
	GlobalWechatApiId, err = GlobalConfigFetcher.GetString(context.Background(), KeyWechatApiId)
	if err != nil {
		log.Fatalf("[Fatal]Couldn't fetch KeyWechatApiId, err:%v", err)
	}
	GlobalWechatApiSecret, err = GlobalConfigFetcher.GetString(context.Background(), KeyWechatApiSecret)
	if err != nil {
		log.Fatalf("[Fatal]Couldn't fetch KeyWechatApiSecret, err:%v", err)
	}

	GlobalChatGptClient = NewChatGptClient(chatApiKey)
}
