/** When your routing table is too long, you can split it into small modules **/

import Layout from '@/layout'

const mydatasRouter = {
  path: '/mydatas',
  component: Layout,
  redirect: 'noRedirect',
  alwaysShow: true,
  name: 'MydataDemo',
  meta: {
    title: 'Mydatas',
    icon: 'qq'
  },
  children: [
    {
      path: 'mydataList',
      component: () => import('@/views/mydata/mydata-table'),
      name: 'mydataList',
      meta: { title: 'mydataList' }
    }, {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/example/edit'),
      name: 'EditArticle',
      meta: { title: 'Edit Article', noCache: true, activeMenu: '/example/list' },
      hidden: true
    }, {
      path: 'create',
      component: () => import('@/views/mydata/mydata-create'),
      name: 'CreateArticle',
      meta: { title: 'Create Article' },
      hidden: true
    },
  ]
}

export default mydatasRouter
