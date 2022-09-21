package go_cms

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/tidwall/gjson"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
	"strings"
)

const listApi = "/api.php/provide/vod/?ac=list"
const detailApi = "/api.php/provide/vod/?ac=detail"

type List struct {
	Code      int64  `json:"-"`
	Msg       string `json:"msg"`
	Page      int64  `json:"-"`
	PageCount int64  `json:"-"`
	Limit     int64  `json:"-"`
	Total     int64  `json:"-"`
}

func (t *List) flx(data []byte) {
	r := gjson.ParseBytes(data)
	t.Code = r.Get("code").Int()
	t.Msg = r.Get("msg").String()
	t.Page = r.Get("page").Int()
	t.PageCount = r.Get("pagecount").Int()
	t.Limit = r.Get("limit").Int()
	t.Total = r.Get("total").Int()
	return
}

type Class struct {
	TypeId   int64  `json:"type_id"`
	TypePid  int64  `json:"type_pid"`
	TypeName string `json:"type_name"`
}

type ClassList struct {
	*List
	Class []*Class `json:"class"`
}

type Video struct {
	VodId       int64  `json:"vod_id"`
	VodName     string `json:"vod_name"`
	TypeId      int64  `json:"type_id"`
	TypeName    string `json:"type_name"`
	VodEn       string `json:"vod_en"`
	VodTime     string `json:"vod_time"`
	VodRemarks  string `json:"vod_remarks"`
	VodPlayFrom string `json:"vod_play_from"`
}

type VideoList struct {
	*List
	VideoList []*Video `json:"list"`
}

type VideoInfo struct {
	VodId   int    `json:"vod_id"`
	TypeId  int    `json:"type_id"`
	TypeId1 int    `json:"type_id_1"`
	GroupId int    `json:"group_id"`
	VodName string `json:"vod_name"`
	VodSub  string `json:"vod_sub"`
	VodEn   string `json:"vod_en"`
	//VodStatus        int    `json:"vod_status"`
	VodLetter        string `json:"vod_letter"`
	VodColor         string `json:"vod_color"`
	VodTag           string `json:"vod_tag"`
	VodClass         string `json:"vod_class"`
	VodPic           string `json:"vod_pic"`
	VodPicThumb      string `json:"vod_pic_thumb"`
	VodPicSlide      string `json:"vod_pic_slide"`
	VodPicScreenshot string `json:"vod_pic_screenshot"`
	VodActor         string `json:"vod_actor"`
	VodDirector      string `json:"vod_director"`
	VodWriter        string `json:"vod_writer"`
	VodBehind        string `json:"vod_behind"`
	VodBlurb         string `json:"vod_blurb"`
	VodRemarks       string `json:"vod_remarks"`
	VodPubdate       string `json:"vod_pubdate"`
	//VodTotal         int    `json:"vod_total"`
	VodSerial      string `json:"vod_serial"`
	VodTv          string `json:"vod_tv"`
	VodWeekday     string `json:"vod_weekday"`
	VodArea        string `json:"vod_area"`
	VodLang        string `json:"vod_lang"`
	VodYear        string `json:"vod_year"`
	VodVersion     string `json:"vod_version"`
	VodState       string `json:"vod_state"`
	VodAuthor      string `json:"vod_author"`
	VodJumpUrl     string `json:"vod_jumpurl"`
	VodTpl         string `json:"vod_tpl"`
	VodTplPlay     string `json:"vod_tpl_play"`
	VodTplDown     string `json:"vod_tpl_down"`
	VodISend       int    `json:"vod_isend"`
	VodLock        int    `json:"vod_lock"`
	VodLevel       int    `json:"vod_level"`
	VodCopyright   int    `json:"vod_copyright"`
	VodPoints      int    `json:"vod_points"`
	VodPointsPlay  int    `json:"vod_points_play"`
	VodPointsDown  int    `json:"vod_points_down"`
	VodHits        int    `json:"vod_hits"`
	VodHitsDay     int    `json:"vod_hits_day"`
	VodHitsWeek    int    `json:"vod_hits_week"`
	VodHitsMonth   int    `json:"vod_hits_month"`
	VodDuration    string `json:"vod_duration"`
	VodUp          int    `json:"vod_up"`
	VodDown        int    `json:"vod_down"`
	VodScore       string `json:"vod_score"`
	VodScoreAll    int    `json:"vod_score_all"`
	VodScoreNum    int    `json:"vod_score_num"`
	VodTime        string `json:"vod_time"`
	VodTimeAdd     int    `json:"vod_time_add"`
	VodTimeHits    int    `json:"vod_time_hits"`
	VodTimeMake    int    `json:"vod_time_make"`
	VodTrySee      int    `json:"vod_trysee"`
	VodDouBanId    int    `json:"vod_douban_id"`
	VodDouBanScore string `json:"vod_douban_score"`
	VodReUrl       string `json:"vod_reurl"`
	VodRelVod      string `json:"vod_rel_vod"`
	VodRelArt      string `json:"vod_rel_art"`
	VodPwd         string `json:"vod_pwd"`
	VodPwdUrl      string `json:"vod_pwd_url"`
	VodPwdPlay     string `json:"vod_pwd_play"`
	VodPwdPlayUrl  string `json:"vod_pwd_play_url"`
	VodPwdDown     string `json:"vod_pwd_down"`
	VodPwdDownUrl  string `json:"vod_pwd_down_url"`
	VodContent     string `json:"vod_content"`
	VodPlayFrom    string `json:"vod_play_from"`
	VodPlayServer  string `json:"vod_play_server"`
	VodPlayNote    string `json:"vod_play_note"`
	VodPlayUrl     string `json:"vod_play_url"`
	VodDownFrom    string `json:"vod_down_from"`
	VodDownServer  string `json:"vod_down_server"`
	VodDownNote    string `json:"vod_down_note"`
	VodDownUrl     string `json:"vod_down_url"`
	VodPlot        int    `json:"vod_plot"`
	VodPlotName    string `json:"vod_plot_name"`
	VodPlotDetail  string `json:"vod_plot_detail"`
	TypeName       string `json:"type_name"`
}

