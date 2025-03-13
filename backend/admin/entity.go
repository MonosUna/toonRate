package admin

type Production struct {
	ID          int64  `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Genre       string `json:"genre"`
	Year        int    `json:"year"`
}

type User struct {
	ID          int64  `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Pfp         string `json:"pfp"`
	Description string `json:"description"`
}

type Collection struct {
	ID      int64  `json:"id"`
	Author  string `json:"author"`
	Topic   string `json:"topic"`
	Message string `json:"message"`
}

type Discussion struct {
	ID           int64  `json:"id"`
	Production   int64  `json:"production"`
	Author       string `json:"author"`
	Topic        string `json:"topic"`
	EntryMessage string `json:"entry_message"`
	Message      string `json:"message"`
}

type Review struct {
	ID         int64  `json:"id"`
	Production int64  `json:"production"`
	Author     string `json:"author"`
	Topic      string `json:"topic"`
	Message    string `json:"message"`
}
