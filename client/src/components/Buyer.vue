<template>
  <div class="buyer">
            <p class="text">Name: <router-link :to="`/buyers/${id_num}`">{{name}}</router-link></p>
            <p class="text">Age: {{age}}</p>
  </div>
</template>

<script>

export default {
    name: 'Buyer',
    props: {
      id_num: String,
      name: String,
      age: Number
    }
}
</script>
