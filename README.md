# Ravigation
## 简介
Ravigation是一个AGV调度系统，其目标是实现一套通用、高效、可扩展的调度系统
## 系统构成
* ravigation： 可执行的调度系统，用来完成实际的AGV及控制节点的状态管理、任务调度、交通管制等任务，对外暴露api接口，便于其他程序的接入
* ravictl： 一个用于人机交互的终端程序，方便对ravigation进行状态查看、配置管理等操作
## 系统依赖
* redis: 用于数据的缓存及持久化存储（TODO: 持久化存储功能使用其他方案替换）
* emqtt: mqtt broker，与AGV及其他控制节点进行通信
## 目录结构
### ravigation：
* apiserver: 对外提供api接口
* service: 为api接口提供服务
* controllerManager: 通过MQTT与所有设备（AGV和控制节点）进行通信，并管理所有设备
* scheduler: 对接到的调度任务进行统一处理
* storage: 数据存储层，未来可扩展为其他存储
* utils: 工具包，包含配置读取以及其他通用
* config: 系统配置文件
* graph: 系统调度基于有向有权图的Dijkstra算法
## 系统设计
1. 设备连接信息： 用于保存agv及控制节点的连接信息，目前只支持ip及端口的设置
2. AGV状态： 用于同步agv的状态，通过设置超时来判断AGV是否连接，当发现不存在AGV状态时，通过mqtt发送设备连接信息
3. 地标信息： 通过唯一序号及坐标标识唯一地标信息，用户通过序号设置控制节点按钮的命令，系统发送给agv时自动规划路线并转换序号为坐标
4. 组命令: 组命令为一组json，未来通过增加导入文件方式来导入所有请求命令
```
{
    "CgName": "movetocmd1",     //组命令名称，作为组命令唯一标识，一个组命令中可包含多个子命令，子命令顺序执行
    "CmdList": [
        {
        "Type": "MoveToCmd",    //子命令类型，目前支持MoveToCmd（移动到某坐标）
        "Index": 1,             //目标地标序号
        "PointList": []         //系统通过寻路，将路径规划为坐标数组，发送给agv
        }
    ]
}
```
5. 控制节点按钮与组命令映射： 为了灵活支持各种控制设备（按钮盒或者移动终端等）,可将每个控制设备看做一个控制节点，该控制节点包含N个btn，通过btn与命令组名称的映射，在接到控制节点的btn消息时，通过映射查询，找到对应的组命令进行相应调度。
6. 控制节点: 每个控制节点可以包含多个btn，每个btn支持一组相应的组命令
## 终端命令
```
//添加设备连接信息
ravictl add connectInfo --type agv --name agv1 --host 127.0.0.1 --port 1234

//添加地标
ravictl add point -i 1 -x 1 -y 2 -z 3

//添加调度任务命令
ravictl add cmd -c '{"CgName":"movetocmd1","CmdList":[{"Type":"MoveToCmd","Index":1,"PointList":[]}]}'

//添加终端按钮命令映射
ravictl add btncmdmap -o node1 -b btn1 -m movetocmd1
```
## 协议
###1. mqtt接口协议
```
type: publish
topic： agv/connect/agv1
Qos： 0
payload:
describe: 周期性检查设备状态，如果发现设备状态不存在则发送设备连接信息

```
```
type: subscribe
topic： agv/state/agv1
Qos： 0
payload: {"Name":"agv1","Timestamp":123456,"X":1,"Y":2,"Z":3,"ActionStatus":2}
describe: 更新agv状态，ttl 20s

```
```
type: subscribe
topic： node/command/node1/btn1
Qos： 2
payload:
describe: 响应控制节点发出的请求命令，根据存储信息，查找控制节点按钮对应的调度请求，并将请求加入待调度队列，等待调度

```
```
type: subscribe
topic： agv/finish/agv1
Qos： 2
payload:
describe: agv当前调度命令执行完毕，将该调度命令从正在执行的队列中清除（TODO: 加入日志系统）

```

###2. api接口协议
TODO