type UrlList struct {
	Name string `json:"name"`
	Url  string `json:"url"`
}

func (v *VideoInfo) GetPlaylist() []*UrlList {
	playlists := make([]*UrlList, 0)
	split1 := strings.Split(v.VodPlayUrl, "#")
	for _, s := range split1 {
		s2 := strings.Split(s, "$")
		playlists = append(playlists, &UrlList{
			Name: s2[0],
			Url:  s2[1],
		})
	}
	return playlists
}
func (v *VideoInfo) GetDownloadList() []*UrlList {
	playlists := make([]*UrlList, 0)
	split1 := strings.Split(v.VodDownUrl, "#")
	for _, s := range split1 {
		s2 := strings.Split(s, "$")
		playlists = append(playlists, &UrlList{
			Name: s2[0],
			Url:  s2[1],
		})
	}
	return playlists
}

type VideoInfoList struct {
	*List
	VideoInfoList []*VideoInfo `json:"list"`
}

type cms struct {
	host      string //服务地址
	listApi   string
	detailApi string
	client    http.Client
}

func NewCMS(host string) *cms {
	c := cms{host: host}
	c.listApi = listApi
	c.detailApi = detailApi
	c.client = http.Client{}
	return &c
}
func (c *cms) SetApiPath(listApi, detailApi string) {
	c.detailApi = detailApi
	c.listApi = listApi
}
func (c *cms) SetProxy(url *url.URL) {
	c.client.Transport = &http.Transport{
		// 设置代理
		Proxy: http.ProxyURL(url),
	}
}
func (c *cms) Get(url1 string) ([]byte, error) {
	request, err := http.NewRequest("GET", url1, nil)
	if err != nil {
		return nil, err
	}
	get, err := c.client.Do(request)
	if err != nil {
		return nil, err
	}
	all, err := ioutil.ReadAll(get.Body)
	return all, err
}

func (c *cms) ClassList() ([]*Class, error) {
	data := make([]*Class, 0)
	get, err := c.Get(fmt.Sprintf("%s%s&ids=1", c.host, c.listApi))
	if err != nil {
		return data, err
	}
	l := ClassList{}
	err = json.Unmarshal(get, &l)
	if err != nil {
		return data, err
	}
	l.flx(get)
	data = l.Class
	return data, nil
}
func (c *cms) NewList(ids []int, Type, page, hour int) (*VideoList, error) {
	data := new(VideoList)
	strIdArr := make([]string, 0)
	if ids != nil {
		for _, id := range ids {
			strIdArr = append(strIdArr, strconv.Itoa(id))
		}
	}
	idsStr := strings.Join(strIdArr, ",")
	get, err := c.Get(fmt.Sprintf("%s%s&ids=%s&t=%d&pg=%d&h=%d", c.host, c.listApi, idsStr, Type, page, hour))
	if err != nil {
		return data, err
	}
	err = json.Unmarshal(get, data)
	if err == nil {
		data.flx(get)
	}
	return data, err
}
func (c *cms) Detail(id int64) (*VideoInfo, error) {
	l := new(VideoInfoList)
	v := new(VideoInfo)
	get, err := c.Get(fmt.Sprintf("%s%s&ids=%d", c.host, c.detailApi, id))
	if err != nil {
		return v, err
	}
	get = bytes.ReplaceAll(get, []byte("\n"), []byte{})
	err = json.Unmarshal(get, l)
	if err == nil {
		l.flx(get)
		if len(l.VideoInfoList) > 0 {
			v = l.VideoInfoList[0]
		}
	}
	return v, err
}

func (c *cms) DetailList(ids []int, Type, page, hour int, key string) (*VideoInfoList, error) {
	strIdArr := make([]string, 0)
	if ids != nil {
		for _, id := range ids {
			strIdArr = append(strIdArr, strconv.Itoa(id))
		}
	}
	idsStr := strings.Join(strIdArr, ",")
	l := new(VideoInfoList)
	get, err := c.Get(fmt.Sprintf("%s%s&ids=%s&t=%d&pg=%d&h=%d&wd=%s", c.host, c.detailApi, idsStr, Type, page, hour, key))
	if err != nil {
		return l, err
	}
	get = bytes.ReplaceAll(get, []byte("\n"), []byte{})
	err = json.Unmarshal(get, l)
	if err == nil {
		l.flx(get)
	}
	return l, err
}
