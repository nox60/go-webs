<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddOrUpdateFunction({ id: 0, forEdit: false })">新建功能点</el-button>

    <el-table
      ref="treeElTable"
      v-loading="listLoading"
      :data="tableData"
      style="width: 100%;margin-bottom: 20px;margin-top:10px;"
      row-key="id"
      border
      lazy
      :load="getFunctions"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}"
    >
      <el-table-column
        prop="id"
        label="编号"
        width="180"
      />
      <el-table-column
        prop="order"
        label="序号（同级菜单内）"
        width="180"
      />
      <el-table-column
        prop="name"
        label="功能点"
        width="180"
      />
      <el-table-column
        prop="path"
        label="路径"
      />

      <el-table-column prop="items" label="页内功能点">
        <template slot-scope="scope">
          <el-button-group v-for="item in scope.row.items" :key="item.itemId">
            <el-button v-if="scope.row.type > 0" size="mini" icon="el-icon-delete" @click="handleDeleteFunctionItem(item.itemId)" />
            <!-- type>0 的菜单类型是用户自己定义的菜单项，可以编辑，反之则是系统核心类型，非常重要，不能编辑 -->
            <el-button v-if="scope.row.type > 0" size="mini" @click="handleAddOrUpdateFunctionItem({ itemId: item.itemId, functionId: scope.row.id })"> {{ item.itemName }} </el-button>
            <el-button v-else size="mini"> {{ item.itemName }} </el-button>
          </el-button-group>
          <el-button v-if="scope.row.type > 0" align="right" type="warning" size="mini" icon="el-icon-circle-plus-outline" @click="handleAddOrUpdateFunctionItem({ itemId: 0, functionId: scope.row.id })" />
        </template>
      </el-table-column>

      <el-table-column label="操作" prop="leaf" align="left" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button v-if="row.type > 0" type="primary" size="mini" @click="handleAddOrUpdateFunction(row)">
            编辑
          </el-button>
          <el-button
            v-if="row.type > 0"
            v-show="row.leaf"
            size="mini"
            type="danger"
            @click="handleDeleteConfirm(row,$index)"
          >
            删除
          </el-button>
        </template>
      </el-table-column>

    </el-table>

    <el-dialog
      v-loading="listLoading"
      :visible.sync="dialogVisible"
      :title="dialogType === 'edit'?'编辑功能点':'新建功能点'"
    >
      <el-form
        ref="functionForm"
        :model="functionForm"
        :rules="rules"
        label-width="120px"
        label-position="left"
      >
        <el-form-item label="编号" prop="id">
          <el-input v-model.number="functionForm.id" :disabled="functionForm.forEdit" placeholder="编号" />
        </el-form-item>
        <el-form-item label="菜单内次序" prop="order">
          <el-input v-model.number="functionForm.order" placeholder="菜单内次序，值越大越靠前" />
        </el-form-item>
        <el-form-item label="菜单名称" prop="name">
          <el-input v-model="functionForm.name" placeholder="Role Name" />
        </el-form-item>
        <el-form-item label="请求路径" prop="path">
          <el-input v-model="functionForm.path" placeholder="/path/1" />
        </el-form-item>

        <el-form-item label="父级菜单" prop="parentId">
          <el-tree
            ref="treeForm"
            :data="treeData"
            :props="treeNodes"
            :load="getTreeNodes"
            lazy
            show-checkbox
            node-key="id"
            check-strictly
            :default-expanded-keys="defaultExpandedNodes"
            :default-checked-keys="defaultSelectedNode"
            @check-change="handleNodeClick"
          />

        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="cancelAddOrEdit">取消</el-button>
        <el-button type="primary" @click="confirmAddOrUpdateFunctionData">确定</el-button>
      </div>
    </el-dialog>

    <el-dialog
      v-loading="listLoading"
      :visible.sync="addItemDialogVisible"
      :title="dialogType==='edit'?'Edit Role':'New Role'"
    >
      <el-form
        ref="itemForm"
        :model="itemForm"
        :rules="rulesForItem"
        label-width="120px"
        label-position="left"
      >
        <el-form-item label="功能点名称" prop="itemName">
          <el-input v-model="itemForm.itemName" placeholder="功能点名称" />
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="cancelAddOrEditItem">取消</el-button>
        <el-button type="primary" @click="confirmAddOrUpdateFunctionItem">确定</el-button>
      </div>
    </el-dialog>

  </div>
</template>

