package middleware

import (
	"fmt"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
	"net/url"
	"time"
)

func ReverseProxy(target *url.URL) *httputil.ReverseProxy {
	fmt.Printf("这个函数需要远程调用,我们接下来输出这个目标路径: %v\n", target)
	// return httputil.NewSingleHostReverseProxy(target)
	defaultRoundTripper := getRoundTripper()
	director := func(req *http.Request) {
		req.URL.Scheme = target.Scheme
		req.URL.Host = target.Host
		req.Host = target.Host
	}
	return &httputil.ReverseProxy{
		Director: director,
		Transport: &transport{
			defaultRoundTripper,
		},
	}
}

func getRoundTripper() http.RoundTripper {
	log.Printf("接下来要进行远程调用,我们需要设置一个默认的RoundTripper")
	var defaultTransport http.RoundTripper = &http.Transport{
		// 设置Proxy
		Proxy: proxyFunc, //用于设置代理的请求方式
		DialContext: (&net.Dialer{
			Timeout:   30 * time.Second,
			KeepAlive: 30 * time.Second,
		}).DialContext,
		MaxIdleConns:          100,
		IdleConnTimeout:       90 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second, //在使用https的时候,定义TLS握手的时间
		ExpectContinueTimeout: 1 * time.Second,
		Dial:                  dialFunc,
	}
	return defaultTransport
}

func dialFunc(network, addr string) (net.Conn, error) {
	conn, err := (&net.Dialer{
		Timeout:   30 * time.Second,
		KeepAlive: 30 * time.Second,
	}).Dial(network, addr)
	if err != nil {
		log.Printf("failed to dial with err(%s)", err.Error())
		return nil, err
	}
	return conn, nil
}
func proxyFunc(req *http.Request) (*url.URL, error) {
	log.Printf("calling proxy,URL:%s,Method:%s", req.URL.String(), req.Method)
	return http.ProxyFromEnvironment(req)
}

type transport struct {
	http.RoundTripper
}
