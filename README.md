## bilibiliscript cmd
> 基于 [bilibiliscript](https://github.com/demoManito/bilibiliscript) 项目的命令行工具

## 初始化
#### (1) 编译好的可执行文件
  - [编译好的可执行文件下载](https://github.com/demoManito/bilibiliscript-cmd/bin)
#### (2) 进入项目根目录执行命令
  1. `make install`
  2. 根据提示将路径写入到 `PATH` 中   
  3. `source ~/.zshrc` 或 `source ~/.bashrc` 即可编译完成

## 使用方法
### 盖楼脚本
- 默认：
  - 参数说明：
    - -i, --article_id: article business id
    - -c, --cookie: cookie in request header
    - -x, --xcsrf: x-csrf in request header
  - 例：
    ```shell
      bs run -i <id> -c <cookies> -x <x-csrf> 
      ```
- 定时功能：
  - 参数说明：
    - -s, --starttime: 盖楼开始时间 
    - -e, --endtime: 盖楼结束时间
      
  - 例：  
  
    ```shell
      bs run -i <id> -c <cookies> -x <x-csrf> \
      -s '2022-07-22 18:00:00' -e '2022-07-22 19:00:00'
      ```
    
### 楼层信息
- 参数说明：
  - -i, --article_id: article business id
  - -c, --cookie: cookie in request header
  - -x, --xcsrf: x-csrf in request header
  - -n, --floor_num: target floor num
- 例：
  ```shell
    bs floor -i <id> -c <cookies> -x <x-csrf> \
    -n 1000
    ```


