package go_cms

import (
	"encoding/json"
	"fmt"
	"net/url"
	"testing"
)

const host1 = "https://cj.lziapi.com"
const host2 = "https://jyzyapi.com"
const host3 = "https://api.guangsuapi.com"

func TestName(t *testing.T) {
	str := `{"code":1,"msg":"数据列表","page":1,"pagecount":1,"limit":"20","total":1,"list":[],"class":[{"type_id":1,"type_pid":0,"type_name":"电影片"},{"type_id":2,"type_pid":0,"type_name":"连续剧"},{"type_id":3,"type_pid":0,"type_name":"综艺片"},{"type_id":4,"type_pid":0,"type_name":"动漫片"},{"type_id":6,"type_pid":1,"type_name":"动作片"},{"type_id":7,"type_pid":1,"type_name":"喜剧片"},{"type_id":8,"type_pid":1,"type_name":"爱情片"},{"type_id":9,"type_pid":1,"type_name":"科幻片"},{"type_id":10,"type_pid":1,"type_name":"恐怖片"},{"type_id":11,"type_pid":1,"type_name":"剧情片"},{"type_id":12,"type_pid":1,"type_name":"战争片"},{"type_id":13,"type_pid":2,"type_name":"国产剧"},{"type_id":14,"type_pid":2,"type_name":"香港剧"},{"type_id":15,"type_pid":2,"type_name":"韩国剧"},{"type_id":16,"type_pid":2,"type_name":"欧美剧"},{"type_id":20,"type_pid":1,"type_name":"记录片"},{"type_id":21,"type_pid":2,"type_name":"台湾剧"},{"type_id":22,"type_pid":2,"type_name":"日本剧"},{"type_id":23,"type_pid":2,"type_name":"海外剧"},{"type_id":24,"type_pid":2,"type_name":"泰国剧"},{"type_id":25,"type_pid":3,"type_name":"大陆综艺"},{"type_id":26,"type_pid":3,"type_name":"港台综艺"},{"type_id":27,"type_pid":3,"type_name":"日韩综艺"},{"type_id":28,"type_pid":3,"type_name":"欧美综艺"},{"type_id":29,"type_pid":4,"type_name":"国产动漫"},{"type_id":30,"type_pid":4,"type_name":"日韩动漫"},{"type_id":31,"type_pid":4,"type_name":"欧美动漫"},{"type_id":32,"type_pid":4,"type_name":"港台动漫"},{"type_id":33,"type_pid":4,"type_name":"海外动漫"},{"type_id":34,"type_pid":1,"type_name":"伦理片"}]}`
	list := ClassList{}
	err := json.Unmarshal([]byte(str), &list)
	println(err)
}
func TestClass(t *testing.T) {
	newCMS := NewCMS(host1)
	class, err := newCMS.ClassList()
	if err == nil {
		for _, c := range class {
			fmt.Println(c.TypeName)
		}
	} else {
		t.Fatal(err)
	}
}
func TestNewList(t *testing.T) {
	newCMS := NewCMS(host1)
	newList, err := newCMS.NewList(nil, 0, 1, 0)
	if err == nil {
		for _, video := range newList.VideoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
}
func TestPath(t *testing.T) {
	newCMS := NewCMS(host2)
	newCMS.SetApiPath("/provide/vod/at/json", "/provide/vod/at/json")
	newList, err := newCMS.NewList(nil, 0, 1, 0)
	if err == nil {
		for _, video := range newList.VideoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
}
func TestDetail(t *testing.T) {
	newCMS := NewCMS(host2)
	newCMS.SetApiPath("/provide/vod/at/json", "/provide/vod/at/json")
	Detail, err := newCMS.Detail(27540)
	if err == nil {
		fmt.Println(Detail.VodName)
	} else {
		t.Fatal(err)
	}
}
func TestDetailList(t *testing.T) {
	newCMS := NewCMS(host3)
	list, err := newCMS.DetailList(nil, 0, 1, 0, "蜡笔小新")
	if err == nil {
		for _, video := range list.VideoInfoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
}
func TestHOSTS(t *testing.T) {
	//天空 https://m3u8.tiankongapi.com/api.php/provide/vod/from/tkm3u8/?ac=list
	newCMS := NewCMS("https://m3u8.tiankongapi.com")
	newCMS.SetApiPath("/api.php/provide/vod/from/tkm3u8/?ac=list", "/api.php/provide/vod/from/tkm3u8/?ac=detail")
	list, err := newCMS.DetailList(nil, 0, 1, 0, "蜡笔小新")
	if err == nil {
		for _, video := range list.VideoInfoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
	//百度 https://api.apibdzy.com/api.php/provide/art/?ac=list
	newCMS = NewCMS("https://api.apibdzy.com")
	list, err = newCMS.DetailList(nil, 0, 1, 0, "蜡笔小新")
	if err == nil {
		for _, video := range list.VideoInfoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
}
func TestProxy(t *testing.T) {
	newCMS := NewCMS(host1)
	parse, _ := url.Parse("http://127.0.0.1:5216")
	newCMS.SetProxy(parse)
	newList, err := newCMS.NewList(nil, 0, 1, 0)
	if err == nil {
		for _, video := range newList.VideoList {
			fmt.Println(video.VodName)
		}
	} else {
		t.Fatal(err)
	}
}
