package utils

import (
	"beeblog/models"
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"log"
	"os"
	"reflect"
	"strconv"
	"time"
)
var client *elastic.Client
var host = "http://47.98.109.5:8209"
const mapping = `
{
	"settings":{
		"number_of_shards": 1,
		"number_of_replicas": 0
	},
	"mappings":{
		"beeblog":{
			"properties":{
				"UserId":{
					"type":"long"
				},
				"BlogHtml":{
					"type":"text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"Ctime":{
					"type":"date"
				},
				"Title":{
					"type":"text",
					"analyzer": "ik_max_word",
					"search_analyzer": "ik_smart"
				},
				"Browses":{
					"type":"long"
				}
			}
		}
	}
}`

func init() {
	var err error
	client, err = elastic.NewClient(
		elastic.SetURL(host),
		elastic.SetSniff(false),
		elastic.SetHealthcheckInterval(10*time.Second),
		elastic.SetGzip(true),
		elastic.SetErrorLog(log.New(os.Stderr, "ELASTIC ", log.LstdFlags)),
		elastic.SetInfoLog(log.New(os.Stdout, "", log.LstdFlags)))

	if err != nil {
		panic(err)
	}
}

func Index()  {
	ctx := context.Background()
	exists, err := client.IndexExists("beeblog").Do(ctx)
	if err != nil {
		panic(err)
	}
	if !exists {
		createIndex, err := client.CreateIndex("beeblog").BodyString(mapping).Do(ctx)
		if err != nil {
			panic(err)
		}
		if !createIndex.Acknowledged {
			fmt.Println("请求未被接收到")
		}
		fmt.Println(createIndex)
	} else {
		fmt.Println("已经存在beeblog索引")
	}
}

func ESSave(blog *models.Blog)  {
	ctx := context.Background()
	id := strconv.FormatInt(blog.Id,10)
	put1, err := client.Index().Index("beeblog").Type("beeblog").Id(id).BodyJson(blog).Do(ctx)
	if err != nil {
		// Handle error
		panic(err)
	}
	fmt.Printf("Indexed beeblog %s to index %s, type %s\n", put1.Id, put1.Index, put1.Type)
}
//删除
func ESDelete(id string) {
	res, err := client.Delete().Index("beeblog").
		Type("beeblog").
		Id(id).
		Do(context.Background())
	if err != nil {
		println(err.Error())
		return
	}
	fmt.Printf("delete result %s\n", res.Result)
}

func Search(key string) (*models.Blog,error) {
	ctx := context.Background()
	query := elastic.NewMultiMatchQuery(key,"Title","BlogHtml")
	searchResult, err := client.Search().
		Index("beeblog").  // 指定index，返回一个*SearchService对象
		//Type("beeblog").
		Query(query).                     // 设置查询体，返回同一个*SearchService对象
		//Sort("user", true).     // 按照user升序排列
		//From(0).Size(10).            // 从第一条数据，找十条，即0-9
		Pretty(true).                   // 使查询request和返回的结果格式美观
		Do(ctx)                               // 返回一个*SearchResult
	if err != nil {
		println("search error",err.Error())
		return nil,err
	}
	fmt.Printf("找到 [%d] 组tweets\n", searchResult.Hits.TotalHits)
	// 查看匹配到多少组数据
	fmt.Printf("找到 [%d] 组tweets\n", searchResult.TotalHits())
	var typ models.Blog
	for _, item := range searchResult.Each(reflect.TypeOf(typ)) { //从搜索结果中取数据的方法
		t := item.(models.Blog)
		fmt.Printf("%#v\n", t)
		return &t,nil
	}
	if err != nil {
		panic(err)
	}
	return nil,nil
}