<script>
import {
  deleteFunction,
  deleteFunctionItem,
  getFunctions,
  addOrUpdateFunction,
  getFunctionById,
  addOrUpdateFunctionItem,
  addOrUpdateRole,
  getFunctionItemById
} from '@/api/role'
import { Message } from 'element-ui'

export default {
  inject: ['reload'],
  data() {
    const checkNumber = (rule, value, callback) => {
      // console.log(value)
      if (!Number.isInteger(value)) {
        callback(new Error('请输入数字值'))
      } else {
        if (value <= 0) {
          callback(new Error('必须大于0'))
        } else if (value > 100000) {
          callback(new Error('不能大于100000'))
        } else {
          callback()
        }
      }
    }
    return {
      typeOptions: [{
        value: 0,
        label: '菜单项目'
      }, {
        value: 1,
        label: '页内控件'
      }],
      functionForm: {
        id: 0,
        number: 0,
        order: 0,
        name: '',
        path: '',
        forEdit: false,
        parentId: 0
        // type: '菜单项目',
      },
      itemForm: {
        itemId: 0,
        itemName: '',
        itemNumber: 0,
        functionId: 0
      },
      defaultExpandedNodes: [],
      defaultSelectedNode: [],
      preParent: '',
      tableData: [],
      treeData: [],
      treeForm: '',
      forEdit: 0,
      listLoading: false,
      dialogVisible: false,
      addItemDialogVisible: false,
      dialogType: 'new',
      checkStrictly: false,
      defaultProps: {
        children: 'children',
        label: 'title'
      },
      rootParent: {
        id: 0
      },
      treeNodes: {
        label: 'name',
        children: 'zones',
        isLeaf: 'leaf'
      },
      rules: {
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        name: [{ required: true, message: '菜单名称不能为空！', trigger: 'blur' }],
        id: [
          { required: true, message: '功能点编号不能为空！', trigger: 'blur' },
          { validator: checkNumber, trigger: 'blur' }
        ],
        order: [
          { required: true, message: '序号值不能为空！', trigger: 'blur' },
          { validator: checkNumber, trigger: 'blur' }
        ]
      },
      rulesForItem: {
      }
    }
  },
  computed: {

  },
  created() {
    getFunctions(0).then(response => { // 表内数据
      this.tableData = response.data
    })
  },
  activated() {

  },
  methods: {
    loadPage() {
      getFunctions(0).then(response => { // 表内数据
        this.tableData = response.data
      })
    },
    initFormData() {
      if (this.functionForm.forEdit === true) { // 编辑数据
        getFunctionById(this.functionForm.id).then(response => {
          setTimeout(() => {
            this.functionForm = response.data
            this.functionForm.forEdit = true
            const defaultNode = new Array(1)
            defaultNode[0] = response.data.parentId

            this.defaultExpandedNodes = response.data.parents
            this.defaultSelectedNode = defaultNode
            this.listLoading = false
          }, 1000)
        })
      } else {
        this.$refs['functionForm'].resetFields()
        this.functionForm.forEdit = false
        const defaultNode = new Array(1)
        defaultNode[0] = 0
        this.defaultSelectedNode = defaultNode
        this.listLoading = false
      }
    },
    initFunctionItemFormData() {
      // this.$refs['itemForm'].resetFields();
      if (this.forEdit === 1) { // 编辑数据
        this.$nextTick(() => {
          // this.itemForm.id = jsonObject.itemId
          getFunctionItemById(this.itemForm.itemId).then(response => {
            setTimeout(() => {
              this.itemForm.itemName = response.data.itemName
              this.itemForm.itemId = response.data.itemId
              this.itemForm.itemNumber = response.data.itemNumber
              this.itemForm.functionId = response.data.functionId
              this.listLoading = false
            }, 1000)
          })
        })
      } else {
        this.listLoading = false
      }
    },
    getFunctions(tree, treeNode, resolve) { // 用于懒加载表内数据
      this.listLoading = true
      getFunctions(tree.id).then(response => {
        setTimeout(() => {
          resolve(
            response.data
          )
          this.listLoading = false
        }, 1000)
      })
    },
    getTreeNodes(node, resolve) { // 新增OR修改菜单中获取树的下级节点数据
      if (node.level === 0) {
        return resolve([{ id: 0, number: 0, order: 0, name: '根节点', path: '/', parentId: 0, hasChildren: true, leaf: false }])
      } else {
        getFunctions(node.data.id).then(response => {
          setTimeout(() => {
            resolve(
              response.data
            )
          }, 1000)
        })
      }
    },
    confirmAddOrUpdateFunctionData() {
      // console.log(this.$refs)
      this.$refs['functionForm'].validate((valid) => {
        if (valid) {
          this.listLoading = true
          console.log(this.functionForm)
          addOrUpdateFunction(this.functionForm).then(response => {
            if (response.data.resultCode !== 1) {
              Message({
                message: '错误，记录已经存在！请重新输入编号！',
                type: 'error',
                duration: 1000
              })
              this.listLoading = false
            } else {
              this.$notify({
                title: 'Success',
                message: '操作成功',
                type: 'success',
                duration: 2000
              })
              this.$refs['functionForm'].resetFields()
              this.reload()
            }
          })
        }
      })
    },
    confirmAddOrUpdateFunctionItem() {
      this.$refs['itemForm'].validate((valid) => {
        if (valid) {
          this.listLoading = true
          addOrUpdateFunctionItem(this.itemForm).then(() => {
            this.$notify({
              title: 'Success',
              message: '操作成功',
              type: 'success',
              duration: 2000
            })
            this.reload()
          })
        }
      })
    },
    confirmAddOrUpdateRole() {
      this.listLoading = true
      // this.roleForm.functions = functions
      // console.log(this.roleForm.id)
      addOrUpdateRole(this.roleForm).then(() => {
        this.$notify({
          title: 'Success',
          message: '操作成功',
          type: 'success',
          duration: 2000
        })
        this.initFormData()
        this.listLoading = false
        this.dialogVisible = false
        this.reload()
      })
    },
    cancelAddOrEdit() {
      this.listLoading = false
      this.dialogVisible = false
      this.$refs['functionForm'].resetFields()
    },
    cancelAddOrEditItem() {
      this.listLoading = false
      this.addItemDialogVisible = false
    },
    handleNodeClick(data, checked, node) { // 新增OR修改权限点时点击树节点
      if (checked === true) {
        this.checkedId = data.id
        this.$refs.treeForm.setCheckedKeys([data.id])
        this.functionForm.parentId = data.id
      } else {
        if (this.checkedId === data.id) {
          this.$refs.treeForm.setCheckedKeys([data.id])
        }
      }
    },
    handleAddOrUpdateFunctionItem(jsonObject) {
      this.itemForm.functionId = jsonObject.functionId
      this.listLoading = true
      this.addItemDialogVisible = true
      if (jsonObject.itemId === 0) { // 新增
        console.log('新增数据')
        this.forEdit = 0
      } else { // 修改
        console.log('修改数据')
        this.forEdit = 1
        this.itemForm.itemId = jsonObject.itemId
      }
      this.$nextTick(() => {
        this.initFunctionItemFormData()
      })
    },
    handleAddOrUpdateFunction(row) {
      this.listLoading = true
      this.dialogVisible = true
      this.functionForm.forEdit = row['forEdit']
      if (row['forEdit']) {
        // console.log('修改数据')
        this.dialogType = 'edit'
        this.functionForm.id = row['id']
      } else {
        // console.log('新增数据')
        this.dialogType = 'new'
        if (this.$refs.tree) {
          this.$refs.tree.setCheckedNodes([])
        }
      }
      this.$nextTick(() => {
        this.initFormData()
      })
    },
    handleDeleteConfirm(row) {
      // console.log(row)
      this.$confirm('确认删除？')
        .then(_ => {
          console.log('点击了确认')
          console.log(row['id'])
          deleteFunction(row['id']).then(() => {
            this.dialogVisible = false
            this.$notify({
              title: 'Success',
              message: '删除数据成功！',
              type: 'success',
              duration: 2000
            })
            this.reload()
          })
          // done();
        })
        .catch(_ => {})
    },
    handleDeleteFunctionItem(itemId) { // 新增OR修改权限点时点击树节点
      this.$confirm('确认删除？')
        .then(_ => {
          deleteFunctionItem(itemId).then(() => {
            this.dialogVisible = false
            this.$notify({
              title: 'Success',
              message: '删除数据成功！',
              type: 'success',
              duration: 2000
            })
            this.reload()
          })
          // done();
        })
        .catch(_ => {})
    },
    handleUpdateFunctionItem(itemId) {
      console.log(itemId)
    }
  }
}
</script>

<style lang="scss" scoped>
  .app-container {
  .roles-table {
    margin-top: 30px;
  }
  .permission-tree {
    margin-bottom: 30px;
  }
  }
</style>
