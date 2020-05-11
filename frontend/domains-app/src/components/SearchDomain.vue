<template>
  <div class="hello">
    <h1>{{ msg }}</h1>
    <p>Here you can get info about the domain that you want!</p>
    <b-container class="bv-example-row">
      <b-row>
        <b-input-group size="sm" prepend="domain" class="mt-3">
          <b-form-input id="inputDomain" v-model="domain"></b-form-input>
          <b-input-group-append>
            <b-button v-on:click="requestDomain(domain)" variant="outline-success">Request info</b-button>
          </b-input-group-append>
        </b-input-group>
      </b-row>

      <h1 v-if="requested">
        <h2>Results</h2>
        <b-container class="bv-example-row">
          <b-row>{{info}}</b-row>
        </b-container>
      </h1>
    </b-container>
  </div>
</template>

<script>
export default {
  name: "SearchDomain",
  data() {
    return {
      domain: "",
      info: "",
      requested: false
    };
  },
  props: {
    msg: String
  },
  methods: {
    requestDomain: function(domain) {
      if (domain == "") alert("The domain must be different from null");
      let self = this;
      this.$http
        .get("http://localhost:8081/domain/" + domain)
        .then(response => {
          this.requested = true;
          this.info = response.body
        });
    }
  }
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
