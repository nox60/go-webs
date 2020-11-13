import { login, getInfo } from '@/api/user'
import { getToken, setToken, removeToken } from '@/utils/auth'
import router, { resetRouter } from '@/router'
const jwt = require('jsonwebtoken')

const state = {
  token: getToken(),
  name: '',
  avatar: '',
  introduction: '',
  roles: [],
  menus: '',
  items: '',
  userStatus: 0,
  userType: 0
}

const mutations = {
  SET_TOKEN: (state, token) => {
    state.token = token
  },
  SET_INTRODUCTION: (state, introduction) => {
    state.introduction = introduction
  },
  SET_NAME: (state, name) => {
    state.name = name
  },
  SET_AVATAR: (state, avatar) => {
    state.avatar = avatar
  },
  SET_ROLES: (state, roles) => {
    state.roles = roles
  },
  SET_MENUS: (state, menus) => {
    state.menus = menus
  },
  SET_ITEMS: (state, items) => {
    state.items = items
  },
  SET_USER_STATUS: (state, userStatus) => {
    state.userStatus = userStatus
  },
  SET_USER_TYPE: (state, userType) => {
    state.userType = userType
  }
}

const actions = {
  // user login
  login({ commit }, userInfo) {
    const { username, password } = userInfo
    return new Promise((resolve, reject) => {
      login({ username: username.trim(), password: password }).then(response => {
        const { data } = response
        commit('SET_TOKEN', data.accountId)
        setToken(data.token)
        console.log('-----------------------------------------111111')
        // console.log(Base64.decode(data.token))
        console.log('+++++++++++++++++++++++++++++++++++++++++++++++')
        resolve()
      }).catch(error => {
        reject(error)
      })
    })
  },
  loginInfo({ commit }, token) {
    // Write vuex
    // console.log('-------------------- write into store')

    console.log(token)

    const resultaa = jwt.decode(token)

    console.log(resultaa.sub)

    const parsedJson = JSON.parse(resultaa.sub)
    console.log('parsed json............', parsedJson)
    // console.log(">>>>>---")
    // console.log(parsedJson)
    // console.log(parsedJson.MenuItems)
    // console.log(parsedJson.pageItems)
    // console.log('-------------------------------------------------------')
    commit('SET_MENUS', parsedJson.menuItems)
    commit('SET_ITEMS', parsedJson.pageItems)
    commit('SET_TOKEN', token)
    commit('SET_USER_STATUS', parsedJson.userStatus)
    commit('SET_USER_TYPE', parsedJson.userType)

    setToken(token)
  },
  // get user info
  getInfo({ commit, state }) {
    console.log('------------------------999999999999999999999999999999')
    return new Promise((resolve, reject) => {
      getInfo().then(response => {
        const { data } = response

        if (!data) {
          reject('Verification failed, please Login again.')
        }

        const { roles, name, avatar, introduction } = data

        // roles must be a non-empty array
        if (!roles || roles.length <= 0) {
          reject('getInfo: roles must be a non-null array!')
        }

        commit('SET_ROLES', roles)
        commit('SET_NAME', name)
        commit('SET_AVATAR', avatar)
        commit('SET_INTRODUCTION', introduction)
        resolve(data)
      }).catch(error => {
        reject(error)
      })
    })
  },

  // user logout
  logout({ commit, state, dispatch }) {
    return new Promise((resolve, reject) => {
      // logout(state.token).then(() => {
      //   commit('SET_TOKEN', '')
      //   commit('SET_ROLES', [])
      //   removeToken()
      //   resetRouter()
      //
      //   // reset visited views and cached views
      //   // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
      //   dispatch('tagsView/delAllViews', null, { root: true })
      //
      //   resolve()
      // }).catch(error => {
      //   reject(error)
      // })
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      commit('SET_MENUS', '')
      commit('SET_ITEMS', '')
      removeToken()
      resetRouter()

      // reset visited views and cached views
      // to fixed https://github.com/PanJiaChen/vue-element-admin/issues/2485
      dispatch('tagsView/delAllViews', null, { root: true })
      resolve()
    })
  },

  // remove token
  resetToken({ commit }) {
    return new Promise(resolve => {
      commit('SET_TOKEN', '')
      commit('SET_ROLES', [])
      removeToken()
      resolve()
    })
  },

  // dynamically modify permissions
  async changeRoles({ commit, dispatch }, role) {
    const token = role + '-token'

    commit('SET_TOKEN', token)
    setToken(token)

    const { roles } = await dispatch('getInfo')

    resetRouter()

    // generate accessible routes map based on roles
    const accessRoutes = await dispatch('permission/generateRoutes', roles, { root: true })
    // dynamically add accessible routes

    router.addRoutes(accessRoutes)

    // reset visited views and cached views
    dispatch('tagsView/delAllViews', null, { root: true })
  }
}

export default {
  namespaced: true,
  state,
  mutations,
  actions
}
