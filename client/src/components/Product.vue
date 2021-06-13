<template>
  <div class="product">
      <h4>{{name}}</h4>
      <p>Price: ${{price}}</p>
  </div>
</template>

<script>

export default {
    name: 'Product',
    props: {
      name: String,
      price: Number
    }
}
</script>