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
      itemId: 5009,
      itemKey: '99933',
      meta: { title: 'mydataLista', itemKey: '5009' }
    }, {
      path: 'edit/:id(\\d+)',
      component: () => import('@/views/example/edit'),
      name: 'EditSample',
      meta: { title: 'Edit Data', noCache: true, activeMenu: '/example/list' },
      hidden: true
    }, {
      // path: 'create/:id(\\d+)',
      path: 'createOrEdit/:itemId(\\d+)',
      component: () => import('@/views/mydata/mydata-createOrEdit'),
      name: 'CreateSample',
      meta: { title: 'Create Data' },
      hidden: true
    },
  ]
}

export default mydatasRouter
