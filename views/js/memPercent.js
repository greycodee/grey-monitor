var memPercentDom = document.getElementById('memPercent');
var memPercentChart = echarts.init(memPercentDom,'vintage');
var memPercentOption;

memPercentOption = {

    series: [{
        type: 'gauge',
        title:{
            show: true
        },
        progress: {
            show: true,
            width: 18
        },
        axisLine: {
            lineStyle: {
                width: 18
            }
        },
        axisTick: {
            show: false
        },
        splitLine: {
            length: 15,
            lineStyle: {
                width: 2,
                color: '#999'
            }
        },
        axisLabel: {
            distance: 25,
            color: '#999',
            fontSize: 20
        },
        anchor: {
            show: true,
            showAbove: true,
            size: 25,
            itemStyle: {
                borderWidth: 10
            }
        },

        detail: {
            valueAnimation: true,
            fontSize: 40,
            offsetCenter: [0, '70%'],
            formatter: '{value}%'

        }
    }]
};
memPercentChart.showLoading();
memPercentChart.setOption(memPercentOption);

const memWS = new WebSocket("ws://"+host+"/service/ws/memPercent")
memWS.onmessage = function (evt) {
    const received_msg = evt.data;
    memPercentChart.hideLoading();
    memPercentChart.setOption({
        series: [
            {
                data: [
                    {value:received_msg,name: '内存使用率(%)'}
                ]
            }
        ]
    })
};