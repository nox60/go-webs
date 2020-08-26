<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddRole">New Role</el-button>

    <el-table
      :data="tableData"
      style="width: 100%;margin-bottom: 20px;"
      row-key="id"
      border
      v-loading="listLoading"
      lazy
      :load="getFunctions"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
      <el-table-column
        prop="id"
        label="id"
        sortable
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="姓名"
        sortable
        width="180">
      </el-table-column>
      <el-table-column
        prop="path"
        label="路径">
      </el-table-column>
    </el-table>


    <el-dialog :visible.sync="dialogVisible" :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form  ref="functionForm" :model="functionForm" :rules="rules"  label-width="120px" label-position="left">
        <el-form-item label="编号" prop="number">
          <el-input v-model.number="functionForm.number"  placeholder="编号" />
        </el-form-item>
        <el-form-item label="菜单内次序">
          <el-input v-model.number="functionForm.order" placeholder="菜单内次序，值越大越靠前" />
        </el-form-item>
        <el-form-item label="菜单名称">
          <el-input v-model="functionForm.name" placeholder="Role Name" />
        </el-form-item>
        <el-form-item label="请求路径">
          <el-input v-model="functionForm.path" placeholder="abc" />
        </el-form-item>
        <el-form-item label="父级菜单">
          <el-tree
            :data="treeData"
            :props="treeNodes"
            :load="getTreeNodes"
            lazy
            show-checkbox
            node-key="id"
            check-strictly
            ref="treeForm"
            @check-change="handleNodeClick">
          </el-tree>

        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="dialogVisible=false">Cancel</el-button>
        <el-button type="primary" @click="addOrUpdateData">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import path from 'path'
  import { deepClone } from '@/utils'
  import { getRoutes, getRoles, addRole, deleteRole, updateRole, getFunctions, getNodes, addOrUpdateFunction } from '@/api/role'

  const defaultRole = {
    key: '',
    name: '',
    description: '',
    routes: []
  }

  export default {
    data() {
      let checkNumber = (rule, value, callback) => {
        console.log(value)
        if(!value) {
          return new Error('必填信息')
        }else {
          if(!Number.isInteger(value)) {
            callback(new Error('请输入数字值'))
          }else{
            if(value < 0) {
              callback(new Error('不能小于0'))
            }else if(value > 20){
              callback(new Error('不能大于20'))
            }else {
              callback()
            }
          }
        }
      };
      return {
        functionForm:{
          name: '',
          number: '',
          order:'',
          description: 'title',
          parentId:''
        },
        tableData:[],
        treeData:[],
        treeForm:'',
        listLoading: false,
        // role: Object.assign({}, defaultRole),
        // routes: [],
        // rolesList: [],
        dialogVisible: false,
        dialogType: 'new',
        checkStrictly: false,
        defaultProps: {
          children: 'children',
          label: 'title'
        },
        rootParent:{
          id: 0
        },
        treeNodes: {
          label: 'name',
          children: 'zones',
          isLeaf: 'leaf'
        },
        rules: {
          timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
          title: [{ required: true, message: 'title is required', trigger: 'blur' }],
          number: [
            {validator: checkNumber, trigger:'blur'}
          ],
        },
      }
    },
    computed: {

    },
    created() {
      getFunctions(0).then(response => {
        this.tableData = response.data
      })

      // let jsonstr = '{"id": 0, "number": 0, "order": 0, "name": "根节点", "path": "/", "parentId": 0, "hasChildren": true, "leaf": false}';
      // this.treeData = JSON.parse(jsonstr);
      // console.log(this.treeData)
    },
    methods: {
      getFunctions(tree, treeNode, resolve) {
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

      getTreeNodes(node, resolve) {

        if (node.level === 0) {
          return resolve([{ id: 0, number: 0, order: 0, name: "根节点", path: "/", parentId: 0, hasChildren: true, leaf: false }]);
        } else {
          getFunctions(node.data.id).then(response => {
            setTimeout(() => {
              // console.log(response.data)
              resolve(
                response.data
              )
            }, 1000)
          })
        }
      },

      addOrUpdateData() {
        this.$refs['functionForm'].validate((valid) => {
          if (valid) {
            this.listLoading = true
            addOrUpdateFunction(this.functionForm).then(() => {
              this.$notify({
                title: 'Success',
                message: '操作成功',
                type: 'success',
                duration: 2000
              })

              getFunctions(0).then(response => {
                this.tableData = response.data
              })

              this.listLoading = false
              this.dialogVisible = false

              // 调用全局挂载的方法,关闭当前标签页
              //this.$store.dispatch("tagsView/delView", 'mydata-createOrEdit');

              // // 返回上一步路由，返回上一个标签页
              // this.$router.go(-1);
              //
              //跳转回到列表界面
             // this.$router.push({path:'/mydatas/mydataList/'})
            })
          }
        })
      },

      handleAddRole() {
        this.role = Object.assign({}, defaultRole)
        if (this.$refs.tree) {
          this.$refs.tree.setCheckedNodes([])
        }
        this.dialogType = 'new'
        this.dialogVisible = true
      },
      handleNodeClick(data, checked, node) {
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
      async confirmRole() {

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
