<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddRole({ id: 0})">新建角色</el-button>
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
        <template slot-scope="scope">
          <el-button type="primary" size="small" @click="handleEdit(scope)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDelete(scope)">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form :model="role" label-width="80px" label-position="left">
        <el-form-item label="Name">
          <el-input v-model="role.name" placeholder="Role Name" />
        </el-form-item>
        <el-form-item label="Desc">
          <el-input
            v-model="role.description"
            :autosize="{ minRows: 2, maxRows: 4}"
            type="textarea"
            placeholder="Role Description"
          />
        </el-form-item>
        <el-form-item label="权限点" prop="">
          <el-tree
            :data="treeData"
            :props="treeNodes"
            :load="getTreeNodes"
            lazy
            show-checkbox
            node-key="id"
            check-strictly
            ref="treeForm"
            :default-expanded-keys="defaultExpandedNodes"
            :default-checked-keys="defaultSelectedNode"
            @check-change="handleNodeClick">
          </el-tree>

        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">取消</el-button>
        <el-button type="primary" @click="confirmRole">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import { getFunctions } from '@/api/role'

const defaultRole = {
  key: '',
  name: '',
  description: '',
  routes: []
}

export default {
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
      role: Object.assign({}, defaultRole),
      routes: [],
      rolesList: [],
      dialogVisible: false,
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
      // if (node.level === 0) {
      //   return resolve([{ id: 0, number: 0, order: 0, name: "根节点", path: "/", parentId: 0, hasChildren: true, leaf: false }]);
      // } else {
        getFunctions(node.data.id).then(response => {
          setTimeout(() => {
            // console.log(response.data)
            resolve(
              response.data
            )
          }, 1000)
        })
      //}
    },
    handleNodeClick(data, checked, node) {//新增OR修改权限点时点击树节点
      if(checked === true) {
        this.checkedId = data.id;
        this.$refs.treeForm.setCheckedKeys([data.id]);
        this.functionForm.parentId = data.id
      } else {
        if (this.checkedId == data.id) {
          this.$refs.treeForm.setCheckedKeys([data.id]);
        }
      }
    },
    confirmRole(){
      console.log('confirm')
    },
    handleAddRole(row){
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
        // this.initFormData()
        //this.dialogVisible = true

      } else { //修改
        console.log('修改数据')
        this.forEdit = 1
        this.functionForm.id = row['id']
      }

      this.$nextTick(()=>{
        console.log(this.$refs)
        //this.$refs.functionForm.resetFields();
        this.initFormData()
      })
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
