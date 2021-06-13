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
          <div class="buyer"
            v-for="(buyer, idx) in buyers"
            :key="buyer.id"
            :item="buyer"
            :index="idx"
          >
          {{buyer.date.slice(0,10)}}
            <a class="text" :href="buyer.id">{{buyer.name}}</a>
            <p class="text">Age: {{buyer.age}}</p>
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
          <div class="products"
            v-for="(product, idx) in products"
            :key="product.id"
            :item="product"
            :index="idx"
          >

            <p class="text">{{product.name}}</p>
            <p class="text">Price: {{`$${product.price}`}}</p>
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
          <div class="transactions"
            v-for="(transaction, idx) in transactions"
            :key="transaction.id"
            :item="transaction"
            :index="idx"
          >

            <p class="text">Transaction ID: {{transaction.id}}</p>
            <p class="text">Buyer ID: {{transaction.buyer.id}}</p>
            <p class="text">Device: {{transaction.device}}</p>
          </div>
        </v-row>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HomeService from '../HomeService'

  export default {
    name: 'HelloWorld',

    data: () => ({
        buyers: [],
        products: [],
        transactions: [],
        error: ''
    }),
    // Lifecycle method rthat runs once the component is created
    async created() {
      try {
        this.buyers = await HomeService.getAllBuyers()
        this.products = await HomeService.getAllProducts()
        this.transactions = await HomeService.getAllTransactions()
      } catch (err) {
        this.error = err.message
      }
    },
  }
</script>

<style scoped>
p.error{
  border: 1px solid #ff5b5f;
  background-color: #ffc5c1;
  padding: 10px;
  margin-bottom: 15px;
}
</style>