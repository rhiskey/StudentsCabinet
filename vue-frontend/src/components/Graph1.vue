<template>
  <div class="graph1">
    <!-- <h1>{{ dataGraph }}</h1> -->
   <!-- TODO: render diagramm d3 -->
   <D3BarChart :config="config" :datum="data"></D3BarChart>
  </div>
</template>

<script>
import {D3BarChart} from 'vue-d3-charts'

export default {
  name: 'Graph1',
    components: {
      D3BarChart
    },
     data() {
        return {
            data :[{
  name: "Lorem",
  total: 40
},{
  name: "Ipsum",
  total: 21
},{
  name: "Dolor",
  total: 20
}],
config: {
  key: "name",
  value: "total",
  color: "steelblue",
}

        };
    },
  props: {
    dataGraph: []
  }
}

// // data
// data = [{
//   name: "Lorem",
//   total: 30
// },{
//   name: "Ipsum",
//   total: 21
// },{
//   name: "Dolor",
//   total: 20
// }]

// // Configuration
// config = {
//   key: "name",
//   value: "total",
//   color: "steelblue",
// }

</script>