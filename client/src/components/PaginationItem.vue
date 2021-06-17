<template>

<div class="pagination-item text-center" 
:class="currentActiveNumber === itemNumber ? 'pagination-item-active' : ''" 
@click="changeToActive(itemNumber)">
    {{itemNumber}}
</div>

</template>

<script>
export default {
    name: 'PaginationItem',
    props: {
        itemNumber: Number,
        currentActiveNumber: Number
    },
    methods: {
        changeToActive(activeNumber){
            // Emit the current number to change offset in request
            this.$emit('change-active-number', activeNumber)
        }
    }
}
</script>

<style lang="scss" scoped>
    .pagination-item, .pagination-item-active{
        cursor: pointer;
        width: 40px;
        height: 40px;
        border-radius: 100%;
        box-shadow: -5px 5px 5px rgb(0, 0, 0, 0.4);
        padding-top: 8px;
        margin-right: 6px;
    }
    .pagination-item{
        background-color: #fff;
        color: blue;
        &-active{
            background-color: blue;
            color: #fff;
        }
    }
</style>
