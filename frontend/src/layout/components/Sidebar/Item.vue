<script>
export default {
  name: 'MenuItem',
  functional: true,
  props: {
    icon: {
      type: String,
      default: ''
    },
    title: {
      type: String,
      default: ''
    },
    accessCode: {
      type: Number,
      default: 999
    }
  },
  render(h, context) {
    const { icon, title } = context.props
    // console.log('parent pro')
    console.log(context.props)
    const vnodes = []

    if (icon) {
      if (icon.includes('el-icon')) {
        vnodes.push(<i class={[icon, 'sub-el-icon']} />)
      } else {
        vnodes.push(<svg-icon icon-class={icon}/>)
      }
    }

    // console.log('------------------------------')
    // console.log(context.props)
    // console.log('----------====================')

    if (title) {
      vnodes.push(<span slot='title'>{(title)} </span>)
    }

    // if (itemKey) {
    //   vnodes.push(<span slot='itemKey'>{(itemKey)} </span>)
    // }
    return vnodes
  }
}
</script>

<style scoped>
.sub-el-icon {
  color: currentColor;
  width: 1em;
  height: 1em;
}
</style>
