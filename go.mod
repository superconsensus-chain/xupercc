module github.com/superconsensus-chain/xupercc

go 1.15

require (
	github.com/astaxie/beego v1.12.2
	github.com/fastly/go-utils v0.0.0-20180712184237-d95a45783239 // indirect
	github.com/gin-contrib/cors v1.3.1
	github.com/gin-gonic/gin v1.6.3
	github.com/jehiah/go-strftime v0.0.0-20171201141054-1d33003b3869 // indirect
	github.com/lestrrat-go/file-rotatelogs v2.4.0+incompatible
	github.com/lestrrat-go/strftime v1.0.3 // indirect
	github.com/natefinch/lumberjack v2.0.0+incompatible
	github.com/rifflock/lfshook v0.0.0-20180920164130-b9218ef580f5
	github.com/sirupsen/logrus v1.7.0
	github.com/superconsensus/matrixcore v1.0.3-0.20220228084652-c1d605f3707c
	github.com/tebeka/strftime v0.1.5 // indirect
	github.com/xuperchain/xuper-sdk-go/v2 v2.0.0
	github.com/xuperchain/xuperchain v0.0.0-20220307080058-ccc024f45456
	go.uber.org/zap v1.16.0
	google.golang.org/genproto v0.0.0-20200608115520-7c474a2e3482 // indirect
	google.golang.org/grpc v1.35.0
	gopkg.in/ini.v1 v1.62.0
	gopkg.in/natefinch/lumberjack.v2 v2.0.0 // indirect
)

//replace github.com/xuperchain/xuper-sdk-go/v2 v2.0.0 => github.com/vantinfy/matrix-sdk-go/v2 v2.0.1-0.20220309100255-278ac065560d
