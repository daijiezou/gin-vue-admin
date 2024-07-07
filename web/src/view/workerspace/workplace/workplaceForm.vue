<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="日期:" prop="workOrderTime">
          <el-date-picker v-model="formData.workOrderTime" type="date" placeholder="选择日期" :clearable="false"></el-date-picker>
       </el-form-item>
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="关联的车:" prop="car">
          <el-input v-model.number="formData.car" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="createTime字段:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="deleteTime字段:" prop="deleteTime">
          <el-input v-model.number="formData.deleteTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工地名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入工地名称" />
       </el-form-item>
        <el-form-item label="工人人数:" prop="needWorkers">
          <el-input v-model.number="formData.needWorkers" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="工地拥有者身份证号码:" prop="ownerIdNumber">
          <el-input v-model="formData.ownerIdNumber" :clearable="true"  placeholder="请输入工地拥有者身份证号码" />
       </el-form-item>
        <el-form-item label="拥有者姓名:" prop="ownerName">
          <el-input v-model="formData.ownerName" :clearable="true"  placeholder="请输入拥有者姓名" />
       </el-form-item>
        <el-form-item label="联系人号码:" prop="ownerPhone">
          <el-input v-model="formData.ownerPhone" :clearable="true"  placeholder="请输入联系人号码" />
       </el-form-item>
        <el-form-item label="需求描述:" prop="requireDescription">
          <el-input v-model="formData.requireDescription" :clearable="true"  placeholder="请输入需求描述" />
       </el-form-item>
        <el-form-item label="updateTime字段:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createWorkplace,
  updateWorkplace,
  findWorkplace
} from '@/api/workerspace/workplace'

defineOptions({
    name: 'WorkplaceForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const formData = ref({
            workOrderTime: new Date(),
            address: '',
            car: undefined,
            createTime: new Date(),
            deleteTime: undefined,
            id: undefined,
            name: '',
            needWorkers: undefined,
            ownerIdNumber: '',
            ownerName: '',
            ownerPhone: '',
            requireDescription: '',
            updateTime: new Date(),
        })
// 验证规则
const rule = reactive({
               workOrderTime : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findWorkplace({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reworkplace
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createWorkplace(formData.value)
               break
             case 'update':
               res = await updateWorkplace(formData.value)
               break
             default:
               res = await createWorkplace(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
