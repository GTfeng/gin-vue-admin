
<template>
  <div>
    <div class="gva-search-box">
      <el-form ref="elSearchFormRef" :inline="true" :model="searchInfo" class="demo-form-inline" :rules="searchRule" @keyup.enter="onSubmit">
        <el-form-item label="插件名称" prop="name">
         <el-input v-model="searchInfo.name" placeholder="搜索条件" />
        </el-form-item>
           <el-form-item label="漏洞类型" prop="vulType">
            <el-select v-model="searchInfo.vulType" clearable placeholder="请选择" @clear="()=>{searchInfo.vulType=undefined}">
              <el-option v-for="(item,key) in stringOptions" :key="key" :label="item.label" :value="item.value" />
            </el-select>
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
            <el-button  type="primary" icon="plus" @click="openDialog()">新增</el-button>
            <el-button  icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="onDelete">删除</el-button>
            
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
        
          <el-table-column align="left" label="插件名称" prop="name" width="120" />
          <el-table-column align="left" label="编码" prop="code" width="120" />
          <el-table-column align="left" label="插件描述" prop="description" width="120" />
        <el-table-column align="left" label="漏洞类型" prop="vulType" width="120">
            <template #default="scope">
                
                {{ filterDict(scope.row.vulType,stringOptions) }}
                
            </template>
        </el-table-column>
          <el-table-column align="left" label="产品类型" prop="proType" width="120" />
          <el-table-column align="left" label="测试目标" prop="objectives" width="120" />
          <el-table-column align="left" label="测试预期" prop="expectations" width="120" />
          <el-table-column align="left" label="测试步骤" prop="procedure" width="120" />
          <el-table-column align="left" label="危害分析" prop="hazard" width="120" />
          <el-table-column align="left" label="危害等级（0低级1中级2高级3超高级）" prop="level" width="120" />
          <el-table-column align="left" label="修复建议" prop="advise" width="120" />
         <el-table-column align="left" label="修改时间" prop="updateTime" width="180">
            <template #default="scope">{{ formatDate(scope.row.updateTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="操作" fixed="right" :min-width="appStore.operateMinWith">
            <template #default="scope">
            <el-button  type="primary" link class="table-button" @click="getDetails(scope.row)"><el-icon style="margin-right: 5px"><InfoFilled /></el-icon>查看</el-button>
            <el-button  type="primary" link icon="edit" class="table-button" @click="updatePayloadInfoFunc(scope.row)">编辑</el-button>
            <el-button   type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
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
    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="dialogFormVisible" :show-close="false" :before-close="closeDialog">
       <template #header>
              <div class="flex justify-between items-center">
                <span class="text-lg">{{type==='create'?'新增':'编辑'}}</span>
                <div>
                  <el-button :loading="btnLoading" type="primary" @click="enterDialog">确 定</el-button>
                  <el-button @click="closeDialog">取 消</el-button>
                </div>
              </div>
            </template>

          <el-form :model="formData" label-position="top" ref="elFormRef" :rules="rule" label-width="80px">
            <el-form-item label="id字段:"  prop="id" >
              <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入id字段" />
            </el-form-item>
            <el-form-item label="插件名称:"  prop="name" >
              <el-input v-model="formData.name" :clearable="true"  placeholder="请输入插件名称" />
            </el-form-item>
            <el-form-item label="编码:"  prop="code" >
              <el-input v-model="formData.code" :clearable="true"  placeholder="请输入编码" />
            </el-form-item>
            <el-form-item label="插件描述:"  prop="description" >
              <el-input v-model="formData.description" :clearable="true"  placeholder="请输入插件描述" />
            </el-form-item>
            <el-form-item label="漏洞类型:"  prop="vulType" >
              <el-select v-model="formData.vulType" placeholder="请选择漏洞类型" style="width:100%" :clearable="true" >
                <el-option v-for="(item,key) in stringOptions" :key="key" :label="item.label" :value="item.value" />
              </el-select>
            </el-form-item>
            <el-form-item label="产品类型:"  prop="proType" >
              <el-input v-model="formData.proType" :clearable="true"  placeholder="请输入产品类型" />
            </el-form-item>
            <el-form-item label="测试目标:"  prop="objectives" >
              <el-input v-model="formData.objectives" :clearable="true"  placeholder="请输入测试目标" />
            </el-form-item>
            <el-form-item label="测试预期:"  prop="expectations" >
              <el-input v-model="formData.expectations" :clearable="true"  placeholder="请输入测试预期" />
            </el-form-item>
            <el-form-item label="测试步骤:"  prop="procedure" >
              <el-input v-model="formData.procedure" :clearable="true"  placeholder="请输入测试步骤" />
            </el-form-item>
            <el-form-item label="危害分析:"  prop="hazard" >
              <el-input v-model="formData.hazard" :clearable="true"  placeholder="请输入危害分析" />
            </el-form-item>
            <el-form-item label="危害等级（0低级1中级2高级3超高级）:"  prop="level" >
              <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入危害等级（0低级1中级2高级3超高级）" />
            </el-form-item>
            <el-form-item label="修复建议:"  prop="advise" >
              <el-input v-model="formData.advise" :clearable="true"  placeholder="请输入修复建议" />
            </el-form-item>
            <el-form-item label="修改时间:"  prop="updateTime" >
              <el-date-picker v-model="formData.updateTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
            </el-form-item>
          </el-form>
    </el-drawer>

    <el-drawer destroy-on-close :size="appStore.drawerSize" v-model="detailShow" :show-close="true" :before-close="closeDetailShow" title="查看">
            <el-descriptions :column="1" border>
                    <el-descriptions-item label="id字段">
                        {{ detailFrom.id }}
                    </el-descriptions-item>
                    <el-descriptions-item label="插件名称">
                        {{ detailFrom.name }}
                    </el-descriptions-item>
                    <el-descriptions-item label="编码">
                        {{ detailFrom.code }}
                    </el-descriptions-item>
                    <el-descriptions-item label="插件描述">
                        {{ detailFrom.description }}
                    </el-descriptions-item>
                    <el-descriptions-item label="漏洞类型">
                        
                        {{ filterDict(detailFrom.vulType,stringOptions) }}
                        
                    </el-descriptions-item>
                    <el-descriptions-item label="产品类型">
                        {{ detailFrom.proType }}
                    </el-descriptions-item>
                    <el-descriptions-item label="测试目标">
                        {{ detailFrom.objectives }}
                    </el-descriptions-item>
                    <el-descriptions-item label="测试预期">
                        {{ detailFrom.expectations }}
                    </el-descriptions-item>
                    <el-descriptions-item label="测试步骤">
                        {{ detailFrom.procedure }}
                    </el-descriptions-item>
                    <el-descriptions-item label="危害分析">
                        {{ detailFrom.hazard }}
                    </el-descriptions-item>
                    <el-descriptions-item label="危害等级（0低级1中级2高级3超高级）">
                        {{ detailFrom.level }}
                    </el-descriptions-item>
                    <el-descriptions-item label="修复建议">
                        {{ detailFrom.advise }}
                    </el-descriptions-item>
                    <el-descriptions-item label="修改时间">
                        {{ detailFrom.updateTime }}
                    </el-descriptions-item>
            </el-descriptions>
        </el-drawer>

  </div>
