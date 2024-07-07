<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="地址" prop="address">
         <el-input v-model="searchInfo.address" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="生日" prop="birthday">
            
           <el-date-picker v-model="searchInfo.birthday" type="datetime" placeholder="搜索条件"></el-date-picker>

        </el-form-item>

        <template v-if="showAllQuery">
          <!-- 将需要控制显示状态的查询条件添加到此范围内 -->
        </template>

        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
          <el-button link type="primary" icon="arrow-down" @click="showAllQuery=true" v-if="!showAllQuery">展开</el-button>
          <el-button link type="primary" icon="arrow-up" @click="showAllQuery=false" v-else>收起</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="id"
        @selection-change="handleSelectionChange"
        @sort-change="sortChange"
        >
        <el-table-column type="selection" width="55" />
        
        <el-table-column align="left" label="地址" prop="address" width="120" />
         <el-table-column sortable align="left" label="年龄" prop="birthday" width="180">
            <template #default="scope">{{ formatBirthday(scope.row.birthday) }}</template>
         </el-table-column>
        <el-table-column align="left" label="身份证号" prop="idNumber" width="120" />
        <el-table-column align="left" label="姓名" prop="name" width="120" />
        <el-table-column align="left" label="昵称" prop="nickName" width="120" />
        <el-table-column align="left" label="手机号" prop="phone" width="120" />
        <el-table-column align="left" label="备注" prop="remarks" width="120" />
        <el-table-column sortable align="left" label="性别" prop="sex" width="120">
            <template #default="scope">
            {{ filterDict(scope.row.sex,genderOptions) }}
            </template>
        </el-table-column>
        <el-table-column align="left" label="状态" prop="status" width="120" />
        <el-table-column align="left" label="总人数" prop="totalPeople" width="120" />
        <el-table-column align="left" label="操作" fixed="right" min-width="240">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateDxWorkerFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-drawer destroy-on-close size="800" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'添加':'修改'}}</span>
                <div>
                  <el-button type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="地址:"  prop="address" >
              <el-input v-model="formData.address" :clearable="true"  placeholder="请输入地址" />
            </el-form-item>
            <el-form-item label="生日:"  prop="birthday" >
              <el-date-picker v-model="formData.birthday" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
            <el-form-item label="身份证号:"  prop="idNumber" >
              <el-input v-model="formData.idNumber" :clearable="true"  placeholder="请输入身份证号" />
            </el-form-item>
            <el-form-item label="姓名:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入姓名" />
            </el-form-item>
            <el-form-item label="昵称:"  prop="nickName" >
              <el-input v-model="formData.nickName" :clearable="true"  placeholder="请输入昵称" />
            </el-form-item>
            <el-form-item label="手机号:"  prop="phone" >
              <el-input v-model="formData.phone" :clearable="true"  placeholder="请输入手机号" />
            </el-form-item>
            <el-form-item label="备注:"  prop="remarks" >
              <el-input v-model="formData.remarks" :clearable="true"  placeholder="请输入备注" />
            </el-form-item>
            <el-form-item label="性别:"  prop="sex" >
              <el-select v-model="formData.sex" placeholder="请选择性别" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in genderOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="状态:"  prop="status" >
              <el-input v-model.number="formData.status" :clearable="true" placeholder="请输入状态" />
            </el-form-item>
            <el-form-item label="总人数:"  prop="totalPeople" >
              <el-input v-model.number="formData.totalPeople" :clearable="true" placeholder="请输入总人数" />
            </el-form-item>
          </el-form>
    </el-drawer>
  </div>
</template>

<script setup>
import {
  createDxWorker,
  deleteDxWorker,
  deleteDxWorkerByIds,
  updateDxWorker,
  findDxWorker,
  getDxWorkerList
} from '@/api/worker/dxWorker'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean,formatBirthday, filterDict ,filterDataSource, ReturnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

defineOptions({
    name: 'DxWorker'
})

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const genderOptions = ref([])
const formData = ref({
        address: '',
        birthday: new Date(),
        id: undefined,
        idNumber: '',
        name: '',
        nickName: '',
        phone: '',
        remarks: '',
        sex: '',
        status: undefined,
        totalPeople: undefined,
        })



// 验证规则
const rule = reactive({
})

const searchRule = reactive({
  createdAt: [
    { validator: (rule, value, callback) => {
      if (searchInfo.value.startCreatedAt && !searchInfo.value.endCreatedAt) {
        callback(new Error('请填写结束日期'))
      } else if (!searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt) {
        callback(new Error('请填写开始日期'))
      } else if (searchInfo.value.startCreatedAt && searchInfo.value.endCreatedAt && (searchInfo.value.startCreatedAt.getTime() === searchInfo.value.endCreatedAt.getTime() || searchInfo.value.startCreatedAt.getTime() > searchInfo.value.endCreatedAt.getTime())) {
        callback(new Error('开始日期应当早于结束日期'))
      } else {
        callback()
      }
    }, trigger: 'change' }
  ],
        birthday : [{ validator: (rule, value, callback) => {
        if (searchInfo.value.startBirthday && !searchInfo.value.endBirthday) {
          callback(new Error('请填写结束日期'))
        } else if (!searchInfo.value.startBirthday && searchInfo.value.endBirthday) {
          callback(new Error('请填写开始日期'))
        } else if (searchInfo.value.startBirthday && searchInfo.value.endBirthday && (searchInfo.value.startBirthday.getTime() === searchInfo.value.endBirthday.getTime() || searchInfo.value.startBirthday.getTime() > searchInfo.value.endBirthday.getTime())) {
          callback(new Error('开始日期应当早于结束日期'))
        } else {
          callback()
        }
      }, trigger: 'change' }],
})

const elFormRef = ref()
const elSearchFormRef = ref()

// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})
// 排序
const sortChange = ({ prop, order }) => {
  const sortMap = {
            birthday: 'birthday',
            sex: 'sex',
  }

  let sort = sortMap[prop]
  if(!sort){
   sort = prop.replace(/[A-Z]/g, match => `_${match.toLowerCase()}`)
  }

  searchInfo.value.sort = sort
  searchInfo.value.order = order
  getTableData()
}

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  elSearchFormRef.value?.validate(async(valid) => {
    if (!valid) return
    page.value = 1
    pageSize.value = 10
    getTableData()
  })
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getDxWorkerList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
    genderOptions.value = await getDictFunc('gender')
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteDxWorkerFunc(row)
        })
    }

// 多选删除
const onDelete = async() => {
  ElMessageBox.confirm('确定要删除吗?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning'
  }).then(async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.id)
        })
      const res = await deleteDxWorkerByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        getTableData()
      }
      })
    }

// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateDxWorkerFunc = async(row) => {
    const res = await findDxWorker({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteDxWorkerFunc = async (row) => {
    const res = await deleteDxWorker({ id: row.id })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}

// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        address: '',
        birthday: new Date(),
        id: undefined,
        idNumber: '',
        name: '',
        nickName: '',
        phone: '',
        remarks: '',
        sex: '',
        status: undefined,
        totalPeople: undefined,
        }
}
// 弹窗确定
const enterDialog = async () => {
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
                closeDialog()
                getTableData()
              }
      })
}

</script>

<style>

</style>
