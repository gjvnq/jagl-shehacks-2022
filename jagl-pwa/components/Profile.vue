<template>
    <b-card :title="profile.name">
        <p v-if="fronting_alters.length != 0"><span style="color: gray">Co-fronting alters:</span> {{fronting_alters.join(', ')}}</p>
        <div style="text-align: center">
            <p style="margin-bottom:0">{{current_time}}</p>
            <p style="color: gray">{{profile.timezone}}</p>
        </div>
        <b-card border-variant="warning" header="Warnings" class="text-center" style="margin-bottom: 1rem">
            <b-card-text>
                <p v-for="notice in profile.notices">
                    {{notice}}
                </p>
            </b-card-text>
        </b-card>
        <div style="display: block">
            <Light v-for="light in profile.lights" :title="light.title" :color="light.color" :desc="light.desc"/>
        </div>
        <br/>
        <b-button variant="outline-primary" href="?edit=true">Edit profile</b-button>
    </b-card>
</template>

<script lang="ts">
import Vue from 'vue'
import { Person } from '~/data_model.ts'

export default Vue.extend({
    name: 'Profile',
    data() {
        return {
            profile: new Person(),
            now: Date.now(),
            date_formatter: new Intl.DateTimeFormat(),
        }
    },
    computed: {
        current_time() {
            return this.date_formatter.format(this.now)
        },
        fronting_alters() {
            return this.profile.alters.filter(alter => alter.fronting).map(alter => alter.title)
        },
    },
    created() {
        var self = this
        setInterval(function () {
            self.now = Date.now()
        }, 1000)
        setInterval(function () {
            self.update_data()
        }, 2000)
    },
    methods: {
        async update_data() {
            const new_profile = await fetch('http://localhost:3001/profiles/'+this.profile.ulid).then(res => res.json())
            if (this.profile.timezone != new_profile.timezone) {
                this.date_formatter = new Intl.DateTimeFormat(['pt-BR', 'en-US'], {
                    timeZone: new_profile.timezone,
                    hour12: false,
                    year: 'numeric',
                    month: 'numeric',
                    day: 'numeric',
                    hour: 'numeric',
                    minute: 'numeric',
                    second: 'numeric',
                    weekday: 'short',
                })
            }
            this.profile = new_profile
        }
    },
    async fetch() {
        this.profile.ulid = this.$route.params.ulid
        await this.update_data()
    },
})
</script>