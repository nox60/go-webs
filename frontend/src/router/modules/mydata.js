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
    }
  ]
}

export default mydatasRouter
