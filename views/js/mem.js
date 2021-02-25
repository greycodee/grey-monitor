var chartDom = document.getElementById('main');
var myChart = echarts.init(chartDom);
var option;

option = {
    title: {
        text: '内存占用',
        subtext: '纯属虚构',
        left: 'center'
    },
    tooltip: {
        trigger: 'item'
    },
    legend: {
        orient: 'vertical',
        left: 'left',
    },
    series: [
        {
            name: '访问来源',
            type: 'pie',
            radius: '50%',
            data: [
                {value: 1048, name: '已使用内存'},
                {value: 735, name: '未使用内存'},

            ],
            emphasis: {
                itemStyle: {
                    shadowBlur: 10,
                    shadowOffsetX: 0,
                    shadowColor: 'rgba(0, 0, 0, 0.5)'
                }
            }
        }
    ]
};

option && myChart.setOption(option);
const host = window.location.host
const ws = new WebSocket("ws://"+host+"/ws")
ws.onmessage = function (evt) {
    const received_msg = evt.data;
    const json = JSON.parse(received_msg);

    myChart.setOption({
        series: [
            {
                data: [
                    {value:json.used,name: '已使用内存'},
                    {value:json.free,name: '剩余内存'},
                ]
            }
        ]
    })
};




