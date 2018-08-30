<template>
  <div class="hello">
    <p>hmmm:{{this.$route.params.userId}}</p>
    <div style="margin:10px 0;">
      <mt-button type="primary" @click="testAjaxGet()">ajax get</mt-button>
      <mt-button type="primary" @click="testAjaxPost()">ajax post</mt-button>
    </div>
  </div>
</template>

<script>
export default {
  name: 'AjaxPage',
  methods: {
    testAjaxGet() {
      this.$ajax.get('/web/shop/get?id=1', this).then((res) => {
        console.log('-----success-----');
        console.log(res['data']['data']);
      }).catch((err) => {
        console.log('-----error-----');
        console.log(err);
      })
    },
    testAjaxPost() {
      let param = {
        "id": "1",
        "jp_name": "",
        "jp_description": "",
        "cn_name": "",
        "cn_description": ""
      }

      this.$ajax.post('/web/shop/create', param, this).then((res) => {
        console.log('-----success-----');
        console.log(res['data']['data']);
      }).catch((err) => {
        console.log('-----error-----');
        console.log(err);
      })
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
