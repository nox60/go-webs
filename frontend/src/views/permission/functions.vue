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

      <el-table-column label="操作" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button  type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button  size="mini" type="danger" @click="handleDeleteConfirm(row,$index)">
            删除
          </el-button>
        </template>
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
            :default-expanded-keys="defaultExpandedNodes"
            :default-checked-keys="defaultSelectedNode"
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
        defaultExpandedNodes:[],
        defaultSelectedNode:[],
        tableData:[],
        treeData:[],
        treeForm:'',
        forEdit:0,
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
      getFunctions(0).then(response => { //表内数据
        this.tableData = response.data
      })
    },
    methods: {
      initFormData(){

        if(this.forEdit == 0){//新增数据
          this.functionForm.parentId = 0
          this.functionForm.description = ''
          this.functionForm.name = ''
          this.functionForm.number = ''
          this.functionForm.order = ''
          this.functionForm.path = ''
        } else {//编辑数据
          this.defaultExpandedNodes = '[0,5]'
          this.defaultSelectedNode = '[4]'
        }


        // 循环请求节点的父级节点，将树渲染完成
        let rootNode = { id: 0, level: 0 }
       // this.getTreeNodes(rootNode)

        this.defaultExpandedNodes = '[0,5]'
        this.defaultSelectedNode = '[4]'

        // new Promise(function(resolve, reject){
        //   //做一些异步操作
        //   setTimeout(function(){
        //     console.log('处理根节点');
        //     //this.$options.methods.getTreeNodes(rootNode, resolve)
        //     console.log(this.treeNodes)
        //     console.log(this.treeData)
        //     console.log('---------------------------------------------------------')
        //   }, 2000);
        // });

        // let node1 = { id: 2, level: 1}
        // this.getTreeNodes( node1)

      },
      initData(){ //初始化表内数据
        getFunctions(0).then(response => {
          this.tableData = response.data
        })
      },
      getFunctions(tree, treeNode, resolve) { //用于懒加载表内数据
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
      getTreeNodes(node, resolve) { //新增OR修改菜单中获取树的下级节点数据
        console.log('---------------------0000')
        console.log(node)
        console.log('---------------------1111')

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

              this.initData()
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
        this.forEdit = 0
        this.initFormData()
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

      handleUpdate(row) {

        console.log(row['id'])

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
