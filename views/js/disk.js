var diskChartDom = document.getElementById('disk');
var disk = echarts.init(diskChartDom, 'dark');
var diskOption;

diskOption = {
    legend: {
        top: 'bottom'
    },
    title: {
        text: '硬盘信息',
        left: 'center',
        top: 20,
        textStyle: {
            color: '#ccc'
        }
    },
    toolbox: {
        show: true,
        feature: {
            mark: {show: true},
            dataView: {show: true, readOnly: true},
            restore: {show: false},
            saveAsImage: {show: false}
        }
    },
    series: [
        {
            name: '面积模式',
            type: 'pie',
            radius: [50, 150],
            center: ['50%', '50%'],
            roseType: 'area',
            itemStyle: {
                borderRadius: 8
            },
            data: [
                {value: 300, name: '已使用(M)'},
                {value: 90, name: '未使用(M)'}
            ]
        }
    ]
};

diskOption && disk.setOption(diskOption);

const http = new XMLHttpRequest()
const diskUrl = "http://"+host+"/api/disk";
http.open("GET",diskUrl)
http.send()
http.onreadystatechange=(e)=>{
    const json = JSON.parse(http.responseText)
    console.log(json)
    for (o in json){
        console.log(o)
        console.log(o)

    }

}