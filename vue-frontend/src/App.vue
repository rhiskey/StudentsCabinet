<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->
    <!-- <HelloWorld msg="Welcome to Your Vue.js App"/> -->
    <TablePagination data="[data: passed]"/>
    <h5>График 1:</h5>
    <Graph1 dataGraph="[x: y, x1:y1]"/>
    <h5>График 2:</h5>
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import TablePagination from './components/TablePagination.vue'
import Graph1 from './components/Graph1.vue'
// import Graph2 from './components/Graph2.vue'

export default {
  name: 'App',
  components: {
    // HelloWorld,
    TablePagination,
    Graph1,
    // Graph2
  }
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>