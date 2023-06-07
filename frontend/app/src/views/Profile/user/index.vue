<template>
  <div class="echart-container">
    <h4>每月工资对比</h4>
    <div id="EchartPie"></div>
  </div>
</template>

<script>
export default {
  data() {
    return {

    }
  },
  methods: {
    changePie() {
      const myEchart = this.$echart.init(document.getElementById('EchartPie'))
      const option = {
        series: [
          {
            type: 'pie',
            data: [
              { value: 900, name: '一月' },
              { value: 895, name: '二月' },
              { value: 896, name: '三月' },
              { value: 870, name: '四月' },
              { value: 971, name: '五月' },
              { value: 990, name: '六月' },
              { value: 961, name: '七月' }
            ]
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
    this.changePie()
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

#EchartPie {
  padding: 2rem;
  height: 40rem;
  width: 60rem;
}
</style>
