<template>
  <div class="echart-container">
    <h4>每月汇报工时人数</h4>
    <div id="EchartBar"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  methods: {
    changeBar() {
      const myEchart = this.$echart.init(document.getElementById('EchartBar'))
      const option = {
        xAxis: {
          data: ['一月', '二月', '三月', '四月', '五月', '六月', '七月']
        },
        yAxis: {},
        series: [
          {
            type: 'bar',
            data: [900, 895, 896, 870, 971, 990, 961]
          }
        ]
      }
      myEchart.setOption(option);
      window.addEventListener('resize', function() {
        myEchart.resize()
      })
    }
  },
  mounted() {
    this.changeBar()
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
#EchartBar {
  padding: 2rem;
  height: 40rem;
  width: 60rem;
}
</style>
