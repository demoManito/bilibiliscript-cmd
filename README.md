## bilibiliscript cmd
> 基于 [bilibiliscript](https://github.com/demoManito/bilibiliscript) 项目的命令行工具

## 初始化
进入项目根目录执行命令：`make install`，即可编译完成，如有问题请提 issues

## 使用方法
- 默认：
```shell
bs -i <id> -c <cookies> -x <x-csrf> 
```
- 定时功能：
  - 参数说明：
    - -s, --starttime: 盖楼开始时间 
    - -e, --endtime: 盖楼结束时间
      
  - 案例：  
    ```shell
      bs -i <id> -c <cookies> -x <x-csrf> \
      -s '2022-07-22 18:00:00' -e '2022-07-22 19:00:00'
      ```


