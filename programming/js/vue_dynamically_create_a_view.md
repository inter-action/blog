# VueJS 动态生成视图 | VueJS - create view dynamically

由于公司项目的原因，这几天在看vue，中间遇到个问题就是如何动态低生成页面元素，由于vue框架本身的原因，让这个特性有些难处理。折腾了几个小时后，有了个初步的解决方案。

需求是这样的，项目中用了 [element-ui](http://element.eleme.io/#/zh-CN/component/dialog), 我需要动态地创建弹出框中的内容，但是组件库中的定义让我没有简单的方式这么做，下面是组件库的使用文档:

```vue
<el-dialog
  title="提示"
  :visible.sync="dialogVisible"
  size="tiny"
  :before-close="handleClose">
  <span>这是一段信息</span>
  <span slot="footer" class="dialog-footer">
    <el-button @click="dialogVisible = false">取 消</el-button>
    <el-button type="primary" @click="dialogVisible = false">确 定</el-button>
  </span>
</el-dialog>

<script>
  export default {
    data() {
      return {
        dialogVisible: false
      };
    },
    methods: {
      handleClose(done) {
        this.$confirm('确认关闭？')
          .then(_ => {
            done();
          })
          .catch(_ => {});
      }
    }
  };
</script>
```

而我要实现的api大概是这样 

```js
  dialog.open('<h1 onclick="">clickme </h1>', function onConfirm(close){close()}, function onCancel(close){close()})
```

我说下这个问题的思路，首先我想能否通过像angular1的那种方式，就是通过compiler将模板编译，然后link到对应的scope上:

```js
// 声明的包装组件
let Dialog = {
  data() {
    return {
      dialogVisible: false,
      tpl: '',
    }
  },
  render(h) {
    return this.compiled.render.apply(this, arguments)
  },
  computed: {
    compiled() {
      let root = `
        <el-dialog title="提示" :visible.sync="dialogVisible" size="tiny" :before-close="handleClose">
          ${this.tpl}
          <span slot="footer" class="dialog-footer">
            <el-button @click="handleCancel">取 消</el-button>
            <el-button type="primary" @click="handleConfirm">确 定</el-button>
          </span>
        </el-dialog>
      `
      return Vue.compile(root)
    }
  },
  methods: {
    open(tpl, onConfirm, onCancel) {
      this.tpl = tpl
      if (onConfirm) this.onConfirm = onConfirm
      if (onCancel) this.onCancel = onCancel
    },
    close() {
      this.dialogVisible = false
    },
    handleClose() {
      this.close();
      // empty
    },
    handleCancel() {
      if (this.onCancel && typeof this.onCancel === 'function') {
        this.onCancel(() => this.dialogVisible = false)
      }
    },
    handleConfirm() {
      if (this.onConfirm && typeof this.onConfirm === 'function') {
        this.onConfirm(() => this.dialogVisible = false)
      }
    }
  }
}


// 使用者

export default{
  data(){
    return {}
  },

  components:{
    'confirm-dialog': Dialog
  },

  methods:{
    open(){
      this.$refs.dialog.open('<h1>hi</hi>', close=>close(), close=>close())
    }
  }
}
```

```vue
<template>
<!-- 使用者的template -->
  <confirm-dialog ref="dialog"></confirm-dialog>
</template>
```

这种思路可以通过动态得改变openDialog中的tpl字符串来改变dialog中的内容，但这种方式的问题是vue并没有angular1中的scope概念, 没有办法将compile生成的对象函数link到
调用者的context(scope)中, 这意味着这种实现没有有效的方式去绑定一些函数的回调，获取dialog中自己插入的form表单的内容等等。

所以，我就在想如果换成vue 中的 createElement 这种方式看看

```js

render(h) {
  return h(
    'el-dialog',
    {
      props: {
        visible: this.dialogVisible,
        size: 'tiny',
        beforeClose: this.handleClose.bind(this)
      }
    },
    [
      this.tpl && this.tpl(h),
      h('span', {
        props: {
          slot: 'footer',
        },
        attrs: {
          class: 'dialog-footer'
        }
      }, [
          h('el-button', {
            on: {
              click: this.handleCancel.bind(this),
            }
          }, '取消'),
          h('el-button', {
            props: {
              type: 'primary'
            },
            on: {
              click: this.handleConfirm.bind(this),
            }
          }, '确定'),
        ])
    ]
  )
},


```

调用方式就变成了 
```js
this.$refs.dialog.open(h=>h('h1', {props:{onClick: ()=>console.debug('i got clicked')}}, 'click me '), close=>close(), close=>close())
```

这种方式解决了问题，但是可以看到动态的内容写法太臃肿，可读性也不高, 换成[jsx](https://github.com/vuejs/babel-plugin-transform-vue-jsx)语法之后


```jsx
  render(h) {
    return (
      <el-dialog title="提示" visible={this.dialogVisible} size="tiny" beforeClose={this.handleClose.bind(this)} onClose={this.close.bind(this)}>
        {this.tpl && this.tpl(h)}
        <span slot="footer" className="dialog-footer">
          <el-button onClick={this.handleCancel.bind(this)}>取消</el-button>
          <el-button type="primary" onClick={this.handleConfirm.bind(this)}>确定</el-button>
        </span>
      </el-dialog>
    )
  }

  methods:{
    open({ title = '提示', template, onConfirm, onCancel }) {
      this.tpl = template;
      this.dialogVisible = true
      if (onConfirm) this.onConfirm = onConfirm
      if (onCancel) this.onCancel = onCancel
    },
  }
```

调用方式优化为

```jsx
  opendialog() {
    this.$refs.dialog.open({
      template: h =>
        (
          <h1 onClick={() => this.reload()}>
            click me
        </h1>
        ),
      onConfirm: close => close(),
      onCancel: close => close()

    })
  }
```

可以看到通过这种方式，解决了弹出框组件自定义内容的问题，但是还有些问题需要优化的。
* 首先，这种api，在引入封装的dialog组件后，需要在调用者components中声明，然后在调用者template插入此组件节点，比较繁琐
* 还有可以看出vue在动态创建组件上，完全没有react简单，可以说比较复杂。
* ref 因为不能接收回调，`ref={v=>this.x=v}`这种形式导致没有有效的方式通过这种新式引用到动态嵌入组件定义的方法。