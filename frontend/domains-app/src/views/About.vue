<template>
  <div class="about">
    <h1>This is an about page</h1>
  <p v-for="item in movies">
    {{ item.price }}
  </p>
  </div>
</template>

<script>
export default {
  data: () => {
    return {
      movies: [{
        name:'spiderman',
        price:200
      },{
                name:'superman',
        price:500
      }]
    }
  }
}
</script>