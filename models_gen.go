// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package hackernewsgraphql

type Item interface {
	IsItem()
}

type Ask struct {
	ID          string    `json:"id"`
	By          string    `json:"by"`
	Time        Timestamp `json:"time"`
	Descendants int       `json:"descendants"`
	Kids        []Comment `json:"kids"`
	Score       int       `json:"score"`
	Text        string    `json:"text"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
}

func (Ask) IsItem() {}

type Comment struct {
	ID     string    `json:"id"`
	By     string    `json:"by"`
	Time   Timestamp `json:"time"`
	Kids   []Comment `json:"kids"`
	Parent Item      `json:"parent"`
	Text   string    `json:"text"`
}

func (Comment) IsItem() {}

type Job struct {
	ID    string    `json:"id"`
	By    string    `json:"by"`
	Time  Timestamp `json:"time"`
	Score int       `json:"score"`
	Text  string    `json:"text"`
	Title string    `json:"title"`
	URL   string    `json:"url"`
}

func (Job) IsItem() {}

type Poll struct {
	ID          string    `json:"id"`
	By          string    `json:"by"`
	Time        Timestamp `json:"time"`
	Descendants int       `json:"descendants"`
	Kids        []Comment `json:"kids"`
	Parts       []PollOpt `json:"parts"`
	Score       int       `json:"score"`
	Text        string    `json:"text"`
	Title       string    `json:"title"`
}

func (Poll) IsItem() {}

type PollOpt struct {
	ID    string    `json:"id"`
	By    string    `json:"by"`
	Time  Timestamp `json:"time"`
	Poll  Poll      `json:"poll"`
	Score int       `json:"score"`
	Text  string    `json:"text"`
}

func (PollOpt) IsItem() {}

type Story struct {
	ID          string    `json:"id"`
	By          string    `json:"by"`
	Time        Timestamp `json:"time"`
	Descendants int       `json:"descendants"`
	Kids        []Comment `json:"kids"`
	Score       int       `json:"score"`
	Title       string    `json:"title"`
	URL         string    `json:"url"`
}

func (Story) IsItem() {}

type Timestamp struct {
	Epoch   int    `json:"epoch"`
	Iso8601 string `json:"iso8601"`
}

type User struct {
	ID        string    `json:"id"`
	About     *string   `json:"about"`
	Created   Timestamp `json:"created"`
	Karma     *int      `json:"karma"`
	Delay     *int      `json:"delay"`
	Submitted []Item    `json:"submitted"`
}
