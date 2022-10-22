<template>
    <div :class="'light light-'+color">
        ⬤ {{title}}
    </div>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'Light',
    props: {
        title: String,
        desc: String,
        color: Number
    },
})
</script>

<style>
.light {
    display: inline-block;
    margin: 0.1rem 0.5rem;
}
.light-0 {
    color: gray;
}
.light-1 {
    color: blue;
}
.light-2 {
    color: green;
}
.light-3 {
    color: orange;
}
.light-4 {
    color: red;
}
</style>