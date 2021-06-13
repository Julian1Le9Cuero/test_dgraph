<template>
  <div class="transaction">
      <h4>ID: {{id_num}}</h4>
      <p>Made on: {{date.slice(0,10)}}</p>
      <p>Device: {{device}}</p>
  </div>
</template>

<script>
export default {
    name: 'Transaction',
    props: {
      date: String,
      device: String,
      id_num: String
    }
}
</script>