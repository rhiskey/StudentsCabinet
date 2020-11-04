<template>
  <div class="graph1">
    <h1>{{ dataGraph }}</h1>
   <!-- TODO: render diagramm d3 -->
  </div>
</template>

<script>
import d3 from ''

export default {
  name: 'Graph1',
  props: {
    dataGraph: []
  }
}
</script>