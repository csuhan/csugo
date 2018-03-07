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

`wget localhost:8080/jwc/myid/mypassword/grade?token=csugo-token`

在请求地址结尾需要添加token

#### 成绩查询接口
`
/jwc/:id/:pwd/grade [get]
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

`/jwc/myid/mypassword/class/term/week [get]`

参数说明:

```
id:学号
pwd:教务系统密码
term:学期．如：2017-2018-2
week:周次．0表示所有周，其他表示周次
```

返回说明：

｀｀｀
返回数据中Class为二维数组，排列方式为:

        周一　周二　　周三　周四　周五　周六　周日
  1-2    1    2     3    4    5    6    7

  3-4    8    9    10    ................

  5-6    ................................

  7-8    ................................

 9-10    ................................

11-12    ...........................   42
｀｀｀


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
```

### System Structure
