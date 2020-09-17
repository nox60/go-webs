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
            v-model="roleForm.functions"
            show-checkbox
            node-key="id"
            check-strictly
            ref="treeForm"
            :default-expanded-keys="defaultExpandedNodes"
            :default-checked-keys="defaultSelectedNode"
            @check="handleClickNode"
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
import Pagination from '@/components/Pagination'

const defaultRole = {
  key: '',
  name: '',
  description: '',
  routes: []
}

export default {
  inject:['reload'],
  components: { Pagination },
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
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    handleClickNode (currentObj, treeStatus) {
      // 用于：父子节点严格互不关联时，父节点勾选变化时通知子节点同步变化，实现单向关联。
      let selected = treeStatus.checkedKeys.indexOf(currentObj.id) // -1未选
      // console.log('------------------------------->>>')
      // console.log(currentObj)
      // console.log(treeStatus)
      // console.log(selected)
      // console.log('===============================>>>')
      // 选中
      if (selected !== -1) {
        // 子节点只要被选中父节点就被选中
        this.selectedParent(currentObj)
        // 统一处理子节点为相同的勾选状态
        this.uniteChildSame(currentObj, true)
      } else {
        // 未选中 处理子节点全部未选中
        console.log(currentObj)

        if (currentObj.childs.length !== 0) {
          this.uniteChildSame(currentObj, false)
        }
      }
    },
    uniteChildSame (treeList, isSelected) {
      this.$refs.treeForm.setChecked(treeList.id, isSelected)
      if (treeList.children) {
        console.log('------------------------------------------??')
        console.log(treeList.children)
        console.log('------------------------------------------!!')

        for (let i = 0; i < treeList.children.length; i++) {
          this.uniteChildSame(treeList.children[i], isSelected)
        }
      }
    },
    // 统一处理父节点为选中
    selectedParent (currentObj) {
      let currentNode = this.$refs.treeForm.getNode(currentObj)
      // console.log('------------------------------->>')
      // console.log(this.$refs)
      // console.log(currentNode)
      // console.log('-------------------------------<<')
      if (currentNode.parent.key !== undefined) {
        this.$refs.treeForm.setChecked(currentNode.parent, true)
        this.selectedParent(currentNode.parent)
      }
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
    handleAddOrEditRole(row){
      this.listLoading = true
      this.dialogVisible = true
      if ( row.roleId === 0 ){ //新增
        console.log('新增数据')
        if (this.$refs.tree) {
          this.$refs.tree.setCheckedNodes([])
        }
        this.dialogType = 'new'
        this.forEdit = 0
      } else { //修改
        console.log('修改数据')
        this.forEdit = 1
        this.roleForm.id = row.roleId
      }
      this.$nextTick(()=>{
        this.initFormData()
      })
    },
    initFormData(){
      this.$refs['roleForm'].resetFields();
      if(this.forEdit == 1) {//编辑数据
        getRoleById(this.roleForm.id).then(response => {
          setTimeout(() => {
            this.roleForm = response.data
            this.defaultSelectedNode = response.data.functions
          }, 1000)
        })
      }
      this.listLoading = false
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
