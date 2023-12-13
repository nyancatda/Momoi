/*
 * @Author: NyanCatda
 * @Date: 2023-12-13 15:19:26
 * @LastEditTime: 2023-12-13 15:34:23
 * @LastEditors: NyanCatda
 * @Description: 伪造UserAgent
 * @FilePath: \Momoi\internal\FakeParameters\UserAgent.go
 */
package FakeParameters

import (
	"strconv"
	"time"

	"github.com/nyancatda/Momoi/v2/tools/Random"
)

/**
 * @description: 随机生成UserAgent
 * @return {string} UserAgent
 */
func RandomUserAgent() string {
	// 随机操作系统
	Platform := Random.Array([]string{"Macintosh", "Windows", "X11"})
	var OS string
	switch Platform {
	case "Macintosh":
		OS = Random.Array([]string{"68K", "PPC", "Intel Mac OS X"})
	case "Windows":
		OS = Random.Array([]string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"})
	case "X11":
		OS = Random.Array([]string{"Linux i686", "Linux x86_64"})
	}

	// 随机浏览器
	Browser := Random.Array([]string{"chrome", "firefox", "ie"})
	switch Browser {
	case "chrome":
		WebkitVersion := strconv.Itoa(Random.Intn(599, 500))
		Version := strconv.Itoa(Random.Intn(99, 0)) + ".0" + strconv.Itoa(Random.Intn(9999, 0)) + "." + strconv.Itoa(Random.Intn(999, 0))
		return "Mozilla/5.0 (" + OS + ") AppleWebKit/" + WebkitVersion + ".0 (KHTML, like Gecko) Chrome/" + Version + " Safari/" + WebkitVersion
	case "firefox":
		NowYear := time.Now().Year()
		Year := strconv.Itoa(Random.Intn(NowYear, 2020))

		RandomMonth := Random.Intn(12, 1)
		var Month string
		if RandomMonth < 10 {
			Month = "0" + strconv.Itoa(RandomMonth)
		} else {
			Month = strconv.Itoa(RandomMonth)
		}

		RandomDay := Random.Intn(28, 1)
		var Day string
		if RandomDay < 10 {
			Day = "0" + strconv.Itoa(RandomDay)
		} else {
			Day = strconv.Itoa(RandomDay)
		}

		GeckoVersion := Year + Month + Day
		Version := strconv.Itoa(Random.Intn(72, 1)) + ".0"
		return "Mozilla/5.0 (" + OS + "; rv:" + Version + ") Gecko/" + GeckoVersion + " Firefox/" + Version
	case "ie":
		Version := strconv.Itoa(Random.Intn(99, 1)) + ".0"
		EngineVersion := strconv.Itoa(Random.Intn(99, 1)) + ".0"
		Option := Random.Array([]bool{true, false})
		var Token string
		if Option {
			Token = Random.Array([]string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}) + "; "
		}

		return "Mozilla/5.0 (compatible; MSIE " + Version + "; " + OS + "; " + Token + "Trident/" + EngineVersion + ")"
	}

	return ""
}
