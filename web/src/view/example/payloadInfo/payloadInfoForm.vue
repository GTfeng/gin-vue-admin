
<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="id字段:" prop="id">
          <el-input v-model.number="formData.id" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="插件名称:" prop="name">
          <el-input v-model="formData.name" :clearable="true"  placeholder="请输入插件名称" />
       </el-form-item>
        <el-form-item label="编码:" prop="code">
          <el-input v-model="formData.code" :clearable="true"  placeholder="请输入编码" />
       </el-form-item>
        <el-form-item label="插件描述:" prop="description">
          <el-input v-model="formData.description" :clearable="true"  placeholder="请输入插件描述" />
       </el-form-item>
        <el-form-item label="漏洞类型:" prop="vulType">
           <el-select v-model="formData.vulType" placeholder="请选择漏洞类型" style="width:100%" :clearable="true" >
              <el-option v-for="(item,key) in stringOptions" :key="key" :label="item.label" :value="item.value" />
           </el-select>
       </el-form-item>
        <el-form-item label="产品类型:" prop="proType">
          <el-input v-model="formData.proType" :clearable="true"  placeholder="请输入产品类型" />
       </el-form-item>
        <el-form-item label="测试目标:" prop="objectives">
          <el-input v-model="formData.objectives" :clearable="true"  placeholder="请输入测试目标" />
       </el-form-item>
        <el-form-item label="测试预期:" prop="expectations">
          <el-input v-model="formData.expectations" :clearable="true"  placeholder="请输入测试预期" />
       </el-form-item>
        <el-form-item label="测试步骤:" prop="procedure">
          <el-input v-model="formData.procedure" :clearable="true"  placeholder="请输入测试步骤" />
       </el-form-item>
        <el-form-item label="危害分析:" prop="hazard">
          <el-input v-model="formData.hazard" :clearable="true"  placeholder="请输入危害分析" />
       </el-form-item>
        <el-form-item label="危害等级（0低级1中级2高级3超高级）:" prop="level">
          <el-input v-model.number="formData.level" :clearable="true" placeholder="请输入" />
       </el-form-item>
        <el-form-item label="修复建议:" prop="advise">
          <el-input v-model="formData.advise" :clearable="true"  placeholder="请输入修复建议" />
       </el-form-item>
        <el-form-item label="修改时间:" prop="updateTime">
          <el-date-picker v-model="formData.updateTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
       </el-form-item>
        <el-form-item>
          <el-button :loading="btnLoading" type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script setup>
import {
  createPayloadInfo,
  updatePayloadInfo,
  findPayloadInfo
} from '@/api/example/payloadInfo'

defineOptions({
    name: 'PayloadInfoForm'
})

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'


const route = useRoute()
const router = useRouter()

// 提交按钮loading
const btnLoading = ref(false)

const type = ref('')
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
               }],
               level : [{
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
      const res = await findPayloadInfo({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
    stringOptions.value = await getDictFunc('string')
}

init()
// 保存按钮
const save = async() => {
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
