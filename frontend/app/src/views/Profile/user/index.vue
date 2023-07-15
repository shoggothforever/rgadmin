<template>
  <div class="echart-container">
    <h4>每月工资走势</h4>
    <div id="EchartCategory"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  methods: {
    changeCategory() {
      const myEchart = this.$echart.init(document.getElementById('EchartCategory'))
      const option = {
        tooltip: {
          trigger: 'item'
        },
        xAxis: {
          type: 'category',
          data: ['一月', '二月', '三月', '四月', '五月', '六月', '七月']
        },
        yAxis: {
          type: 'value'
        },
        series: [
          {
            data: [7000, 7500, 6700, 8000, 9000, 10000, 11000],
            type: 'line',
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            }
          }
        ]
      };
      myEchart.setOption(option);
      window.addEventListener('resize', function() {
        myEchart.resize()
      })
    }
  },
  mounted() {
    this.changeCategory()
  }
}
</script>

<style lang="scss" scoped>
.echart-container {
  flex-direction: column;
  display: flex;
  justify-content: center;
  align-items: center;
}

#EchartCategory {
  padding: 2rem;
  height: 40rem;
  width: 60rem;
}
</style>
