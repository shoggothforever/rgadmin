<template>
    <div class="work-container">
        <el-form :model="form" label-width="120px">
            <el-form-item label="用户姓名">
                <el-input v-model="form.name" />
            </el-form-item>
            <el-form-item label="工时数">
                <el-input v-model="form.workTime" />
            </el-form-item>
            <el-form-item label="职位">
                <!-- 'admissions staff', 'assistant', 'cleaner', 'computer administrator', 'counsellor', 'driver', 'filing clerk', 'hr', 'librarian', 'principal', 'restaurant staff', 'security bureau', 'teacher', 'vice-principal' -->
                <el-select v-model="form.region" placeholder="请选择你的职位">
                    <el-option label="招生负责人" value="admissions staff" />
                    <el-option label="助理" value="assistant" />
                    <el-option label="清洁工" value="cleaner" />
                    <el-option label="电脑管理员" value="computer administrator" />
                    <el-option label="辅导员" value="counsellor" />
                    <el-option label="司机" value="driver" />
                    <el-option label="档案员" value="filing clerk" />
                    <el-option label="人力资源" value="hr" />
                    <el-option label="图书管理员" value="librarian" />
                    <el-option label="校长" value="principal" />
                    <el-option label="餐厅工作人员" value="restaurant staff" />
                    <el-option label="保安" value="security bureau" />
                    <el-option label="教师" value="teacher" />
                    <el-option label="副校长" value="vice-principal" />
                </el-select>
            </el-form-item>
            <el-form-item label="上报时间">
                <el-col :span="11">
                    <el-date-picker v-model="form.date1" type="date" placeholder="选择时间" style="width: 100%" />
                </el-col>
            </el-form-item>
            <el-form-item label="是否是编制">
                <el-switch v-model="form.delivery" />
            </el-form-item>
            <el-form-item label="详细描述">
                <el-input v-model="form.desc" type="textarea" />
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
  name: '',
  workTime: '',
  region: '',
  date1: '',
  delivery: false,
  desc: ''
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
