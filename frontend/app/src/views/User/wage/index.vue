<template>
    <div class="work-container">
        <el-form :model="form" label-width="120px">
            <el-form-item label="工时数">
                <el-input v-model="form.workTime" />
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="onSubmit">提交</el-button>
                <el-button>取消</el-button>
            </el-form-item>
        </el-form>
        <!--用element-plus的el-form写一个工资条,要求：第一行是一个标题，叫工资条，然后第二行开始是名称，月份，金额-->
        <el-table :data="wageList" style="width: 100%">
            <el-table-column
            label="工资条">
              <el-table-column
                prop="name"
                label="名称">
              </el-table-column>
              <el-table-column
                prop="month"
                label="月份">
              </el-table-column>
              <el-table-column
                prop="amountBeforeTax"
                label="税前工资">
              </el-table-column>
              <el-table-column
                prop="amount"
                label="实际工资">
              </el-table-column>
            </el-table-column>
        </el-table>
    </div>
</template>

<script lang="ts" setup>
import { reactive } from 'vue'
import { uploadWorkTime } from '../../../api/user';
import { ElMessage } from 'element-plus';

// do not use same name with ref
const form = reactive({
  workTime: ''
})

const wageList = reactive([
  {
    name: '',
    month: '',
    amount: '',
    amountBeforeTax: ''
  }
])

const onSubmit = () => {
  const { workTime } = form
  const dataTime = +workTime
  console.log(typeof dataTime)
  //   const tempTime = new Date(date1)
  //   const year = tempTime.getFullYear()
  //   const month = tempTime.getMonth() + 1
  uploadWorkTime({ workTime: dataTime }).then((res) => {
    wageList[0].name = res.name
    wageList[0].month = res.month
    wageList[0].amount = res.actualWage
    wageList[0].amountBeforeTax = res.wageBeforeTax
    ElMessage.success('提交成功')
  })
}
</script>

<style lang="scss" scoped>
.work-container {
  padding: 20px;
}
</style>
