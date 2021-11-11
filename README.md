# Breachkit - 一个漏洞快速扫描工具

## 🏸 What
Breachkit 的定位是一个漏洞扫描工具，它能够对漏洞进行快速扫描。Breachkit 对POC编写非常友好，提供Json Poc的解析以及多种逻辑的支持。

得力于Golang的并发优势，Breachkit 能够在极短的时间里完成对目标主机的漏洞扫描。

## 🏂 Run
Breachkit 对 Linux、MacOS、Windows 均提供了二进制可执行文件，前往 [Release](https://github.com/binganao/breachkit/releases) 下载对应版本即可运行:
```Bash
# Linux and MacOS
❯ chmod +x breachkit
❯ ./breachkit

# Windows
C:\Users\bingan\Desktop> breachkit.exe
```  

![](https://github.com/binganao/breachkit/blob/main/breachkit.png)

## 🐚 Options
```Bash
-np
  	Set all hosts to live [eg. -np]
  	设置所有主机为存活状态
-output string
  	Save the scan results to the specified file (default "output.txt")
  	设置一个输出文件
-target string
  	Specify a target [eg. -target https://example.com/file.php]
  	设置单个目标
-targets string
  	Reads all targets from the specified file [eg. -targets filename.txt]
  	从文件中获取多个目标
-timeout int
  	Response timeout time [eg. -timeout 10] (default 5)
  	超时设置
```

## 🐬 PocRule

Breachkit 提供简单的POC编写方式，在此之前，需要在 Breachkit 同目录中新建一个 "pocs" 目录，并在 "pocs" 目录下创建 all.json 文件。

> Breachkit 提供body、bodyor、bodyand、header、headeror、headerand、bodyheaderand七种验证逻辑，分别对返回包的头部和内容进行检测

```python
[
    {
        "Rank":"1",
        "Name":"Example 1",
        "Method":"POST",
        "Path":"/",
        "Header":null,
        "Body":"",
        "Logic":"bodyor",
        "StrVerity":{
            "BodyStr":["",""],
            "HeaderStr":null
        }
    },
    {
      "Rank":"1",
      "Name":"Example 2",
      "Method":"POST",
      "Path":"/",
      "Header":null,
      "Body":"",
      "Logic":"header",
      "StrVerity":{
        "BodyStr":null,
        "HeaderStr":{
          "os": "pwd"
        }
      }
    },
    {
      "Rank":"1",
      "Name":"Example 3",
      "Method":"GET",
      "Path":"/?cmd=pwd",
      "Header":null,
      "Body":"",
      "Logic":"header",
      "StrVerity":{
        "BodyStr":["/root/"],
        "HeaderStr":null
      }
    }
]
```

## 💊 Future

> 增加fofa支持