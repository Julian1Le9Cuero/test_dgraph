<template>
<div class="pagination mt-6" >
    <span v-if="skip > 1" class="showLess" @click="updatePagItemsDown">
        ...
    </span>
    <paginationItem @change-active-number="updateActiveNumber" v-for="(number, idx) in Array.from({length: pagItemsToShow}, (_, i) => i + increaseAmount)"
            :key="idx"
            :itemNumber="number"
            :currentActiveNumber="currentActiveNumber"/>
            <!-- Show more if skips done are less than maximum number of skips -->
    <span class="showMore"
    v-if="skip <= Math.ceil(itemsLength/(maxPagItemsToShow*itemsToShow) - 1)" 
    @click="updatePagItemsUp">...</span>
</div>
</template>

<script>
import PaginationItem from './PaginationItem.vue'

export default {
    name: 'Pagination',
    data() {
        return { 
            pagItemsToShow: 0,
            maxPagItemsToShow: 10,
            increaseAmount: 1,
            currentActiveNumber: 1,
            skip: 1,
        }
    },
    components:{
        PaginationItem
    },
    props: {
        itemsLength: Number,
        itemsToShow: Number,
    },
    methods: {
        updatePagItemsUp(){
            // Update numbers shown
            this.increaseAmount = (this.maxPagItemsToShow * this.skip) + 1
            // Update skips
            this.skip++
            // Check if user cannot skip to update PaginationItems
            const {itemsLength, itemsToShow} = this.$props
            if (this.skip > Math.ceil(itemsLength/(this.maxPagItemsToShow*itemsToShow) - 1)) {
                this.pagItemsToShow = Math.ceil(itemsLength/itemsToShow) - this.maxPagItemsToShow*(this.skip-1)
            }
            // if(itemsLength < Math.ceil(itemsLength/itemsToShow)*this.maxPagItemsToShow*this.skip){
            //    this.pagItemsToShow = Math.ceil(itemsLength/itemsToShow) - this.maxPagItemsToShow*(this.skip-1)
            // } 
        
        },
        updatePagItemsDown(){
            this.increaseAmount -= this.maxPagItemsToShow
            this.skip--
            const {itemsLength, itemsToShow} = this.$props
            // Check if user can skip to show 10 items again
            if (this.skip <= Math.ceil(itemsLength/(this.maxPagItemsToShow*itemsToShow) - 1)) {
                this.pagItemsToShow = this.maxPagItemsToShow
            }
            // if(itemsLength > Math.ceil(itemsLength/itemsToShow)*this.maxPagItemsToShow*this.skip){
            //    this.pagItemsToShow = this.maxPagItemsToShow
            // } 
        },
        updateActiveNumber(activeNumber){
            this.$emit('change-active-number', activeNumber)
            this.currentActiveNumber = activeNumber
        }
    },
    created(){
        const {itemsLength, itemsToShow} = this.$props
        if (Math.ceil(itemsLength/itemsToShow) < this.maxPagItemsToShow) {
            this.pagItemsToShow = Math.ceil(itemsLength/itemsToShow)
        } else{
            this.pagItemsToShow = this.maxPagItemsToShow
        }
    }
}
</script>

<style scoped>
    .pagination{
        display: flex;
        flex-direction: row;
        position: absolute;
        top: 50%;
        left: 50%;
        transform: translate(-50%, -50%);
    }
    .showMore, .showLess{
        font-size: 20px;
        color: rgb(168, 162, 162);
        letter-spacing: 1px;
        cursor: pointer;
    }
    .showLess{
        margin-right: 6px;
    }
</style>