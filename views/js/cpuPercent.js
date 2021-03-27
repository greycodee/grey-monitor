var cpuPercentDom = document.getElementById('cpuPercent');
var cpuPercentChart = echarts.init(cpuPercentDom,'vintage');
var cpuPercentOption;

cpuPercentOption = {

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
cpuPercentChart.showLoading();
cpuPercentChart.setOption(cpuPercentOption);

// const host = window.location.host
// const host = "127.0.0.1:8988"
const cpuWS = new WebSocket("ws://"+host+"/service/ws/cpuPercentSingle")
cpuWS.onmessage = function (evt) {
    const received_msg = evt.data;

    const json = JSON.parse(received_msg);
    Math.ceil(json[0])
    cpuPercentChart.hideLoading();
    cpuPercentChart.setOption({
        series: [
            {
                data: [
                    {value:Math.ceil(json[0]),name: 'CPU使用率(%)'}
                ]
            }
        ]
    })
};