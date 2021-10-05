package main

/////////////////////////////////////////////
/* Import */

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/PuerkitoBio/goquery"
)

/////////////////////////////////////////////
/* Declaration */
// number

// string
	const STR_BASE_URL string = "https://www.jobkorea.co.kr/Search/?stext=frontend&local=I000&tabType=recruit"

  // error - string

  // error - error

/////////////////////////////////////////////
/* Logic */

// Main Logic

func main() {

	c_amount_pages := make(chan int)
	go fn_get__amount_pages(c_amount_pages)
	int_amount := <- c_amount_pages

	fn_get__job_infos(int_amount)
}

// Getting the Info of the Jobs

func fn_get__job_infos(_page_amount int) {

	for i := 1; i <= _page_amount; i = i + 1 {
		str_full_url := fn_get__full_url(i)
		fmt.Println(str_full_url)
	}
}

func fn_get__full_url(_page_num int) string {

	str_now_page := "&Page_No=" + strconv.Itoa(_page_num)
	ret_full_url := STR_BASE_URL + str_now_page
	return ret_full_url
}

// Getting Amount of Pages Part

func fn_get__amount_pages(_c_amount_pages chan int) {

	res := fn_get__http()

	defer res.Body.Close()

	doc, err := goquery.NewDocumentFromReader(res.Body)

	fn_check__error(err)

	// first page + extra pages
	int_pages := doc.Find(".tplPagination > ul a").Length() + 1

	_c_amount_pages <- int_pages
}

func fn_get__http() *http.Response {

	res, err := http.Get(STR_BASE_URL)

	fn_check__error(err)
	fn_check__status_code(res)

	return res
}

// Exception Checking Part

func fn_check__error(_err error) {
	if _err != nil {
		log.Fatalln(_err)
	}
}

func fn_check__status_code(_res *http.Response) {
	if _res.StatusCode != 200 {
		log.Fatalln(fmt.Sprintf("Request failed with status code: %d", _res.StatusCode))
	}
}