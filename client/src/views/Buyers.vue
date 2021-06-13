<template>
  <v-container>
    <v-row class="text-center">

      <v-col class="mb-4">
        <h1 class="display-2 font-weight-bold mb-3">
          Take a look at the buyers
        </h1>

        <p class="error" v-if="error">{{error}}</p>
      </v-col>

      <v-col
        class="mb-5"
        cols="12"
      >

        <v-row justify="center">
          <div class="item-grid">
          <div class="buyers"
            v-for="(buyer, idx) in buyers"
            :key="buyer.id"
            :item="buyer"
            :index="idx"
          >
            <p>Member since {{buyer.date.slice(0,10)}}</p>
            <p class="text">Name: <a class="text" :href="`buyers/${buyer.id}`">{{buyer.name}}</a></p>
            <p class="text">Age: {{buyer.age}}</p>
          </div>
          </div>
        </v-row>
      </v-col>

    </v-row>
  </v-container>
</template>

<script>
import HomeService from '../HomeService'

  export default {
    name: 'Buyers',
    data: () => ({
        buyers: [],
        error: ''
    }),
    // Lifecycle method rthat runs once the component is created
    async created() {
      try {
        if (this.buyers.length > 0) {
          this.buyers = this.buyers
        } else{
          this.buyers = await HomeService.getBuyers()
        }
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