</template>

<script setup>
import {
  createPayloadInfo,
  deletePayloadInfo,
  deletePayloadInfoByIds,
  updatePayloadInfo,
  findPayloadInfo,
  getPayloadInfoList
} from '@/api/example/payloadInfo'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict ,filterDataSource, returnArrImg, onDownloadFile } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import { useAppStore } from "@/pinia"




defineOptions({
    name: 'PayloadInfo'
})

// 提交按钮loading
const btnLoading = ref(false)
const appStore = useAppStore()

// 控制更多查询条件显示/隐藏状态
const showAllQuery = ref(false)

// 自动化生成的字典（可能为空）以及字段
const stringOptions = ref([])
const formData = ref({
            id: undefined,
            name: '',
            code: '',
            description: '',
            vulType: '',
            proType: '',
            objectives: '',
            expectations: '',
            procedure: '',
            hazard: '',
            level: undefined,
            advise: '',
            updateTime: new Date(),
        })



// 验证规则
const rule = reactive({
               name : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
               {
                   whitespace: true,
                   message: '不能只输入空格',
                   trigger: ['input', 'blur'],
              }
              ],
               level : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               },
              ],
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
  const table = await getPayloadInfoList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
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
    stringOptions.value = await getDictFunc('vulType')
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
            deletePayloadInfoFunc(row)
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
      const res = await deletePayloadInfoByIds({ ids })
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
const updatePayloadInfoFunc = async(row) => {
    const res = await findPayloadInfo({ id: row.id })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data
        dialogFormVisible.value = true
    }
}


// 删除行
const deletePayloadInfoFunc = async (row) => {
    const res = await deletePayloadInfo({ id: row.id })
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
        id: undefined,
        name: '',
        code: '',
        description: '',
        vulType: '',
        proType: '',
        objectives: '',
        expectations: '',
        procedure: '',
        hazard: '',
        level: undefined,
        advise: '',
        updateTime: new Date(),
        }
}
// 弹窗确定
const enterDialog = async () => {
     btnLoading.value = true
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return btnLoading.value = false
              let res
              switch (type.value) {
                case 'create':
                  res = await createPayloadInfo(formData.value)
                  break
                case 'update':
                  res = await updatePayloadInfo(formData.value)
                  break
                default:
                  res = await createPayloadInfo(formData.value)
                  break
              }
              btnLoading.value = false
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


const detailFrom = ref({})

// 查看详情控制标记
const detailShow = ref(false)


// 打开详情弹窗
const openDetailShow = () => {
  detailShow.value = true
}


// 打开详情
const getDetails = async (row) => {
  // 打开弹窗
  const res = await findPayloadInfo({ id: row.id })
  if (res.code === 0) {
    detailFrom.value = res.data
    openDetailShow()
  }
}


// 关闭详情弹窗
const closeDetailShow = () => {
  detailShow.value = false
  detailFrom.value = {}
}


</script>

<style>

</style>
