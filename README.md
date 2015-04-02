# segment_count
中文分词词频统计工具
## 功能
- 对mysql数据表中的源数据表的源字段进行分词，然后把词频统计结果存到目标数据表

## 应用场景
- 分析搜索短语中最热门的词

## 使用说明
- 使用 github.com/huichen/sego 作为中文分词库
- 修改config.go，配置数据库
- 参照 bin/main.go 构建程序 segment_count
- 输入格式: segment_count -srctable 源数据表名 -srcfield 源数据表字段名 -disttable 目标数据表名
