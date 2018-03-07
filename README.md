### Introduction
本项目是GO语言框架beego开发的简易API服务，提供中南大学校内的几个API接口服务，旨在练习GO WEB开发．

包括：
* 教务：成绩查询
* 教务：排名查询
* 教务：课表查询
* 招聘：招聘信息
* 校车：校车查询
### Usage
```
git clone https://github.com/csuhan/csugo.git

cd csugo

bee run

```

### Service
> 请在请求参数后加上token=csugo-token

如成绩查询:

`wget localhost:8080/api/v1/jwc/myid/mypassword/grade?token=csugo-token`

在请求地址结尾需要添加token

#### 成绩查询接口
`
youwebsite/api/v1/jwc/:id/:pwd/grade [get]
`

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
      "入学以来": {
        "TotalScore": "139.5",
        "ClassRank": "7",
        "AverScore": "85.58"
      }
    },
    {
      "2017-2018-1": {
        "TotalScore": "23.5",
        "ClassRank": "6",
        "AverScore": "90.66"
      }
    },
    {
      "2017": {
        "TotalScore": "23.5",
        "ClassRank": "6",
        "AverScore": "90.66"
      }
    },
    {
      "2016-2017-2": {
        "TotalScore": "28",
        "ClassRank": "6",
        "AverScore": "86.88"
      }
    },
    {
      "2016-2017-1": {
        "TotalScore": "31",
        "ClassRank": "7",
        "AverScore": "86.63"
      }
    },
    {
      "2016": {
        "TotalScore": "59",
        "ClassRank": "7",
        "AverScore": "86.75"
      }
    },
    {
      "2015-2016-2": {
        "TotalScore": "31",
        "ClassRank": "5",
        "AverScore": "83.26"
      }
    },
    {
      "2015-2016-1": {
        "TotalScore": "26",
        "ClassRank": "14",
        "AverScore": "81.12"
      }
    },
    {
      "2015": {
        "TotalScore": "57",
        "ClassRank": "8",
        "AverScore": "82.28"
      }
    }
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

`youwebsite/api/v1/job/:type/:pageindex/:pagesize/:hastime [get]`

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
    {
      "Link": "http://jobsky.csu.edu.cn/Home/ArticleDetails/10209",
      "Title": "上海泛微网络科技股份有限公司",
      "Time": "2018.03.23",
      "Place": "中南大学校本部科教南楼 206"
    },
    {
      "Link": "http://jobsky.csu.edu.cn/Home/ArticleDetails/10204",
      "Title": "北京新七天电子商务技术股份有限公司",
      "Time": "2018.03.22",
      "Place": "中南大学校本部科教北楼 207"
    }
  ]
}
```
### System Structure

项目组成:

```
├── conf
│   └── app.conf 配置文件
├── controllers　控制器
│   ├── bus.go　校车
│   ├── default.go　首页
│   ├── error.go　错误处理
│   ├── job.go　招聘
│   └── jwc.go　教务
├── logs　日志
│   └── project.log
├── main.go　入口文件
├── middleware　中间件
│   └── apiauth.go　token认证
├── models　模型
│   ├── bus.go　校车
│   ├── job.go　招聘
│   └── jwc.go　教务
├── routers 路由
│   ├── commentsRouter_controllers.go
│   └── router.go
├── static　静态文件
│   ├── css
│   ├── img
│   └── js
├── tests　测试
│   └── default_test.go
├── utils
│   └── errors.go
└── views　视图
    ├── errors
    │   └── 404.html
    └── index.html
```