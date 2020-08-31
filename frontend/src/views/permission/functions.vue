<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddOrUpdate({ id: 0})">New Role</el-button>

    <el-table
      :data="tableData"
      style="width: 100%;margin-bottom: 20px;"
      row-key="id"
      border
      v-loading="listLoading"
      lazy
      ref="treeElTable"
      :load="getFunctions"
      :tree-props="{children: 'children', hasChildren: 'hasChildren'}">
      <el-table-column
        prop="number"
        label="编号"
        sortable
        width="180">
      </el-table-column>
      <el-table-column
        prop="name"
        label="功能点"
        sortable
        width="180">
      </el-table-column>
      <el-table-column
        prop="path"
        label="路径">
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


    <el-dialog :visible.sync="dialogVisible"
               v-loading="listLoading"
               :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form  ref="functionForm"
                :model="functionForm"
                :rules="rules"
                label-width="120px"
                label-position="left">
        <el-form-item label="编号" prop="number">
          <el-input v-model.number="functionForm.number"  placeholder="编号" />
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
        <el-button type="danger" @click="cancelAddOrEdit">Cancel</el-button>
        <el-button type="primary" @click="addOrUpdateData">Confirm</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import path from 'path'
  import { deepClone } from '@/utils'
  import { deleteFunction, getFunctions, addOrUpdateFunction, getFunctionById } from '@/api/role'

  const defaultRole = {
    key: '',
    name: '',
    description: '',
    routes: []
  }

  export default {
    inject:['reload'],
    data() {
      let checkNumber = (rule, value, callback) => {
        //console.log(value)
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
          id: '',
          number: '',
          order:'',
          name: '',
          path:'',
          parentId:'',
        },
        defaultExpandedNodes:[],
        defaultSelectedNode:[],
        preParent:'',
        tableData:[],
        treeData:[],
        treeForm:'',
        forEdit:0,
        listLoading: false,
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


      // this.$nextTick(()=>{
      //   console.log(this.$refs)
      // })

    },
    activated() {

    },
    methods: {
      loadPage(){
        getFunctions(0).then(response => { //表内数据
          this.tableData = response.data
        })
      },
      initFormData(){

        // this.functionForm.id = 0
        // this.functionForm.parentId = 0
        // this.functionForm.name = ''
        // this.functionForm.number = ''
        // this.functionForm.order = ''
        // this.functionForm.path = ''
        // this.dialogVisible = true
        console.log('--------------------------------------')
        console.log(this.$refs)

        // this.$refs['functionForm'].resetFields();

        if(this.forEdit == 1) {//编辑数据
          getFunctionById(this.functionForm.id).then(response => {
            setTimeout(() => {
              this.functionForm = response.data
              let defaultNode = new Array(1);
              defaultNode[0] = response.data.parentId

              this.defaultExpandedNodes = response.data.parents
              this.defaultSelectedNode = defaultNode

            }, 1000)
          })
        }
        this.dialogVisible = true
        this.listLoading = false
      },
      initData(){ //初始化表内数据
        this.tableData = []
        this.treeData = []
        this.defaultExpandedNodes = []
        this.defaultSelectedNode = []
        this.treeForm = ''
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
        console.log(this.$refs)
        this.$refs['functionForm'].validate((valid) => {
          if (valid) {
            this.listLoading = true

            getFunctionById(this.functionForm.id).then(response => {
              setTimeout(() => {
                            //console.log('cuurrent parent................................')
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
      cancelAddOrEdit() {
              this.listLoading = false
              this.dialogVisible = false
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
      handleAddOrUpdate(row) {
        this.listLoading = true
        this.dialogVisible = false
        if ( row['id'] === 0 ){ //新增
          console.log('新增数据')
          //this.role = Object.assign({}, defaultRole)
          if (this.$refs.tree) {
            this.$refs.tree.setCheckedNodes([])
          }
          this.dialogType = 'new'
          this.forEdit = 0
          this.initFormData()
        } else { //修改
          console.log('修改数据')
          this.forEdit = 1
          this.functionForm.id = row['id']
          this.initFormData()
        }
      },
      handleDeleteConfirm(row) {
        console.log(row)
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
            done();
          })
          .catch(_ => {});
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
