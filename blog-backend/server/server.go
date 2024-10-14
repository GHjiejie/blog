package server

import (
	"blog-backend/pkg/db"
	"log"
	"net/http"
)

type BlogServer struct {
	DBEngine   db.Handle
	httpServer *http.Server
}

// Start 函数启动 HTTP 服务器
func NewBlogServer() (*BlogServer, error) {
	// 首先要进行数据库的连接,使用的是
	dbEngine, err := db.NEWHandler()
	if err != nil {
		log.Printf("failed to create db engine handler with err(%s)", err.Error())
		return nil, err
	}
	s := &BlogServer{
		DBEngine: dbEngine,
	}
	if err := s.prepareServer(); err != nil {
		log.Printf("failed to prepare server with err(%s)", err.Error())
		return nil, err
	}
	return s, nil

}

func (s *BlogServer) Start() error {
	log.Printf("BlogServer start")
	go func() {
		if err := s.httpServer.ListenAndServe(); err != nil {
			log.Printf("httpServer.ListenAndServe() error(%v)", err)
			panic(err)
		}
	}()
	log.Printf("BlogServer start success，listening on %s", s.httpServer.Addr)
	return nil
}

// func (s *BlogServer) Stop() error {

// }

func (s *BlogServer) prepareServer() error {
	return s.prepareNetServer()
}

func (s *BlogServer) prepareNetServer() error {
	log.Printf("prepareNetServer")
	// HTTP
	httpmux := http.NewServeMux()
	httpmux.Handle("/v1/", s.Tracing(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, World!"))
	})))
	httpmux.Handle("/v2/", s.Tracing(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, World!"))
	})))
	httpmux.Handle("/v3/", s.Tracing(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, World!"))
	})))
	httpmux.Handle("/v4/", s.Tracing(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

		w.Write([]byte("Hello, World!"))
	})))
	s.httpServer = &http.Server{
		Addr:    ":8080",
		Handler: httpmux,
	}

	return nil
}
