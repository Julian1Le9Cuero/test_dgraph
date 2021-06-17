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
         </div>

        <div class="ml-10">
         <p class="filterByDate-btn"><strong>Filter by date:</strong></p>
          <v-date-picker
           v-model="date"
           year-icon="mdi-calendar-blank"
           prev-icon="mdi-skip-previous"
           next-icon="mdi-skip-next"
           @click:date="handleDateClick"
          ></v-date-picker>
         <p class="selectedDate" v-if="selectedDate !== null"><strong>Selected date: {{selectedDate}}</strong></p>
      </div>
      </v-col>
      <!-- Add filter date -->
      <v-col
        class="mb-5"
        cols="12"
      >
      <div class="text-center">
        <h1 class="headline font-weight-bold mb-6">
          Buyers
        </h1>
      </div>
            <Buyers :itemsToShow="itemsToShow" :date="filterDate" 
            :itemsLimit="180" :showDate="false" :key="filterDate"/>
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
            <Products :itemsToShow="itemsToShow" :date="filterDate"
            :key="filterDate"/>
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
        <Transactions :itemsToShow="itemsToShow" :date="filterDate"
        :key="filterDate"/>
      </v-col>
    </v-row>
  </v-container>
</template>

<script>
import Products from '../components/Products.vue'
import Buyers from '../components/Buyers.vue'
import Transactions from '../components/Transactions.vue'

  export default {
    name: 'Home',
    components:{
      Products,
      Buyers,
      Transactions
    },
    data: () => ({
        itemsToShow: 20,
        filterDate: null,
        selectedDate: null
    }),
    methods: {
      handleDateClick(date){
          this.selectedDate = date
          // Convert date to unix format
          this.filterDate = new Date(date).getTime() / 1000 + 24*60*60
      }
    },  
  }
</script>
