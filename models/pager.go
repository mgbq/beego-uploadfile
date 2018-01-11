package models

import (
	"bytes"
	"fmt"
	"math"
	"strings"
)

type Pager struct {
	Page     int
	Totalnum int
	Pagesize int
	urlpath  string
	urlquery string
	nopath   bool
}

func NewPager(page, totalnum, pagesize int, url string, nopath ...bool) *Pager {
	p := new(Pager)
	p.Page = page
	p.Totalnum = totalnum
	p.Pagesize = pagesize

	arr := strings.Split(url, "?")
	p.urlpath = arr[0]
	if len(arr) > 1 {
		p.urlquery = "?" + arr[1]
	} else {
		p.urlquery = ""
	}

	if len(nopath) > 0 {
		p.nopath = nopath[0]
	} else {
		p.nopath = false
	}

	return p
}

func (c *Pager) url(page int) string {
	if c.nopath { //不使用目录形式
		if c.urlquery != "" {
			return fmt.Sprintf("%s%s&page=%d", c.urlpath, c.urlquery, page)
		} else {
			return fmt.Sprintf("%s?page=%d", c.urlpath, page)
		}
	} else {
		return fmt.Sprintf("%s/page/%d%s", c.urlpath, page, c.urlquery)
	}
}

func (c *Pager) ToString() string {
	if c.Totalnum <= c.Pagesize {
		return ""
	}

	var buf bytes.Buffer
	var from, to, linknum, offset, totalpage int

	offset = 5
	linknum = 10

	totalpage = int(math.Ceil(float64(c.Totalnum) / float64(c.Pagesize)))

	if totalpage < linknum {
		from = 1
		to = totalpage
	} else {
		from = c.Page - offset
		to = from + linknum
		if from < 1 {
			from = 1
			to = from + linknum - 1
		} else if to > totalpage {
			to = totalpage
			from = totalpage - linknum + 1
		}
	}

	buf.WriteString("<div class=\"clearfix\"><ul class=\"pagination pagination-sm pull-right\">")
	if c.Page > 1 {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&laquo;</a></li>", c.url(c.Page-1)))
	} else {
		buf.WriteString("<li class=\"disabled\"><span>&laquo;</span></li>")
	}

	if c.Page > linknum {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">1...</a></li>", c.url(1)))
	}

	for i := from; i <= to; i++ {
		if i == c.Page {
			buf.WriteString(fmt.Sprintf("<li class=\"active\"><span>%d</span></li>", i))
		} else {
			buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">%d</a></li>", c.url(i), i))
		}
	}

	if totalpage > to {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">...%d</a></li>", c.url(totalpage), totalpage))
	}

	if c.Page < totalpage {
		buf.WriteString(fmt.Sprintf("<li><a href=\"%s\">&raquo;</a></li>", c.url(c.Page+1)))
	} else {
		buf.WriteString(fmt.Sprintf("<li class=\"disabled\"><span>&raquo;</span></li>"))
	}
	buf.WriteString("</ul></div>")

	return buf.String()
}
