<template>
  <v-container>
    <v-row>

      <v-col class="mb-4">
        <div class="text-center">
        <h1 class="display-2 font-weight-bold mb-3">
          Welcome to Dgraph Shopper
        </h1>

        <p class="subheading font-weight-regular">
          Take a look at the
          <br>
          buyers, products and transactions
        </p>
        <p class="error" v-if="error">{{error}}</p>
         </div>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
      <div class="text-center">
        <h1 class="headline font-weight-bold mb-6">
          Buyers
        </h1>
      </div>

          <div class="item-grid">
            <buyer v-for="(buyer, idx) in response.buyers"
            :key="buyer.id"
            :id_num="buyer.id"
            :name="buyer.name"
            :age="buyer.age"
            :date="buyer.date"
            :showDate="false"
            :index="idx"/>
          </div>

      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
      <div class="text-center">
        <h1 class="headline font-weight-bold mb-6">
          Products
        </h1>
      </div>
          <div class="item-grid">
            <product v-for="(product, idx) in response.products"
            :key="product.id"
            :name="product.name"
            :price="product.price"
            :index="idx"/>
          </div>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >
      <div class="text-center">
        <h1 class="headline font-weight-bold mb-6">
          Transactions
        </h1>
      </div>

        <div class="item-grid">
            <transaction v-for="(transaction, idx) in response.transactions"
            :key="transaction.id"
            :id_num="transaction.id"
            :date="transaction.date"
            :device="transaction.device"
            :buyer_id="transaction.buyer.id"
            :showDate="false"
            :index="idx"/>
        </div>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import HomeService from '../HomeService'
import Product from '../components/Product.vue'
import Buyer from '../components/Buyer.vue'
import Transaction from '../components/Transaction.vue'

  export default {
    name: 'Welcome',
    components:{
      Product,
      Buyer,
      Transaction
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
