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
               @close="handleClose"
               :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form :model="roleForm"
               ref="roleForm"
               :modal-append-to-body='true'
               v-if='dialogVisible'
               label-width="100px"
               label-position="left">
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
<!--          <el-tree-->
<!--            :data="treeData"-->
<!--            :props="defaultProps"-->
<!--            v-model="roleForm.functions"-->
<!--            show-checkbox-->
<!--            default-expand-all-->
<!--            node-key="id"-->
<!--            ref="treeForm"-->
<!--            check-strictly-->
<!--            :default-checked-keys="defaultSelectedNode"-->
<!--            @check="handleClickNode"-->
<!--            >-->
<!--          </el-tree>-->


          <el-table
            :data="tableData"
            style="width: 100%;margin-bottom: 20px;margin-top:10px;"
            row-key="id"
            border
            default-expand-all
            :tree-props="{children: 'children', hasChildren: ''}"> /* 这里有坑*/
            <el-table-column
              prop="number"
              label="编号"
              width="180">
            </el-table-column>

            <el-table-column
              prop="name"
              label="功能点"
              width="180">
            </el-table-column>
            <el-table-column
              prop="path"
              label="路径">
            </el-table-column>

            <el-table-column prop="items" label="页内功能点">
              <template slot-scope="scope">
                <el-button-group v-for="item in scope.row.items"  >
                  <el-button  size="mini" icon="el-icon-delete" @click="handleDeleteFunctionItem(item.itemId)"></el-button>
                  <el-button  size="mini" @click="handleAddOrUpdateFunctionItem({ itemId: item.itemId, functionId: scope.row.id })">{{item.itemName}}</el-button>
                </el-button-group>
                <el-button align="right" type="warning" size="mini" icon="el-icon-circle-plus-outline" @click="handleAddOrUpdateFunctionItem({ itemId: 0, functionId: scope.row.id })">
                </el-button>
              </template>
            </el-table-column>

          </el-table>


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
import {addOrUpdateRole,
  getFunctions,
  listRoleData,
  deleteRole,
  getRoleById,
  getAllFuncs
} from '@/api/role'
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
      tableData:[
      ],
      treeForm:'',
      forEdit:0,
      defaultProps: {
        label: 'name',
        children: 'children',
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
  mounted() {
    this.changeCss();
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
      if (selected !== -1) {
        // 子节点只要被选中父节点就被选中
        this.selectedParent(currentObj)
        // 统一处理子节点为相同的勾选状态
        // this.uniteChildSame(currentObj, true)
      } else {
        // 未选中 处理子节点全部未选中
        if (currentObj.children != null && currentObj.children.length !== 0) {
          this.uniteChildSame(currentObj, false)
        }
      }
    },
    uniteChildSame (treeList, isSelected) {
      this.$refs.treeForm.setChecked(treeList.id, isSelected)
      if (treeList.children) {
        for (let i = 0; i < treeList.children.length; i++) {
          this.uniteChildSame(treeList.children[i], isSelected)
        }
      }
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
    // https://blog.csdn.net/qq_41612675/article/details/86612840
    // 横排样式

    // 统一处理父节点为选中
    selectedParent (currentObj) {
      console.log('父节点被选中')
      let currentNode = this.$refs.treeForm.getNode(currentObj)

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
      // this.dialogVisible = false
      // 获取所有功能点
      getAllFuncs().then(response => {
        setTimeout(() => {
          this.treeData = response.data
        }, 1000)
      })
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
        console.log('----------------------------------------------------------->>>')
        console.log(row)
        this.listLoading = true
        getAllFuncs().then(response => {
          setTimeout(() => {
            // this.tableData = response.data
            this.tableData =

              [{"id":1,"number":11,"order":33,"name":"esta","path":"/eksk/fa","parentId":0,"hasChildren":true,"leaf":false,"parents":null,"children":[{"id":2,"number":2,"order":3,"name":"aa","path":"s","parentId":1,"hasChildren":true,"leaf":false,"parents":null,"children":[{"id":3,"number":2,"order":22,"name":"asdf","path":"/test/case","parentId":2,"hasChildren":false,"leaf":true,"parents":null,"children":null,"items":[],"ItemStr":""},{"id":6,"number":2,"order":3,"name":"bb","path":"bbbb","parentId":2,"hasChildren":false,"leaf":true,"parents":null,"children":null,"items":[{"itemId":8,"itemName":"f11112222","itemNumber":0,"functionId":0}],"ItemStr":"8|!|f11112222"},{"id":7,"number":4,"order":3,"name":"2","path":"asdf","parentId":2,"hasChildren":false,"leaf":true,"parents":null,"children":null,"items":[],"ItemStr":""}],"items":[{"itemId":7,"itemName":"232222","itemNumber":0,"functionId":0}],"ItemStr":"7|!|232222"},{"id":5,"number":1,"order":2,"name":"a","path":"ss","parentId":1,"hasChildren":false,"leaf":true,"parents":null,"children":null,"items":[],"ItemStr":""}],"items":[{"itemId":6,"itemName":"qwf11111","itemNumber":0,"functionId":0}],"ItemStr":"6|!|qwf11111"},{"id":4,"number":2,"order":3,"name":"sadf","path":"/fkeka","parentId":0,"hasChildren":false,"leaf":true,"parents":null,"children":null,"items":[{"itemId":1,"itemName":"ffffffsssss","itemNumber":0,"functionId":0}],"ItemStr":"1|!|ffffffsssss"}]          }, 1000)



        })
      }
      this.$nextTick(()=>{
        this.initFormData()
      })
    },
    handleClose(){
      this.$refs['roleForm'].resetFields();
    },
    initFormData(){
      if(this.forEdit == 1) {//编辑数据
        getRoleById(this.roleForm.id).then(response => {
          setTimeout(() => {
            this.dialogVisible = true
            this.$nextTick(() => {
              this.$refs['roleForm'].resetFields();
              this.roleForm = response.data
              this.defaultSelectedNode = response.data.functions
              this.listLoading = false
            })
          }, 1000)
        })
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


    renderContent(h, { node, data, store }) {//树节点的内容区的渲染 Function
      console.log('------------------------>>>>>>>>>>>>>')
      let classname = "";
      // 由于项目中有三级菜单也有四级级菜单，就要在此做出判断
      console.log(node)
      console.log(node.isLeaf)
      if (node.isLeaf === true) {
        console.log('fffffffffffffffffffffffffffffffffffff')
        classname = "foo";
      }
      return h(
        "p",
        {
          class: 'foo'
        },
        node.label
      );
    },
    changeCss() {
      console.log('into change css method---------------------------------------------')
      var levelName = document.getElementsByClassName("foo"); // levelname是上面的最底层节点的名字
      console.log(levelName)
      console.log(levelName.length)
      for (var i = 0; i < levelName.length; i++) {
        console.log('------------------------------------------fund')
        // cssFloat 兼容 ie6-8  styleFloat 兼容ie9及标准浏览器
        levelName[i].parentNode.style.cssFloat = "left"; // 最底层的节点，包括多选框和名字都让他左浮动
        levelName[i].parentNode.style.styleFloat = "left";
        levelName[i].parentNode.onmouseover = function() {
          this.style.backgroundColor = "#fff";
        };
      }
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
