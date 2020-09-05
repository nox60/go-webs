<template>
  <div class="app-container">
    <div class="filter-container">
      <el-input v-model="listQuery.title" placeholder="Title" style="width: 200px;" class="filter-item" @keyup.enter.native="handleFilter" />
      <el-select v-model="listQuery.importance" placeholder="Imp" clearable style="width: 90px" class="filter-item">
        <el-option v-for="item in importanceOptions" :key="item" :label="item" :value="item" />
      </el-select>
      <el-select v-model="listQuery.type" placeholder="Type" clearable class="filter-item" style="width: 130px">
        <el-option v-for="item in typeValuesArray" :key="item.key" :label="item.display_name+'('+item.key+')'" :value="item.key" />
      </el-select>
      <el-select v-model="listQuery.sort" style="width: 140px" class="filter-item" @change="handleFilter">
        <el-option v-for="item in sortOptions" :key="item.key" :label="item.label" :value="item.key" />
      </el-select>
      <el-button v-waves class="filter-item" type="primary" icon="el-icon-search" @click="handleFilter">
        Search
      </el-button>
      <el-button class="filter-item" style="margin-left: 10px;" type="primary" icon="el-icon-edit" @click="handleAdd">
          新建文章
      </el-button>
      <el-button v-waves :loading="downloadLoading" class="filter-item" type="primary" icon="el-icon-download" @click="handleDownload">
        Export
      </el-button>
      <el-checkbox v-model="showReviewer" class="filter-item" style="margin-left:15px;" @change="tableKey=tableKey+1">
        reviewer
      </el-checkbox>
    </div>

    <el-table
      :key="tableKey"
      v-loading="listLoading"
      :data="list"
      border
      fit
      highlight-current-row
      style="width: 100%;"
      @sort-change="sortChange"
    >
      <el-table-column label="ID" prop="id" sortable="custom" align="center" width="80" :class-name="getSortClass('id')">
        <template slot-scope="{row}">
          <span>{{ row.itemId }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Date" width="150px" align="center">
        <template slot-scope="{row}">
          <span>{{ row.createTime | parseTime('{y}-{m}-{d}') }}</span>
        </template>
      </el-table-column>
      <el-table-column label="Title" min-width="150px">
        <template slot-scope="{row}">
          <span>{{ row.itemTitle }}</span>
        </template>
      </el-table-column>
      <el-table-column label="类型" width="110px" align="center" :formatter="formatType">

      </el-table-column>
<!--      <el-table-column v-if="showReviewer" label="Reviewer" width="110px" align="center">-->
<!--        <template slot-scope="{row}">-->
<!--          <span style="color:red;">{{ row.reviewer }}</span>-->
<!--        </template>-->
<!--      </el-table-column>-->
<!--      <el-table-column label="Imp" width="80px">-->
<!--        <template slot-scope="{row}">-->
<!--          <svg-icon v-for="n in + row.importance" :key="n" icon-class="star" class="meta-item__icon" />-->
<!--        </template>-->
<!--      </el-table-column>-->
      <el-table-column label="星级" align="center" width="95">
        <template slot-scope="scope" >
          <!-- <el-rate v-model="scope.row.evaValue" :allow-half="true"  disabled show-score text-color="#ff9900" score-template="{value}"></el-rate> -->
          <el-rate v-model="scope.row.itemStar" :allow-half="true"  disabled text-color="#ff9900"></el-rate>
        </template>
      </el-table-column>
      <el-table-column label="状态" class-name="status-col" width="100" :formatter="statusType">

      </el-table-column>
      <el-table-column label="Actions" align="center" width="230" class-name="small-padding fixed-width">
        <template slot-scope="{row,$index}">
          <el-button v-if="row.itemStatus=='0'" type="primary" size="mini" @click="handleUpdate(row)">
            编辑
          </el-button>
          <el-button v-if="row.itemStatus=='0'" size="mini" type="success" @click="handleModifyStatus(row,'published')">
            申请审核
          </el-button>
          <el-button v-if="row.itemStatus=='0'" size="mini" type="danger" @click="handleDeleteConfirm(row,$index)">
            删除
          </el-button>
        </template>
      </el-table-column>
    </el-table>

    <pagination v-show="total>0" :total="total" :page.sync="listQuery.page" :limit.sync="listQuery.limit" @pagination="getList" />

    <el-dialog :title="textMap[dialogStatus]" :visible.sync="dialogFormVisible">
      <el-form ref="dataForm" :rules="rules" :model="itemForm" label-position="left" label-width="100px" style="width: 400px; margin-left:50px;">
        <el-form-item label="类别" prop="itemType">
          <el-select v-model="itemForm.itemType" class="filter-item" placeholder="Please select">
            <el-option v-for="item in typeValuesArray" :key="item.typeValue" :label="item.typeName" :value="item.typeValue" />
          </el-select>
        </el-form-item>
        <el-form-item label="createTime" prop="createTime">
          <el-date-picker v-model="itemForm.createTime" type="date" value-format="yyyy-MM-dd" placeholder="Please pick a date" />
        </el-form-item>
        <el-form-item label="Title" prop="itemTitle">
          <el-input v-model="itemForm.itemTitle" />
        </el-form-item>
        <el-form-item label="Status">
          <el-select v-model="itemForm.itemStatus" class="filter-item" placeholder="Please select">
            <el-option v-for="item in statusOptions" :key="item.statusValue" :label="item.statusName" :value="item.statusValue" />
          </el-select>
        </el-form-item>
        <el-form-item label="星级">
          <el-rate v-model="itemForm.itemStar" :colors="['#99A9BF', '#F7BA2A', '#FF9900']" :max="3" style="margin-top:8px;" />
        </el-form-item>
        <el-form-item label="Content">
          <el-input v-model="itemForm.itemContent" :autosize="{ minRows: 2, maxRows: 4}" type="textarea" placeholder="Please input" />
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">
          Cancel
        </el-button>
        <el-button type="primary" @click="dialogStatus==='create'?createData():updateData()">
          Confirm
        </el-button>
      </div>
    </el-dialog>

    <el-dialog :visible.sync="dialogPvVisible" title="Reading statistics">
      <el-table :data="pvData" border fit highlight-current-row style="width: 100%">
        <el-table-column prop="key" label="Channel" />
        <el-table-column prop="pv" label="Pv" />
      </el-table>
      <span slot="footer" class="dialog-footer">
        <el-button type="primary" @click="dialogPvVisible = false">Confirm</el-button>
      </span>
    </el-dialog>
  </div>
</template>

<script>
import { getSampleData, deleteItem } from '@/api/data-list'
import waves from '@/directive/waves' // waves directive
import { parseTime } from '@/utils'
import Pagination from '@/components/Pagination' // secondary package based on el-pagination

const typeValuesArray = [
  { typeValue: 0, typeName: '小说' },
  { typeValue: 1, typeName: '散文' },
  { typeValue: 2, typeName: '科技论文' },
  { typeValue: 3, typeName: '其他' }
]

const statusOptions = [
  { statusValue: 0, statusName: '编辑中' },
  { statusValue: 1, statusName: '审核中' },
  { statusValue: 2, statusName: '已发布' }
]

// // arr to obj, such as { CN : "China", US : "USA" }
// const itemTypeKeyValue = typeValuesArray.reduce((acc, cur) => {
//   acc[cur.key] = cur.typeName
//   return acc
// }, {})

export default {
  name: 'ComplexTable',
  components: { Pagination },
  directives: { waves },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    },
    // typeFilter(type) {
    //   return itemTypeKeyValue[type]
    // }
  },
  data() {
    return {
      tableKey: 0,
      list: null,
      total: 0,
      listLoading: false,
      listQuery: {
        page: 1,
        limit: 20,
        importance: undefined,
        title: undefined,
        type: undefined,
        sort: '+id'
      },
      importanceOptions: [1, 2, 3],
      typeValuesArray,
      statusOptions,
      sortOptions: [{ label: 'ID Ascending', key: '+id' }, { label: 'ID Descending', key: '-id' }],
      showReviewer: false,
      itemForm: {
        id: undefined,
        importance: 1,
        itemDesc: '',
        createTime: new Date(),
        itemTitle: '',
        itemType: 0,
        itemStatus: 'published',
        itemStar: 0,
        itemContent: '',
      },
      dialogFormVisible: false,
      dialogStatus: '',
      textMap: {
        update: 'Edit',
        create: 'Create'
      },
      dialogPvVisible: false,
      pvData: [],
      rules: {
        type: [{ required: true, message: 'type is required', trigger: 'change' }],
        timestamp: [{ type: 'date', required: true, message: 'timestamp is required', trigger: 'change' }],
        title: [{ required: true, message: 'title is required', trigger: 'blur' }]
      },
      downloadLoading: false
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.listLoading = true
      getSampleData(this.listQuery).then(response => {
        this.list = response.data.dataLists
        this.total = response.data.totalCounts
        console.log(this.list)
        // Just to simulate the time of the request
        setTimeout(() => {
          this.listLoading = false
        }, 1.5 * 1000)
      })
    },
    formatType(row, column) {
      return typeValuesArray[row['itemType']].typeName
    },
    statusType(row){
      return statusOptions[row['itemStatus']].statusName
    },
    handleFilter() {
      this.listQuery.page = 1
      this.getList()
    },
    handleModifyStatus(row, status) {
      this.$message({
        message: '操作Success',
        type: 'success'
      })
      row.status = status
    },
    sortChange(data) {
      const { prop, order } = data
      if (prop === 'id') {
        this.sortByID(order)
      }
    },
    sortByID(order) {
      if (order === 'ascending') {
        this.listQuery.sort = '+id'
      } else {
        this.listQuery.sort = '-id'
      }
      this.handleFilter()
    },
    resetTemp() {
      this.temp = {
        id: undefined,
        importance: 1,
        remark: '',
        timestamp: new Date(),
        title: '',
        status: 'published',
        type: ''
      }
    },
    handleCreate() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    handleCreateNew() {
      this.resetTemp()
      this.dialogStatus = 'create'
      this.dialogFormVisible = true
      this.$nextTick(() => {
        this.$refs['dataForm'].clearValidate()
      })
    },
    // createData() {
    //   this.$refs['dataForm'].validate((valid) => {
    //     if (valid) {
    //       this.itemForm.id = parseInt(Math.random() * 100) + 1024 // mock a id
    //       this.itemForm.author = 'vue-element-admin'
    //       console.log(this.itemForm)
    //
    //       addItem(this.itemForm).then(() => {
    //         //this.list.unshift(this.itemForm)
    //         this.dialogFormVisible = false
    //
    //         this.$notify({
    //           title: 'Success',
    //           message: '新增成功',
    //           type: 'success',
    //           duration: 2000
    //         })
    //         this.getList()
    //
    //       })
    //     }
    //   })
    // },
    handleAdd() {
      console.log('handler addd........................................................')
      let editUrl = '/mydatas/createOrEdit/0'
      this.$router.push({path:editUrl})
    },

    handleUpdate(row) {
      // this.temp = Object.assign({}, row) // copy obj
      // this.temp.timestamp = new Date(this.temp.timestamp)
      // this.dialogStatus = 'update'
      // this.dialogFormVisible = true
      // this.$nextTick(() => {
      //   this.$refs['dataForm'].clearValidate()
      // })

      console.log("--------------------------------->>>")
      console.log(row['itemId'])

      //this.$router.push({path:'/mydatas/mydataList/'})

      let editUrl = '/mydatas/createOrEdit/'+row['itemId']

      this.$router.push({path:editUrl})

      //const title = 'Update Sample Data'
      //const route = Object.assign({}, this.tempRoute, { title: `${title}` })

      //先关闭已经打开的
      // this.$store.dispatch('tagsView/delView', route)
      //this.$store.dispatch('tagsView/updateVisitedView', route)
      //this.$router.push({path:'/mydatas/mydataList/'})


    },
    // updateData() {
    //   this.$refs['dataForm'].validate((valid) => {
    //     if (valid) {
    //       const tempData = Object.assign({}, this.temp)
    //       tempData.timestamp = +new Date(tempData.timestamp) // change Thu Nov 30 2017 16:41:05 GMT+0800 (CST) to 1512031311464
    //       updateArticle(tempData).then(() => {
    //         const index = this.list.findIndex(v => v.id === this.temp.id)
    //         this.list.splice(index, 1, this.temp)
    //         this.dialogFormVisible = false
    //         this.$notify({
    //           title: 'Success',
    //           message: 'Update Successfully',
    //           type: 'success',
    //           duration: 2000
    //         })
    //       })
    //     }
    //   })
    // },
    handleDeleteConfirm(row) {
      this.$confirm('确认删除？')
        .then(_ => {
          console.log('点击了确认')
          console.log(row['itemId'])
          deleteItem(row['itemId']).then(() => {
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
    // handleFetchPv(pv) {
    //   fetchPv(pv).then(response => {
    //     this.pvData = response.data.pvData
    //     this.dialogPvVisible = true
    //   })
    // },
    handleDownload() {
      this.downloadLoading = true
      import('@/vendor/Export2Excel').then(excel => {
        const tHeader = ['timestamp', 'title', 'type', 'importance', 'status']
        const filterVal = ['timestamp', 'title', 'type', 'importance', 'status']
        const data = this.formatJson(filterVal)
        excel.export_json_to_excel({
          header: tHeader,
          data,
          filename: 'table-list'
        })
        this.downloadLoading = false
      })
    },
    formatJson(filterVal) {
      return this.list.map(v => filterVal.map(j => {
        if (j === 'timestamp') {
          return parseTime(v[j])
        } else {
          return v[j]
        }
      }))
    },
    getSortClass: function(key) {
      const sort = this.listQuery.sort
      return sort === `+${key}` ? 'ascending' : 'descending'
    }
  }
}
</script>
