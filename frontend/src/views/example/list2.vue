<template>
  <div class="app-container">
    <el-table v-loading="listLoading" :data="list" border fit highlight-current-row style="width: 100%">
      safasfasf
    </el-table>

  </div>
</template>

<script>
import { getGolangData } from '@/api/article'
import Pagination from '@/components/Pagination' // Secondary package based on el-pagination

export default {
  name: 'ArticleList',
  components: { Pagination },
  filters: {
    statusFilter(status) {
      const statusMap = {
        published: 'success',
        draft: 'info',
        deleted: 'danger'
      }
      return statusMap[status]
    }
  },
  data() {
    return {
      list: null,
      total: 0,
      listLoading: true,
      listQuery: {
        page: 1,
        limit: 20
      }
    }
  },
  created() {
    this.fetchUpdateLaboratoryData()
  },
  methods: {


    fetchUpdateLaboratoryData() {
      new Promise((resolve, reject) => {
        getGolangData()
          .then(res => {
            if (res.code == "200") {
              console.log(res.data)
            }
            resolve();
          })
          .catch(error => {
            console.log(error)
            reject(error);
          });
      }).catch(error => {
        console.log(error)
      });
    },

    getList() {
      this.listLoading = true
      getGolangData().then(response => {
        console.log(response)
        this.listLoading = false
      })
    }
  }
}
</script>

<style scoped>
.edit-input {
  padding-right: 100px;
}
.cancel-btn {
  position: absolute;
  right: 15px;
  top: 10px;
}
</style>
