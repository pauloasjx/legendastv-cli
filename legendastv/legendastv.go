package legendastv

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"os"
	"regexp"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/net/publicsuffix"
)

type Subtitle struct {
	Title  string
	Link   string
	Author string
	//Downloads int64
	//Grade     float64
	//Date      time.Time
}

func Login(login string, password string) http.Client {
	_url := "http://legendas.tv/login"

	jar, _ := cookiejar.New(&cookiejar.Options{PublicSuffixList: publicsuffix.List})

	client := http.Client{Jar: jar}
	resp, _ := client.PostForm(_url, url.Values{
		"data[User][username]": {login},
		"data[User][password]": {password},
	})
	defer resp.Body.Close()

	//body, _ := ioutil.ReadAll(resp.Body)

	return client
}

func Search(client http.Client, query string) []Subtitle {
	var subtitles []Subtitle
	_url := fmt.Sprintf("http://legendas.tv/legenda/busca/%s", query)

	resp, _ := client.Get(_url)
	defer resp.Body.Close()

	doc, _ := goquery.NewDocumentFromResponse(resp)
	//fmt.Println(doc.Text())

	r_downloads, _ := regexp.Compile(`(\d*) downloads`)
	r_grade, _ := regexp.Compile(`nota (\d*)`)
	//r_data, _ := regex.Compile(`(\d{2}/\d{2}/\d{4}).*`)

	doc.Find(".f_left").Each(func(i int, container *goquery.Selection) {
		wrapper := container.Find("p:not([class])")

		title := wrapper.Text()
		link, _ := wrapper.Find("a").Attr("href")

		wrapper = container.Find("p[class='data']")

		author := wrapper.Find("a").Text()

		downloads := r_downloads.FindString(wrapper.Text())
		downloads = downloads[:len(downloads)-10]

		grade := r_grade.FindString(wrapper.Text())
		grade = grade[5:]

		subtitles = append(subtitles, Subtitle{title, link, author})
	})

	//fmt.Println(subtitles[0])

	return subtitles
}

func Download(client http.Client, subtitle Subtitle) {
	_url := fmt.Sprintf("http://legendas.tv%s", subtitle.Link)

	resp, _ := client.Get(_url)
	defer resp.Body.Close()

	r_download, _ := regexp.Compile(`/downloadarquivo/\w*`)
	body, _ := ioutil.ReadAll(resp.Body)
	link_download := fmt.Sprintf("http://legendas.tv%s", r_download.FindString(string(body)))

	file, _ := os.Create(fmt.Sprintf("%s.rar", subtitle.Title))
	defer file.Close()

	resp, _ = client.Get(link_download)
	defer resp.Body.Close()

	_, _ = io.Copy(file, resp.Body)

}
