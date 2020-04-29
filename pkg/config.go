package pkg

import (
	"fmt"
	"os"

	"github.com/fsnotify/fsnotify"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"
)

var (
	Config     *viper.Viper
	configname string
	configpath []string
	typed      string
	debug      bool
)

func init() {
	// viper也提供了读取Command Line参数的功能：
	pflag.StringVarP(&configname, "configname", "n", ".config.yaml", "读取的配置文件")
	pflag.StringSliceVarP(&configpath, "configpath", "p", []string{}, "添加读取的配置文件路径")
	pflag.StringVarP(&typed, "typed", "t", "yaml", "设置配置文件类型")
	pflag.BoolVarP(&debug, "debug", "d", false, "日志debug输出")
	pflag.Parse()

	viper.BindPFlags(pflag.CommandLine)

	if debug {
		log.SetLevel(log.DebugLevel)
	} else {
		log.SetLevel(log.InfoLevel)
	}
	// log.Info(configpath)
	// log.Info(configname)
	// 判断文件
	var path string
	home, err := Home()
	if err != nil {
		panic(err)
	}

	if len(configpath) == 0 {
		path = fmt.Sprintf("%s/%s", home, configname)
	} else {
		path = fmt.Sprintf("%s%s", configpath[0], configname)
	}

	log.Infof("读取配置文件 %s", path)
	if !IsPathExists(path) {
		log.Warnf("配置文件不存在，动态创建中...")
		f, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY, 0600)
		defer f.Close()
		if err != nil {
			panic(err)
		} else {
			_, err = f.Write([]byte(`ll:
  desc: 快速查看文件
  alias: ls -l
gt:
  desc: 文件记录
  alias: hostname

ssh:
  desc: ssh hostn manage
  alias: ssh root@127.0.0.1
  common:
    root: high priviage
    tom: little baby
  flags:
   root:
   - -p, port
   - -u, username
   - -t, this is for root
   tom:
   - -t, this is for tom
  command:
    root:
    - root@192.168.0.1,work1
    - root@10.12.251.1,aws1
    tom:
    - tom@172.16.0.5,docker

git:
  desc: "代码管理"
  alias: git status
  flags: # 参数List
    default:
    - -v,--verbose
    - -v, --verbose         冗长输出
    status:
    - -s, --short           以简洁的格式显示状态
    - -b, --branch          显示分支信息
    - --show-stash,显示贮藏区信息
    - --ahead-behind,计算完整的领先/落后值
    - --porcelain[=<版本>],  机器可读的输出
    log:
    - --long,                以长格式显示状态（默认）
    - -z, --null            条目以 NUL 字符结尾
    - -u, --untracked-files[=<模式>]
    - --ignored,[=<模式>]    显示已忽略的文件，可选模式：traditional、matching、no。（默认：traditional）
    - --ignore-submodules,[=<何时>]
    - --column[=<风格>],     以列的方式显示未跟踪的文件
    - --no-renames,          不检测重命名
    init:
    - -t1,test1
    - -t2,test2
    - -t3,test3
    - -t4,test4
    g1:
    - -g1,g1
    - -a2,g2
    - -c3,g3
  common: # 子目录说明字段
    submodules: 依赖说明
    add: 添加说明
    go: go说明
    test: test说明
    foreach: "便利说明"
  command: # 命令tree结构
    status: 状态 # 字典 map[string]string
    submodules: # 字段数组 map[string]interface{}
      unknow: 
      - unknowTarge,未配置说明 
      add:  # 数组 interface{}.([]string)
      - branch,分支
      - force,强制
      - reference,不知道
      - path,路径
      init: 初始化 # interface{}.(string)
      deinit: de初始化
      update: 更新
      summary: 统计
      foreach: 
        go:
        - g1,G1
        - g2,G2
        - g3,G3
        test:
        - t1,g1
        - t3,-t2
        - cc,gc1`))
		}

	}

	// 读取yaml文件
	Config = viper.New()
	//添加读取的配置文件路径
	Config.SetConfigName(configname)
	if len(configpath) > 0 {
		for _, x := range configpath {
			Config.AddConfigPath(x)
		}
	} else {
		Config.AddConfigPath(home)
	}

	// 设置配置文件类型
	Config.SetConfigType(typed)
	if err := Config.ReadInConfig(); err != nil {
		panic(err)
	}

	// 很多时候，我们服务器启动之后，如果临时想修改某些配置参数，需要重启服务器才能生效，但是viper提供了监听函数，可以免重启修改配置参数，非常的实用：
	//创建一个信道等待关闭（模拟服务器环境）
	// ctx, _ := context.WithCancel(context.Background())
	//cancel可以关闭信道
	//ctx, cancel := context.WithCancel(context.Background())
	//设置监听回调函数
	Config.OnConfigChange(func(e fsnotify.Event) {
		log.Infof("配置文件发生变化 :%s", e.String())
		initCommands()
		// cancel()
	})
	//开始监听
	Config.WatchConfig()
	//信道不会主动关闭，可以主动调用cancel关闭
	// <-ctx.Done()

	initCommands()
}

func Test() {
	// TimeStamp: "2018-10-18 10:09:22"
	// Address: "Chongqing"
	// Postcode: 518000
	// CompanyInfomation:
	// Change: 1
	// Name: "Sunny"
	// MarketCapitalization: 50000000
	// EmployeeNum: 200
	// Department:
	// - "Finance"
	// - "Design"
	// - "Program"
	// - "Sales"
	// IsOpen: false
	// ssh:
	// - root@172.0.0.1,dev
	// - tom@192.168.0.1,prod
	// git:
	// - status,状态
	// - log,日志
	fmt.Printf(
		`
    TimeStamp:%s
    CompanyInfomation.Name:%s
    CompanyInfomation.Department:%s `,
		Config.Get("TimeStamp"),
		Config.Get("CompanyInfomation.Name"),
		Config.Get("CompanyInfomation.Department"),
	)
	tmp := Config.GetStringSlice("ssh")

	for a, b := range tmp {
		fmt.Println(a, b)
	}
	x := make(chan int, 0)
	<-x
}
