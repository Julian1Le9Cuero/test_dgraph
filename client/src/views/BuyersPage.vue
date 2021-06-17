<template>
  <v-container>

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Take a look at the buyers
        </h1>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >

        <v-row justify="center">
            <Buyers :itemsToShow="itemsToShow" :date="null" 
            :itemsLimit="550" :showDate="true"/>
        </v-row>
      </v-col>

  </v-container>
</template>

<script>
import Buyers from '../components/Buyers.vue'

  export default {
    name: 'BuyersPage',
    components: {
      Buyers
    },
    data: () => ({
        itemsToShow: 20
    })
    // Lifecycle method rthat runs once the component is created
    // async created() {
    //   try {
    //     if (this.buyers.length > 0) {
    //       this.buyers = this.buyers
    //     } else{
    //       this.buyers = await HomeService.getBuyers()
    //     }
    //   } catch (err) {
    //     this.error = err.message
    //   }
    // },
  }
</script>
