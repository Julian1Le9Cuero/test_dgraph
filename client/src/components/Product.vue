<template>

  <v-card
    class="mx-auto"
    width="244"
  >
    <v-card-text>
      <p class="product-price">Price: ${{price}}</p>
      <div class="text--primary">
        {{name}}
      </div>
    </v-card-text>

  </v-card>
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