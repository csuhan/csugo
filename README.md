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
    ......
    ]

账号或者密码错误:

{
  "StateCode": -1,
  "Error": "账号或者密码错误,请重新输入",
  "Grades": []
}

```

### System Structure
