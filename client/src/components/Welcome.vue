<template>
  <v-container>
    <v-row class="text-center">
      <v-col cols="12">
        <v-img
          :src="require('../assets/logo.svg')"
          class="my-3"
          contain
          height="200"
        />
      </v-col>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to Dgraph Shopper
        </h1>

        <p class="subheading font-weight-regular">
          Take a look at the
          <br>
          buyers, products and transactions
        </p>
        <p class="error" v-if="error">{{error}}</p>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Buyers
        </h2>

        <v-row justify="center">
          <div class="item-grid">
            <buyer v-for="(buyer, idx) in response.buyers"
            :key="buyer.id"
            :id_num="buyer.id"
            :name="buyer.name"
            :age="buyer.age"
            :index="idx"/>
          </div>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Products
        </h2>

        <v-row justify="center">
          <div class="item-grid">
            <product v-for="(product, idx) in response.products"
            :key="product.id"
            :name="product.name"
            :price="product.price"
            :index="idx"/>
          </div>
        </v-row>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
        <h2 class="headline font-weight-bold mb-3">
          Transactions
        </h2>

        <v-row justify="center">
          <div class="item-grid">
          <div class="transactions"
            v-for="(transaction, idx) in response.transactions"
            :key="transaction.id"
            :item="transaction"
            :index="idx"
          >
            <p class="text"><strong>Transaction ID:</strong> {{transaction.id}}</p>
            <p class="text">Buyer ID: {{transaction.buyer.id}}</p>
            <p class="text">Device: {{transaction.device}}</p>
          </div>
          </div>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HomeService from '../HomeService'
import Product from '../components/Product.vue'
import Buyer from '../components/Buyer.vue'

  export default {
    name: 'Welcome',
    components:{
      Product,
      Buyer
    },
    data: () => ({
        response: {},
        error: ''
    }),
    // Lifecycle method that runs once the component is created
    async created() {
      try {
        this.response = await HomeService.getAllData()
      } catch (err) {
        this.error = err.message
      }
    },
  }
</script>
