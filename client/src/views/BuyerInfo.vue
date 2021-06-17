<template>
   <div class="about pa-6">
      <p class="error" v-if="error">{{error}}</p>
      <h1 class="headline font-weight-bold mb-3">
          Your transactions
      </h1>
    <!-- Loop over transactions -->
    <div class="item-grid">
          <transaction v-for="(transaction, idx) in response.transactions"
            :key="transaction.id"
            :id_num="transaction.id"
            :date="transaction.date"
            :device="transaction.device"
            :buyer_id="transaction.buyer.id"
            :showDate="true"
            :index="idx"/>
          </div>
          <!-- Loop over products -->
    <h1 class="headline font-weight-bold mt-6 mb-3">Products bought</h1>
          <div class="item-grid">
          <product v-for="(product, idx) in response.products"
            :key="product.id"
            :name="product.name"
            :price="product.price"
            :index="idx"/>
          </div>
          <!-- Loop over buyers -->
    <h1 class="headline font-weight-bold mt-6 mb-3">Buyers similar to you</h1>
      <div class="item-grid">
            <buyer v-for="(buyer, idx) in response.buyers"
            :key="buyer.id"
            :id_num="buyer.id"
            :name="buyer.name"
            :age="buyer.age"
            :date="buyer.date"
            :showDate="true"
            :index="idx"/>
      </div>
      <!-- Loop over product recommendations -->
    <h1 class="headline font-weight-bold mt-6 mb-3">Products you may like</h1>
      <div class="item-grid">
            <product v-for="(product, idx) in response.productRecommendations"
            :key="product.id"
            :name="product.name"
            :price="product.price"
            :index="idx"/>
      </div>
  </div>
</template>

<script>
import HomeService from '../HomeService'
import Buyer from '../components/Buyer.vue'
import Product from '../components/Product.vue'
import Transaction from '../components/Transaction.vue'

export default {
  name: 'BuyersInfo',
  components: {
      Buyer,
      Product,
      Transaction
    },
    data: () => ({
        response: {},
        error: '',
    }),
     created(){
        this.fetchData()
    },
    watch: {
        // Call again the method if route is changed
        '$route': 'fetchData'
    },
    methods: {
        async fetchData(){
          const fetchedId = this.$route.params.id
            try {
               this.response = await HomeService.getBuyerById(fetchedId)
            } catch (err) {
               this.error = err.message
            }
        }
    },
}

</script>