<template>
  <div>
    <el-form :model="form" label-width="120px">
        <el-form-item label="用户姓名">
            <el-input v-model="form.name" />
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
        <el-form-item label="查询全年工资">
            <el-switch v-model="form.delivery" />
        </el-form-item>
        <el-form-item>
            <el-button type="primary" @click="onSubmit">提交</el-button>
            <el-button>取消</el-button>
        </el-form-item>
    </el-form>
    <el-table :data="wageList" style="width: 100%">
        <el-table-column
        label="查询工资列表">
            <el-table-column
            prop="wageBeforeTax"
            label="税前工资">
            </el-table-column>
            <el-table-column
            prop="tax"
            label="收税">
            </el-table-column>
            <el-table-column
            prop="actualWage"
            label="实际工资">
            </el-table-column>
        </el-table-column>
    </el-table>
  </div>
</template>

<script setup>
import { reactive } from 'vue'
import { queryWage } from '@/api/user.js'

const form = reactive({
  name: '',
  region: '',
  date1: '',
  delivery: false,
  desc: ''
})

var wageList = reactive([])

const onSubmit = () => {
  const tmpData = new Date(form.date1)
  const year = tmpData.getFullYear()
  var month = tmpData.getMonth() + 1
  if (form.delivery) {
    month = -1;
  }
  console.log(year, month)
  queryWage({ year: year, month: month }).then((res) => {
    wageList.push(res)
  })
}
</script>

<style lang="scss" scoped>

</style>
