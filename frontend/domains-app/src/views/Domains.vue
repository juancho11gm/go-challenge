<template>
  <div class="about">
    <h1>Here you can see the requested domains</h1>
    <main>
      <div>
        <b-container class="bv-example-row">
          <b-row>
            <b-table striped hover :items="info"></b-table>
          </b-row>
        </b-container>
      </div>
    </main>
  </div>
</template>

<script>
//https://es.vuejs.org/v2/cookbook/using-axios-to-consume-apis.html
export default {
  data() {
    return {
      info: []
    };
  },
  methods: {},
  created() {
    //function exec when te component it's called
    this.$http
      .get("http://localhost:8081/")
      .then(response => (this.info = response.body.items));
  }
};
</script>