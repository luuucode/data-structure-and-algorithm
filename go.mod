module data-structure-and-algorithm

require (
	github.com/360EntSecGroup-Skylar/excelize v1.4.1 // indirect
	github.com/DataDog/zstd v1.3.5 // indirect
	github.com/PuerkitoBio/goquery v1.5.1 // indirect
	github.com/Shopify/sarama v1.20.0
	github.com/TruthHun/html2md v0.0.0-20190507142218-8352cc68f88e
	github.com/afex/hystrix-go v0.0.0-20180502004556-fa1af6a1f4f5
	github.com/alibaba/sentinel-golang v0.3.0
	github.com/aliyun/alibaba-cloud-sdk-go v0.0.0-20190827030439-84d9962c10d3
	github.com/aliyun/aliyun-oss-go-sdk v2.0.1+incompatible
	github.com/andybalholm/cascadia v1.2.0 // indirect
	github.com/antchfx/htmlquery v1.2.3 // indirect
	github.com/antchfx/xmlquery v1.2.4 // indirect
	github.com/antchfx/xpath v1.1.8 // indirect
	github.com/astaxie/beego v1.12.0 // indirect
	github.com/axgle/mahonia v0.0.0-20180208002826-3358181d7394
	github.com/codahale/hdrhistogram v0.0.0-20161010025455-3a0bb77429bd // indirect
	github.com/coreos/etcd v3.3.18+incompatible
	github.com/coreos/go-systemd v0.0.0-20191104093116-d3cd4ed1dbcf // indirect
	github.com/dchest/captcha v0.0.0-20170622155422-6a29415a8364
	github.com/denisenkom/go-mssqldb v0.0.0-20190515213511-eb9f6a1743f3 // indirect
	github.com/dgrijalva/jwt-go v3.2.0+incompatible
	github.com/eclipse/paho.mqtt.golang v1.2.0
	github.com/erikstmartin/go-testdb v0.0.0-20160219214506-8d10e4a1bae5 // indirect
	github.com/fatih/camelcase v1.0.0 // indirect
	github.com/gin-gonic/gin v1.6.3
	github.com/go-redis/redis v6.15.6+incompatible
	github.com/go-session/session v3.1.2+incompatible
	github.com/go-sql-driver/mysql v1.5.0 // indirect
	github.com/go-vgo/robotgo v0.0.0-20191216133555-c86926da97a5
	github.com/gobwas/glob v0.2.3 // indirect
	github.com/gocolly/colly v1.2.0
	github.com/gofrs/uuid v3.2.0+incompatible // indirect
	github.com/goinggo/mapstructure v0.0.0-20140717182941-194205d9b4a9
	github.com/golang/mock v1.3.1
	github.com/golang/protobuf v1.4.2 // indirect
	github.com/google/go-cmp v0.4.0
	github.com/google/uuid v1.1.1
	github.com/gorilla/mux v1.6.2
	github.com/grpc-ecosystem/go-grpc-middleware v1.1.0
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway v1.12.1 // indirect
	github.com/hashicorp/consul/api v1.1.0
	github.com/jinzhu/copier v0.0.0-20190625015134-976e0346caa8
	github.com/jinzhu/gorm v1.9.2
	github.com/jinzhu/inflection v0.0.0-20180308033659-04140366298a // indirect
	github.com/jinzhu/now v1.0.0 // indirect
	github.com/kennygrant/sanitize v1.2.4 // indirect
	github.com/koding/multiconfig v0.0.0-20171124222453-69c27309b2d7
	github.com/lib/pq v1.1.1 // indirect
	github.com/lxn/walk v0.0.0-20191128110447-55ccb3a9f5c1
	github.com/lxn/win v0.0.0-20191128105842-2da648fda5b4 // indirect
	github.com/mattn/go-colorable v0.1.6
	github.com/mdp/qrterminal/v3 v3.0.0
	github.com/microcosm-cc/bluemonday v1.0.2
	github.com/mozillazg/go-pinyin v0.18.0
	github.com/nfnt/resize v0.0.0-20180221191011-83c6a9932646
	github.com/nguyenthenguyen/docx v0.0.0-20181031033525-8cb697a41e43
	github.com/nicksnyder/go-i18n/v2 v2.0.3
	github.com/nsqio/go-nsq v1.0.7
	github.com/olivere/elastic v6.2.18+incompatible
	github.com/opentracing/opentracing-go v1.1.0
	github.com/panjf2000/ants/v2 v2.3.1
	github.com/pibigstar/go-sdk v0.0.0-20190727082016-c4f5d238d8f5
	github.com/pierrec/lz4 v2.0.7+incompatible // indirect
	github.com/pkg/errors v0.8.1
	github.com/qianlnk/pgbar v0.0.0-20190929032005-46c23acad4ed
	github.com/qianlnk/to v0.0.0-20191230085244-91e712717368 // indirect
	github.com/robfig/cron v0.0.0-20180505203441-b41be1df6967
	github.com/russross/blackfriday v2.0.0+incompatible
	github.com/saintfish/chardet v0.0.0-20120816061221-3af4cd4741ca // indirect
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/shirou/gopsutil v2.19.12+incompatible
	github.com/shurcooL/sanitized_anchor_name v1.0.0 // indirect
	github.com/skip2/go-qrcode v0.0.0-20190110000554-dc11ecdae0a9
	github.com/smallnest/rpcx v0.0.0-20170601021420-329cf0078113
	github.com/smartwalle/alipay/v3 v3.0.9
	github.com/sony/sonyflake v1.0.0
	github.com/streadway/amqp v0.0.0-20190404075320-75d898a42a94
	github.com/temoto/robotstxt v1.1.1 // indirect
	github.com/uber/jaeger-client-go v2.23.1+incompatible
	github.com/uber/jaeger-lib v2.2.0+incompatible // indirect
	github.com/unidoc/unioffice v1.2.0
	github.com/varstr/uaparser v0.0.0-20170929040706-6aabb7c4e98c
	github.com/xdg/scram v0.0.0-20180814205039-7eeb5667e42c // indirect
	github.com/xdg/stringprep v1.0.0 // indirect
	github.com/xuri/excelize v1.4.0
	github.com/yanyiwu/gojieba v1.1.2
	github.com/youzan/go-nsq v1.3.1
	go.mongodb.org/mongo-driver v1.2.0
	go.uber.org/ratelimit v0.1.0
	go.uber.org/zap v1.13.0 // indirect
	golang.org/x/crypto v0.0.0-20200221231518-2aa609cf4a9d
	golang.org/x/net v0.0.0-20200602114024-627f9648deb9 // indirect
	golang.org/x/oauth2 v0.0.0-20190604053449-0f29369cfe45
	golang.org/x/sync v0.0.0-20190911185100-cd5d95a43a6e
	golang.org/x/sys v0.0.0-20200515095857-1151b9dac4a9 // indirect
	golang.org/x/text v0.3.2
	golang.org/x/time v0.0.0-20190308202827-9d24e82272b4
	google.golang.org/appengine v1.6.6 // indirect
	google.golang.org/grpc v1.27.0
	google.golang.org/protobuf v1.24.0 // indirect
	gopkg.in/Knetic/govaluate.v3 v3.0.0 // indirect
	gopkg.in/go-oauth2/redis.v3 v3.2.1
	gopkg.in/go-playground/pool.v3 v3.1.1
	gopkg.in/oauth2.v3 v3.12.0
	gopkg.in/yaml.v2 v2.2.8
	sigs.k8s.io/yaml v1.1.0 // indirect
)

go 1.13
