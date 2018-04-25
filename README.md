### Introduction
本项目是基于GO语言Web框架beego开发的API服务，提供中南大学校内聚合查询API接口服务，目前已用在微信小程序[we中南](https://lovesmg.cn/2018/04/25/wecsu.html)中,旨在练习GO WEB开发.
演示站点:[https://csugo.lovesmg.cn](https://csugo.lovesmg.cn)
功能列表：
1. 成绩查询
2. 排名查询
3. 考试查询(TODO)
4. 图书续借
5. 图书查询(TODO)
6. 自习教室查询
7. 校车查询
8. 招聘查询
9. 四六级查询
10. 计算机等级查询(TODO)
11. 课表查询
12. 校内通知查询

### Usage
> 请保证已经安装了beego和bee,查看[beego安装教程](https://beego.me/quickstart)

```shell
git clone https://github.com/csuhan/csugo.git
cd csugo
bee run
```

### Keys
本项目基本思路是爬去校内网站信息,经过格式化后输出为json格式,为客户端提供便捷的API接口,具体实现过程可以参见代码实现.
以成绩查询为例,具体实现思路:
1. 模拟登陆[http://csujwc.its.csu.edu.cn/jsxsd/](http://csujwc.its.csu.edu.cn/jsxsd/)(此地址绕过了验证码)
2. 抓取原始成绩页[http://csujwc.its.csu.edu.cn/jsxsd/kscj/yscjcx_list](http://csujwc.its.csu.edu.cn/jsxsd/kscj/yscjcx_list)
3. 解析网页,将成绩信息解析出来,转为GO的struct,然后格式化输出为json格式
4. 对应API服务路由https://csugo.lovesmg.cn/api/v1/jwc/:id/:pwd/grade

### Service

> 请在请求参数后加上token=csugo-token

如成绩查询:

`wget localhost:8080/api/v1/jwc/myid/mypassword/grade?token=csugo-token`

在请求地址结尾需要添加token

下面是部分接口数据示例

#### 成绩查询接口
`youwebsite/api/v1/jwc/:id/:pwd/grade [get]`

参数说明:

```
id:学号

pwd:教务系统密码
```

返回内容说明

```
成功：
{
  "StateCode": 1,
  "Error": "",
  "Grades": [
    {
      "ClassNo": 1,
      "FirstTerm": "2015-2016-1",
      "GottenTerm": "2015-2016-1",
      "ClassName": "[010001T1]新生课",
      "MiddleGrade": "85",
      "FinalGrade": "85",
      "Grade": "85",
      "ClassScore": "1",
      "ClassType": "必修",
      "ClassProp": "信息技术类"
    },
    {
      "ClassNo": 2,
      "FirstTerm": "2015-2016-1",
      "GottenTerm": "2015-2016-1",
      "ClassName": "[080203X1]工程制图基础",
      "MiddleGrade": "84",
      "FinalGrade": "84",
      "Grade": "84",
      "ClassScore": "4",
      "ClassType": "必修",
      "ClassProp": "信息技术类"
    },
    部分内容省略......
    ]

账号或者密码错误:
{
  "StateCode": -1,
  "Error": "账号或者密码错误,请重新输入",
  "Grades": []
}

```
#### 课表查询

`youwebsite/api/v1/jwc/:id/:pwd/class/:term/:week [get]`

参数说明:

```
id:学号
pwd:教务系统密码
term:学期．如：2017-2018-2
week:周次．0表示所有周，其他表示周次
```

返回说明：

```
返回数据中Class为二维数组，长度为42,排列方式为:

        周一　周二　　周三　周四　周五　周六　周日
  1-2    1    2     3    4    5    6    7

  3-4    8    9    10    ................

  5-6    ................................

  7-8    ................................

 9-10    ................................

11-12    ...........................   42
```

返回数据：

```
成功：
{
  "StateCode": 1,
  "Error": "",
  "Class": [
    [
      {
        "ClassName": "雷达干涉测量（双语）",
        "Teacher": "李志伟教授",
        "Weeks": "3-12(周)",
        "Place": "B座309"
      }
    ],
    [
      {
        "ClassName": "马克思主义基本原理",
        "Teacher": "罗会钧教授",
        "Weeks": "1-16(周)",
        "Place": "A座208"
      }
    ],
    [
      {
        "ClassName": "",
        "Teacher": "",
        "Weeks": "",
        "Place": ""
      }
    ],
    [
      {
        "ClassName": "遥感应用与专题制图",
        "Teacher": "陶超副教授",
        "Weeks": "9-16(周)",
        "Place": "C座204"
      },
      {
        "ClassName": "微波遥感",
        "Teacher": "汪长城副教授",
        "Weeks": "1-8(周)",
        "Place": "B座218"
      }
    ],
    [
      {
        "ClassName": "地理信息系统原理及应用（双语）",
        "Teacher": "邹滨副教授",
        "Weeks": "5-16(周)",
        "Place": "B座210"
      }
    ],
    部分内容省略...
  ]
}
账号或者密码错误:
{
  "StateCode": -1,
  "Error": "账号或者密码错误,请重新输入",
  "Class": []
}
```
#### 排名查询

`youwebsite/api/v1/jwc/:id/;pwd/rank [get]`

参数说明：

```
id:学号
pwd:教务系统密码
```

返回数据：

```
{
  "StateCode": 1,
  "Error": "",
  "Rank": [
    {
      "Term": "入学以来",
      "TotalScore": "139.5",
      "ClassRank": "7",
      "AverScore": "85.58"
    },
    {
      "Term": "2017-2018-1",
      "TotalScore": "23.5",
      "ClassRank": "6",
      "AverScore": "90.66"
    },
    {
      "Term": "2017",
      "TotalScore": "23.5",
      "ClassRank": "6",
      "AverScore": "90.66"
    },
    {
      "Term": "2016-2017-2",
      "TotalScore": "28",
      "ClassRank": "6",
      "AverScore": "86.88"
    },
	...省略部分内容
  ]
}
```

#### 校车查询

`youwebsite/api/v1/bus/search/:start/:end/:time`
参数说明：

```
start:起点
end:终点
time:出发时间
站点包括：['校本部图书馆前坪','南校区一教学楼前坪','升华学生公寓大门',
      	 '新校区教学楼D座南坪','新校区艺术楼','湘雅医学院老校区',
      	 '湘雅医学院新校区','湘雅医学院新校区大门','铁道校区办公楼前坪',
      	 '铁道校区图书馆前坪','科教新村','东塘']
时间包括：['7:00-7:59','8:00-8:59','9:00-9:59',
         '10:00-10:59','11:00-11:59','13:00-13:59',
         '14:00-14:59','15:00-15:59','17:00-17:59',
         '18:00-18:59','20:00-20:59']
```
返回数据说明：
```
{
  "StateCode": 1,
  "Error": "",
  "Buses": [
    {
      "StartTime": "7:30",
      "Start": "校本部图书馆前坪",
      "End": "新校区艺术楼",
      "RunTime": "周一至周五",
      "Num": "2",
      "Seat": "45",
      "Stations": [
        "校本部图书馆前坪",
        "升华学生公寓大门",
        "新校区教学楼D座南坪",
        "新校区艺术楼",
        "铁道校区图书馆前坪"
      ]
    },
    {
      "StartTime": "7:40",
      "Start": "校本部图书馆前坪",
      "End": "新校区艺术楼",
      "RunTime": "星期六",
      "Num": "1",
      "Seat": "26",
      "Stations": [
        "校本部图书馆前坪",
        "升华学生公寓大门",
        "新校区教学楼D座南坪",
        "新校区艺术楼"
      ]
    }
  ]
}
```
#### 招聘查询

`youwebsite/api/v1/job/:typeid/:pageindex/:pagesize/:hastime [get]`

参数说明：
```
type:招聘类型,1-本部招聘,2-湘雅招聘,3-铁道招聘,4-在线招聘,5-事业招考
pageindex:页码
pagesize:页面信息条数
hastime:是否包含招聘会时间：0-不包含,1-包含(会大大增加请求耗费时间)
```
返回数据说明:

```
{
  "StateCode": 1,
  "Error": "",
  "Jobs": [
    {
      "Link": "http://jobsky.csu.edu.cn/Home/ArticleDetails/10219",
      "Title": "明基BenQ友达集团",
      "Time": "2018.04.02",
      "Place": "中南大学校本部 立言厅"
    },
    {
      "Link": "http://jobsky.csu.edu.cn/Home/ArticleDetails/10231",
      "Title": "三七互娱2018春季校园招聘",
      "Time": "2018.03.28",
      "Place": "中南大学校本部科教南楼301"
    },
    {
      "Link": "http://jobsky.csu.edu.cn/Home/ArticleDetails/9983",
      "Title": "TCL多媒体科技控股有限公司",
      "Time": "2018.03.27",
      "Place": "中南大学校本部科教南楼 407"
    },
	...省略部分内容
  ]
}
```

### System Structure
项目采用beego的MVC模式,由于仅提供API服务,因此没有包含视图,每一个功能为一个模块,如controllers/cet.go和models/cet.go组成cet模块.

项目组成:

```
├── conf
│   └── app.conf
├── controllers
│   ├── bus.go
│   ├── cet.go
│   ├── classroom.go
│   ├── dangke.go
│   ├── default.go
│   ├── error.go
│   ├── job.go
│   ├── jwc.go
│   ├── lib.go
│   ├── news.go
│   └── wxuser.go
├── csugo
├── csugo.tar.gz
├── data
│   ├── classes.db //自习教室数据库
│   └── wxapp.db //用户信息数据库
├── logs
│   ├── project.2018-04-24.log
│   └── project.log
├── main.go
├── middleware
│   └── apiauth.go
├── models
│   ├── bus.go
│   ├── cet.go
│   ├── classroom.go
│   ├── dangke.go
│   ├── db.go
│   ├── job.go
│   ├── jwc.go
│   ├── lib.go
│   ├── news.go
│   └── wxuser.go
├── README.md
├── routers
│   ├── commentsRouter_controllers.go
│   └── router.go
├── static
│   └── js
│       └── reload.min.js
├── tests
│   └── default_test.go
├── utils
│   ├── errors.go
│   └── urls.go
└── views
    ├── errors
    │   └── 404.html
    └── index.html
```