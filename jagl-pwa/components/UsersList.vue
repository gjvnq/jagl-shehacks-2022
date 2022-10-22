<template>
    <b-container>
        <b-row v-for="user in users" v-bind:key="user.ulid" style="margin-bottom: 1rem">
            <b-col>
                <b-card :title="user.name">
                    <b-link v-bind:href="'/profiles/' + user.ulid" class="card-link">See profile</b-link>
                </b-card>
            </b-col>
        </b-row>
    </b-container>
</template>

<script lang="ts">
import Vue from 'vue'

export default Vue.extend({
    name: 'UsersList',
    data() {
        return {
            users: []
        }
    },
    async fetch() {
        this.users = await fetch('http://localhost:3001/users/list').then(res => res.json())
    },
})
</script>
