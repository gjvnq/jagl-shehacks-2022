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

export default Vue.extend({
    name: 'Profile',
    data() {
        return {
            profile: {
                alters: [],
                lights: [],
            },
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
    },
    async fetch() {
        const profile_ulid = this.$route.params.ulid;
        this.profile = await fetch('http://localhost:3001/profiles/'+profile_ulid).then(res => res.json())
        console.log(this.profile)
        this.date_formatter = new Intl.DateTimeFormat([], {
            timeZone: this.profile.timezone,
            year: 'numeric',
            month: 'numeric',
            day: 'numeric',
            hour: 'numeric',
            minute: 'numeric',
            second: 'numeric',
        });

    },
})
</script>