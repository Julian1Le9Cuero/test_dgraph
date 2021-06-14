<template>

  <v-card
    class="mx-auto"
    width="244"
  >
    <v-card-text>
      <p class="text-h5 text--primary">
        Transaction ID: <br> 
        <span class="trnsact-id">{{id_num}}</span> 
      </p>
      <div class="text--primary">
        {{showDate ? ` Made on: ${date.slice(0,10)}` : ""}} <br>
        {{!showDate ? ` Buyer ID: ${buyer_id}` : ""}} <br>
        Device: {{device}}
      </div>
    </v-card-text>

  </v-card>
</template>

<script>
export default {
    name: 'Transaction',
    props: {
      date: String,
      device: String,
      id_num: String,
      buyer_id: String,
      showDate: Boolean
    }
}
</script>