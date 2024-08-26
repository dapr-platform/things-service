#上传json格式

{
"name":"gw01",
"type":"gateway",
"model":"aaa",
"devices":[
{
"name":"dev1",
"type":"thsensor",
"model":"test001",
"attributes":[
{
"name":"temperature",
"tag":"ao1",
"value":$tagValue(ao1,%.2lf)
},
{
"name":"humidity",
"tag":"ao2",
"value":$tagValue(ao2,%.2lf)
}
]},
{
"name":"dev2",
"type":"sensor1",
"model":"test002",
"attributes":[
{
"name":"attr1",
"tag":"do1",
"value":$tagValue(test_dev2:do1,%.2lf)
}
]
}]
}


#发送消息

Command Topic: 选填项，指定用于接收命令的主题。从云服务端往该主题发布数据可以修改设备上的Tag点值，数据的格式如下示例，Tag点及点值由户自定义，写封包中可以没有时间戳数据（即ts）。此项不填则该设备将不会接受云服务修改点值的命令。

修改点值的封包举例如下：以下封包将会将AO_1的值写为12.88，AO_2的值写为18.76

{
    "w":[
        {
            "tag":"AO_1",
            "value":12.88
        },
        {
            "tag":"AO_2",
            "value":18.76
        }
    ],
    "ts":"2017-12-28T12:22:21+0000"
}
