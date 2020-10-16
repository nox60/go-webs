<template>
  <div class="app-container">
    <el-button type="primary" @click="handleAddOrEditUser({ accountId: 0})">新建用户</el-button>
    <el-table
      :data="usersList"
      highlight-current-row
      v-loading="listLoading"
      style="width: 100%;margin-top:10px;"
      border>
      <el-table-column align="center" label="用户ID" width="220" prop="accountId">
      </el-table-column>
      <el-table-column align="center" label="用户名" width="220" prop="userName">
      </el-table-column>
      <el-table-column align="center" label="姓名" width="220" prop="realName">
      </el-table-column>
      <el-table-column align="center" label="Operations">
        <template slot-scope="{row,$index}">
          <el-button type="primary" size="small" @click="handleAddOrEditUser(row)">编辑</el-button>
          <el-button type="danger" size="small" @click="handleDeleteConfirm(row)">删除</el-button>
        </template>
      </el-table-column>

    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

    <el-dialog
      :visible.sync="dialogVisible"
      @close="handleClose"
      :title="dialogType==='edit'?'Edit Role':'New Role'">
      <el-form
        :model="userForm"
        ref="userForm"
        :modal-append-to-body='true'
        v-if='dialogVisible'
        label-width="100px"
        label-position="left">
        <el-form-item label="用户名" prop="userName">
          <el-input
            v-model="userForm.userName"
            placeholder="用户名" />
        </el-form-item>
        <el-form-item label="用户姓名" prop="realName">
          <el-input
            v-model="userForm.realName"
            placeholder="用户姓名"
          />
        </el-form-item>

        <el-form-item label="用户角色" prop="roles">
        <template>
            <el-checkbox
              v-model="userForm.roleIds"
              v-for="itemObj in allRoles"
              :label="itemObj.roleId"
              :key="itemObj.roleId"
              border
              size="mini">
              {{itemObj.name}}
            </el-checkbox>
        </template>
        </el-form-item>
      </el-form>
      <div style="text-align:right;">
        <el-button type="danger" @click="cancelAddOrEdit">取消</el-button>
        <el-button type="primary" @click="confirmAddOrUpdateUser">确认</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import {listUserData,addOrUpdateUser, deleteUser} from '@/api/user'
  import {listRoleData} from '@/api/role'
  import Pagination from '@/components/Pagination'

  const defaultRole = {
    key: '',
    name: '',
    description: '',
    routes: []
  }

  export default {
    inject: ['reload'],
    components: { Pagination },
    data() {
      return {
        forEdit: 0,
        userForm: {
          accountId: -1,
          userName: '',
          realName: '',
          roleIds: []
        },
        total: 0,
        listQuery: {
          page: 1,
          limit: 20,
          accountId: -1,
          importance: undefined,
          title: undefined,
          type: undefined,
          sort: '+id'
        },
        listLoading: false,
        dialogVisible: false,
        role: Object.assign({}, defaultRole),
        routes: [],
        usersList: [],
        allRoles: [],
        dialogType: 'new',
        checkStrictly: false
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
    },
    methods: {
      getList() {
        this.listLoading = true
        listUserData(this.listQuery).then(response => {
          this.usersList = response.data.dataLists
          this.total = response.data.totalCounts
          setTimeout(() => {
            this.listLoading = false
          }, 1.5 * 1000)
        })

        listRoleData({"page":1,"limit":999999,"sort":"+id"}).then(response => {
          this.allRoles = response.data.dataLists
        })
      },

      confirmAddOrUpdateUser() {
        this.listLoading = true
        console.log(this.userForm)
        addOrUpdateUser(this.userForm).then(() => {
          this.$notify({
            title: 'Success',
            message: '操作成功',
            type: 'success',
            duration: 2000
          })
          // this.initFormData()
          this.$refs['userForm'].resetFields()
          this.listLoading = false
          this.dialogVisible = false
          this.reload()
        })
      },
      handleAddOrEditUser(row) {
        this.listLoading = true
        if (row.accountId === 0) { // 新增
          console.log('新增数据')
          this.dialogType = 'new'
          this.forEdit = 0
        } else { // 修改
          console.log('修改数据')
          this.forEdit = 1
          this.userForm.accountId = row.accountId
          // getAllFuncs().then(response => {
          //   setTimeout(() => {
          //     this.tableData = response.data
          //   }, 1000)
          // })
        }
        this.$nextTick(() => {
          this.initFormData()
        })
      },
      handleClose() {
        console.log('call handleClose.....')
        this.$refs['userForm'].resetFields()
      },
      initFormData() {
        if (this.forEdit === 1) { // 编辑数据
          listUserData({"page":1,"limit":1,"accountId":this.userForm.accountId}).then(response => {
            setTimeout(() => {
              this.dialogVisible = true
              this.$nextTick(() => {
                this.$refs['userForm'].resetFields()
                this.userForm = response.data.dataLists[0]
                this.defaultSelectedNode = response.data.functions
                this.listLoading = false
              })
            }, 1000)
          })
        } else {
          this.dialogVisible = true
          this.$nextTick(() => {
            this.$refs['userForm'].resetFields()
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
            console.log(row.accountId)
            deleteUser(row.accountId).then(() => {
              this.dialogFormVisible = false
              this.$notify({
                title: 'Success',
                message: '删除用户成功！',
                type: 'success',
                duration: 2000
              })
              this.getList()
            })
            // done();
          })
          .catch(_ => {})
      },
      handleSelectFunction(checked, obj) {
        // var type=Object.prototype.toString.call(this.userForm.items);
        if (checked) {
          if (obj.parentIds) {
            obj.parentIds.forEach((item, index, array) => {
              // 要判断已经被选中，如果没有被选中才选中
              if (this.userForm.functions.indexOf(item) === -1 && item !== 0 && item !== -1) {
                this.userForm.functions.push(item)
              }
            })
          }
          // 参数：value数组中的当前项, index当前项的索引, array原始数组；
        } else {
          console.log(obj)

          if (obj.childIds) {
            obj.childIds.forEach((item, index, array) => {
              // 首先要判断该孩子节点是否已经被选中，如果已经被选择了，才撤销选中状态
              if (this.userForm.functions.indexOf(item) > -1) {
                // this.userForm.functions.push(item)
                const index = this.userForm.functions.indexOf(item)
                this.userForm.functions.splice(index, 1)
              }
            })
          }
          if (obj.childItems) {
            obj.childItems.forEach((item, index, array) => {
              // 首先要判断该孩子节点是否已经被选中，如果已经被选择了，才撤销选中状态
              if (this.userForm.items.indexOf(item) > -1) {
                // this.userForm.functions.push(item)
                const index = this.userForm.items.indexOf(item)
                this.userForm.items.splice(index, 1)
              }
            })
          }
        }
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
