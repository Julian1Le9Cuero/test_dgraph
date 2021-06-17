<template>
    <div>
    <div class="text-center">
     <p class="error" v-if="error">{{error}}</p>
     <p v-if="resp.buyer_query.length === 0">There are no buyers to show for the selected date</p>
    </div>
    <div class="item-grid" v-if="!isLoading & resp.buyer_query.length > 0">
        <buyer v-for="(buyer, idx) in resp.buyer_query.slice(offset, offset+itemsToShow)"
            :key="buyer.id"
            :id_num="buyer.id"
            :name="buyer.name"
            :age="buyer.age"
            :date="buyer.date"
            :showDate="showDate"
            :index="idx"/>
    </div>
    <div class="container mt-6 mb-6">
    <Pagination :itemsToShow="itemsToShow" 
        :itemsLength="resp.buyer_query.length"
        @change-active-number="updateOffset"/>
    </div>
    </div>
</template>

<script>
import HomeService from '../HomeService'
import Buyer from '../components/Buyer.vue'
import Pagination from '../components/Pagination.vue'

export default {
    name: 'Buyers',
    data: () => ({
        resp: {},
        error: '',
        offset: 0,
        isLoading: true
    }),
    components: {
        Buyer,
        Pagination
    },
    props: {
        date: Number,
        itemsToShow: Number,
        itemsLimit: Number,
        showDate: Boolean
    },
    // Lifecycle method that runs once the component is created
    async created() {
      try {
        const filters = {date: this.$props.date, limit: this.$props.itemsLimit, offset: this.offset, model: 'Buyer'}
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
