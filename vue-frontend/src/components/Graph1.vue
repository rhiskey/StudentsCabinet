<template>
  <div class="graph1">
    <!-- <h1>{{ dataGraph }}</h1> -->
    <!-- TODO: render diagramm d3 -->
  </div>
</template>

<script>
// import d3 from ''
const axios = require("axios");

export default {
  name: "Graph1",
  props: {
    dataGraph: [],
    startDate: Date,
    endDate: Date,
  },
  components: {},

  beforeMount(){
    this.mounted();
  },

  methods: {
    mounted() {
      // Pass data from date selector
      // POST request using axios with error handling
      const datePeriod = { startDate: "2020-08-01", endDate: "2020-11-05" };
      axios
        .post("http://127.0.0.1:3000/graph1", datePeriod)
        .then((response) => (this.articleId = response.data.id))
        .catch((error) => {
          this.errorMessage = error.message;
          console.error("There was an error!", error);
        });
    },

    // // Fire on date change
    // async sendData() {
    //   //const { data } = await axios.get("http://127.0.0.1:3000/graph1");
    //   //this.data = data;

    //   const { data } = await axios({
    //     method: "POST",
    //     url: "http://127.0.0.1:3000/graph1",
    //     data: data,
    //     headers: { "content-type": "text/plain" },
    //   })
    //     .then((result) => {
    //       // this.response = result.data;
    //       this.add = result.data["add"];
    //       this.mul = result.data["mul"];
    //       this.sub = result.data["sub"];
    //       this.div = result.data["div"];
    //       /*eslint-disable*/
    //       console.log(result.data);
    //       /*eslint-enable*/
    //     })
    //     .catch((error) => {
    //       /*eslint-disable*/
    //       console.error(error);
    //       /*eslint-enable*/
    //     });
    //   this.data = data;
    // },
  },
};
</script>