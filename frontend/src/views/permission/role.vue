<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddOrEditRole({ id: 0})">新建角色</el-button>
    <el-table
      :data="rolesList"
      style="width: 100%;margin-top:10px;" border>
      <el-table-column align="center" label="角色代码" width="220" prop="">

      </el-table-column>
      <el-table-column align="center" label="角色名称" width="220">

      </el-table-column>
      <el-table-column align="header-center" label="角色说明">

      </el-table-column>
      <el-table-column align="header-center" label="角色状态">

      </el-table-column>
      <el-table-column align="center" label="Operations">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="small" @click="handleAddOrEditRole(row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDeleteConfirm(row)">删除</el-button>
        </template>
      </el-table-column>

      <el-table-column label="操作" prop="leaf" align="left"  width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button  type="primary" size="mini" @click="handleAddOrUpdate(row)">
            编辑
          </el-button>
          <el-button
            size="mini"
            v-show="row.leaf"
            type="danger"
            @click="handleDeleteConfirm(row,$index)">
            删除
          </el-button>
        </template>
      </el-table-column>

    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form :model="roleForm"
               ref="roleForm"
               label-width="100px" label-position="left">
        <el-form-item label="角色名称">
          <el-input
            v-model="roleForm.name"
            placeholder="权限点名称" />
        </el-form-item>
        <el-form-item label="角色编码">
          <el-input
            v-model="roleForm.code"
            placeholder="权限点编码"
          />
        </el-form-item>
        <el-form-item label="权限点" prop="">
          <el-tree
            :data="treeData"
            :props="treeNodes"
            :load="getTreeNodes"
            v-model="roleForm.functions"
            lazy
            show-checkbox
            node-key="id"
            check-strictly
            ref="treeForm"
            :default-expanded-keys="defaultExpandedNodes"
            :default-checked-keys="defaultSelectedNode"
            >
          </el-tree>

        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="confirmAddOrUpdateRole">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {addOrUpdateRole, getFunctionById, getFunctions} from '@/api/role'

const defaultRole = {
  key: '',
  name: '',
  description: '',
  routes: []
}

export default {
  inject:['reload'],
  data() {
    return {
      defaultExpandedNodes:[],
      defaultSelectedNode:[],
      treeData:[],
      treeForm:'',
      treeNodes: {
        label: 'name',
        children: 'zones',
        isLeaf: 'leaf'
      },
      roleForm: {
        name: '',
        code: '',
        functions: [],
      },
      listLoading: false,
      dialogVisible: false,
      role: Object.assign({}, defaultRole),
      routes: [],
      rolesList: [],
      dialogType: 'new',
      checkStrictly: false,
      defaultProps: {
        children: 'children',
        label: 'title'
      }
    }
  },
  computed: {
    routesData() {
      return this.routes
    }
  },
  created() {

  },
  methods: {

    getTreeNodes(node, resolve) { //新增OR修改菜单中获取树的下级节点数据
        getFunctions(node.data.id).then(response => {
          setTimeout(() => {
            // console.log(response.data)
            resolve(
              response.data
            )
          }, 1000)
        })
    },

    confirmAddOrUpdateRole(){
      let functions = new Array()

      this.$refs.treeForm.getCheckedNodes().forEach((data, index, array) => {
        functions[index] = data.id
      });
      this.listLoading = true
      this.roleForm.functions = functions
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

    addOrUpdateData() {
      console.log(this.$refs)
      this.$refs['roleForm'].validate((valid) => {
        if (valid) {
          this.listLoading = true

          getFunctionById(this.functionForm.id).then(response => {
            setTimeout(() => {
              this.preParent = response.data.parentId
              addOrUpdateFunction(this.functionForm).then(() => {
                this.$notify({
                  title: 'Success',
                  message: '操作成功',
                  type: 'success',
                  duration: 2000
                })

                this.listLoading = false
                this.dialogVisible = false
                this.functionForm.id = 0
                this.functionForm.parentId = 0
                this.functionForm.name = ''
                this.functionForm.number = ''
                this.functionForm.order = ''
                this.functionForm.path = ''

                this.reload()
              })
            }, 1000)
          })
        }
      })
    },

    handleAddOrEditRole(row){
      this.listLoading = true
      this.dialogVisible = true
      if ( row['id'] === 0 ){ //新增
        console.log('新增数据')
        //this.role = Object.assign({}, defaultRole)
        if (this.$refs.tree) {
          this.$refs.tree.setCheckedNodes([])
        }
        this.dialogType = 'new'
        this.forEdit = 0
      } else { //修改
        console.log('修改数据')
        this.forEdit = 1
        this.functionForm.id = row['id']
      }

      this.$nextTick(()=>{
        console.log(this.$refs)
        this.initFormData()
      })
    },

    initFormData(){
      console.log('reset fields')
      this.$refs.roleForm.resetFields();
      if(this.forEdit == 1) {//编辑数据

      } else {

      }
    },

    handleDeleteConfirm(row) {
      console.log(row)
      this.$confirm('确认删除？')
        .then(_ => {
          console.log('点击了确认')
          console.log(row['id'])

          done();
        })
        .catch(_ => {});
    },

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
