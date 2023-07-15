<template>
  <div class="admin-container">
    <h1>Welcome!</h1>
    <h4>欢迎来到财务管理系统</h4>
    <p>您的身份是<strong style="color:#489ef7">管理员</strong></p>
    <p>用户姓名:{{ name }}</p>
    <br />
    <br />
    <br class="line" />
    <br />
    <br />
    <div id="EchartPie"></div>
  </div>
</template>

<script>
import { useUserStore } from '../../../stores/user'
import { defineComponent } from 'vue'
import { lookUp } from '../../../api/admin';
const userStore = useUserStore()
export default defineComponent({
  data() {
    return {
      name: userStore.name,
      userNumber: 0,
      allNumber: 0
    }
  },
  methods: {
    getUsersNumber() {
      return new Promise((resolve, reject) => {
        lookUp().then((res) => {
          this.userNumber = res.currentCount;
          this.allNumber = res.allCount
          resolve(true)
        })
          .catch((err) => {
            reject(err)
          })
      })
    },
    changePie() {
      const myEchart = this.$echart.init(document.getElementById('EchartPie'))
      const option = {
        tooltip: {
          trigger: 'item'
        },
        series: [
          {
            type: 'pie',
            label: {
              show: true,
              position: 'outside'
            },
            itemStyle: {
              borderRadius: 10,
              borderColor: '#fff',
              borderWidth: 2
            },
            labelLine: {
              show: true
            },
            data: [
              {
                value: this.userNumber,
                name: '已上报工时人数'
              },
              {
                value: this.allNumber - this.userNumber,
                name: '未上报工时人数'
              }
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
    this.getUsersNumber().then(() => {
      this.changePie()
    })
  }
})
</script>

<style lang="scss" scoped>
.admin-container {
    display: flex;
    justify-content: center;
    flex-direction: column;
    align-items: center;
    padding: 5rem 0;

    h1 {
        font-size: 1.5rem;
    }

    h4 {
        font-size: 1.2rem;
    }

    p {
        font-size: 1rem;
    }

    .user-number {
        color:#489ef7;
        border-top: 1px solid lightgray;
        width: 100%;
        display: flex;
        justify-content: center;
        padding: 3rem 0;
    }

    .all-number {
        color: #f76260;
    }

    .line {
      display: block; /* 将<br>标签显示为块级元素 */
      content: ""; /* 添加一个空内容 */
      border-top: 1px solid lightgray; /* 添加上边框，设置为浅灰色 */
      margin: 10px 0; /* 可选：为了给线条留出一些间距，可以调整边距值 */
    }
}
#EchartPie {
  padding: 2rem;
  height: 40rem;
  width: 60rem;
}
</style>
