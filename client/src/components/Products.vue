<template>
    <div>
        <div class="text-center">
            <p class="error" v-if="error">{{error}}</p>
            <p v-if="resp.product_query.length === 0">There are no products to show for the selected date</p>
        </div>
        <div class="item-grid" v-if="!isLoading">
        <product v-for="(product, idx) in resp.product_query.slice(offset, offset+itemsToShow)"
            :key="product.id"
            :name="product.name"
            :price="product.price"
            :index="idx"/>
        </div>
        <div class="container mt-6 mb-6">
        <Pagination :itemsToShow="itemsToShow" 
        :itemsLength="resp.product_query.length"
        @change-active-number="updateOffset"/>
    </div>
    </div>
</template>

<script>
import HomeService from '../HomeService'
import Product from '../components/Product.vue'
import Pagination from '../components/Pagination.vue'

export default {
    name: 'Products',
    components: {
        Product,
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
    // Lifecycle method that runs once the component is created
    async created() {
      try {
        const filters = {date: this.$props.date, limit: 180, offset: this.offset, model: 'Product'}
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
