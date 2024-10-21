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
	// 后面还要判断https的情况,在这里暂时不进行处理
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

// 接下来实现RoundTripper接口的RoundTrip方法,为什么要实现这个方法呢?
// 因为我们需要在请求的时候,对请求进行一些处理,比如设置请求头,设置请求体等等
// 这个方法是在请求的时候被调用的
// 结果测试,我们发现这个RoundTrip方法会优先于我们的代理方法被调用,所以我们可以在这个方法里面对请求进行一些处理
// 但是为什么会优先于代理方法被调用呢?这个是因为我们在创建ReverseProxy的时候,我们设置了Transport属性,这个属性是一个RoundTripper接口,所以在请求的时候,会优先调用这个方法
func (t *transport) RoundTrip(req *http.Request) (*http.Response, error) {
	log.Printf("calling RoundTrip,URL:%s,Method:%s", req.URL.String(), req.Method)
	// 在这里我们可以对请求进行一些处理,比如设置请求头,设置请求体等等
	resp, err := t.RoundTripper.RoundTrip(req)
	if err != nil {
		log.Printf("failed to RoundTrip with err(%s)", err.Error())
		return nil, err
	}
	log.Printf("输出请求的响应: %v", resp)

	return resp, nil
}
