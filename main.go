package main

import (
	"github.com/freelifer/gohelper/transport/restful"

	_ "github.com/freelifer/gohelper/pkg/settings"
)

// http://www.cnblogs.com/CloudMan6/p/7770916.html

// docker run -p 5601:5601 -p 9200:9200 -p 5044:5044 -e ES_JAVA_OPTS="-Xms256m -Xmx256m" -it --name elk sebp/elk

// docker run -d -p 9200:9200 -p 5601:5601 -e ES_JAVA_OPTS="-Xms256m -Xmx256m" nshou/elasticsearch-kibana

// $

// docker run --init -d --name elasticsearch -p 9200:9200 blacktop/elasticsearch
// $ docker run --init -d --name kibana --link elasticsearch -p 5601:5601 blacktop/kibana

// docker run --init -e ES_JAVA_OPTS="-Xms256m -Xmx256m" --name elasticsearch -p 9200:9200 blacktop/elasticsearch

// go test models/*.go -v -o .
// @title Golang GoHelper API
// @version 1.0
// @description An system of gohelper
func main() {
	restful.Run()
}

// X-RateLimit-Limit: 用户每个小时允许发送请求的最大值
// X-RateLimit-Remaining：当前时间窗口剩下的可用请求数目
// X-RateLimit-Rest: 时间窗口重置的时候，到这个时间点可用的请求数量就会变成X-RateLimit-Limit 的值

// acecleaner.apk
// ('===>>>', ['loopme', 'admob', 'mopub', 'appmonet', 'facebook'])

// emojiflash.apk
// ('===>>>', ['admob', 'facebook', 'inmobi', 'mopub', 'facebook'])

// flashlighta.apk
// ('===>>>', ['tapjoy', 'admob', 'facebook', 'mopub', 'appcloudbox'])
