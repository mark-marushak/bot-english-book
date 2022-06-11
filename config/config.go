package config

import (
	"fmt"
	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/file"
	"github.com/mark-marushak/bot-english-book/logger"
	"github.com/pkg/errors"
	"log"
	"net/http"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	config             = koanf.New(".")
	parser             = yaml.Parser()
	TokenNotFoundError = errors.New("token not found")
	Token              string
	WG                 *sync.WaitGroup
)

func configFolder(configFile string) string {
	_, b, _, ok := runtime.Caller(0)

	if !ok {
		log.Fatal("[ERR]: configFolder ")
	}

	return fmt.Sprintf("%s/%s", filepath.Dir(b), configFile)
}

func RequestTelegramBot(method string, options map[string][]string) (*http.Response, error) {
	if len(Token) <= 0 {
		return nil, TokenNotFoundError
	}

	return http.PostForm(fmt.Sprintf("https://api.telegram.org/bot%s/%s", Token, method), options)
}

func NewConfig() *koanf.Koanf {
	if err := config.Load(file.Provider(configFolder("config.yml")), parser); err != nil {
		logger.Get().Error("error loading config: %v", err)
	}

	if err := Get().Unmarshal("telegram.bot-api", &Token); err != nil {
		logger.Get().Error("[ERR] Initalization config: %v", err)
	}

	return config
}

func Get() *koanf.Koanf {
	return config
}

type ResponseBody struct {
	OK     bool                   `json:"ok"`
	Result map[string]interface{} `json:"result"`
}

func SetWaitGroup(wg *sync.WaitGroup) *sync.WaitGroup {
	WG = wg
	return WG
}

func GetWaitGroup() *sync.WaitGroup {
	return WG
}

//func IsExceptionError(err error) bool {
//	for _, exception := range Exceptions {
//		result := ComparePercentage(exception, err.Error()) > 90.00
//	}
//}
//
//func ComparePercentage(s1, s2 string) float64 {
//	if len(s1) > len(s2) {
//		s1, s2 = s2, s1
//	}
//
//	var count float64
//	for i := 0; i < len(s1); i++ {
//		if s1[i] == s2[i] {
//			count++
//		}
//	}
//
//	return count / float64(len(s1))
//}
