package examples

import (
	"fmt"
	"github.com/xu42/youzan-sdk-go/auth"
)

func testToken() {
	// 获取自用型AccessToken
	result, err := auth.GenSelfToken("CLIENT_ID", "CLIENT_SECRET", 110)
	fmt.Println(result, result.AccessToken, err)
}
