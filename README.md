# segment_count
中文分词词频统计工具
## 功能
- 把mysql数据表中的源数据表的源字段进行分词，把词频统计结果存到目标数据表

## 应用场景
- 分析搜索短语中最热门的词

## 使用方式
- 修改config.go，配置数据库
- 参照 bin/main.go ，然后执行 go build segment_count ，创建入口文件
- 输入格式: segment_count -srctable 源数据表名 -srcfield 源数据表字段名 -disttable 目标数据表名
