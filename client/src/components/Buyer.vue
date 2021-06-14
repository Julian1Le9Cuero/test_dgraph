<template>
  <v-card
    class="mx-auto"
    max-width="344"
  >
    <v-card-text>
      <div>{{showDate ? `Member since ${date.slice(0,10)}` : ""}}</div>
      <p class="text-h4 text--primary">
        {{name}}
      </p>
      <p>Age: {{age}}</p>
      <div class="text--primary">
        relating to or dependent on charity; charitable.<br>
        "an eleemosynary educational institution."
      </div>
    </v-card-text>
    <v-card-actions>
      <v-btn
        text
        color="teal accent-4"
        @click="sendTobuyer(id_num)"
      >
        Learn More
      </v-btn>
      </v-card-actions>

  </v-card>

</template>

<script>

export default {
    name: 'Buyer',
    props: {
      id_num: String,
      name: String,
      age: Number,
      date: String,
      showDate: Boolean
    },
    methods: {
      sendTobuyer(id_num){
        this.$router.push(`/buyers/${id_num}`)
      }
    },
}
</script>
