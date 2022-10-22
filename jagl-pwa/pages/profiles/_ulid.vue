<template>
    <Profile/>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'ProfilePage',
    mounted() {
        console.log(this.$route.params);
    }
})
</script>

<style>
:root {
    max-width: 360px;
    max-height: 800px;
    margin: auto;
}
</style>
