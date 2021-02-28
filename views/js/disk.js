var diskDom = document.getElementById('disk');
var diskChart = echarts.init(diskDom,'vintage');
var diskOption;

diskOption = {
    title: {
        text: '磁盘使用情况',
        subtext: 'dd',
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
            name: '使用情况',
            type: 'pie',
            radius: '50%',
            data: [
                {value: 1048, name: '已使用'},
                {value: 735, name: '未使用'},
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

diskChart.setOption(diskOption);

// const diskUrl = "http://"+host+"/api/disk";
// http.open("GET",diskUrl)
// http.send()
// http.onreadystatechange=(e)=>{
//     const json = JSON.parse(http.responseText)
//     console.log(json)
// }

const diskUseUrl = "http://"+host+"/api/diskPath";
http.open("POST", diskUseUrl);
http.setRequestHeader('Content-Type', 'application/json');
var diskPath = {path:'/home'}
http.onreadystatechange=(e)=>{
    let j=http.response
    let json = JSON.parse(j)
    console.log(http.response)
    console.log(json.total)
    diskChart.setOption({
        title: {
            subtext: '磁盘: '+json.path
        },
        series:[{
            data:[
                {value: json.used, name: '已使用'},
                {value: json.free, name: '未使用'},
            ]
        }]
    });
}
http.send(JSON.stringify(diskPath))