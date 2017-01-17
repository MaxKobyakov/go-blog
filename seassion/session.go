package seasion 

type sessionData struct {
	username string
}

type Session struct {
	data map[string]*sessionData
}

func NewSassion () *Session {
	s := new (Seassion)
	s.data = make(map[string]*sessionData)
return s
}