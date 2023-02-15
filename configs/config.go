package configs

import (
	"fmt"

	"github.com/kelseyhightower/envconfig"
)

// Package constants definition.
const (
	Prefix = ""
)

// AppConfig app config
var AppConfig App

// New returns a new instance of App configuration.
func New() (*App, error) {
	if err := envconfig.Process(Prefix, &AppConfig); err != nil {
		return nil, err
	}
	return &AppConfig, nil
}

// App represents all configuration of qr-service application.
type App struct {
	Host                       string `default:"0.0.0.0" envconfig:"HOST"`
	Port                       int    `default:"8080" envconfig:"PORT"`
	RunMode                    string `default:"debug" envconfig:"RUN_MODE"`
	CDN                        string `default:"cdn.vinid.net" envconfig:"CDN"`
	DB                         MySQL
	Redis                      Redis
	Kafka                      Kafka
	Keys                       Keys `envconfig:"KEYS"`
	BasicAuth                  BasicAuthAccounts
	Env                        string `default:"DEV" envconfig:"ENV"`
	ServiceName                string `envconfig:"SERVICE_NAME"`
	ProjectID                  string `envconfig:"PROJECT_ID"`
	ConsumerWorkerNumber       int    `default:"10" envconfig:"CONSUMER_WORKER_NUMBER"`
	Slack                      Slack
	Gamification               Gamification
	User                       User
	CacheTimeout               int `default:"3600" envconfig:"CACHE_TIMEOUT"`
	Notification               Notification
	RedLock                    RedLock
	Transaction                Transaction
	GRPC                       GRPC
	AesUtil                    AesUtil
	VinShop                    VinShop
	UploadedFileExpirationTime uint32 `default:"259200" envconfig:"UPLOADED_FILE_EXPIRATION_TIME"`
	MakerChecker               MakerChecker
	AdminService               AdminService
	Auth0Service               Auth0Service
	MerchantInformation        MerchantInformation
	NewRelicConfig             NewRelicConfig
}

type NewRelicConfig struct {
	AppName string `default:"gami-default" envconfig:"NEW_RELIC_APP_NAME"`
	License string `default:"" envconfig:"NEW_RELIC_LICENSE"`
}

type MerchantInformation struct {
	BaseURL   string `default:"https://api-merchant-qc.int.vinid.dev" envconfig:"MERCHANT_BASE_URL"`
	Endpoint  string `default:"/ops/v1/internal/merchant/info" envconfig:"MERCHANT_ENDPOINT"`
	BasicAuth string `default:"" envconfig:"MERCHANT_TOKEN"`
}

// AddressListener returns address listener of HTTP server.
func (c *App) AddressListener() string {
	return fmt.Sprintf("%v:%v", c.Host, c.Port)
}

// User represents merchant service configuration.
type User struct {
	OauthEndpoint     string `default:"https://oauth.vinid.dev/oauth2" envconfig:"OAUTH_ENDPOINT"`
	AccountEndpoint   string `default:"https://api-dev.int.vinid.dev/account/v1" envconfig:"USER_ENDPOINT"`
	ProfileEndpoint   string `default:"https:/api-dev.int.vinid.dev/user-profile/v1" envconfig:"PROFILE_ENDPOINT"`
	UserScope         string `default:"internal.profile" envconfig:"USER_SCOPE"`
	NotificationScope string `default:"notification.internal" envconfig:"NOTIFICATION_SCOPE"`
	IntegrationHeader struct {
		Authorization string `default:"" envconfig:"USER_AUTHORIZATION"`
		GrantType     string `default:"client_credentials" envconfig:"USER_GRANT_TYPE"`
	} `envconfig:"INTEGRATION_HEADER"`
}

// Keys contain jwt authentication information
type Keys struct {
	App struct {
		PublicKeyOPS    string `default:"./keys/app/ops-key" envconfig:"OPS_PUBLIC_KEY"`
		PublicKeyMobile string `default:"./keys/app/mobile-key" envconfig:"MOBILE_PUBLIC_KEY"`
		VinShopKey      string `default:"./keys/app/vinshop-key" envconfig:"VINSHOP_PUBLIC_KEY"`
	} `envconfig:"APP"`
}

