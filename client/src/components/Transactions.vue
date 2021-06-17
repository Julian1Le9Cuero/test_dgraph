<template>
    <div>
        <div class="text-center">
            <p class="error" v-if="error">{{error}}</p>
            <p v-if="resp.trnsact_query.length === 0">There are no transactions to show for the selected date</p>
        </div>
        <div class="item-grid" v-if="!isLoading"> 
        <transaction v-for="(transaction, idx) in resp.trnsact_query.slice(offset, offset+itemsToShow)"
            :key="transaction.id"
            :id_num="transaction.id"
            :date="transaction.date"
            :device="transaction.device"
            :buyer_id="transaction.buyer.id"
            :showDate="false"
            :index="idx"/>
        </div>
        <div class="container mt-6 mb-6">
        <Pagination :itemsToShow="itemsToShow" 
        :itemsLength="resp.trnsact_query.length"
        @change-active-number="updateOffset"/>
        </div>
    </div>
</template>

<script>
import HomeService from '../HomeService'
import Transaction from '../components/Transaction.vue'
import Pagination from '../components/Pagination.vue'

export default {
    name: 'Transactions',
    components: {
        Transaction,
        Pagination
    },
    data: () => ({
        resp: {},
        error: '',
        offset: 0,
        isLoading: true
    }),
    props: {
        date: Number,
        itemsToShow: Number
    },
    async created() {
      try {
        const filters = {date: this.$props.date, limit: 1000, offset: this.offset, model: 'Transaction'}
        this.resp = await HomeService.getAllData(filters)
        if (this.resp !== undefined) {
            this.isLoading = false
        }
      } catch (err) {
        this.error = err.message
      }
    },
    methods: {
        updateOffset(activeNumber){
            this.offset = (activeNumber - 1) * this.$props.itemsToShow
        }
    },
}
</script>

