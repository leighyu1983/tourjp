<template>
  <div class="hello">
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaa</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>
    <p>bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb</p>

    <el-upload
        class='ensure ensureButt'
        :action="importFileUrl"
        :data="upLoadData"
        name="importfile"
        :onError="uploadError"
        :onSuccess="uploadSuccess"
        :beforeUpload="beforeAvatarUpload"
        >
      <el-button size="large" type="primary" style="flex:none;width:200px;">点击上传</el-button>
    </el-upload>
    
    <div style="margin:10px 0;">
      <mt-button type="primary" @click="goAjaxPage()">primary</mt-button>
    </div>
    
    <h1>{{ msg }}</h1>
    <h2>Essential Links</h2>
    <ul>
      <li>
        <a
          href="https://vuejs.org"
          target="_blank"
        >
          Core Docs
        </a>
      </li>
      <li>
        <a
          href="https://forum.vuejs.org"
          target="_blank"
        >
          Forum
        </a>
      </li>
      <li>
        <a
          href="https://chat.vuejs.org"
          target="_blank"
        >
          Community Chat
        </a>
      </li>
      <li>
        <a
          href="https://twitter.com/vuejs"
          target="_blank"
        >
          Twitter
        </a>
      </li>
      <br>
      <li>
        <a
          href="http://vuejs-templates.github.io/webpack/"
          target="_blank"
        >
          Docs for This Template
        </a>
      </li>
    </ul>
    <h2>Ecosystem</h2>
    <ul>
      <li>
        <a
          href="http://router.vuejs.org/"
          target="_blank"
        >
          vue-router
        </a>
      </li>
      <li>
        <a
          href="http://vuex.vuejs.org/"
          target="_blank"
        >
          vuex
        </a>
      </li>
      <li>
        <a
          href="http://vue-loader.vuejs.org/"
          target="_blank"
        >
          vue-loader
        </a>
      </li>
      <li>
        <a
          href="https://github.com/vuejs/awesome-vue"
          target="_blank"
        >
          awesome-vue
        </a>
      </li>
    </ul>
  </div>
</template>

<script>
export default {
  name: 'HelloWorld',
  data () {
    return {
      importFileUrl: 'http://localhost:805/web/shop/image/upload',
      upLoadData: {
          shop_id: 'eae945d4-fc74-4ecd-81db-47592ce126d2', 
      },
      msg: 'Welcome to Your Vue.js App'
    }
  },
  methods: {
    goAjaxPage() {
      this.$router.push({ name: 'AjaxPage', params: { userId: 123 }});
    }
  }
}
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h1, h2 {
  font-weight: normal;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