// BasicAuthAccounts contains accounts for basic authenticating.
type BasicAuthAccounts struct {
	Internal struct {
		Username string `default:"" envconfig:"AUTH_USERNAME"`
		Password string `default:"" envconfig:"AUTH_PASSWORD"`
	}
}

// Notification represents merchant service configuration.
type Notification struct {
	Endpoint   string `default:"" envconfig:"NOTI_ENDPOINT"`
	EndpointV2 string `default:"" envconfig:"NOTIV2_ENDPOINT"`
	ApiKey     string `default:"" envconfig:"NOTI_API_KEY"`
	AppID      string `default:"" envconfig:"APP_ID"`
	TemplateID string `default:"" envconfig:"TEMPLATE_ID"`
	GroupID    int    `default:"2" envconfig:"NOTI_GROUP"`
	ChannelID  string `default:"1" envconfig:"NOTI_CHANNEL"`
}

// Gamification contain jwt authentication information
type Gamification struct {
	Endpoint          string `default:"" envconfig:"GAMIFICATION_ENDPOINT"`
	IntegrationHeader struct {
		Authorization string `default:"" envconfig:"GAMIFICATION_AUTHORIZATION"`
	}
	RedeemTopic string `default:"" envconfig:"GAMI_REDEEM_TOPIC"`
}

type Slack struct {
	Endpoint string `default:"https://hooks.slack.com/services" envconfig:"SLACK_ENDPOINT"`
	Channel  string `default:"TBFDUP13L/BNYFR0UTS" envconfig:"SLACK_CHANNEL"`
	Token    string `default:"" envconfig:"SLACK_TOKEN"`
}

type Transaction struct {
	Endpoint      string `default:"https://api-int.vinid.dev/internal/transaction-history/v1" envconfig:"TRANSACTION_ENDPOINT"`
	Authorization string `default:"" envconfig:"TRANSACTION_AUTHORIZATION"`
}

type AesUtil struct {
	Key   string `default:"" envconfig:"AES_KEY"`
	Nonce string `default:"" envconfig:"AES_NONCE"`
	IV    string `default:"" envconfig:"AES_IV"`
}

type VinShop struct {
	NtfAuth         string `default:"" envconfig:"VINSHOP_NTF_AUTH"` // base64 from internaldev:2f3asfrhqAGT
	NtfBaseURL      string `default:"https://api-merchant-qc.int.vinid.dev/notification-service/v2" envconfig:"VINSHOP_NTF_URL"`
	NtfSendEndpoint string `default:"/notifications" envconfig:"VINSHOP_NTF_SEND_ENDPOINT"`
	NtfTopicPrefix  string `default:"merchantCode_" envconfig:"VINSHOP_NTF_TOPIC_PREFIX"`
	NtfType         string `default:"PROMOTION" envconfig:"VINSHOP_NTF_TYPE"`
}

// MakerChecker represents struct config.
type MakerChecker struct {
	Endpoint      string `envconfig:"MAKER_CHECKER_ENDPOINT"`
	Authorization string `default:"" envconfig:"MAKER_CHECKER_AUTHOR"`
	Username      string `default:"maker_checker" envconfig:"MAKER_CHECKER_USERNAME"`
	Key           string `default:"" envconfig:"MAKER_KEY"`
}

type AdminService struct {
	BaseURL   string `default:"https://api-merchant-qc.int.vinid.dev/tcb-int/v1" envconfig:"ADMIN_SERVICE_URL"`
	Endpoint  string `default:"/token/info" envconfig:"ADMIN_SERVICE_ENDPOINT"`
	BasicAuth string `envconfig:"ADMIN_SERVICE_BASIC_AUTH"`
}

type Auth0Service struct {
	Audience string `default:"https://api-qc.vinid.dev/level-1" envconfig:"AUTH0_AUDIENCE"`
	Domain   string `default:"auth-dev.onemount.dev" envconfig:"AUTH0_DOMAIN"`
	JwksUri  string `default:"https://auth-dev.onemount.dev/.well-known/jwks.json" envconfig:"AUTH0_JWKS_URI"`
}
