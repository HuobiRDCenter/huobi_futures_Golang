package reqbuilder

import (
	"fmt"
	"net/url"
	"strings"
	"time"
)

type PrivateUrlBuilder struct {
	host    string
	akKey   string
	akValue string
	smKey   string
	smValue string
	svKey   string
	svValue string
	tKey    string

	signer *Signer
}

func (p *PrivateUrlBuilder) Init(accessKey string, secretKey string, host string) *PrivateUrlBuilder {
	p.akKey = "AccessKeyId"
	p.akValue = accessKey
	p.smKey = "SignatureMethod"
	p.smValue = "HmacSHA256"
	p.svKey = "SignatureVersion"
	p.svValue = "2"
	p.tKey = "Timestamp"

	p.host = host
	p.signer = new(Signer).Init(secretKey)

	return p
}

func (p *PrivateUrlBuilder) Build(method string, path string, request *GetRequest) string {
	time := time.Now().UTC()

	return p.BuildWithTime(method, path, time, request)
}

// BuildWithQuery 与 Build 相同，但额外把调用方已拼好的 query 串（形如 "?a=1&b=2"，
// 也容忍前导 "&"）一并纳入签名与最终 URL。
//
// 存在的理由：V5 的私有 GET 接口要求业务参数参与签名。若先 Build 签名、再把参数
// 拼到 URL 尾部，服务端会返回 401 Verification failure（已用生产 API 实测确认）。
// 本方法让存量代码保持"先拼 option 串"的写法不变，只把收尾的 `url += option` 换成本方法。
func (p *PrivateUrlBuilder) BuildWithQuery(method string, path string, request *GetRequest, extraQuery string) string {
	time := time.Now().UTC()

	return p.BuildWithQueryAndTime(method, path, time, request, extraQuery)
}

// BuildWithQueryAndTime 是 BuildWithTime 的带附加 query 版本，extraQuery 传空即与原行为一致。
func (p *PrivateUrlBuilder) BuildWithQueryAndTime(method string, path string, utcDate time.Time, request *GetRequest, extraQuery string) string {
	extra := strings.TrimPrefix(extraQuery, "?")
	if extra != "" {
		if values, err := url.ParseQuery(extra); err == nil {
			merged := new(GetRequest).InitFrom(request)
			for key, vals := range values {
				for _, val := range vals {
					merged.AddParam(key, val)
				}
			}
			request = merged
		}
	}

	time := utcDate.Format("2006-01-02T15:04:05")

	req := new(GetRequest).InitFrom(request)
	req.AddParam(p.akKey, p.akValue)
	req.AddParam(p.smKey, p.smValue)
	req.AddParam(p.svKey, p.svValue)
	req.AddParam(p.tKey, time)

	parameters := req.BuildParams()

	signature := p.signer.Sign(method, p.host, path, parameters)

	url := fmt.Sprintf("https://%s%s?%s&Signature=%s", p.host, path, parameters, url.QueryEscape(signature))

	return url
}

func (p *PrivateUrlBuilder) BuildWithTime(method string, path string, utcDate time.Time, request *GetRequest) string {
	return p.BuildWithQueryAndTime(method, path, utcDate, request, "")
}
