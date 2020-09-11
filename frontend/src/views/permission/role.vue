<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddOrEditRole({ roleId: 0})">新建角色</el-button>
    <el-table
      :data="rolesList"
      highlight-current-row
      v-loading="listLoading"
      style="width: 100%;margin-top:10px;"
      border>
      <el-table-column align="center" label="角色代码" width="220" prop="code">
      </el-table-column>
      <el-table-column align="center" label="角色名称" width="220" prop="name">
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

    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

    <el-dialog :visible.sync="dialogVisible"
               v-loading="listLoading"
               :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form :model="roleForm"
               ref="roleForm"
               label-width="100px" label-position="left">
        <el-form-item label="角色名称" prop="name">
          <el-input
            v-model="roleForm.name"
            placeholder="权限点名称" />
        </el-form-item>
        <el-form-item label="角色编码" prop="code">
          <el-input
            v-model="roleForm.code"
            placeholder="权限点编码"
          />
        </el-form-item>
        <el-form-item label="权限点" prop="functions">
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
        <el-button type="danger" @click="cancelAddOrEdit">取消</el-button>
        <el-button type="primary" @click="confirmAddOrUpdateRole">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {addOrUpdateRole, getFunctions, listRoleData, deleteRole, getRoleById} from '@/api/role'

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
      forEdit:0,
      treeNodes: {
        label: 'name',
        children: 'zones',
        isLeaf: 'leaf'
      },
      roleForm: {
        roleId: 0,
        name: '',
        code: '',
        functions: [],
      },
      total: 0,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
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
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      listRoleData(this.listQuery).then(response => {
        this.rolesList = response.data.dataLists
        this.total = response.data.totalCounts
        console.log(this.list)
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
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
      console.log(this.roleForm.id)
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
    handleAddOrEditRole(row){
      this.listLoading = true
      this.dialogVisible = true
      // console.log(row.roleId)
      if ( row.roleId === 0 ){ //新增
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
        this.roleForm.id = row.roleId
        // console.log(row['id'])
        // console.log(this.roleForm.id)
      }

      this.$nextTick(()=>{
        console.log(this.$refs)
        this.initFormData()
      })
    },
    initFormData(){
      console.log(this.$refs['roleForm'])
      this.$refs['roleForm'].resetFields();
      //console.log(this.roleForm.id)
      if(this.forEdit == 1) {//编辑数据
        //getRoleById
        getRoleById(this.roleForm.id).then(response => {
          setTimeout(() => {
            this.roleForm = response.data
            // let defaultNode = new Array(1);
            // defaultNode[0] = response.data.parentId

            // this.defaultExpandedNodes = response.data.parents
            // this.defaultSelectedNode = defaultNode
            this.listLoading = false
          }, 1000)
        })

      } else {
        console.log('reset fields')
        // this.$refs.roleForm.resetFields();
        // this.$refs['roleForm'].resetFields();
        this.listLoading = false

      }
    },
    cancelAddOrEdit() {
      this.listLoading = false
      this.dialogVisible = false
    },
    handleDeleteConfirm(row) {
      console.log(row)
      this.$confirm('确认删除？')
        .then(_ => {
          console.log('点击了确认')
          console.log(row.roleId)
          deleteRole(row.roleId).then(() => {
            this.dialogFormVisible = false
            this.$notify({
              title: 'Success',
              message: '删除数据成功！',
              type: 'success',
              duration: 2000
            })
            this.getList()
          })
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
