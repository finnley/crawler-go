# crawler-go

## 总体算法

### 发现用户

* 通过城市列表->城市->(下一页)->用户

* 通过用户->猜你喜欢

* 通过已有用户id+1来猜测用户id

### 实现步骤

单任务版 -> 并发版(多任务) -> 分布式

## 单任务版爬虫

* 获取并打印所有城市第一页用户的详细信息

### 乱码处理

下载包

go get -v golang.org/x/text

自动获取网站编码

go get -v golang.org/x/net/html

### 正则表达式

* css选择器

$('#cityList'')

$('.city-list')

$('.city-list>dd>a')

* 正则表达式

### 解析器 Parser

* 输入: utf-8编码的文本

* 输出: Request{URL,对应Parser}列表，Item列表