package db

type Handle interface {
}

func NEWHandler() (Handle, error) {
	return NEWSQLDB()
}
