<template>
  <div id="app">
    <!-- <img alt="Vue logo" src="./assets/logo.png"> -->

    <!-- Pick date and load specified table and graphs -->

       <!-- <date-range-picker
            ref="picker"
            :opens="opens"
            :locale-data="{ firstDay: 1, format: 'DD-MM-YYYY HH:mm:ss' }"
            :minDate="minDate" :maxDate="maxDate"
            :singleDatePicker="singleDatePicker"
            :timePicker="timePicker"
            :timePicker24Hour="timePicker24Hour"
            :showWeekNumbers="showWeekNumbers"
            :showDropdowns="showDropdowns"
            :autoApply="autoApply"
            v-model="dateRange"
            @update="updateValues"
            @toggle="checkOpen"
            :linkedCalendars="linkedCalendars"
            :dateFormat="dateFormat"
    > -->

       <date-range-picker
            ref="picker"
            :opens="center"
            :locale-data="{ firstDay: 1, format: 'DD-MM-YYYY HH:mm:ss' }"
        >

        <template v-slot:input="picker" style="min-width: 350px;">
            {{ picker.startDate | date }} - {{ picker.endDate | date }}
        </template>
    </date-range-picker>

    <TablePagination data="[data: passed]"/>
    <h5>Количество обучающихся в разрезе типа обучающегося (выпускник/студент):</h5>
    <Graph1 dataGraph="[x: y, x1:y1]"/>
    <h5>Количество обучающихся в разрезе вида аттетации (ГИА/Промежуточная аттестация):</h5>
  </div>
</template>

<script>
// import HelloWorld from './components/HelloWorld.vue'
import TablePagination from './components/TablePagination.vue'
import Graph1 from './components/Graph1.vue'
// import Graph2 from './components/Graph2.vue'

import DateRangePicker from 'vue2-daterange-picker'
import 'vue2-daterange-picker/dist/vue2-daterange-picker.css'

export default {
  name: 'App',
  components: {
    // HelloWorld,
    TablePagination,
    Graph1,
    DateRangePicker,
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