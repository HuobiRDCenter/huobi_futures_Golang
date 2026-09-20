// Package v5test 是 linear-swap（U 本位合约）V5 Open API 的实测用例集，
// 覆盖 43 个 REST 接口与 12 个 WS 接口（7 推送 + 5 交易指令）。
//
// 判定口径：只要请求发出了、拿到了可解析的应答、且应答码不是
// 鉴权失败（401/403）或路由错误（404/405），即记为 PASS。
// 业务类错误（余额不足、订单不存在、价格超限……）同样算 PASS——
// 它证明路径、HTTP 方法、签名口径都对，接口确实存在。
//
// 密钥优先取环境变量（便于在仓库外注入，避免提交密钥），
// 未设置时才回落到 config.AccessKey / config.SecretKey。
//
// 注意：WS 用例每条各自建一次连接（与仓库既有 ws 测试写法一致），一轮共 12 次建连。
// 线上对单 IP 的 WS 建连频率有限制，连续跑第二轮时后续几条会直接
// "read: connection reset by peer"（连 TCP 都没建起来），表现为若干条 WS 用例超时失败。
// 已实测确认：间隔 5 分钟以上再跑即全绿，与 SDK 代码无关。要连续压测请改为复用一条连接。
package v5test

import (
	"fmt"
	"os"
	"sort"

	"github.com/HuobiRDCenter/huobi_futures_Golang/config"
)

var (
	passCount int
	failCount int
	skipped   []string
	failures  []string
)

// accessKey 取环境变量，未设置则回落到仓库 config
func accessKey() string {
	if v := os.Getenv("HTX_ACCESS_KEY"); v != "" {
		return v
	}
	return config.AccessKey
}

// secretKey 取环境变量，未设置则回落到仓库 config
func secretKey() string {
	if v := os.Getenv("HTX_SECRET_KEY"); v != "" {
		return v
	}
	return config.SecretKey
}

// host 取环境变量，未设置则回落到仓库 config
func host() string {
	if v := os.Getenv("HTX_HOST"); v != "" {
		return v
	}
	return config.Host
}

// ok 记一次通过
func ok(name string, code int64, message string) {
	passCount++
	fmt.Printf("  PASS  %-42s code=%-6d %s\n", name, code, truncate(message, 60))
}

// bad 记一次失败
func bad(name string, format string, args ...interface{}) {
	failCount++
	detail := fmt.Sprintf(format, args...)
	failures = append(failures, fmt.Sprintf("%s: %s", name, detail))
	fmt.Printf("  FAIL  %-42s %s\n", name, detail)
}

// skip 记一次跳过
func skip(name string, reason string) {
	skipped = append(skipped, fmt.Sprintf("%s: %s", name, reason))
	fmt.Printf("  SKIP  %-42s %s\n", name, reason)
}

// isRouteOrAuthError V5 的 401/403/404/405 说明路径、方法或签名有问题，不算通过
func isRouteOrAuthError(code int64) bool {
	switch code {
	case 401, 403, 404, 405:
		return true
	}
	return false
}

// check 校验一个 REST 应答。ts 为应答时间戳，用于识别"解析失败导致全零"的假应答。
// 各 DTO 的 Code 字段有 int / int32 / int64 多种写法，调用方统一用 int64(…) 收口。
func check(name string, code int64, ts int64, message string) {
	if code == 0 && ts == 0 {
		bad(name, "应答无法解析（code=0 且 ts=0），原始报文可能不是 JSON")
		return
	}
	if isRouteOrAuthError(code) {
		bad(name, "路由/鉴权失败 code=%d message=%s", code, message)
		return
	}
	ok(name, code, message)
}

// truncate 截断过长文本，避免刷屏
func truncate(s string, n int) string {
	if len(s) <= n {
		return s
	}
	return s[:n] + "..."
}

// Section 打印小节标题
func Section(title string) {
	fmt.Printf("\n===== %s =====\n", title)
}

// Report 打印小结
func Report() {
	fmt.Printf("\n===== V5 实测小结 =====\n")
	fmt.Printf("PASS: %d   FAIL: %d   SKIP: %d\n", passCount, failCount, len(skipped))
	if len(failures) > 0 {
		sort.Strings(failures)
		fmt.Printf("\n失败明细：\n")
		for _, f := range failures {
			fmt.Printf("  - %s\n", f)
		}
	}
	if len(skipped) > 0 {
		sort.Strings(skipped)
		fmt.Printf("\n跳过明细：\n")
		for _, s := range skipped {
			fmt.Printf("  - %s\n", s)
		}
	}
}

// RunAllExamples 跑全部 V5 实测（REST 43 + WS 12）
func RunAllExamples() {
	if accessKey() == "" || secretKey() == "" {
		fmt.Println("V5 实测需要密钥：请设置 HTX_ACCESS_KEY / HTX_SECRET_KEY 环境变量，" +
			"或在 config/key.go 中填写 AccessKey / SecretKey")
		return
	}

	RunRestAccountMarket()
	RunRestTradePositionAlgo()
	RunWs()

	Report()
}
