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

      <div class="results" v-if="requested">
        <h2>Results about {{domain}}</h2>
        <b-img  :src="info[0].logo" width="500px" alt="Organization logo"></b-img>
        <b-container class="bv-example-row">
          <b-row>
              <b-table size="sm" stacked :fields="fields" :items="info"></b-table>
          </b-row>
          <h3>Endpoints</h3>
           <b-row>
              <b-table size="sm" stacked :items="info[0].Endpoints"></b-table>
          </b-row>
        </b-container>
      </div>
    </b-container>
  </div>
</template>

<script>
export default {
  name: "SearchDomain",
  data() {
    return {
      domain: "",
      fields: [
          'ServersChanged',
          'ssl_grade',
          'previous_ssl_grade',
          'logo',
          'title',
          'Status',
        ],
      info: [
      ],
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
          this.info.push(response.body);
          this.requested = true;
        });
    },
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>

a {
  color: #42b983;
}

h3 {
  margin: 40px 0 0;
}

li {
  display: inline-block;
  margin: 0 10px;
}

ul {
  list-style-type: none;
  padding: 0;
}

.results {
  margin: 20px;
}

</style>
