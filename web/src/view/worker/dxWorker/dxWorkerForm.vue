<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="地址:" prop="address">
          <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
       </el-form-item>
        <el-form-item label="生日:" prop="birthday">
          <el-date-picker v-model="formData.birthday" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="createTime字段:" prop="createTime">
          <el-date-picker v-model="formData.createTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item label="deleteTime字段:" prop="deleteTime">
          <el-input v-model.number="formData.deleteTime" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="id:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="身份证号:" prop="idNumber">
          <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入身份证号" />
       </el-form-item>
        <el-form-item label="姓名:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
       </el-form-item>
        <el-form-item label="昵称:" prop="nickName">
          <el-input v-model="formData.nickName" :clearable="true"  placeholder="请输入昵称" />
       </el-form-item>
        <el-form-item label="手机号:" prop="phone">
          <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入手机号" />
       </el-form-item>
        <el-form-item label="备注:" prop="remarks">
          <el-input v-model="formData.remarks" :clearable="true"  placeholder="请输入备注" />
       </el-form-item>
        <el-form-item label="性别:" prop="sex">
           <el-select v-model="formData.sex" placeholder="请选择性别" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="状态:" prop="status">
          <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="总人数:" prop="totalPeople">
          <el-input v-model.number="formData.totalPeople" :clearable="true" placeholder="请输入" />
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
  createDxWorker,
  updateDxWorker,
  findDxWorker
} from '@/api/worker/dxWorker'

defineOptions({
    name: 'DxWorkerForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'

const route = useRoute()
const router = useRouter()

const type = ref('')
const genderOptions = ref([])
const formData = ref({
            address: '',
            birthday: new Date(),
            createTime: new Date(),
            deleteTime: undefined,
            id: undefined,
            idNumber: '',
            name: '',
            nickName: '',
            phone: '',
            remarks: '',
            sex: '',
            status: undefined,
            totalPeople: undefined,
            updateTime: new Date(),
        })
// 验证规则
const rule = reactive({
})

const elFormRef = ref()

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findDxWorker({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.redxWorker
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    genderOptions.value = await getDictFunc('gender')
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createDxWorker(formData.value)
               break
             case 'update':
               res = await updateDxWorker(formData.value)
               break
             default:
               res = await createDxWorker(formData.value)
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
