<template>
  <div class="pagination-table">
    <h3>Данные из БД:</h3>
    <div id="app">{{data}}</div>
    <h5>Тестовые данные (хардкодинг):</h5>
<!-- 
    <ul id="example-1">
    <li v-for="item in data" :key="item.id">
        {{ item.name }}
    </li>
    </ul> -->

    <ul id="v-for-object" class="demo">
    <li v-for="value in demoData"  v-bind:key="value.title">
        {{ value }}
    </li>
    </ul>

    <!-- <grid :cols="cols" :rows="rows"></grid> -->
    <h5>Пример таблицы с пагинацией и сортировкой:</h5>
    <grid
    :auto-width="autoWidth"
    :cols="cols"
    :from="from"
    :language="language"
    :pagination="pagination"
    :rows="rows"
    :search="search"
    :server="server"
    :sort="sort"
    :width="width"
    ></grid>

    <!-- TODO: render dynamic pagination table 20 rows on page -->
  </div>
</template>

<script>
import Grid from 'gridjs-vue'

const axios = require("axios");

export default {
  name: "TablePagination",
    components: {
      Grid
    },
  data() {
    return {
        // cols: ['Make', 'Model', 'Year', 'Color'],
        // rows: [
        //   ['Ford', 'Fusion', '2011', 'Silver'],
        //   ['Chevrolet', 'Cruz', '2018', 'White']
        // ],
        
        // An array containing strings of column headers
        cols: ['id','Name', 'Surname', 'Email'],

        // AND EITHER an array containing row data
        rows: [
          ['1', 'Иван', 'Пупкин', 'some@mail.ru'],
          ['2', 'Митрофанов', 'Илья', 'some2@mail.ru'],
        ],

        data: [],
        demoData: {
            object: {
                title: 'How to do lists in Vue',
                author: 'Jane Doe',
                publishedAt: '2016-04-10'
            },
            object2: {
                title: 'Example2',
                author: 'Vasya',
                publishedAt: '2020-08-15'
            }
        },
        
        //       // OR a server settings object
        // server() ({
        //   url: 'https://api.com/search?q=my%20query',
        //   then: res => res.data.map(col => [col1.data, col2.data]),
        //   handle: res => res.status === 404
        //     ? { data: [] } : res.ok
        //     ? res.json() : new Error('Something went wrong')
        // }),

        pagination: true || {},
        sort: true,
        autoWidth: true,
    };
  },
  beforeMount() {
    this.getData();
  },

  methods: {
    async getData() {
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
