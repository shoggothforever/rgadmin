<template>
    <el-table :data="tableData" style="width: 100%">
      <el-table-column fixed prop="year" label="年份" />
      <el-table-column prop="month" label="月份" />
      <el-table-column fixed="right" label="下载" style="margin-left:auto !important;">
        <template #default="scope">
          <el-button link type="primary" size="large" @click="getWageDetil({'year': +scope.row.year, 'month': +scope.row.month})">下载</el-button>
        </template>
      </el-table-column>
    </el-table>
</template>

<script setup>
import { checkDetil } from '@/api/admin.js'
import { reactive } from 'vue'

const tableData = reactive([])

const getWageDetil = (detil) => {
  checkDetil(detil).then((res) => {
    const url = res.downloadUrl
    window.open(url)
  })
}

const currentYear = new Date().getFullYear();
const currentMonth = new Date().getMonth() + 1;

for (let year = currentYear - 2; year <= currentYear; year++) {
  for (let month = 1; month <= 12; month++) {
    if (year === currentYear && month > currentMonth) {
      break;
    }
    tableData.push({
      year: year.toString(),
      month: month.toString()
    });
  }
}
</script>

<style lang="scss">
::v-deep(.el-table td .cell) {
  line-height: 35px !important;
  font-size: 0.7rem !important;
}
</style>
