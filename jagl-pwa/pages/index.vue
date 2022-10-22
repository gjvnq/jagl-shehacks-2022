<template>
  <ol>
    <li><NuxtLink to="/users/list">List Users</NuxtLink></li>
    <!-- <li><NuxtLink to="/profiles/me">My Profile</NuxtLink></li> -->
  </ol>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
  name: 'IndexPage'
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
