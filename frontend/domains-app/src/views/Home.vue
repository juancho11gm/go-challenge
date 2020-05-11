<template>
  <div class="home">
    <SearchDomain msg="Welcome to the domain App"/>
  </div>
</template>

<script>
// @ is an alias to /src
import SearchDomain from '@/components/SearchDomain.vue'

export default {
  name: 'Home',
  components: {
    SearchDomain
  }
}


</script>
