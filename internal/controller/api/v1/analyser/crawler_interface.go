package analyser

type Crawler interface {
	FromURL(url string) (map[string]interface{}, error)
}
