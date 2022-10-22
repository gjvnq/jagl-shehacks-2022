<template>
    <Profile v-if="!edit_mode"/>
    <EditProfile v-else-if="edit_mode"/>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'ProfilePage',
    computed: {
        edit_mode() {
            return (this.$route.query.edit === 'true')
        }
    },
    mounted() {
        console.log('route params', this.$route.params)
        console.log('query', this.$route.query)
    },
})
</script>

<style>
:root {
    padding: 1rem;
    max-width: 360px;
    max-height: 800px;
    margin: auto;
}
</style>
