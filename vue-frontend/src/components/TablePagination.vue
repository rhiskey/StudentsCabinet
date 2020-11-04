<template>
  <div class="pagination-table">
    <div id="app">{{data}}</div>

    <ul id="example-1">
    <li v-for="item in lists" :key="item.id">
        {{ item.name }}
    </li>
    </ul>
    <!-- TODO: render dynamic pagination table 20 rows on page -->
  </div>
</template>

<script>
// import ''
const axios = require("axios");

export default {
  name: "TablePagination",
  //   props: {
  //     data: [],
  //   },

  data() {
    return {
    //   loading: false,
    //   post: null,
    //   error: null,
      data: [],
    };
  },
  beforeMount() {
    this.getData();
  },
//   created() {
//     // fetch the data when the view is created and the data is
//     // already being observed
//     this.fetchData();
//   },
//   watch: {
//     // call again the method if the route changes
//     $route: "fetchData",
//   },

  methods: {
    // fetchData() {
    //   this.error = this.post = null;
    //   this.loading = true;
    //   const fetchedId = this.$route.params.id;
    //   // replace `getPost` with your data fetching util / API wrapper
    //   axios(fetchedId, (err, post) => {
    //     // make sure this request is the last one we did, discard otherwise
    //     if (this.$route.params.id !== fetchedId) return;
    //     this.loading = false;
    //     if (err) {
    //       this.error = err.toString();
    //     } else {
    //       this.post = post;
    //     }
    //   });
    // },
    async getData() {
    //   const res = await fetch("https://api.agify.io/?name=michael");
    //   const data = await res.json();
    //   this.data = data;
    const { data } = await axios.get("http://127.0.0.1:3000/table");
      this.data = data;
      
    //   axios({ method: "POST", url: "http://127.0.0.1:8090/calc", data: data, headers: {"content-type": "text/plain" } }).then(result => { 
    //       // this.response = result.data;
    //       this.add = result.data['add']
    //       this.mul = result.data['mul']
    //       this.sub = result.data['sub']
    //       this.div = result.data['div']
    //       /*eslint-disable*/
    //       console.log(result.data) 
    //       /*eslint-enable*/
    //     }).catch( error => {
    //         /*eslint-disable*/
    //         console.error(error);
    //         /*eslint-enable*/
    //   });
    },
  },
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
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
